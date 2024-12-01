package main

import (
	"enkhalifapro/persons/build"
	"enkhalifapro/persons/cmd"
)

var Version string

func main() {
	build.Version = Version
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
