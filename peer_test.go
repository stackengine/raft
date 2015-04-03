package raft

import (
	"github.com/stackengine/selog"
	"io/ioutil"
	"net"
	"os"
	"testing"
)

func TestPeerSetup(t *testing.T) {
	SeLog = selog.Register("raft", 0)
	SeLog.Enable()
}

func TestJSONPeers(t *testing.T) {
	// Create a test dir
	dir, err := ioutil.TempDir("", "raft")
	if err != nil {
		t.Fatalf("err: %v ", err)
	}
	defer os.RemoveAll(dir)

	// Create the store
	_, trans := NewInmemTransport()
	store := NewJSONPeers(dir, trans)

	// Try a read, should get nothing
	peers, err := store.Peers()
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(peers) != 0 {
		t.Fatalf("peers: %v", peers)
	}

	// Initialize some peers
	newPeers := []net.Addr{NewInmemAddr(), NewInmemAddr(), NewInmemAddr()}
	if err := store.SetPeers(newPeers); err != nil {
		t.Fatalf("err: %v", err)
	}

	// Try a read, should peers
	peers, err = store.Peers()
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(peers) != 3 {
		t.Fatalf("peers: %v", peers)
	}
}

func TestJSONInit(t *testing.T) {
	// Create a test dir
	dir, err := ioutil.TempDir("", "raft")
	if err != nil {
		t.Fatalf("err: %v ", err)
	}
	defer os.RemoveAll(dir)

	// Create the store
	_, trans := NewInmemTransport()
	store := NewJSONPeers(dir, trans)

	// Try a read, should get nothing
	peers, err := store.Peers()
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(peers) != 0 {
		t.Fatalf("peers: %v", peers)
	}

	// Initialize some peers
	newPeers := []net.Addr{NewInmemAddr(), NewInmemAddr(), NewInmemAddr()}
	if err := store.SetPeers(newPeers); err != nil {
		t.Fatalf("err: %v", err)
	}

	// Try a read, should peers
	peers, err = store.Peers()
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(peers) != 3 {
		t.Fatalf("peers: %v", peers)
	}

	// now Init.
	err = store.Init()
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	// Try a read, should get nothing
	peers, err = store.Peers()
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(peers) != 0 {
		t.Fatalf("peers: %v", peers)
	}

}
