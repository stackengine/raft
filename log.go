package raft

import (
	"fmt"
	"net"
)

// LogType describes various types of log entries.
type LogType uint8

const (
	// LogCommand is applied to a user FSM.
	LogCommand LogType = iota

	// LogNoop is used to assert leadership.
	LogNoop

	// LogAddPeer is used to add a new peer.
	LogAddPeer

	// LogRemovePeer is used to remove an existing peer.
	LogRemovePeer

	// LogBarrier is used to ensure all preceeding operations have been
	// applied to the FSM. It is similar to LogNoop, but instead of returning
	// once committed, it only returns once the FSM manager acks it. Otherwise
	// it is possible there are operations committed but not yet applied to
	// the FSM.
	LogBarrier
)

var logTypeName = map[LogType]string{
	LogCommand:    "LogCmd",
	LogNoop:       "LogNoop",
	LogAddPeer:    "LogAddPeer",
	LogRemovePeer: "LogRemovePeer",
	LogBarrier:    "LogBarrier",
}

func (t LogType) String() string {
	if t > LogBarrier {
		return fmt.Sprintf("Unknown LogType: %d", t)
	}
	return logTypeName[t]
}

// Log entries are replicated to all members of the Raft cluster
// and form the heart of the replicated state machine.
type Log struct {
	Index uint64
	Term  uint64
	Type  LogType
	Data  []byte

	// Peer is not exported since it is not transmitted, only used
	// internally to construct the Data field.
	peer net.Addr
}

// LogStore is used to provide an interface for storing
// and retrieving logs in a durable fashion
type LogStore interface {
	// Returns the first index written. 0 for no entries.
	FirstIndex() (uint64, error)

	// Returns the last index written. 0 for no entries.
	LastIndex() (uint64, error)

	// Gets a log entry at a given index
	GetLog(index uint64, log *Log) error

	// Stores a log entry
	StoreLog(log *Log) error

	// Stores multiple log entries
	StoreLogs(logs []*Log) error

	// Deletes a range of log entries. The range is inclusive.
	DeleteRange(min, max uint64) error
}

type Logger interface {
	Enable()
	Disable()

	Printf(format string, v ...interface{})
	Println(v ...interface{})

	Errorf(format string, v ...interface{})

	Infof(format string, v ...interface{})
	Infoln(v ...interface{})
}
