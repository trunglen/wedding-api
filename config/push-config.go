package config

import (
	"seed/x/fcm"
)

type pushConfig struct {
	APIKey string `toml:"api_key"`
}

var PushConfig pushConfig

func (p pushConfig) initPush() {
	fcm.NewFCM(p.APIKey)
}
