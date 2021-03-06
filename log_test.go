package raft

import "github.com/Sirupsen/logrus"

// logs is a mock logger implementation used for raft testing
type logs struct {
	*logrus.Logger
}

func (logs) Enable()  {}
func (logs) Disable() {}

func (a logs) Printf(format string, v ...interface{}) {
	a.Logger.Infof(format, v...)
}
func (a logs) Println(v ...interface{}) {
	a.Logger.Println(v...)
}

func (a logs) Errorf(format string, v ...interface{}) {
	a.Logger.Errorf(format, v...)
}

func (a logs) Infof(format string, v ...interface{}) {
	a.Logger.Infof(format, v...)
}

func (a logs) Infoln(v ...interface{}) {
	a.Logger.Infoln(v...)
}
