package app

import (
	"context"
	"errors"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// New return a new firebase app instance from the provided config and client options
func New(ctx context.Context, config *firebase.Config, opts ...option.ClientOption) (*firebase.App, error) {

	app, err := firebase.NewApp(ctx, config, opts...)
	if err != nil {
		return nil, errors.New("error initializing firebase app")
	}

	return app, nil
}
