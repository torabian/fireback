package mcore

import "fmt"

// represents a js/ts argument signature, "name?: string | null"
type JsFnArgument struct {
	Ts  string
	Js  string
	Key string
}

func NewJsArgument(dto JsFnArgument) JsFnArgument {
	return dto
}

func (x JsFnArgument) CompileJs() string {
	return fmt.Sprintf("%v", x.Js)
}

func (x JsFnArgument) CompileTs() string {
	return fmt.Sprintf("%v", x.Ts)
}
