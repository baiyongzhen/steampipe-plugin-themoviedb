package themoviedb

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-themoviedb",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"themoviedb_nowplaying": tableNowPlaying(),
			"themoviedb_popular": tablePopular(),
			"themoviedb_toprated": tableMovieTopRated(),
			"themoviedb_upcoming": tableMovieUpcoming(),
		},
	}
	return p
}