package workspaces

import (
	"fmt"
	"log"
	"testing"

	"github.com/fatih/color"
)

type WorkspaceCreationCtx struct {
	F QueryDSL
}

type Test struct {
	Name     string
	Function func(t *TestContext) error
}

type TestContext struct {
	testing.T
	F QueryDSL
	// Include any necessary fields or methods for test context
}

func (tc *TestContext) ErrorLn(message string, args ...interface{}) {
	// Implement error reporting mechanism
	c := color.New(color.FgRed)
	c.Print("Test Has failed unfortunately:")
	log.Fatalln(message, args)

}

func (tc *TestContext) Log(message string, args ...interface{}) {
	// Implement error reporting mechanism
	c := color.New(color.Faint)
	c.Println(message, args)

}

func (tc *TestContext) Errorf(message string, args ...interface{}) {
	// Implement error reporting mechanism
	c := color.New(color.FgRed)
	c.Print("Test Has failed unfortunately:")
	log.Fatalln(message, args)

}
func (tc *TestContext) Error(message string, args ...interface{}) {
	// Implement error reporting mechanism
	c := color.New(color.FgRed)
	c.Println(message, args)

}

func RunTests(F QueryDSL) {
	// testContext := &TestContext{F: F} // Initialize test context

	// for _, test := range tests {
	// 	err := test.Function(testContext)
	// 	if err == nil {
	// 		c := color.New(color.FgGreen)
	// 		fmt.Print("\u2713 Test \"")
	// 		c.Print(test.Name)
	// 		fmt.Print("\" Has passed successfully")
	// 	}
	// 	fmt.Println("")

	// }
}

func TestRunner(ctx *TestContext, tests []Test) {
	for _, test := range tests {
		err := test.Function(ctx)
		if err == nil {
			c := color.New(color.FgGreen)
			fmt.Print("\u2713 Test \"")
			c.Print(test.Name)
			fmt.Print("\" Has passed successfully")
		}
		fmt.Println("")

	}
}
