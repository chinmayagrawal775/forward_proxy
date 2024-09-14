package pool

import (
	"net"
	"sync"
)

type Pool struct {
	workQueue chan net.Conn
	wg        sync.WaitGroup
}

func InitializeThredaPool(poolSize int, connectionHandler func(net.Conn)) *Pool {
	pool := &Pool{
		workQueue: make(chan net.Conn),
	}

	pool.wg.Add(poolSize)

	for i := 0; i < poolSize; i++ {
		go func() {
			defer pool.wg.Done()
			for connection := range pool.workQueue {
				connectionHandler(connection)
			}
		}()
	}

	return pool
}

func (p *Pool) AddNewConnection(connection net.Conn) {
	p.workQueue <- connection
}

func (p *Pool) Close() {
	close(p.workQueue)
	p.wg.Wait()
}
