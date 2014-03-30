package goredis_proxy

import (
	. "GoRedis/goredis"
	"GoRedis/libs/stdlog"
	"strings"
)

// 配置Proxy
// > config master localhost:6379
// > config slave localhost:6379
// > config mode rw/rrw
func (server *GoRedisProxy) OnCONFIG(session *Session, cmd *Command) (reply *Reply) {
	if cmd.Len() < 3 {
		return ErrorReply("bad command")
	}
	field := strings.ToUpper(cmd.StringAtIndex(1))
	var err error
	switch field {
	case "MASTER":
		host := cmd.StringAtIndex(2)
		err = server.resetMaster(host)
	case "SLAVE":
		host := cmd.StringAtIndex(2)
		err = server.resetSlave(host)
	case "MODE":
		mode := cmd.StringAtIndex(2)
		err = server.resetMode(mode)
	default:
		return ErrorReply("not support")
	}

	if err == nil {
		reply = StatusReply("OK")
	} else {
		reply = ErrorReply(err)
	}
	return
}

func (server *GoRedisProxy) resetMaster(host string) (err error) {
	server.Suspend()
	defer server.Resume()

	stdlog.Println("CONFIG master", host)
	if server.master != nil {
		server.master.Close()
	}
	server.options.MasterHost = host
	server.master, err = NewRemoteSession(server.options.MasterHost, server.options.PoolSize)
	return
}

func (server *GoRedisProxy) resetSlave(host string) (err error) {
	server.Suspend()
	defer server.Resume()

	stdlog.Println("CONFIG slave", host)
	if server.slave != nil {
		server.slave.Close()
	}
	server.options.SlaveHost = host
	server.slave, err = NewRemoteSession(server.options.SlaveHost, server.options.PoolSize)
	return
}

func (server *GoRedisProxy) resetMode(mode string) (err error) {
	server.Suspend()
	defer server.Resume()

	stdlog.Println("CONFIG mode", mode)
	server.options.Mode = mode
	return
}
