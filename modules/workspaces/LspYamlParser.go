package workspaces

import (
	"fmt"
	"slices"
	"strings"
)

type Context struct {
	Path []string
}

func (ctx *Context) UpdateContext(line string, indentChar rune, indentSize int) {
	currentIndent := 0
	for _, char := range line {
		if char == indentChar {
			currentIndent++
		} else {
			break
		}
	}

	level := currentIndent / indentSize
	key := strings.TrimSpace(line)

	if len(ctx.Path) <= level {
		ctx.Path = append(ctx.Path, key)
	} else {
		ctx.Path[level] = key
		ctx.Path = ctx.Path[:level+1]
	}
}

func countLeadingSpaces(line string) int {
	return len(line) - len(strings.TrimLeft(line, " "))
}

func isArrayItem(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "-")
}

func cleanKey(line string) string {

	index := strings.Index(line, ":")
	if index > -1 {
		line = line[:index]
	}

	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(line, "-", ""), ":", ""))
}

type YamlContext struct {
	Label   string
	Index   int
	Indent  int
	IsArray bool
}

func compileToString(items []*YamlContext) string {
	result := []string{}

	for index, item := range items {
		if index == len(items)-1 {
			continue
		}
		if item.IsArray {
			result = append(result, item.Label+"/"+fmt.Sprintf("%v", item.Index)+"")
		} else {
			result = append(result, item.Label)
		}
	}
	result = slices.DeleteFunc(result, func(input string) bool {
		return strings.TrimSpace(input) == ""
	})

	return strings.Join(result, "/")
}

func compileToStringForEmpty(items []*YamlContext) string {
	result := []string{}

	for index, item := range items {
		if item.IsArray && index != len(items)-1 {
			result = append(result, item.Label+"/"+fmt.Sprintf("%v", item.Index)+"")
		} else {
			result = append(result, item.Label)
		}
	}
	result = slices.DeleteFunc(result, func(input string) bool {
		return strings.TrimSpace(input) == ""
	})

	return strings.Join(result, "/")
}

func GetContextByLineAndCol(file string, lineNo int, colNo int) string {

	num := 0
	depth := 0

	context := []*YamlContext{
		{
			Label: "$document",
		},
	}

	normalizedString := strings.ReplaceAll(file, "\r\n", "\n")
	lines := strings.Split(normalizedString, "\n")

	// index := -1
	for _, line := range lines {
		num++

		if len(strings.TrimSpace(line)) == 0 {
			if num == lineNo {
				return compileToStringForEmpty(context[:(colNo / 2)])
			}
			continue
		}

		lineKey := cleanKey(line)
		// fmt.Printf("Line: %v %s\n", num, line)
		lineIndent := countLeadingSpaces(line)

		// fmt.Println("indent:", lineIndent)

		currentDepth := lineIndent / 2

		if currentDepth == depth {
			context[len(context)-1] = &YamlContext{
				Label: lineKey,
			}

			if isArrayItem(line) {

				if context[len(context)-2] != nil {
					context[len(context)-2].Index++
				}
			}

		} else {
			if currentDepth < depth {
				actualLeft := depth - currentDepth + 1
				context = context[:len(context)-actualLeft]
				if len(context) > 1 && context[len(context)-1] != nil && context[len(context)-1].IsArray {
					context[len(context)-1].Index++
				}
			} else {
				if len(context) > 0 && isArrayItem(line) {
					context[len(context)-1].IsArray = true
				}
			}

			context = append(context, &YamlContext{
				Label: lineKey,
			})
		}

		// after that, let's set the depth for next line
		depth = currentDepth

		// fmt.Println(line)
		if num == lineNo {
			return compileToString(context)
		}
	}

	return ""
}
