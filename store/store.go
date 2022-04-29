package store

import (
	chatErrors "chat-server/errors"
	"fmt"
	"time"
)

//alice -> [bob : [" ", " "]] [clare: [" " , " "]]
//bob ->
//clare ->

type Message struct {
	Msg       string    `json:"Message"`
	TimeStamp time.Time `json:"ReceivedTime"`
}
type userMessage struct {
	UserMsg map[string][]Message `json:"UserMessages"`
}

type store struct {
	messages map[string]userMessage
}

type Store interface {
	AddMessage(fromUser, toUser, msg string) error
	GetMessageForUser(userName string) (userMessage, error)
}

func NewStore() Store {
	return &store{
		messages: make(map[string]userMessage),
	}
}

func (s *store) AddMessage(fromUser, toUser, msg string) error {

	fmt.Println(s.messages)
	_, ok := s.messages[toUser]
	if ok {
		currentMessages := s.messages[toUser].UserMsg[fromUser]
		s.messages[toUser].UserMsg[fromUser] = append(currentMessages, Message{
			Msg:       msg,
			TimeStamp: time.Now(),
		})
		fmt.Println("After Storing", s.messages)
	} else {
		newMessage := new(userMessage)
		newMessage.UserMsg = make(map[string][]Message)
		newMessage.UserMsg[fromUser] = []Message{{
			Msg:       msg,
			TimeStamp: time.Now(),
		}}
		s.messages[toUser] = *newMessage
		fmt.Print(newMessage)
	}

	return nil
}

func (s *store) GetMessageForUser(userName string) (userMessage, error) {
	msgs, ok := s.messages[userName]
	if !ok {
		fmt.Print(msgs)
		return userMessage{}, chatErrors.ErrUserNotFound
	}
	fmt.Println(msgs)
	s.messages[userName] = userMessage{}
	return msgs, nil
}
