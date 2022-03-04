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
		browsers := []string{"Chrome", "Safari", "Firefox", "Brave", "Internet Explorer"}
		return getStringValue(browsers), nil
	})
	_ = faker.AddProvider("device", func(v reflect.Value) (interface{}, error) {
		devices := []string{"Mac", "Windows"}
		return getStringValue(devices), nil
	})
	_ = faker.AddProvider("deviceBrandName", func(v reflect.Value) (interface{}, error) {
		brands := []string{"Apple", "Microsoft", "Samsung", "LG", "Dell"}

		return getStringValue(brands), nil

	})
	_ = faker.AddProvider("operatingSystem", func(v reflect.Value) (interface{}, error) {
		os := []string{"Mac OS X", "Windows 10", "Windows 11", "Windows 95", "iOS", "Android X"}

		return getStringValue(os), nil
	})
	_ = faker.AddProvider("osVersion", func(v reflect.Value) (interface{}, error) {
		versions := []string{"10.15.7", "10.15.6", "10.15.5", "15.1", "15.2", "14.1", "14.2", "14.3"}
		return getStringValue(versions), nil
	})
	_ = faker.AddProvider("appName", func(v reflect.Value) (interface{}, error) {
		appNames := []string{"Watch NCAA", "Watch TBS", "Watch TRU", "Watch TNT", "Watch Eleague"}
		return getStringValue(appNames), nil
	})
	_ = faker.AddProvider("appVersion", func(v reflect.Value) (interface{}, error) {
		versions := []string{"1.0.0", "0.0.1"}
		return getStringValue(versions), nil
	})
	_ = faker.AddProvider("brand", func(v reflect.Value) (interface{}, error) {
		brands := []string{"NCAA", "TBS", "TNT", "TRU", "ELEAGUE"}
		return getStringValue(brands), nil
	})
	_ = faker.AddProvider("platform", func(v reflect.Value) (interface{}, error) {
		platforms := []string{"Web", "Mobile", "Connected"}
		return getStringValue(platforms), nil
	})
	_ = faker.AddProvider("subBrand", func(v reflect.Value) (interface{}, error) {
		subs := []string{"ncaa.com/march-madness-live"}
		return getStringValue(subs), nil
	})
	_ = faker.AddProvider("city", func(v reflect.Value) (interface{}, error) {
		cities := []string{"NEW YORK", "ATLANTA", "DALLAS", "DETROIT", "GAINESVILLE"}
		return getStringValue(cities), nil
	})
	_ = faker.AddProvider("country", func(v reflect.Value) (interface{}, error) {
		countries := []string{"US"}
		return getStringValue(countries), nil
	})
	_ = faker.AddProvider("state", func(v reflect.Value) (interface{}, error) {
		states := []string{"NY", "FL", "GA", "KY", "TN", "SC", "CA"}
		return getStringValue(states), nil
	})
	_ = faker.AddProvider("cdn", func(v reflect.Value) (interface{}, error) {
		cdns := []string{"akamai", "cloudfront", "level3"}
		return getStringValue(cdns), nil
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

func getStringValue(values []string) string {
	idx := getRandomNumber(len(values))
	return values[idx]
}

func getRandomNumber(upperBound int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Intn(upperBound)

}
