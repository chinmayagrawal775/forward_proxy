package threadpool

import (
	"net"
	"sync"
)

type ThreadPool struct {
	workQueue chan net.Conn
	wg        sync.WaitGroup
}

func InitializeThreadPool(poolSize int, connectionHandler func(net.Conn)) *ThreadPool {
	pool := &ThreadPool{
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

func (p *ThreadPool) AddNewConnection(connection net.Conn) {
	p.workQueue <- connection
}

func (p *ThreadPool) Close() {
	close(p.workQueue)
	p.wg.Wait()
}
