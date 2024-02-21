package workspaces

import (
	"sort"
	"strconv"
	"strings"
)

type language2 struct {
	name    string
	quality float64
}

type languageSlice []language2

func (ls languageSlice) SortByQuality() {
	sort.Sort(ls)
}

func (s languageSlice) Len() int {
	return len(s)
}

func (s languageSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s languageSlice) Less(i, j int) bool {
	return s[i].quality > s[j].quality
}

// ParseAcceptLanguage returns RFC1766 language codes parsed and sorted from
// languages.
//
// If supportedLanguages is not empty, the returned codes will be filtered
// by its contents.
func GetExactLanguageFromAcceptLanguage(languages string, supportedLanguages []string) string {
	d := ParseAcceptLanguage2(languages, supportedLanguages)
	if len(d[0]) != 2 {
		return d[0][0:2]
	}
	return d[0]
}
func ParseAcceptLanguage2(languages string, supportedLanguages []string) []string {
	preferredLanguages := strings.Split(languages, ",")
	preferredLanguagesLen := len(preferredLanguages)

	// Preallocate processed languages, as we know the maximum possible.
	langsCap := preferredLanguagesLen
	if len(supportedLanguages) > 0 {
		langsCap = len(supportedLanguages)
	}
	langs := make(languageSlice, 0, langsCap)

	for i, rawPreferredLanguage := range preferredLanguages {
		// Format strings.
		preferredLanguage := strings.Replace(strings.ToLower(strings.TrimSpace(rawPreferredLanguage)), "_", "-", 0)

		if preferredLanguage == "" {
			continue
		}

		// Split out quality factor.
		parts := strings.SplitN(preferredLanguage, ";", 2)

		// If supported languages are given, return only the langs that fit.
		supported := len(supportedLanguages) == 0
		for _, supportedLanguage := range supportedLanguages {
			if supported = supportedLanguage == parts[0]; supported {
				break
			}
		}

		if !supported {
			continue
		}

		lang := language2{parts[0], 0}
		if len(parts) == 2 {
			q := parts[1]

			if strings.HasPrefix(q, "q=") {
				q = strings.SplitN(q, "=", 2)[1]
				var err error
				if lang.quality, err = strconv.ParseFloat(q, 64); err != nil {
					// Default value (1) if quality is empty.
					lang.quality = 1
				}
			}
		}

		// Use order of items if no quality is given.
		if lang.quality == 0 {
			lang.quality = float64(preferredLanguagesLen - i)
		}

		langs = append(langs, lang)

	}

	langs.SortByQuality()

	// Filter quality string.
	langString := make([]string, 0, len(langs))
	for _, lang := range langs {
		langString = append(langString, lang.name)
	}

	if len(langString) == 0 {
		langString = append(langString, "en")
	}
	return langString

}
