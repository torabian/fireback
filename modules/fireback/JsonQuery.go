package fireback

type VM struct {
	Value     string `json:"value"`
	Operation string `json:"operation"`
}

// func Escape(sql string) string {
// 	dest := make([]byte, 0, 2*len(sql))
// 	var escape byte
// 	for i := 0; i < len(sql); i++ {
// 		c := sql[i]

// 		escape = 0

// 		switch c {
// 		case 0: /* Must be escaped for 'mysql' */
// 			escape = '0'
// 			break
// 		case '\n': /* Must be escaped for logs */
// 			escape = 'n'
// 			break
// 		case '\r':
// 			escape = 'r'
// 			break
// 		case '\\':
// 			escape = '\\'
// 			break
// 		case '\'':
// 			escape = '\''
// 			break
// 		case '"': /* Better safe than sorry */
// 			escape = '"'
// 			break
// 		case '\032': //十进制26,八进制32,十六进制1a, /* This gives problems on Win32 */
// 			escape = 'Z'
// 		}

// 		if escape != 0 {
// 			dest = append(dest, '\\', escape)
// 		} else {
// 			dest = append(dest, c)
// 		}
// 	}

// 	return string(dest)
// }
