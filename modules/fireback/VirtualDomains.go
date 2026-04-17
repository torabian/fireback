package fireback

import (
	"os"
	"runtime"
	"strings"
)

func hostsPath() string {
	switch runtime.GOOS {
	case "windows":
		return `C:\Windows\System32\drivers\etc\hosts`
	default:
		return "/etc/hosts"
	}
}

// This can be even changed per project
const markerStart = "# local-fb-domain-sim-start"
const markerEnd = "# local-fb-domain-sim-end"

func EnableDomain(domain string) error {
	path := hostsPath()

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	content := string(data)

	domains := extractDomains(content)
	domains[domain] = true

	newBlock := buildBlock(domains)

	content = removeBlock(content)
	content = strings.TrimSpace(content) + "\n\n" + newBlock + "\n"

	return os.WriteFile(path, []byte(content), 0644)
}

func DisableDomain(domain string) error {
	path := hostsPath()

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	content := string(data)

	domains := extractDomains(content)
	delete(domains, domain)

	newBlock := buildBlock(domains)

	content = removeBlock(content)
	content = strings.TrimSpace(content) + "\n\n" + newBlock + "\n"

	return os.WriteFile(path, []byte(content), 0644)
}

func extractDomains(content string) map[string]bool {
	lines := strings.Split(content, "\n")

	inBlock := false
	domains := map[string]bool{}

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == markerStart {
			inBlock = true
			continue
		}
		if line == markerEnd {
			inBlock = false
			continue
		}

		if inBlock && line != "" {
			parts := strings.Fields(line)
			if len(parts) == 2 {
				domains[parts[1]] = true
			}
		}
	}

	return domains
}

func buildBlock(domains map[string]bool) string {
	var sb strings.Builder

	sb.WriteString(markerStart + "\n")
	for d := range domains {
		sb.WriteString("127.0.0.1 " + d + "\n")
	}
	sb.WriteString(markerEnd)

	return sb.String()
}

func removeBlock(content string) string {
	lines := strings.Split(content, "\n")
	var out []string

	skip := false
	for _, line := range lines {
		if strings.Contains(line, markerStart) {
			skip = true
			continue
		}
		if strings.Contains(line, markerEnd) {
			skip = false
			continue
		}
		if !skip {
			out = append(out, line)
		}
	}

	return strings.Join(out, "\n")
}
