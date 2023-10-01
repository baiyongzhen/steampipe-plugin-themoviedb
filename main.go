package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/baiyongzhen/steampipe-plugin-themoviedb/themoviedb"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: themoviedb.Plugin})
}