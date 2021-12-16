package config

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
)

const (
	AppName = "jsch"
	EnvProd = "production"
)

type (
	Config struct {
		Env   string `required:"true"`
		Mongo MongoConfig
		HTTP  HTTPConfig
		Auth  AuthConfig
	}

	MongoConfig struct {
		URI      string `required:"true"`
		User     string `required:"true"`
		Password string `required:"true"`
		Database string `required:"true" default:"jobSchedule"`
	}

	AuthConfig struct {
		JWT                    JWTConfig
		PasswordSalt           string `envconfig:"auth_password_salt" required:"true"`
		VerificationCodeLength int    `envconfig:"auth_verification_code_length" required:"true"`
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration `envconfig:"jwt_access_token_ttl" default:"2h"`
		RefreshTokenTTL time.Duration `envconfig:"jwt_refresh_token_ttl" default:"720h"`
		SigningKey      string        `envconfig:"jwt_signing_key" required:"true"`
	}

	HTTPConfig struct {
		Host               string        `envconfig:"http_host"`
		Port               string        `ignored:"true" default:"8085"`
		ReadTimeout        time.Duration `envconfig:"http_read_timeout" default:"10s"`
		WriteTimeout       time.Duration `envconfig:"http_write_timeout" default:"10s"`
		MaxHeaderMegabytes int           `envconfig:"http_max_header_megabytes" default:"1"`
	}
)

func Init() (*Config, error) {
	var cfg Config
	if err := envconfig.Process(AppName, &cfg); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) validate() error {
	if cfg.Env != "local" && cfg.Env != "production" {
		errMsg := fmt.Sprintf("key %s_%s accepts only the values of local, %s", strings.ToUpper(AppName), strings.ToUpper("env"), EnvProd)

		return errors.New(errMsg)
	}

	return nil
}
