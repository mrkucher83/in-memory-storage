package compute

import (
	"regexp"
	"strings"
)

type Parser struct{}

func (p *Parser) Parse(input string) (string, string, string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return "", "", ""
	}

	re := regexp.MustCompile(`^[\w*/\-_]+$`)
	for _, arg := range parts {
		if !re.MatchString(arg) {
			return "INVALID", "", ""
		}
	}
	cmd := strings.ToUpper(parts[0])

	switch cmd {
	case "SET":
		return cmd, parts[1], parts[2]
	case "GET", "DEL":
		return cmd, parts[1], ""
	default:
		return "INVALID", "", ""
	}
}
