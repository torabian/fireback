package fireback

import (
	"encoding/json"
	"fmt"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/urfave/cli"
)

func VapidCmd() cli.Command {

	return cli.Command{
		Name:        "vapid",
		Description: "VAPID web push notification public/private key generator",
		Usage:       `VAPID web push notification public/private key generator`,
		Action: func(c *cli.Context) error {
			privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
			if err != nil {
				return err
			}

			fmt.Println("private: ", privateKey)
			fmt.Println("public: ", publicKey)

			return nil
		},
	}
}

func AutoConfig() cli.Command {

	return cli.Command{
		Name:        "autoconfig",
		Description: "Creates private/public VAPID keys, and saves them in the environment",
		Usage:       `Creates private/public VAPID keys, and saves them in the environment`,
		Action: func(c *cli.Context) error {
			privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
			if err != nil {
				return err
			}

			fmt.Println("private: ", privateKey)
			fmt.Println("public: ", publicKey)

			config.VapidPrivateKey = privateKey
			config.VapidPublicKey = publicKey

			if err2 := config.Save(".env"); err2 != nil {
				return err2
			}
			fmt.Println("Saved successfully.")

			return nil
		},
	}
}

func SendWebPush() cli.Command {

	return cli.Command{
		Name:        "push",
		Description: "Sends a push notification",
		Usage:       `Sends a push notification`,
		Action: func(c *cli.Context) error {
			data, _, err := WebPushConfigEntityStream(QueryDSL{})
			if err != nil {
				return err
			}

			for wpcs := range data {
				for _, wpc := range wpcs {

					// Decode subscription
					s := &webpush.Subscription{}
					json.Unmarshal(wpc.Subscription.Bytes(), s)

					// Send Notification
					resp, err := webpush.SendNotification([]byte("akshdakjshdajshdakshdajksdhaskdhakjdhakjshd"), s, &webpush.Options{
						Subscriber:      "example@example.com",
						VAPIDPublicKey:  config.VapidPublicKey,
						VAPIDPrivateKey: config.VapidPrivateKey,
						TTL:             30,
						Topic:           "topics1",
					})

					if err != nil {
						// TODO: Handle error
						return err
					}
					defer resp.Body.Close()
				}
			}

			return nil
		},
	}
}

var PushNotificationCmd cli.Command = cli.Command{

	Name:  "pushnot",
	Usage: "Push notification (web-push) config and actions",
	Subcommands: []cli.Command{
		VapidCmd(),
		AutoConfig(),
		SendWebPush(),
		WebPushConfigCliFn(),
	},
}
