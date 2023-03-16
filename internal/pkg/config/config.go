package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	MongoDbURI     string         `mapstructure:"mongodb_uri"`
	MongoDbName    string         `mapstructure:"mongodb_database"`
	Port           uint16         `mapstructure:"port"`
	Host           string         `mapstructure:"host"`
	ServerURL      string         `mapstructure:"server_url"`
	ApplicationURL string         `mapstructure:"application_url"`
	Auth           Auth           `mapstructure:"auth"`
	Jwt            Jwt            `mapstructure:"jwt"`
	Smtp           Smtp           `mapstructure:"smtp"`
	TeamMembership TeamMembership `mapstructure:"team_membership"`
}

type Auth struct {
	CookieHashKey  string         `mapstructure:"hash_key"`
	CookieBlockKey string         `mapstructure:"block_key"`
	GenericOauth   []GenericOauth `mapstructure:"generic_oauth"`
}

type GenericOauth struct {
	Name         string   `mapstructure:"name"`
	ClientID     string   `mapstructure:"client_id"`
	ClientSecret string   `mapstructure:"client_secret"`
	Scopes       []string `mapstructure:"scopes"`
	AuthURL      string   `mapstructure:"auth_url"`
	TokenURL     string   `mapstructure:"token_url"`
	UserInfoURL  string   `mapstructure:"user_info_url"`
	// This is for UI
	IconColor string `mapstructure:"icon_color"`
	// userinfo data
	EmailPath       string `mapstructure:"email_path" `
	UidPath         string `mapstructure:"uid_path"`
	DisplayNamePath string `mapstructure:"name_path"`
}

type Jwt struct {
	IssuerName string `mapstructure:"issuer_name"`
	SigningKey string `mapstructure:"signing_key"`
}

type Smtp struct {
	From     string `mapstructure:"from"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"passoword"`
	UserName string `mapstructure:"username"`
}

type TeamMembership struct {
	InviteeExpiry time.Duration `mapstructure:"invitee_expiry"`
}

func New(path string) (*Config, error) {
	var config Config
	// load config paths
	viper.AddConfigPath("~/")
	viper.AddConfigPath(path)
	viper.SetConfigType("toml")
	viper.SetConfigName("apic")

	// set defaults
	viper.SetDefault("mongodb_database", "api-catalog")
	viper.SetDefault("auth.hash_key", "4Wpob%Up26^%2rqx3Z88TW9LucjFMuh%")
	viper.SetDefault("auth.block_key", "%ynb%nk3GX7HMuP%9H*m2F#5h2%DV8x*")
	viper.SetDefault("port", 4200)
	viper.SetDefault("host", "0.0.0.0")
	viper.SetDefault("server_url", "http://localhost:4200")
	viper.SetDefault("application_url", "http://localhost:3000")
	viper.SetDefault("auth.generic_oauth.email_path", "email")
	viper.SetDefault("auth.generic_oauth.uid_path", "id")
	viper.SetDefault("auth.generic_oauth.name_path", "name")
	viper.SetDefault("jwt.issuer_name", "api-catalog")
	viper.SetDefault("team_membership.invitee_expiry", 24*time.Hour)
	viper.SetDefault("jwt.signing_key", "X%$agpa9^yK2c*mUtCy2ocmrFGYGpAu$5pJtuVH$!zL4vdS5QTT44949k!gdRXgieB7Csir4U8J@m%*G")
	viper.SetDefault("smtp.from", "no-reply@api-catalog.com")
	// load from env if found
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	// set default for slice generic auth
	for i, val := range config.Auth.GenericOauth {
		if val.UidPath == "" {
			val.UidPath = "id"
		}
		if val.EmailPath == "" {
			val.EmailPath = "email"
		}
		if val.DisplayNamePath == "" {
			val.DisplayNamePath = "name"
		}
		config.Auth.GenericOauth[i] = val
	}

	if err != nil {
		return nil, err
	}

	return &config, nil
}
