package helper

import (
	"log"
	"sirloinapi/config"

	"firebase.google.com/go/v4/messaging"
)

func PushNotification(title, msg, token string) error {
	client, ctx, err := config.InitFCMClient()
	if err != nil {
		log.Println("error initializing FCM client: " + err.Error())
	}
	// Define the message to be sent
	message := messaging.Message{
		Data: map[string]string{
			"title": title,
			"body":  msg,
		},
		Notification: &messaging.Notification{
			Title: title,
			Body:  msg,
		},
		Token: token,
	}

	// Send the message to the device
	response, err := client.Send(ctx, &message)
	// log.Println(response)
	if err != nil {
		log.Printf("error sending message: %v\n", err)
		log.Println("response: ", response)
		return err
	}

	return nil
}
