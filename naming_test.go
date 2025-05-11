//go:build nettest
// +build nettest

package gosam

import (
	"fmt"
	"testing"
)

func TestClientLookupInvalid(t *testing.T) {
	var err error

	client, err := NewClientFromOptions(SetDebug(false))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}

	addr, err := client.Lookup(`!(@#)`)
	if addr != "" || err == nil {
		t.Error("client.Lookup() should throw an error.")
	}

	repErr, ok := err.(ReplyError)
	if !ok {
		t.Fatalf("client.Lookup() should return a ReplyError")
	}
	if repErr.Result != ResultKeyNotFound && repErr.Result != ResultInvalidKey {
		t.Errorf("client.Lookup() should either throw an ResultKeyNotFound(i2p) or ResultInvalidKey(i2pd) error.\nGot:%+v%s%s\n", repErr, "!=", ResultKeyNotFound)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestClientLookupValid(t *testing.T) {
	client, err := NewDefaultClient()
	if err != nil {
		fmt.Printf("NewDefaultClient() should not throw an error.\n%s\n", err)
		return
	}

	addr, err := client.Lookup("zzz.i2p")
	if err != nil {
		fmt.Printf("client.Lookup() should not throw an error.\n%s\n", err)
		return
	}

	if addr == `GKapJ8koUcBj~jmQzHsTYxDg2tpfWj0xjQTzd8BhfC9c3OS5fwPBNajgF-eOD6eCjFTqTlorlh7Hnd8kXj1qblUGXT-tDoR9~YV8dmXl51cJn9MVTRrEqRWSJVXbUUz9t5Po6Xa247Vr0sJn27R4KoKP8QVj1GuH6dB3b6wTPbOamC3dkO18vkQkfZWUdRMDXk0d8AdjB0E0864nOT~J9Fpnd2pQE5uoFT6P0DqtQR2jsFvf9ME61aqLvKPPWpkgdn4z6Zkm-NJOcDz2Nv8Si7hli94E9SghMYRsdjU-knObKvxiagn84FIwcOpepxuG~kFXdD5NfsH0v6Uri3usE3XWD7Pw6P8qVYF39jUIq4OiNMwPnNYzy2N4mDMQdsdHO3LUVh~DEppOy9AAmEoHDjjJxt2BFBbGxfdpZCpENkwvmZeYUyNCCzASqTOOlNzdpne8cuesn3NDXIpNnqEE6Oe5Qm5YOJykrX~Vx~cFFT3QzDGkIjjxlFBsjUJyYkFjBQAEAAcAAA==` {
		t.Log("Success")
	} else {
		t.Errorf("Address of zzz.i2p != \nGKapJ8koUcBj~jmQzHsTYxDg2tpfWj0xjQTzd8BhfC9c3OS5fwPBNajgF-eOD6eCjFTqTlorlh7Hnd8kXj1qblUGXT-tDoR9~YV8dmXl51cJn9MVTRrEqRWSJVXbUUz9t5Po6Xa247Vr0sJn27R4KoKP8QVj1GuH6dB3b6wTPbOamC3dkO18vkQkfZWUdRMDXk0d8AdjB0E0864nOT~J9Fpnd2pQE5uoFT6P0DqtQR2jsFvf9ME61aqLvKPPWpkgdn4z6Zkm-NJOcDz2Nv8Si7hli94E9SghMYRsdjU-knObKvxiagn84FIwcOpepxuG~kFXdD5NfsH0v6Uri3usE3XWD7Pw6P8qVYF39jUIq4OiNMwPnNYzy2N4mDMQdsdHO3LUVh~DEppOy9AAmEoHDjjJxt2BFBbGxfdpZCpENkwvmZeYUyNCCzASqTOOlNzdpne8cuesn3NDXIpNnqEE6Oe5Qm5YOJykrX~Vx~cFFT3QzDGkIjjxlFBsjUJyYkFjBQAEAAcAAA==\n, check to see if it changed, %s", addr)
	}

	fmt.Println("Address of zzz.i2p:")
	// Addresses change all the time
	fmt.Println(addr)

	// Output:
	//Address of zzz.i2p:
	//
}
