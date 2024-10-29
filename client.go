package main

import (
	"bufio"
	"fmt"
	"local/models"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	ServerUrl := "ws://localhost:8080/ws"
	var UserName string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your name!")
	if scanner.Scan() {
		UserName = scanner.Text()
	}

	conn, _, err := websocket.DefaultDialer.Dial(ServerUrl, nil)
	if err != nil {
		log.Fatal("Websocket dial error:", err)
	}
	defer conn.Close()

	// Listen for messages on a goroutine to prevent blocking
	go func() {
		for {
			message := &models.Message{}
			err := conn.ReadJSON(message)
			if err != nil {
				fmt.Println("Read message from server error:", err)
				return
			}

			fmt.Printf("%s: %s \n", message.Sender, message.Content)
		}
	}()

	for scanner.Scan() {
		text := scanner.Text()
		if text == "exit" {
			break
		}

		message := models.NewMessage(UserName, text)
		err := conn.WriteJSON(message)
		if err != nil {
			log.Fatal("Sending message error:", err)
		}
	}
}
