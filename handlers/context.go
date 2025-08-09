package handlers

import (
	"context"
	"goauthic/utils"
)

type contextKey string

const userClaimsKey contextKey = "userClaims"

func WithUserClaims(ctx context.Context, claims *utils.Claims) context.Context {
	return context.WithValue(ctx, userClaimsKey, claims)
}

func GetUserClaims(ctx context.Context) (*utils.Claims, bool) {
	claims, ok := ctx.Value(userClaimsKey).(*utils.Claims)
	return claims, ok
}
