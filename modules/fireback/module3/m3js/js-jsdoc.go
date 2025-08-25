package m3js

type JsdocComment struct {
	content []string
	intend  string
}

func (x *JsdocComment) Add(line string) *JsdocComment {
	x.content = append(x.content, "* "+line)

	return x
}

func (x *JsdocComment) String() string {
	data := []byte{'/', '*', '*', '\r', '\n'}

	for _, line := range x.content {
		data = append(data, x.intend+line...)
		data = append(data, []byte("\r\n")...)
	}

	data = append(data, []byte(x.intend+"**/\r\n")...)

	return string(data)
}

func NewJsDoc(intend string) *JsdocComment {

	return &JsdocComment{
		intend: intend,
	}
}
