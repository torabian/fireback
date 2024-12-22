package workspaces

import (
	"fmt"
	"strings"

	"github.com/lnquy/cron"
)

type DescribeContext struct {

	// Only include specific module names
	IncludeOnly []string
}

func generateMarkdownTable(headers []string, rows [][]string) string {
	var sb strings.Builder

	// Add headers
	sb.WriteString("| " + strings.Join(headers, " | ") + " |\n")

	// Add separator
	sb.WriteString("|" + strings.Repeat(" --- |", len(headers)) + "\n")

	// Add rows
	for _, row := range rows {
		sb.WriteString("| " + strings.Join(row, " | ") + " |\n")
	}

	return sb.String()
}

// Focuses on the yaml module2 content
func DescribeModule2(m *Module2, item *ModuleProvider) string {
	content := []string{}

	// Let's describe the general information about the module first

	content = append(content, "## "+ToUpper(m.Name))
	if m.Description != "" {
		content = append(content, "Description: "+m.Description)
	}
	if m.Namespace != "" {
		content = append(content, "Namespace: "+m.Namespace)
	}

	// let's give some info about the entities which are in this module

	if len(m.Entities) > 0 {

		content = append(content, "\r\n")
		content = append(content, "### "+ToUpper(m.Name)+" Entities")

		{
			headers := []string{"Name", "Usage", "Main data"}
			rows := [][]string{}

			for _, entity := range m.Entities {
				mainData := []string{}

				for _, field := range entity.Fields {
					mainData = append(mainData, ToUpper(field.Name))
				}

				rows = append(rows, []string{ToUpper(entity.Name), entity.Description, strings.Join(mainData, ", ")})
			}

			table := generateMarkdownTable(headers, rows)
			content = append(content, table+"\r\n")
		}

	}
	if len(m.Tasks) > 0 {
		content = append(content, "\r\n")
		content = append(content, "### "+ToUpper(m.Name)+" Tasks")
		{
			headers := []string{"Name", "Usage", "Triggers"}
			rows := [][]string{}

			for _, task := range m.Tasks {
				triggerInfo := []string{}
				for _, trigger := range task.Triggers {
					if trigger.Cron != nil {
						v := *trigger.Cron

						exprDesc, _ := cron.NewDescriptor()

						desc, err := exprDesc.ToDescription(*trigger.Cron, cron.Locale_en)
						// "Every minute"
						if err == nil {
							v += " (" + desc + ")"
						}

						triggerInfo = append(triggerInfo, v)
					}
				}

				rows = append(rows, []string{task.Name, task.Description, strings.Join(triggerInfo, ", ")})
			}

			table := generateMarkdownTable(headers, rows)
			content = append(content, table+"\r\n")
		}
	}
	// actions in the definition - this is not enough. A lot of actions are virtual
	// and injected in the ModuleProvider

	if item.ActionsBundle != nil {

		content = append(content, "\r\n")
		content = append(content, "### "+ToUpper(m.Name)+" actions ("+fmt.Sprintf("%v", len(item.ActionsBundle.Actions))+")")

		for _, action := range item.ActionsBundle.Actions {
			content = append(content, "#### **"+action.Name+"**")
			content = append(content, action.Description)
			content = append(content, "*Url:*: "+action.Url+" ("+action.Method+")")
			content = append(content, "\r\n\r\n")
		}
	}
	content = append(content, "\r\n\r\n")

	return strings.Join(content, "\r\n")
}

func Describe(xapp *XWebServer, ctx *DescribeContext) string {

	content := []string{}

	// Let's add overall information about the project
	content = append(content, "# "+xapp.Title)

	stat := CountXappModules(xapp)
	content = append(content, fmt.Sprintf("Total modules: %v", stat.TotalModules))
	content = append(content, fmt.Sprintf("Modules overview: %s", strings.Join(stat.ModuleNames, ", ")))

	for _, item := range xapp.Modules {
		if item.Definitions == nil {
			continue
		}

		if len(ctx.IncludeOnly) > 0 {
			if !Contains(ctx.IncludeOnly, item.Name) {
				continue
			}
		}

		defFile, err := GetSeederFilenames(item.Definitions, "")

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		for _, path := range defFile {
			var mod2 Module2
			ReadYamlFileEmbed(item.Definitions, path, &mod2)

			content = append(content, DescribeModule2(&mod2, item))

		}

	}

	return strings.Join(content, "\r\n")
}

type XAppModuleInfo struct {
	TotalModules int
	ModuleNames  []string
}

// Counts modules in an xwebserver app.
// it also goes through children and count them as well
func CountXappModules(xapp *XWebServer) XAppModuleInfo {
	stat := XAppModuleInfo{
		TotalModules: 0,
	}

	var counter func(m []*ModuleProvider)
	counter = func(m []*ModuleProvider) {
		for _, item := range m {
			stat.TotalModules++
			stat.ModuleNames = append(stat.ModuleNames, item.Name)
			if len(item.Children) > 0 {
				counter(item.Children)
			}
		}
	}

	counter(xapp.Modules)

	return stat
}
