package main

import (
	"enkhalifapro/connections/build"
	"enkhalifapro/connections/cmd"
)

var Version string

func main() {
	build.Version = Version
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
