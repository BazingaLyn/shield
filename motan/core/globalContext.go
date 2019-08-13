package core

import (
	cfg "shield/motan/config"
)

type Context struct {
	ConfigFile       string
	Config           cfg.Config
	RegistryURLs     map[string]*URL
	RefersURLs       map[string]*URL
	HTTPClientURLs   map[string]*URL
	BasicReferURLs   map[string]*URL
	ServiceURLs      map[string]*URL
	BasicServiceURLs map[string]*URL
	AgentURL         *URL
	ClientURL        *URL
	ServerURL        *URL

	application string
	pool        string
}
