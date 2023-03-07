package auth

import (
	"context"
	"net/http"

	"github.com/1-platform/api-catalog/internal/pkg/config"
	"github.com/akhilmhdh/authy"
	oauth2Core "github.com/akhilmhdh/authy/modules/oauth2/core"
	"github.com/akhilmhdh/authy/modules/oauth2/rest"
)

type Auth struct {
	authy *authy.Authy
	store store
}

type store interface {
	GetUserByPID(ctx context.Context, pid string) (authy.User, error)
	NewOauth2User(ctx context.Context, provider string, userInfo map[string]any) (oauth2Core.Oauth2User, error)
	SaveOauth2User(ctx context.Context, user oauth2Core.Oauth2User) error
}

func New(store store, cfg *config.Config) *Auth {
	session := authy.NewCookieSessionMaker([]byte(cfg.Auth.CookieHashKey), []byte(cfg.Auth.CookieBlockKey))

	a := authy.New(store, session)
	a.Config.ServerURL = cfg.ServerURL
	a.Config.ApplicationURL = cfg.ApplicationURL

	session.AutoConfig(a.Config.ServerURL, a.Config.ApplicationURL)

	oauth2Cfg := oauth2.NewConfig()
	om := oauth2.New(a, store, oauth2Cfg)

	for _, generic_oauth := range cfg.Auth.GenericOauth {
		om.RegisterProvider(generic_oauth.Name, &oauth2Core.OAuth2Strategy{
			ClientID:     generic_oauth.ClientID,
			ClientSecret: generic_oauth.ClientSecret,
			Scopes:       generic_oauth.Scopes,
			Endpoint: oauth2Core.Endpoint{
				AuthURL:  generic_oauth.AuthURL,
				TokenURL: generic_oauth.TokenURL,
			},
			GetUserInfo: func(ctx context.Context, token string) (map[string]any, error) {
				return GenericOauthUserInfo(token, generic_oauth.UserInfoURL,
					generic_oauth.EmailPath, generic_oauth.UidPath, generic_oauth.DisplayNamePath)
			},
		})
	}

	return &Auth{store: store, authy: a}
}

func (a *Auth) InjectMiddleware() func(handler http.Handler) http.Handler {
	return a.authy.Middlewares
}
