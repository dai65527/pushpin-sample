package main

import (
	"bytes"
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/stream/:channel", handleGetChatStream)
	e.POST("/message/:channel", handlePostChatStream)
	e.Server.Addr = ":8080"
	e.Logger.Fatal(e.Start(":8080"))
}

func handleGetChatStream(c echo.Context) error {
	channel := c.Param("channel")
	c.Response().Header().Set("Grip-Hold", "stream")
	c.Response().Header().Set("Grip-Channel", channel)
	return c.String(200, "stream opened")
}

type Message struct {
	SenderName string `json:"sender_name"`
	Message    string `json:"message"`
}

type PublishHttpStream struct {
	Content string `json:"content"`
}

type PublishFormats struct {
	HttpStream PublishHttpStream `json:"http-stream"`
}

type PublishItem struct {
	Channel string         `json:"channel"`
	Formats PublishFormats `json:"formats"`
}

type PublishRequest struct {
	Items []PublishItem `json:"items"`
}

func handlePostChatStream(c echo.Context) error {
	channel := c.Param("channel")
	message := new(Message)
	if err := c.Bind(message); err != nil {
		return err
	}
	messageJson, err := json.Marshal(message)
	if err != nil {
		return err
	}
	publishRequest := PublishRequest{
		Items: []PublishItem{
			{
				Channel: channel,
				Formats: PublishFormats{
					HttpStream: PublishHttpStream{
						Content: message.Message,
					},
				},
			},
		},
	}
	publishRequestJson, err := json.Marshal(publishRequest)
	if err != nil {
		return err
	}
	fmt.Println(string(publishRequestJson))
	resp, err := http.Post("http://pushpin:5561/publish/", "application/x-www-form-urlencoded", bytes.NewReader(publishRequestJson))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	if resp.StatusCode != 200 {
		return c.String(500, "failed to send message")
	}
	return c.JSON(200, messageJson)
}
