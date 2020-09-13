// +build nettest

package goSam

import "testing"

import (
	"fmt"
	"time"
	//"log"
	"net/http"
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
	server, err := NewClientFromOptions(SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	listener, err := server.Listen()
	if err != nil {
		t.Fatalf("Listener() Error: %q\n", err)
	}
	http.HandleFunc("/", HelloServer)
	go http.Serve(listener, nil)
	time.Sleep(time.Second * 15)

	client, err = NewClientFromOptions(SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	tr := &http.Transport{
		Dial: client.Dial,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("http://" + server.Base32() + ".b32.i2p")
	if err != nil {
		t.Fatalf("Get Error: %q\n", err)
	}
	defer resp.Body.Close()
	t.Log("Get returned ", resp)
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
