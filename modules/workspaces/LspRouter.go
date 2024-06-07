package workspaces

import "github.com/sourcegraph/go-lsp"

var actions = []LspHandler{
	{
		Path:    "root",
		Handler: MainHandler,
	},
	{
		Path:    "root/actions",
		Handler: ActionsHandler,
	},
	{
		Path:    "root/actions/:index/method",
		Handler: MethodHandler,
	},
	{
		Path:    "root/entities",
		Handler: EntityItemsHandler,
	},
	{
		Path:    "root/entities/:index/fields",
		Handler: ActionsFieldsHandler,
	},
	{
		Path:    "root/actions/:line/in",
		Handler: ActionsInHandler,
	},
	{
		Path:    "root/actions/:line/in/fields",
		Handler: ActionsFieldsHandler,
	},
	{
		Path:    "root/actions/:line/out/fields",
		Handler: ActionsFieldsHandler,
	},
}

func ActionsFieldsHandler() []lsp.CompletionItem {

	return []lsp.CompletionItem{

		{
			Label: "linkedTo",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "description",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "name",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "type",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "primitive",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "target",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "rootClass",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "validate",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "excerptSize",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "default",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "translate",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "unsafe",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "allowCreate",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "module",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "json",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "of",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "yaml",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "idFieldGorm",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "computedType",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "computedTypeClass",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "matches",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "gorm",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "gormMap",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "sql",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "fullName",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "fields",
			Kind:  lsp.CIKClass,
		},
	}
}
func ActionsInHandler() []lsp.CompletionItem {

	return []lsp.CompletionItem{

		{
			Label: "fields",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "dto",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "entity",
			Kind:  lsp.CIKClass,
		},
	}
}

func ActionsRootHandler() []lsp.CompletionItem {
	return []lsp.CompletionItem{
		{
			Label:      "New Action",
			Kind:       lsp.CIKSnippet,
			InsertText: "- name: actionName\r\n  path: /action/path",
		},
	}
}

func EntitiesRootHandler() []lsp.CompletionItem {
	return []lsp.CompletionItem{
		{
			Label:      "New Entity",
			Kind:       lsp.CIKSnippet,
			InsertText: "- name: entityName\r\n  fields: \r\n    - name: field1 \r\n      type: string",
		},
	}
}

func MethodHandler() []lsp.CompletionItem {
	return []lsp.CompletionItem{
		{
			Label: "get",
			Kind:  lsp.CIKConstant,
		},
		{
			Label: "post",
			Kind:  lsp.CIKConstant,
		},
		{
			Label: "delete",
			Kind:  lsp.CIKConstant,
		},
		{
			Label: "patch",
			Kind:  lsp.CIKConstant,
		},
		{
			Label: "reactive",
			Kind:  lsp.CIKConstant,
		},
	}
}

func ActionsHandler() []lsp.CompletionItem {

	// fmt.Println(100)
	return []lsp.CompletionItem{

		{
			Label: "actionName",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "cliName",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "actionAliases",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "name",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "url",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "method",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "fn",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "description",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "group",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "format",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "in",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "out",
			Kind:  lsp.CIKClass,
		},
		{
			Label: "security",
			Kind:  lsp.CIKClass,
		},
	}
}

func MainHandler() []lsp.CompletionItem {
	// fmt.Println(100)
	return []lsp.CompletionItem{
		{
			Label:      "actions list",
			InsertText: "actions:",
			Kind:       lsp.CIKProperty,
			Detail:     "Actions at module level",
		},
		{
			Label:      "entities",
			InsertText: "entities:\r\n  ",
			Kind:       lsp.CIKProperty,
			Detail:     "Modules entities",
		},
		{
			Label:      "dto",
			InsertText: "dto:",
			Kind:       lsp.CIKProperty,
			Detail:     "Module dtos",
		},
		{
			Label:      "path",
			InsertText: "path:",
			Kind:       lsp.CIKProperty,
			Detail:     "Go/Java module path",
		},
		{
			Label:      "name",
			InsertText: "name:",
			Kind:       lsp.CIKProperty,
			Detail:     "Go/Java module name",
		},
	}

}

func EntityItemsHandler() []lsp.CompletionItem {
	// fmt.Println(100)
	return []lsp.CompletionItem{

		{
			Label: "permissions",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "name",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "distinctBy",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "prependScript",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "prependCreateScript",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "prependUpdateScript",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "noQuery",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "access",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "queryScope",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "security",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "http",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "patch",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "queries",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "get",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "gormMap",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "query",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "post",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "importList",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "fields",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "c",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "cliName",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "cliShort",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "cliDescription",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "cte",
			Kind:  lsp.CIKProperty,
		},
		{
			Label: "postFormatter",
			Kind:  lsp.CIKProperty,
		},
	}
}
