package themoviedb

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type themoviedbConfig struct {
	Token     *string `cty:"token"`
	Language  *string `cty:"language"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
	"language": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &themoviedbConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) themoviedbConfig {
	if connection == nil || connection.Config == nil {
		return themoviedbConfig{}
	}
	config, _ := connection.Config.(themoviedbConfig)
	return config
}