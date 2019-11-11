package main

type clientInfo struct {
	LatestHTTPVersionSupported string `json:"latest_http_version_supported"`
}

func getClientInfo() *clientInfo {
	info := new(clientInfo)
	info.LatestHTTPVersionSupported = "2.0"
	return info
}
