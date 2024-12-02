package main

import (
	"enkhalifapro/locations/build"
	"enkhalifapro/locations/cmd"
)

var Version string

func main() {
	build.Version = Version
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
