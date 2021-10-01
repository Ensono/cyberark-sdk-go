package config

import (
	"net/http"

	"github.com/amido/cyberark-sdk-go/cyberark"
)

// A mode for logging on to CA,
type LogonMode int

// Mode for parser behaviour
const (
	CyberArk LogonMode = iota // default mode
	AML
	LDAP
	Radius // radius
)

type Config struct {
	Username          *string
	Password          *string
	LogDebug          *bool
	MaskAuthorization *bool
	Endpoint          *string
	Token             *string
	LogonMode         *LogonMode

	// The HTTP client to use when sending requests. Defaults to
	// `http.DefaultClient`.
	HTTPClient *http.Client
}

func NewConfig() *Config {

	return &Config{
		MaskAuthorization: cyberark.Bool(true),
		LogDebug:          cyberark.Bool(false),
		HTTPClient:        http.DefaultClient,
	}
}

func (c *Config) WithPassword(password string) *Config {
	c.Password = cyberark.String(password)
	return c
}

func (c *Config) WithUsername(username string) *Config {
	c.Username = cyberark.String(username)
	return c
}

func (c *Config) WithToken(token string) *Config {
	c.Token = cyberark.String(token)
	return c
}

func (c *Config) WithEndpoint(endpoint string) *Config {
	c.Endpoint = cyberark.String(endpoint)
	return c
}

func (c *Config) WithDebug(debug bool) *Config {
	c.LogDebug = cyberark.Bool(debug)
	return c
}

func (c *Config) WithLogonMode(mode LogonMode) *Config {
	c.LogonMode = &mode
	return c
}

func (c *Config) WithMaskAuthorization(debug bool) *Config {
	c.MaskAuthorization = cyberark.Bool(debug)
	return c
}

func (c *Config) setTokenOnConfig() *Config {
	
}