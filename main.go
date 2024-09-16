package main

import (
	"log"

	"os"
	"os/signal"
	"syscall"

	"github.com/chinmayagrawal775/forward_proxy/config"
	"github.com/chinmayagrawal775/forward_proxy/pkg/proxy"
	"github.com/chinmayagrawal775/forward_proxy/pkg/threadpool"
	"github.com/chinmayagrawal775/forward_proxy/server"
	"github.com/chinmayagrawal775/forward_proxy/utils"
)

// this will initialize the list of restricted host & words for this proxy server
func init() {
	utils.LoadFile("config/restricted-host.txt", &config.RestrictedHosts)
	utils.LoadFile("config/restricted-words.txt", &config.RestrictedWords)
}

func main() {

	// this will start profiling server in seperate go routine , so that profiling server will not blocked by proxy server
	server.InitProfilingServer()

	// booting up proxy server
	proxyServer := server.InitProxyServer()

	// threadpooling: Number of connections capped at 50
	threadpool := threadpool.InitializeThreadPool(50, proxy.ConnectionHandler)

	shutdownSignal := make(chan os.Signal, 1) // channel for singanlling the server closing
	loopExitSignal := make(chan bool, 1)      // channel to notify to break out from loop

	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-shutdownSignal
		log.Println("Closing Proxy Server.")
		proxyServer.ShutdownServer()
		threadpool.Close()
		loopExitSignal <- true // signalling 'loopExitSignal' channel to break out of the loop
		log.Println("Server Shutdown Gracefully. Bye..!!")
		os.Exit(0)
	}()

	for {
		conn := proxyServer.AcceptNewConnection()
		if conn == nil {
			<-loopExitSignal
			break
		}

		// go proxy.ConnectionHandler(conn) // this fork new thread each time a new connection came. Will spike the goroutines in case of huge traffic.
		threadpool.AddNewConnection(conn)
	}
}
