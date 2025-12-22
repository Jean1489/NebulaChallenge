package utils

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

// ValidateHost valida que el host sea válido
func ValidateHost(host string) error {
	if host == "" {
		return fmt.Errorf("host cannot be empty")
	}

	host = strings.TrimSpace(host)

	// Remover protocolo si lo tiene
	host = strings.TrimPrefix(host, "http://")
	host = strings.TrimPrefix(host, "https://")

	// Remover path si lo tiene
	if idx := strings.Index(host, "/"); idx != -1 {
		host = host[:idx]
	}

	// Remover puerto si lo tiene
	if idx := strings.Index(host, ":"); idx != -1 {
		host = host[:idx]
	}

	// Verificar si es una IP válida
	if net.ParseIP(host) != nil {
		return nil
	}

	// Verificar si es un dominio válido
	if !isValidDomain(host) {
		return fmt.Errorf("invalid hostname format: %s", host)
	}

	return nil
}

// isValidDomain verifica si es un dominio válido
func isValidDomain(domain string) bool {
	// Debe tener al menos un punto
	if !strings.Contains(domain, ".") {
		return false
	}

	// Regex para dominio válido
	domainRegex := regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)
	return domainRegex.MatchString(domain)
}

// SanitizeHost limpia el host de protocolos, paths y puertos
func SanitizeHost(host string) string {
	host = strings.TrimSpace(host)
	host = strings.TrimPrefix(host, "http://")
	host = strings.TrimPrefix(host, "https://")

	if idx := strings.Index(host, "/"); idx != -1 {
		host = host[:idx]
	}

	if idx := strings.Index(host, ":"); idx != -1 {
		host = host[:idx]
	}

	return host
}
