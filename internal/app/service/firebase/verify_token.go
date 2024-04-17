package firebase

import (
	"context"

	"firebase.google.com/go/auth"
)

func (s Service) VerifyToken(ctx context.Context, tokenId string) (*auth.Token, error) {
	return s.FirebaseClient.VerifyToken(ctx, tokenId)
}
