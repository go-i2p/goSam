package main

import (
	"fmt"
	"log"

	"github.com/eyedeekay/goSam"
)

/**
THIS is a freestanding test for SAMv3.2 AUTH commands using goSam. It's
intended to be run separate from the other tests so that you don't accidentally end
up setting SAM session passwords and leaving them in the PasswordManager if a test
fails for some reason before you can remove them.
**/

func main() {
	client, err := goSam.NewClientFromOptions()
	if err != nil {
		client, err = goSam.NewClientFromOptions(
			goSam.SetUser("user"),
			goSam.SetPass("password"),
		)
		fmt.Println("Looks like you restarted the I2P router before sending AUTH DISABLE.")
		fmt.Println("This probably means that your SAM Bridge is in a broken state where it can't")
		fmt.Println("accept HELLO or AUTH commands anymore. You should fix this by removing the")
		fmt.Println("sam.auth=true entry from sam.config.")
		err = client.TeardownAuth()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(err)
		panic(err)
	}
	err = client.SetupAuth("user", "password")
	if err != nil {
		log.Println(err)
	}
	client2, err := goSam.NewDefaultClient()
	if err != nil {
		log.Println(err)
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
