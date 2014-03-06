package goredis_server

import (
	"time"
)

// 包装info的输出
type Info struct {
	server              *GoRedisServer
	ops_per_sec         int64
	last_total_commands int64
}

func NewInfo(server *GoRedisServer) *Info {
	v := &Info{
		server: server,
	}
	go v.secondTicker()
	return v
}

func (i *Info) secondTicker() {
	ticker := time.NewTicker(time.Second * 1)
	for _ = range ticker.C {
		total := i.total_commands_processed()
		i.ops_per_sec, i.last_total_commands = total-i.last_total_commands, total
	}
}

func (i *Info) Version() string {
	return VERSION
}

func (i *Info) connected_clients() int64 {
	return i.server.counters.Get("connection").Count()
}

func (i *Info) instantaneous_ops_per_sec() int64 {
	return i.ops_per_sec
}

func (i *Info) total_commands_processed() int64 {
	return i.server.cmdCateCounters.Get("total").Count()
}
