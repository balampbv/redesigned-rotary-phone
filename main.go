package main

import (
	chat "chat-server/controllers"
	"chat-server/store"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	store := store.NewStore()

	chatCtrl := chat.NewChatController(store)

	e.POST("/send", chatCtrl.SendMessages)
	e.GET("/messages/:username", chatCtrl.GetAllMessages)

	e.Logger.Fatal(e.Start(":1323"))
}
