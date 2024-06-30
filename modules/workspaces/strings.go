package workspaces

import "strings"

func PickFirstNWords(s string, n int) string {
	count := 0
	out := []string{}

	strRange := strings.Split(s, " ")

	for _, item := range strRange {
		count++

		if count >= n {
			break
		}

		out = append(out, item)
	}

	affix := ""

	if len(strRange) > n {
		affix = "…"
	}
	return strings.Join(out, " ") + affix

}
