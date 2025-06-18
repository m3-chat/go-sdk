package main

import (
	"log"

	"github.com/m3-chat/go-sdk/client"
	"github.com/m3-chat/go-sdk/types"
)

func main() {
	c := client.NewClient(&types.ClientOptions{
		Stream: true,
	})

	err := c.GetResponse(types.RequestParams{
		Model:   "mistral",
		Content: "Hello, how are you?",
	})

	if err != nil {
		log.Fatal(err)
	}
}
