//go:build !wasm

package fireback

import (
	"fmt"
	"strings"

	"github.com/denisbrodbeck/machineid"
)

var UNIQUE_MACHINE_ID string = ""
var UNIQUE_APP_PUBLIC_KEY string = ""

func DashedString(input string) string {
	out := ""
	for index, char := range input {
		out += string(char)
		if index != 0 && (index+1)%4 == 0 && len(input) != index+1 {
			out += "-"
		}
	}

	return out
}
func IssueMachineId(name string) {
	id, err := machineid.ProtectedID(name)

	if err != nil {
		fmt.Println("We could not determine your system unique id. Licensing might fail, and you will access less features")
	}

	UNIQUE_MACHINE_ID = DashedString(strings.ToUpper(id[:16]))

}
