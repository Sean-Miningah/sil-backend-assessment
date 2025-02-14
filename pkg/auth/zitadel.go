package auth

import (
	"context"

	"github.com/sean-miningah/sil-backend-assessment/pkg/config"
	"github.com/zitadel/zitadel-go/v3/pkg/authentication"
	openid "github.com/zitadel/zitadel-go/v3/pkg/authentication/oidc"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

func NewZitadelMiddleware(cfg *config.Config) (*authentication.Interceptor, error) {
	ctx := context.Background()

	// Initialize Zitadel client
	conf := zitadel.New(cfg.ZitadelIssuerURL)

	auth, err := authentication.New(ctx, conf, cfg.ZitadelClientSecret,
		openid.DefaultAuthentication(cfg.ZitadelClientID, cfg.ZitadelRedirectURI, cfg.ZitadelClientSecret),
	)
	if err != nil {
		return nil, err
	}

	mw := authentication.Middleware(auth)

	return mw, nil

}
