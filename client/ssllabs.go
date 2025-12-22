package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"NebulaChallenge/models"
)

const (
	BaseURL   = "https://api.ssllabs.com/api/v2"
	UserAgent = "Nebula-Challenge-SSLLabs-Client/1.0"
)

// Client representa el cliente HTTP para SSL Labs API
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient crea una nueva instancia del cliente
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: BaseURL,
	}
}

// RateLimitInfo contiene información de rate limiting
type RateLimitInfo struct {
	MaxAssessments     int
	CurrentAssessments int
}

// GetRateLimitInfo extrae información de rate limiting de los headers
func (c *Client) GetRateLimitInfo(resp *http.Response) *RateLimitInfo {
	maxStr := resp.Header.Get("X-Max-Assessments")
	currentStr := resp.Header.Get("X-Current-Assessments")

	info := &RateLimitInfo{}

	if max, err := strconv.Atoi(maxStr); err == nil {
		info.MaxAssessments = max
	}

	if current, err := strconv.Atoi(currentStr); err == nil {
		info.CurrentAssessments = current
	}

	return info
}

// GetInfo obtiene información del servicio SSL Labs
func (c *Client) GetInfo() (*models.Info, error) {
	endpoint := fmt.Sprintf("%s/info", c.baseURL)

	var info models.Info
	if err := c.doRequest(endpoint, &info); err != nil {
		return nil, fmt.Errorf("error getting info: %w", err)
	}

	return &info, nil
}

// StartAnalysis inicia un nuevo análisis
func (c *Client) StartAnalysis(host string, publish bool) (*models.Host, error) {
	params := url.Values{}
	params.Add("host", host)
	params.Add("startNew", "on")
	params.Add("all", "done")

	if publish {
		params.Add("publish", "on")
	} else {
		params.Add("publish", "off")
	}

	endpoint := fmt.Sprintf("%s/analyze?%s", c.baseURL, params.Encode())

	var hostResult models.Host
	if err := c.doRequest(endpoint, &hostResult); err != nil {
		return nil, fmt.Errorf("error starting analysis: %w", err)
	}

	return &hostResult, nil
}

// CheckAnalysis verifica el estado del análisis
func (c *Client) CheckAnalysis(host string) (*models.Host, error) {
	params := url.Values{}
	params.Add("host", host)
	params.Add("all", "done")

	endpoint := fmt.Sprintf("%s/analyze?%s", c.baseURL, params.Encode())

	var hostResult models.Host
	if err := c.doRequest(endpoint, &hostResult); err != nil {
		return nil, fmt.Errorf("error checking analysis: %w", err)
	}

	return &hostResult, nil
}

// doRequest realiza una petición HTTP GET y parsea la respuesta JSON
func (c *Client) doRequest(endpoint string, result interface{}) error {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("User-Agent", UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	// Verificar código de estado
	if resp.StatusCode != http.StatusOK {
		return c.handleErrorResponse(resp.StatusCode, body)
	}

	// Parsear JSON
	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	return nil
}

// handleErrorResponse maneja respuestas de error de la API
func (c *Client) handleErrorResponse(statusCode int, body []byte) error {
	switch statusCode {
	case 400:
		return fmt.Errorf("invalid parameters (400): %s", string(body))
	case 429:
		return fmt.Errorf("rate limit exceeded (429): too many requests")
	case 500:
		return fmt.Errorf("internal server error (500): %s", string(body))
	case 503:
		return fmt.Errorf("service unavailable (503): please retry in 15 minutes")
	case 529:
		return fmt.Errorf("service overloaded (529): please retry in 30 minutes")
	default:
		return fmt.Errorf("unexpected status code %d: %s", statusCode, string(body))
	}
}
