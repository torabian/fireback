package fireback

import (
	"encoding/json"
	"strings"
)

type VM struct {
	Value     string `json:"value"`
	Operation string `json:"operation"`
}

func Escape(sql string) string {
	dest := make([]byte, 0, 2*len(sql))
	var escape byte
	for i := 0; i < len(sql); i++ {
		c := sql[i]

		escape = 0

		switch c {
		case 0: /* Must be escaped for 'mysql' */
			escape = '0'
			break
		case '\n': /* Must be escaped for logs */
			escape = 'n'
			break
		case '\r':
			escape = 'r'
			break
		case '\\':
			escape = '\\'
			break
		case '\'':
			escape = '\''
			break
		case '"': /* Better safe than sorry */
			escape = '"'
			break
		case '\032': //十进制26,八进制32,十六进制1a, /* This gives problems on Win32 */
			escape = 'Z'
		}

		if escape != 0 {
			dest = append(dest, '\\', escape)
		} else {
			dest = append(dest, c)
		}
	}

	return string(dest)
}

func JsonQueryToSql(ma string) string {
	if ma == "{}" {
		return ""
	}
	m := map[string]VM{}

	err := json.Unmarshal([]byte(ma), &m)

	if err != nil {
		return ""
	}

	where := []string{}
	for field := range m {
		snake := ToSnakeCase(field)
		if m[field].Operation == "contains" {
			where = append(where, "`"+Escape(snake)+"` like '%"+Escape(m[field].Value)+"%'")
		}
	}

	if len(where) > 0 {
		sql := strings.Join(where, " and ")
		return "and " + sql
	}

	return ""
}
