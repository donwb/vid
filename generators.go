package main

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/bxcodec/faker/v3"
)

// CustomGenerator ...
func CustomGenerator() {
	_ = faker.AddProvider("customIdFaker", func(v reflect.Value) (interface{}, error) {
		return int64(43), nil
	})
	_ = faker.AddProvider("browser", func(v reflect.Value) (interface{}, error) {
		browsers := [5]string{"Chrome", "Safari", "Firefox", "Brave", "Internet Explorer"}
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		idx := r1.Intn(5)

		retVal := browsers[idx]
		return retVal, nil
	})

	_ = faker.AddProvider("device", func(v reflect.Value) (interface{}, error) {
		devices := [2]string{"Mac", "Windows"}
		idx := getRandomNumber(2)

		return devices[idx], nil

	})

	// _ = faker.AddProvider("gondoruwo", func(v reflect.Value) (interface{}, error) {
	// 	obj := Gondoruwo{
	// 		Name:       "Power",
	// 		Locatadata: 324,
	// 	}
	// 	return obj, nil
	// })

	// _ = faker.AddProvider("customUUID", func(v reflect.Value) (interface{}, error) {
	// 	s := CustomUUID{
	// 		0, 8, 7, 2, 3,
	// 	}
	// 	return s, nil
	// })
}

func getRandomNumber(upperBound int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Intn(upperBound)

}
