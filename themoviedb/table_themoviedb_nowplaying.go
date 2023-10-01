package themoviedb

import (
	"context"
	"os"
	"strconv"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

/*
* NowPlaying
* https://developer.themoviedb.org/reference/movie-now-playing-list
* tmdb api function: GetMovieNowPlaying
*/
func tableNowPlaying() *plugin.Table {
	return &plugin.Table{
		Name: "themoviedb_nowplaying",
		Description: "Get a list of movies that are currently in theatres.",
		List: &plugin.ListConfig{
			Hydrate: listNowPlaying,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getNowPlaying,
		},
		Columns: []*plugin.Column{
			{Name: "poster_path", Type: proto.ColumnType_STRING, Description: "poster_path"},
			{Name: "adult", Type: proto.ColumnType_BOOL, Description: "adult"},
			{Name: "overview", Type: proto.ColumnType_STRING, Description: "overview"},
			{Name: "release_date", Type: proto.ColumnType_STRING, Description: "release_date"},
			{Name: "genres", Type: proto.ColumnType_JSON, Description: "genres"},
			{Name: "id", Type: proto.ColumnType_INT, Description: "id"},
			{Name: "original_title", Type: proto.ColumnType_STRING, Description: "original_title"},
			{Name: "original_language", Type: proto.ColumnType_STRING, Description: "original_language"},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "title"},
			{Name: "backdrop_path", Type: proto.ColumnType_STRING, Description: "backdrop_path"},
			{Name: "popularity", Type: proto.ColumnType_DOUBLE, Description: "popularity"},
			{Name: "vote_count", Type: proto.ColumnType_INT, Description: "vote_count"},
			{Name: "video", Type: proto.ColumnType_BOOL, Description: "video"},
			{Name: "vote_average", Type: proto.ColumnType_DOUBLE, Description: "vote_average"},
		},		
	}
}

func listNowPlaying(ctx context.Context, 
	d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	logger := plugin.Logger(ctx)
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	// Get TheMovieDB Language
	language := os.Getenv("THEMOVIEDB_LANGUAGE")
	themoviedbConfig := GetConfig(d.Connection)
	if themoviedbConfig.Language != nil {
		language = *themoviedbConfig.Language
	}

	page := 1
	options := make(map[string]string)
	options["page"] = strconv.Itoa(page)
	options["language"] = language

	for {
		movies, err := conn.GetMovieNowPlaying(options)
		if err != nil {
			logger.Trace("GetMovieNowPlaying Error:", err.Error())
			return nil, err
		}
		for _, t := range movies.Results {
			d.StreamListItem(ctx, t)
		}
		if int64(page) > movies.TotalPages {
			break
		}
		page++
		options["page"] = strconv.Itoa(page)
	}

	return nil, nil
}

func getNowPlaying(ctx context.Context, 
	d *plugin.QueryData, 
	h *plugin.HydrateData) (interface{}, error) {
	/*
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}*/
	logger := plugin.Logger(ctx)
	quals := d.EqualsQuals
	//steampipe logger
	logger.Warn("getNowPlaying", "quals", quals)
	id := quals["id"].GetInt64Value()
	logger.Warn("getNowPlaying", "id", id)
	return nil, nil
}