package firebaseclient

import (
	"context"
	"kindercastle_backend/internal/pkg/logging"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type firebaseClient struct {
	firebaseApp *firebase.App
	firebaseErr error
}

func (fc *firebaseClient) VerifyToken(ctx context.Context, tokenId string) (*auth.Token, error) {
	if fc.firebaseErr != nil {
		logging.Logger().Error().Err(fc.firebaseErr).Msg("Firebase setup error")
		return nil, fc.firebaseErr
	}

	client, err := fc.firebaseApp.Auth(ctx)
	if err != nil {
		logging.Logger().Error().Err(err).Msg("Failed to initialize Firebase Auth client")
		return nil, err
	}

	token, err := client.VerifyIDToken(ctx, tokenId)
	if err != nil {
		logging.Logger().Error().Err(err).Msg("Invalid or expired ID token")
		return nil, err
	}

	return token, nil
}

func NewFirebaseClient(credPath string) FirebaseClient {
	opt := option.WithCredentialsFile(credPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logging.Logger().Error().Err(err).Msg("Error initializing Firebase app")
	}
	return &firebaseClient{
		firebaseApp: app,
		firebaseErr: err,
	}
}
