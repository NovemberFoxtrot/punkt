package main

import (
	"fmt"
	"log"

	"github.com/iwanbk/gobeanstalk"
	"github.com/sirsean/go-mailgun/mailgun"
)

func main() {
	conn, err := gobeanstalk.Dial("localhost:11300")

	if err != nil {
		log.Printf("connect failed")
		log.Fatal(err)
	}

	mg_client := mailgun.NewClient("", "")

	for {
		j, err := conn.Reserve()

		if err != nil {
			log.Println("reserve failed")
			log.Fatal(err)
		}

		err = conn.Delete(j.Id)

		if err != nil {
			log.Fatal(err)
		}

		message1 := mailgun.Message{
			FromName:    "Spacely Sprockets",
			FromAddress: "spacely@spacelysprockets.com",
			ToAddress:   "roco@selfdestruct.jp",
			Subject:     "Go Mailgun sample message",
			Body:        string(j.Body),
		}

		fmt.Println("Attempting to send to ", mg_client.Endpoint(message1))

		body, err := mg_client.Send(message1)

		if err != nil {
			fmt.Println("Got an error:", err)
		} else {
			fmt.Println(body)
		}
	}
}
