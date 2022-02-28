package sender

import (
	"context"
	"sync"

	"firebase.google.com/go/messaging"
)

type Sender struct {
	msgClient   *messaging.Client
	mu          sync.Mutex
	pushCounter int
}

// New returns an instance of firebase messaging sender.
func New(client *messaging.Client) (*Sender, error) {

	s := Sender{
		msgClient: client,
	}

	return &s, nil
}

// Sender send a firebase messaging.Message and return the response or error
func (s *Sender) SendPush(ctx context.Context, msg *messaging.Message) (string, error) {

	s.mu.Lock()

	res, err := s.msgClient.Send(ctx, msg)
	if err != nil {
		s.mu.Unlock()
		return "", err
	}

	s.pushCounter++
	s.mu.Unlock()

	return res, nil
}

// PushCount return the number of messages sent
func (s *Sender) PushCount() int {
	return s.pushCounter
}
