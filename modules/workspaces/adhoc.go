package workspaces

import "github.com/urfave/cli"

var AdhocTools cli.Command = cli.Command{

	Name:  "adhoc",
	Usage: "Different tools for converting files, which are not really related to fireback but handy",
	Subcommands: []cli.Command{
		{
			Name: "webp",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "input",
					Required: true,
					Usage:    "Image source file, jpg, png, tiff, etc...",
				},
				&cli.StringFlag{
					Name:     "output",
					Required: true,
					Usage:    "Output directiry, which will be a list of different sizes of the image",
				},
				&cli.Int64Flag{
					Name:  "width",
					Usage: "Image width. Default is 0",
				},
				&cli.Int64Flag{
					Required: true,
					Usage:    "Image height. Default is 0",
				},
			},
			Action: func(c *cli.Context) error {

				ConvertToWebp(c.String("input"), c.String("output"), ImageCropSize{
					Width:  int(c.Int64("width")),
					Height: int(c.Int64("height")),
				})
				return nil
			},
		},
	},
}
