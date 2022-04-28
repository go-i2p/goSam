package main

import (
	"fmt"

	"github.com/eyedeekay/goSam"
)

/**
THIS is a freestanding test for SAMv3.2 AUTH commands using goSam. It's
intended to be run separate from the other tests so that you don't accidentally end
up setting SAM session passwords and leaving them in the PasswordManager if a test
fails for some reason before you can remove them.
**/

func main() {
	client, err := goSam.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	err = client.SetupAuth("user", "password")
	if err != nil {
		panic(err)
	}
	client2, err := goSam.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	conn, err := client2.Dial("", "idk.i2p")
	if err != nil {
		fmt.Println(err)
	}
	conn.Close()
	err = client.RemoveAuthUser("user")
	if err != nil {
		panic(err)
	}
	//fmt.Println(r)
	err = client.TeardownAuth()
	if err != nil {
		panic(err)
	}
	//r, err = client.NewDestination()
}
