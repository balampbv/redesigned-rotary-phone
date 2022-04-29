package controllers

import (
	"fmt"
	"net/http"

	"chat-server/store"

	"github.com/labstack/echo/v4"
)

type commonResponse struct {
	Message string `json:"message"`
}

type chatController struct {
	chatStore store.Store
}

type ChatController interface {
	SendMessages(cxt echo.Context) error
	GetAllMessages(cxt echo.Context) error
}

func NewChatController(store store.Store) ChatController {
	return &chatController{
		chatStore: store,
	}
}

type Message struct {
	FromUser string `json:"FromUser"`
	ToUser   string `json:"ToUser"`
	Message  string `json:"Message"`
}

func (c *chatController) SendMessages(cxt echo.Context) error {
	//to decode the body cxt.Bind() - retrun bad request if not valid
	msg := new(Message)
	err := cxt.Bind(msg)
	if err != nil {
		return cxt.JSON(http.StatusBadRequest, commonResponse{
			Message: fmt.Sprintf("Invalid send message request received, error: %s", err.Error()),
		})
	}
	//fmt.Print("From USer: ", msg.FromUser)
	//store the msgs
	err = c.chatStore.AddMessage(msg.FromUser, msg.ToUser, msg.Message)
	if err != nil {
		return cxt.JSON(http.StatusInternalServerError, commonResponse{
			Message: fmt.Sprintf("Internal server error received: %s", err.Error()),
		})
	}

	//send http ok
	return cxt.JSON(http.StatusOK, commonResponse{
		Message: "success",
	})
}

func (c *chatController) GetAllMessages(cxt echo.Context) error {
	//query the user details
	userName := cxt.Param("username")
	//check the unread msgs and return
	message, err := c.chatStore.GetMessageForUser(userName)
	if err != nil {
		return cxt.JSON(http.StatusInternalServerError, commonResponse{
			Message: fmt.Sprintf("Internal server error received: %s %s", userName, err.Error()),
		})
	}
	//send success
	return cxt.JSON(http.StatusOK, message)
}
