package main

import (
	"fmt"
	"os"

	"github.com/torabian/fireback/modules/workspaces"
)

func scan() {
	filename := "fireback-configuration.yml"
	if file, err := os.ReadFile(filename); err != nil {
		fmt.Println(err)

		fmt.Println(workspaces.GetContextByLineAndCol(string(file), 10, 2))
	}

}
