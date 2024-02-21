//go:build !wasm

package workspaces

import (
	"fmt"
	"strings"

	"github.com/denisbrodbeck/machineid"
)

func IssueMachineId(name string) {
	id, err := machineid.ProtectedID(name)

	if err != nil {
		fmt.Println("We could not determine your system unique id. Licensing might fail, and you will access less features")
	}

	UNIQUE_MACHINE_ID = DashedString(strings.ToUpper(id[:16]))

}
