package workspaces

import (
	"github.com/urfave/cli"
)

var GsmProviderTestCmd cli.Command = cli.Command{

	Name:  "sms",
	Usage: "Sends the text message via gsm provider id",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "id",
			Value:    "",
			Usage:    "Provider which you want to use for the message",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "message",
			Value:    "",
			Usage:    "Message content",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "to",
			Value:    "",
			Usage:    "Message recipient",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {
		message := c.String("message")
		result, err := GsmSendSMS(c.String("id"), message, []string{c.String("to")})
		HandleActionInCli(c, result, err, map[string]map[string]string{})

		return nil
	},
}

func init() {
	GsmProviderCliCommands = append(GsmProviderCliCommands, GsmProviderTestCmd)
}
