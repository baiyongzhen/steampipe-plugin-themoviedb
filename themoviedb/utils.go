package themoviedb

/**
* https://github.com/cyruzin/golang-tmdb
* This is a Golang wrapper for working with TMDb API. It aims to support version 3.
* An API Key is required. To register for one, head over to themoviedb.org.
* This product uses the TMDb API but is not endorsed or certified by TMDb.
**/
import (
	"context"
	"os"

	"github.com/cyruzin/golang-tmdb"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*tmdb.Client, error) {
	// Get TheMovieDB API Token
	token := os.Getenv("THEMOVIEDB_TOKEN")
	themoviedbConfig := GetConfig(d.Connection)
	if themoviedbConfig.Token != nil {
		token = *themoviedbConfig.Token
	}

	tmdbClient, err := tmdb.Init(token)	
	if err != nil {
		return nil, err
	}
	// Enabling auto retry functionality.
	tmdbClient.SetClientAutoRetry()
	return tmdbClient, nil
}