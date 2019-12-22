package main

import (
	"net/http"
)

type clientInfo struct {
	LatestHTTPVersionSupported string `json:"latest_http_version_supported"`
	SupportsHTTP2ServerPush    bool   `json:"supports_http2_server_push"`
}

func getClientInfo(w http.ResponseWriter) *clientInfo {
	info := new(clientInfo)
	info.LatestHTTPVersionSupported = "2.0"
	info.SupportsHTTP2ServerPush = false

	if _, ok := w.(http.Pusher); ok {
		info.SupportsHTTP2ServerPush = true
	}

	return info
}
