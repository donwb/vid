package main

import (
	"fmt"
	"reflect"
)

func prettyPrinter(vids VideoCallStruct) {

	startSection("VIDEO INFORMATION ----")

	reflectPrinter(vids)

	startSection("Times")
	fmt.Println("Client Time: \t" + vids.ClientTime.ClientEventTimestamp.String())

	startSection("IDs")
	reflectPrinter(vids.Ids)

	startSection("MVPDs")
	reflectPrinter(vids.Mvpd)

	startSection("Player")
	reflectPrinter(vids.Player)

	startSection("ProductProperties")
	reflectPrinter(vids.ProductProperties)

	startSection("Session")
	reflectPrinter(vids.Session)

	startSection("Video")
	reflectPrinter(vids.Video)

	startSection("VideoSession")
	reflectPrinter(vids.VideoSession)

	startSection("Location")
	reflectPrinter(vids.Location)

	startSection("Library")
	reflectPrinter(vids.Library)

}

func dumpObject(vids VideoCallStruct) {
	fmt.Printf("%+v", vids)
}

func startSection(sectionName string) {
	fmt.Println()
	fmt.Println(sectionName)
}

func reflectPrinter(typeToPrint interface{}) {
	tValues := reflect.ValueOf(typeToPrint)
	tType := tValues.Type()

	for i := 0; i < tValues.NumField(); i++ {
		if tValues.Field(i).Kind() != reflect.Struct {
			fmt.Print("---- " + tType.Field(i).Name)
			fmt.Print(": \t")
			fmt.Println(tValues.Field(i).Interface())
		}

	}
}
