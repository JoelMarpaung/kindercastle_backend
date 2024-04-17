package firebaseclient

import (
	"context"

	"firebase.google.com/go/auth"
)

type FirebaseClient interface {
	VerifyToken(ctx context.Context, tokenId string) (*auth.Token, error)
}
