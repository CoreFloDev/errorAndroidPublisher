package main

import (
	"context"

	_ "github.com/99designs/gqlgen/handler"
	_ "github.com/go-chi/chi/v5"
	_ "github.com/go-webauthn/webauthn/webauthn"
	_ "github.com/google/wire"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/nicksnyder/go-i18n/v2/i18n"
	_ "github.com/vektah/gqlparser/v2"
	_ "golang.org/x/net/http2"
	"google.golang.org/api/androidpublisher/v3"
)

func main() {
	ctx := context.Background()
	androidpublisher.NewService(ctx)
}