package WhatMac

import (
	"bytes"
	"io/ioutil"
	"log"

	"howett.net/plist"
)

var MacOSInfo macOSVersion

func Decode() {
	initialFile, err := ioutil.ReadFile("/System/Library/CoreServices/SystemVersion.plist")
	fileToReader := bytes.NewReader(initialFile)
	if err != nil {
		log.Panic(err)
	}
	decoder := plist.NewDecoder(fileToReader)
	err = decoder.Decode(&MacOSInfo)
}
