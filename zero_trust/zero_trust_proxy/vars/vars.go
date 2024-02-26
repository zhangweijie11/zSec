package vars

import (
	"crypto/tls"
	"github.com/zhangweijie11/zSec/zero_trust/zero_trust_proxy/config"
)

const (
	CookieName = "secProxy_Authorization"
	HeaderName = "secProxy-Jwt-Assertion"
)

var (
	ConfigPath   string
	ConfigFile   = "config.yaml"
	Conf         *config.Config
	TlsConfig    *tls.Config
	DebugMode    bool
	CurDir       string
	CaKey        string
	CaCert       string
	CallbackPath = "/.xsec/callback"
	LogoutPath   = "/.xsec/logout"
)
