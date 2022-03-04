package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/bxcodec/faker/v3"
)

func main() {
	fmt.Println("Starting.....")

	CustomGenerator()

	vids1 := createNewVideoSession()
	vids2 := createNewVideoSession()
	vids3 := createNewVideoSession()

	//prettyPrinter(finalVids)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		publish(vids1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		publish(vids2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		publish(vids3)
	}()
	wg.Wait()

	fmt.Println("Done!!")
}

func createNewVideoSession() VideoCallStruct {
	ids := Ids{}

	err := faker.FakeData(&ids)
	if err != nil {
		fmt.Println(err)
	}

	v := VideoCallStruct{}
	ct := ClientTime{}

	faker.FakeData(&v)
	faker.FakeData(&ct)

	v.Ids = ids
	v.ClientTime = ct

	manualVids := doManualFields(v)
	sessionVids := createSession(manualVids)
	finalVids := setVideoDetails(sessionVids)
	return finalVids
}

func doManualFields(vids VideoCallStruct) VideoCallStruct {
	vids.BrowserVersion = "98.0.4758"

	// do time stamp also
	vids.Model = vids.Device

	m := Mvpd{
		MvpdID:       "null",
		AuthRequired: "Not Logged in",
		AuthState:    "preview",
		DisplayName:  "null",
	}
	p := Player{PlayerName: "Top"}

	l := Library{
		Name:    "Prism JS",
		Version: "2.11.0",
	}

	vids.Mvpd = m
	vids.Player = p
	vids.Library = l

	return vids
}

func createSession(vids VideoCallStruct) VideoCallStruct {
	// need to establish a sessoin associate it w/a ukid and some other stuff
	// What is the start event called???
	vids.EventName = "start"
	vids.EventType = "video-qos"
	vids.ClientTime.ClientEventTimestamp = time.Now()
	vids.IPAddress = faker.IPv4()

	ids := Ids{
		Kruxid: faker.UUIDHyphenated(),
		Cdpid:  faker.UUIDHyphenated(),
		Wmukid: faker.UUIDHyphenated(),
	}
	vids.Ids = ids
	sessionID := faker.UUIDHyphenated()
	vids.Session.SessionID = sessionID
	vids.VideoSession.SessionID = sessionID

	return vids
}

func setVideoDetails(vids VideoCallStruct) VideoCallStruct {
	vids.Video.AssetID = "223 - Iowa vs Cincinnati"
	vids.Video.AssetType = "live"
	vids.Video.CurrentMediaState = "pending"
	vids.Video.CurrentVideoPosition = "0"
	vids.Video.StreamType = "live"
	vids.Video.VideoTitleID = 223

	return vids
}

type Tester struct {
	Domain    string `faker:"device"`
	IPv4      string `faker:"ipv4"`
	IPv6      string `faker:"ipv6"`
	Timestamp string `faker:"timestamp"`
	UKID      string `faker:"uuid_digit"`
	Name      string `faker:"first_name"`
	Anow      Another
}

type Another struct {
	Name string `faker:"first_name_male"`
	Last string
}
