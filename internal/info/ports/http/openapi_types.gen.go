// Package http provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package http

// SongDetail defines model for SongDetail.
type SongDetail struct {
	Link        string `json:"link"`
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
}

// GetInfoParams defines parameters for GetInfo.
type GetInfoParams struct {
	Group string `form:"group" json:"group"`
	Song  string `form:"song" json:"song"`
}
