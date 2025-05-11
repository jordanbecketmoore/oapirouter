package oapirouter

import (
	"regexp"
	"strings"
)

// ToKubernetesResourceName converts an arbitrary string into a valid Kubernetes resource name.
func ToKubernetesResourceName(input string) string {
	// Convert to lowercase
	name := strings.ToLower(input)

	// Replace invalid characters with "-"
	name = regexp.MustCompile(`[^a-z0-9.-]`).ReplaceAllString(name, "-")

	// Trim leading and trailing invalid characters
	name = strings.Trim(name, "-.")

	// Ensure the name is no longer than 253 characters
	if len(name) > 253 {
		name = name[:253]
	}

	return name
}
