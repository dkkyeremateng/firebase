package msg

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

// New returns a new instance of messaging client.
func New(ctx context.Context, app *firebase.App) (*messaging.Client, error) {

	fcmCli, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("messaging client: %w", err)
	}

	return fcmCli, nil
}
