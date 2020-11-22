// +build nettest

package goSam

import "testing"

import (
	"fmt"
	"time"
	//"log"
	"net/http"

	"github.com/eyedeekay/sam3/helper"
	"github.com/eyedeekay/sam3/i2pkeys"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

var client *Client

func setup(t *testing.T) {
	var err error

	// these tests expect a running SAM brige on this address
	client, err = NewClientFromOptions(SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
}

func TestCompositeClient(t *testing.T) {
	listener, err := sam.I2PListener("testservice", "127.0.0.1:7656", "testkeys")
	if err != nil {
		t.Fatalf("Listener() Error: %q\n", err)
	}
	http.HandleFunc("/", HelloServer)
	go http.Serve(listener, nil)

	listener2, err := sam.I2PListener("testservice2", "127.0.0.1:7656", "testkeys2")
	if err != nil {
		t.Fatalf("Listener() Error: %q\n", err)
	}
	//	http.HandleFunc("/", HelloServer)
	go http.Serve(listener2, nil)

	listener3, err := sam.I2PListener("testservice3", "127.0.0.1:7656", "testkeys3")
	if err != nil {
		t.Fatalf("Listener() Error: %q\n", err)
	}
	//	http.HandleFunc("/", HelloServer)
	go http.Serve(listener3, nil)

	client, err = NewClientFromOptions(SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	tr := &http.Transport{
		Dial: client.Dial,
	}
	client := &http.Client{Transport: tr}
	time.Sleep(time.Second * 30)
	go func() {
		resp, err := client.Get("http://" + listener.Addr().(i2pkeys.I2PAddr).Base32())
		if err != nil {
			t.Fatalf("Get Error: %q\n", err)
		}
		defer resp.Body.Close()
		t.Log("Get returned ", resp)
	}()
	go func() {
		resp, err := client.Get("http://" + listener2.Addr().(i2pkeys.I2PAddr).Base32())
		if err != nil {
			t.Fatalf("Get Error: %q\n", err)
		}
		defer resp.Body.Close()
		t.Log("Get returned ", resp)
	}()
	go func() {
		resp, err := client.Get("http://" + listener3.Addr().(i2pkeys.I2PAddr).Base32())
		if err != nil {
			t.Fatalf("Get Error: %q\n", err)
		}
		defer resp.Body.Close()
		t.Log("Get returned ", resp)
	}()

	time.Sleep(time.Second * 15)
}

func teardown(t *testing.T) {
	if err := client.Close(); err != nil {
		t.Fatalf("client.Close() Error: %q\n", err)
	}
}

func TestClientHello(t *testing.T) {
	setup(t)
	t.Log(client.Base32())
	teardown(t)
}

func TestNewDestination(t *testing.T) {
	setup(t)
	t.Log(client.Base32())
	if _, err := client.NewDestination(SAMsigTypes[3]); err != nil {
		t.Error(err)
	}
	teardown(t)
}
