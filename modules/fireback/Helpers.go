package fireback

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"text/template"
)

func ToUpper(t string) string {
	if t == "" {
		return ""
	}
	return strings.ToUpper(t[0:1]) + t[1:]
}
func ToLower(t string) string {
	if t == "" {
		return ""
	}
	return strings.ToLower(t[0:1]) + t[1:]
}

func ToCamelCaseClean(input string) string {
	splitBySpecial := regexp.MustCompile("[^A-Za-z0-9]+")
	words := splitBySpecial.Split(input, -1)

	var result string
	for _, word := range words {
		// Convert each word to camel case
		word = strings.ToLower(word)
		if word == "" {
			continue
		}
		word = strings.ToUpper(word[0:1]) + word[1:]
		result += word
	}

	// Remove non-alphanumeric characters
	nonAlphaNumeric := regexp.MustCompile("[^A-Za-z0-9]")
	result = nonAlphaNumeric.ReplaceAllString(result, "")

	return ToLower(result)
}

var CommonMap = template.FuncMap{
	"endsWithDto": func(s string) bool {
		return strings.HasSuffix(s, "Dto")
	},
	"last": func(x int, a interface{}) bool {
		return x == reflect.ValueOf(a).Len()-1
	},
	"goComment":         goComment,
	"until":             generateRange,
	"typescriptComment": typescriptComment,
	"join":              strings.Join,
	"trim":              strings.TrimSpace,
	"upper":             ToUpper,
	"lower":             ToLower,
	"snakeUpper":        ToSnakeUpper,
	"escape":            EscapeDoubleQuotes,
	"safeIndex":         SafeIndex,
	"hasSuffix":         strings.HasSuffix,
	"regex":             regexReplace,
	"arr":               func(els ...any) []any { return els },
	"inc": func(i int) int {
		return i + 1
	},
	"fx": func(fieldName string, depth int) string {
		return fieldName + "[index" + fmt.Sprintf("%v", depth) + "]."
	},
}

func mergeMaps(map1, map2 map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})

	// Copy values from map1 to merged
	for key, value := range map1 {
		merged[key] = value
	}

	// Copy values from map2 to merged, overwriting existing keys
	for key, value := range map2 {
		merged[key] = value
	}

	return merged
}

func generateRange(start, end int) []int {
	result := make([]int, end-start+1)
	for i := range result {
		result[i] = i + start
	}
	return result
}

func SafeIndex(slice []interface{}, index int) bool {
	if index < 0 || index >= len(slice) {
		return false
	}
	return true
}

func EscapeDoubleQuotes(input string) string {
	return strings.ReplaceAll(input, `"`, `\"`)
}

func goComment(comment string) string {
	// Escape problematic characters and split into lines
	lines := strings.Split(comment, "\n")
	for i, line := range lines {
		lines[i] = "// " + strings.ReplaceAll(line, "*/", "* /") // Escape `*/`
	}
	return strings.Join(lines, "\n")
}

func typescriptComment(comment string) string {
	// Escape problematic characters and split into lines
	lines := strings.Split(comment, "\n")
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, "*/", "* /") // Escape `*/`
	}
	return strings.Join(lines, "\n")
}

func regexReplace(input, pattern, replacement string) (string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	return re.ReplaceAllString(input, replacement), nil
}

func toCamelCase(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

// dotToCamelCase converts "person.FirstName.LastName" to "person.firstName.lastName"
func dotToCamelCase(input string) string {
	parts := strings.Split(input, ".")
	for i, part := range parts {
		parts[i] = toCamelCase(part)
	}
	return strings.Join(parts, ".")
}
