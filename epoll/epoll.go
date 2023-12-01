package epoll

import (
	"net"
	"sync"

	"golang.org/x/sys/unix"
)

/*
1. 使用mutex + map进行, 瞬时高并发时map写很多
*/

type Epoller interface {
}

func NewEpoller() (Epoller, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}

	return &epoller{
		fd:          fd,
		connections: make(map[int]net.Conn),
	}, nil
}

type epoller struct {
	fd          int
	connections map[int]net.Conn
	mtx         sync.RWMutex
}
