package main

import "time"

type VideoCallStruct struct {
	BrowserName       string            `json:"browserName" faker:"browser"`
	BrowserVersion    string            `json:"browserVersion"`
	ClientTime        ClientTime        `json:"clientTime"`
	Device            string            `json:"device" faker:"device"`
	DeviceBrandName   string            `json:"deviceBrandName" faker:"deviceBrandName"`
	EventName         string            `json:"eventName"`
	EventType         string            `json:"eventType"`
	Ids               Ids               `json:"ids"`
	IPAddress         string            `json:"ipAddress"`
	Model             string            `json:"model"`
	Mvpd              Mvpd              `json:"mvpd"`
	OperatingSystem   string            `json:"operatingSystem" faker:"operatingSystem"`
	OsVersion         string            `json:"osVersion" faker:"osVersion"`
	Player            Player            `json:"player"`
	ProductProperties ProductProperties `json:"productProperties"`
	Session           Session           `json:"session"`
	Video             Video             `json:"video"`
	VideoSession      VideoSession      `json:"videoSession"`
	Location          Location          `json:"location"`
	Cdn               string            `json:"cdn"  faker:"cdn"`
	Library           Library           `json:"library"`
}
type ClientTime struct {
	ClientEventTimestamp time.Time `json:"clientEventTimestamp"`
}
type Ids struct {
	Cdpid  string `json:"cdpid"`
	Kruxid string `json:"kruxid"`
	Wmukid string `json:"wmukid"`
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
	AppName    string `json:"appName" faker:"appName"`
	AppVersion string `json:"appVersion" faker:"appVersion"`
	Brand      string `json:"brand" faker:"brand"`
	Platform   string `json:"platform" faker:"platform"`
	SubBrand   string `json:"subBrand" faker:"subBrand"`
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
	SessionID string `json:"sessionId"`
}
type Location struct {
	City     string `json:"city" faker:"city"`
	Country  string `json:"country" faker:"country"`
	State    string `json:"state" faker:"state"`
	TimeZone string `json:"timeZone" faker:"timezone"`
}
type Library struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
