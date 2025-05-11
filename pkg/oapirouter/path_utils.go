package oapirouter

import "strings"

// Take a path string and replace all path parameters
// (e.g., {id}) with a regular expression pattern
// that matches any non-slash character sequence.
func ToRegularExpressionPath(path string) string {
	pathComponents := strings.Split(path, "/")
	for i, component := range pathComponents {
		if strings.HasPrefix(component, "{") && strings.HasSuffix(component, "}") {
			pathComponents[i] = "[^/]+"
		}
	}

	return strings.Join(pathComponents, "/")
}
