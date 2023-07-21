package main

import (
	lkit_go "github.com/xlkness/lkit-go"
	gmtool2 "joytool/apps/gmtool"
	user2 "joytool/apps/user"
)

func main() {
	scd := lkit_go.NewScheduler()

	user := user2.NewApp()
	gmtool := gmtool2.NewApp()

	scd.CreateApp(user, gmtool)

	err := scd.Run()
	if err != nil {
		panic(err)
	}
}
