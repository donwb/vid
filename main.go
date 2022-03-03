package main

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

func main() {
	fmt.Println("here!")

	CustomGenerator()

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

	finalVids := doManualFields(v)

	fmt.Printf("%+v", finalVids)

}

func doManualFields(vids VideoCallStruct) VideoCallStruct {
	vids.BrowserVersion = "98.0.4758"

	m := Mvpd{
		MvpdID:       "null",
		AuthRequired: "Not Logged in",
		AuthState:    "preview",
		DisplayName:  "null",
	}
	vids.Mvpd = m

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
