// +build nettest

package goSam

import "testing"

import (
	"fmt"
	"time"
	//"log"
	"net/http"

	"github.com/eyedeekay/sam3"
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
	sam, err := sam3.NewSAM("127.0.0.1:7656")
	if err != nil {
		t.Fatalf("Listener() Error: %q\n", err)
	}
	keys, err := sam.NewKeys()
	if err != nil {
		t.Fatalf("Listener() Error: %q\n", err)
	}
	stream, err := sam.NewStreamSession("serverTun", keys, sam3.Options_Medium)
	if err != nil {
		t.Fatalf("Listener() Error: %q\n", err)
	}
	listener, err := stream.Listen()
	if err != nil {
		t.Fatalf("Listener() Error: %q\n", err)
	}
	http.HandleFunc("/", HelloServer)
	go http.Serve(listener, nil)

	client, err = NewClientFromOptions(SetDebug(true))
	if err != nil {
		t.Fatalf("NewDefaultClient() Error: %q\n", err)
	}
	tr := &http.Transport{
		Dial: client.Dial,
	}
	client := &http.Client{Transport: tr}
	time.Sleep(time.Second * 30)
	resp, err := client.Get("http://" + keys.Addr().Base32())
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
