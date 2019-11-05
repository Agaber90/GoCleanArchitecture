package models

//SitesRequest to accept list of URLS
type SitesRequest struct {
	URLS[] string `json:"urls"`
}
