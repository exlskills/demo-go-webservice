package config

import (
	"fmt"
	"github.com/exlinc/golang-utils/envconfig"
	"github.com/sirupsen/logrus"
)

// The envconfig struct tag is used to explicitly name the var, set defaults, and flag required values
type Config struct {
	DBPath          string   `envconfig:"DB_PATH" required:"true"`
	Mode            string   `envconfig:"MODE" default:"production"`
	ListenAddress   string   `envconfig:"LISTEN_ADDRESS" default:"0.0.0.0"`
	ListenPort      string   `envconfig:"LISTEN_PORT" default:"3333"`
	AllowedOrigins  []string `envconfig:"ALLOWED_ORIGINS" default:"*"`
	ServiceAPIKey   string   `envconfig:"SERVICE_API_KEY" default:"insecure"`
	TokenCookieName string   `envconfig:"TOKEN_COOKIE_NAME" default:"auth_tkn"`
}

var conf *Config

const (
	DebugMode      = "debug"
	ProductionMode = "production"
)

// This function gets called automatically when the package is loaded
func init() {
	conf = &Config{}
	// This prefix means our variables will be in the form of GDEMO_MODE (for example)
	err := envconfig.Process("GDEMO", conf)
	if err != nil {
		fmt.Println("Fatal error processing configuration")
		panic(err)
	}
	l := conf.GetLogger()

	// Sanity check
	if !conf.IsDebugMode() && !conf.IsProductionMode() {
		l.Fatal("Invalid GDEMO_MODE variable, it must be either `debug` or `production`")
	}
}

// Cfg returns the configuration - will panic if the config has not been loaded or is nil (which shouldn't happen as that's implicit in the package init)
func Cfg() *Config {
	if conf == nil {
		panic("Config is nil")
	}
	return conf
}

func (cfg *Config) GetLogger() *logrus.Logger {
	var l = logrus.New()
	l.Formatter = &logrus.JSONFormatter{}
	return l
}

func (cfg *Config) IsDebugMode() bool {
	return cfg.Mode == DebugMode
}

func (cfg *Config) IsProductionMode() bool {
	return cfg.Mode == ProductionMode
}
