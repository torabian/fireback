package fireback

import (
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v3"
)

func isYaml(headerValue string) bool {

	return slices.Contains([]string{
		"application/x-yaml",
		"application/yaml",
		"text/yaml",
		"yaml",
		"yml",
	}, headerValue)

}
func isCsv(headerValue string) bool {

	return slices.Contains([]string{
		"text/csv",
		"csv",
	}, headerValue)

}

func IsYaml(c *gin.Context) bool {
	return isYaml(c.GetHeader("Accept"))
}

func IsYamlCli(c *cli.Command) bool {
	return isYaml(c.String("x-accept"))
}

func IsCsvCli(c *cli.Command) bool {
	return isCsv(c.String("x-accept"))
}
