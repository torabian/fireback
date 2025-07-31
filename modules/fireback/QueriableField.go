package fireback

import (
	"encoding/json"
	"fmt"
	reflect "reflect"

	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/ast"
)

func ReplaceIdentifier(node ast.Node, from, to string) ast.Node {
	switch n := node.(type) {
	case *ast.IdentifierNode:
		if n.Value == from {
			n.Value = to
		}
	case *ast.BinaryNode:
		n.Left = ReplaceIdentifier(n.Left, from, to)
		n.Right = ReplaceIdentifier(n.Right, from, to)
	case *ast.UnaryNode:
		n.Node = ReplaceIdentifier(n.Node, from, to)
	case *ast.CallNode:
		for i := range n.Arguments {
			n.Arguments[i] = ReplaceIdentifier(n.Arguments[i], from, to)
		}
	}
	return node
}

func ExprToSQL(node ast.Node) string {

	switch n := node.(type) {
	case *ast.BinaryNode:
		left := ExprToSQL(n.Left)
		right := ExprToSQL(n.Right)

		op := n.Operator
		switch op {
		case "==":
			op = "=="
		case "!=":
			op = "<>"
		case "||":
			op = "OR"
		case "&&":
			op = "AND"
		}

		return fmt.Sprintf("(%s %s %s)", left, op, right)

	case *ast.CallNode:
		if fn, ok := n.Callee.(*ast.IdentifierNode); ok && fn.Value == "substring" && len(n.Arguments) == 3 {
			str := ExprToSQL(n.Arguments[0])
			start := ExprToSQL(n.Arguments[1])
			length := ExprToSQL(n.Arguments[2])
			return fmt.Sprintf("SUBSTRING(%s FROM %s + 1 FOR %s)", str, start, length)
		}
		return fmt.Sprintf("[unsupported function: %T]", n)

	case *ast.IdentifierNode:
		return ("`" + n.Value + "`")
		// return ("`" + n.Value + "`") This is needed when we are getting pure sql

	case *ast.StringNode:
		return fmt.Sprintf(`'%s'`, n.Value)

	case *ast.IntegerNode:
		return fmt.Sprintf(`%d`, n.Value)

	case *ast.FloatNode:
		return fmt.Sprintf(`%f`, n.Value)

	case *ast.UnaryNode:
		right := ExprToSQL(n.Node)
		return fmt.Sprintf("(%s %s)", n.Operator, right)

	default:
		return fmt.Sprintf("[unsupported: %T]", node)
	}
}

func ExprToJQ(node ast.Node) string {
	switch n := node.(type) {
	case *ast.BinaryNode:
		left := ExprToJQ(n.Left)
		right := ExprToJQ(n.Right)

		op := n.Operator
		switch op {
		case "==":
			op = "=="
		case "!=":
			op = "!="
		case "||":
			op = "or"
		case "&&":
			op = "and"
		case "<>":
			op = "!="
		}

		return fmt.Sprintf("(%s %s %s)", left, op, right)

	case *ast.CallNode:
		if fn, ok := n.Callee.(*ast.IdentifierNode); ok && fn.Value == "substring" && len(n.Arguments) == 3 {
			str := ExprToJQ(n.Arguments[0])
			start := ExprToJQ(n.Arguments[1])
			length := ExprToJQ(n.Arguments[2])
			return fmt.Sprintf("(%s | .[%s:(%s+%s)])", str, start, start, length)
		}
		return fmt.Sprintf("[unsupported function: %T]", n)

	case *ast.IdentifierNode:
		return fmt.Sprintf(".%s", n.Value)

	case *ast.StringNode:
		return fmt.Sprintf(`"%s"`, n.Value)

	case *ast.IntegerNode:
		return fmt.Sprintf(`%d`, n.Value)

	case *ast.FloatNode:
		return fmt.Sprintf(`%f`, n.Value)

	case *ast.UnaryNode:
		right := ExprToJQ(n.Node)
		return fmt.Sprintf("(%s %s)", n.Operator, right)

	default:
		return fmt.Sprintf("[unsupported: %T]", node)
	}
}

func InitQueryStruct(s any) any {
	val := reflect.ValueOf(s).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		if field.Kind() == reflect.Struct && field.Type().Name() == "QueriableField" {

			q := QueriableField{
				FieldName: fieldType.Tag.Get("column"),
				FieldQs:   fieldType.Tag.Get("qs"),
				// Add rest of the meta tags here will be needed
			}
			field.Set(reflect.ValueOf(q))
		}
	}

	return s
}

type QueriableField struct {
	// Query     string `json:"query"`
	// Operation string `json:"operation"`
	UserInput string `json:"userInput"`

	FieldName string `json:"fieldName"`

	FieldQs string `json:"fieldQs"`
}

func (m QueriableField) AsSql() (string, error) {

	if m.UserInput == "" {
		return "", nil
	}

	env := map[string]any{
		"this": "",
	}

	options := []expr.Option{
		expr.Env(env),
		expr.Function("substring", func(args ...any) (any, error) {
			// Dummy implementation for compile-time only
			return nil, nil
		}),
	}

	program, err := expr.Compile(m.UserInput, options...)
	if err != nil {
		return "", err
	}

	ReplaceIdentifier(program.Node(), "this", m.FieldName)
	sql := ExprToSQL(program.Node())

	return sql, nil
}

func (m QueriableField) AsJq() (string, error) {

	if m.UserInput == "" {
		return "", nil
	}

	env := map[string]any{
		"this": "",
	}

	options := []expr.Option{
		expr.Env(env),
		expr.Function("substring", func(args ...any) (any, error) {
			// Dummy implementation for compile-time only
			return nil, nil
		}),
	}

	program, err := expr.Compile(m.UserInput, options...)
	if err != nil {
		return "", err
	}

	ReplaceIdentifier(program.Node(), "this", m.FieldQs)
	sql := ExprToJQ(program.Node())

	return sql, nil
}

func (m QueriableField) MarshalJSON() ([]byte, error) {
	if res, err := m.AsSql(); err != nil {
		return json.Marshal("Cannot cast the query to sql:")
	} else {
		return json.Marshal(res)
	}
}

func (q QueriableField) String() string {
	return ""
}

// func (q QueriableField) String() string {
// 	if q.Query == "" {
// 		return ""
// 	}

// 	// Now we need to check, if there are no special signs, such as
// 	// =, < > ! , {", then user intended to be equal

// 	autoPrefix := "="
// 	if len(q.Query) > 0 {
// 		if q.Query[0] == '<' || q.Query[0] == '>' || q.Query[0] == '%' || q.Query[0] == '!' || q.Query[0] == '=' || q.Query[0] == '{' {
// 			// Then it's assumed user completed the expq.Queryssion,
// 			autoPrefix = ""
// 		}
// 	}

// 	// I need the reflect tag column from itself
// 	return "$col " + autoPrefix + q.Query
// }
