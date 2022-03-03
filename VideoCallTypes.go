package main

import "time"

type VideoCallStruct struct {
	BrowserName       string            `json:"browserName" faker:"browser"`
	BrowserVersion    string            `json:"browserVersion"`
	ClientTime        ClientTime        `json:"clientTime"`
	Device            string            `json:"device" faker:"device"`
	DeviceBrandName   string            `json:"deviceBrandName"`
	EventName         string            `json:"eventName"`
	EventType         string            `json:"eventType"`
	Ids               Ids               `json:"ids"`
	IPAddress         string            `json:"ipAddress"`
	Model             string            `json:"model"`
	Mvpd              Mvpd              `json:"mvpd"`
	OperatingSystem   string            `json:"operatingSystem"`
	OsVersion         string            `json:"osVersion"`
	Player            Player            `json:"player"`
	ProductProperties ProductProperties `json:"productProperties"`
	Session           Session           `json:"session"`
	Video             Video             `json:"video"`
	VideoSession      VideoSession      `json:"videoSession"`
	Location          Location          `json:"location"`
	Cdn               string            `json:"cdn"`
	Library           Library           `json:"library"`
}
type ClientTime struct {
	ClientEventTimestamp time.Time `json:"clientEventTimestamp" faker:timestamp`
}
type Ids struct {
	Cdpid  string `json:"cdpid" faker:"uuid_hyphenated"`
	Kruxid string `json:"kruxid" faker:"uuid_hyphenated"`
	Wmukid string `json:"wmukid" faker:"uuid_hyphenated"`
}
type Mvpd struct {
	AuthRequired string `json:"authRequired"`
	AuthState    string `json:"authState"`
	DisplayName  string `json:"displayName"`
	MvpdID       string `json:"mvpdId"`
}
type Player struct {
	PlayerName string `json:"playerName"`
}
type ProductProperties struct {
	AppName    string `json:"appName"`
	AppVersion string `json:"appVersion"`
	Brand      string `json:"brand"`
	Platform   string `json:"platform"`
	SubBrand   string `json:"subBrand"`
}
type Session struct {
	SessionID string `json:"sessionId"`
}
type Video struct {
	AssetID              string `json:"assetId"`
	AssetType            string `json:"assetType"`
	CurrentMediaState    string `json:"currentMediaState"`
	CurrentVideoPosition string `json:"currentVideoPosition"`
	StreamType           string `json:"streamType"`
	VideoTitleID         int    `json:"videoTitleId"`
}
type VideoSession struct {
	SessionID time.Time `json:"sessionId"`
}
type Location struct {
	City     string `json:"city"`
	Country  string `json:"country"`
	State    string `json:"state"`
	TimeZone string `json:"timeZone"`
}
type Library struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
