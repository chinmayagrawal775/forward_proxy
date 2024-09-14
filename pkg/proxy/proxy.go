package proxy

import (
	"bufio"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/chinmayagrawal775/forward_proxy/utils"
)

func ConnectionHandler(connection net.Conn) {
	defer connection.Close()

	buff := bufio.NewReader(connection)
	proxiedRequest, readErr := http.ReadRequest(buff)
	if readErr != nil {
		log.Printf("Error while reading request: %s", readErr.Error())
		return
	}

	log.Printf("Client: %s Request URL: %s\n", connection.RemoteAddr().String(), proxiedRequest.Host)

	if utils.IsRestrictedHost(proxiedRequest.Host) {
		connection.Write([]byte("HTTP/1.1 403 Forbidden\r\n\r\n Access to the site blocked!!\r\n"))
		log.Printf("Client: %s \t403 Forbidden\n", connection.RemoteAddr().String())
		return
	}

	targetConnAddr := proxiedRequest.Host
	if !strings.Contains(targetConnAddr, ":443") {
		targetConnAddr = proxiedRequest.Host + ":80"
	}

	targetConn, err := net.Dial("tcp", targetConnAddr)
	if err != nil {
		log.Println("Error connecting to target server:", err)
		connection.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n"))
		return
	}

	defer targetConn.Close()

	if proxiedRequest.Method == http.MethodConnect { // for HTTPs
		// this will create a tunnel between the client and target server to allow encrypted communication
		connection.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
	} else {
		//  forward the client's HTTP request to the target server
		if err := proxiedRequest.Write(targetConn); err != nil {
			log.Println("Error writing request to target server:", err)
			return
		}
	}

	go io.Copy(targetConn, connection)
	io.Copy(connection, targetConn)

	log.Printf("Client: %s \t200 OK\n", connection.RemoteAddr().String())

	/*
		Below code is for checking if the response contained the restricted words
		Commented this flow, as this works only for HTTP, not HTTPs
		Because HTTPs use TLS, and to read response from this, we need to break encryption
		Which is something we should not do.
		But if you want to try it out, I have mentioned the detailed steps here.
	*/

	// STEP-1: READ RESPONSE FROM PROXIED REQUEST
	// resp, err := http.ReadResponse(bufio.NewReader(targetConn), proxiedRequest)
	// if err != nil {
	// 	log.Println("Error reading response from target server:", err)
	// 	return
	// }
	// defer resp.Body.Close()

	// STEP-2: READ BODY BUFFER FROM RESPONSE
	// bodyBuffer := new(bytes.Buffer)
	// if _, err = bodyBuffer.ReadFrom(resp.Body); err != nil {
	// 	log.Println("Error reading response body:", err)
	// 	return
	// }

	// STEP-3: CHECK IF THE BODY BUFFER(ACTUAL RESPONSE) CONTAINS THE RESTRICED WORD
	// if utils.IsRestrictedWord(string(bodyBuffer.Bytes())) {
	// 	connection.Write([]byte("HTTP/1.1 403 Forbidden\r\n\r\n Cannot show response, as it contains restricted words\r\n"))
	// 	log.Printf("Client: %s \t403 Forbidden\n", connection.RemoteAddr().String())
	// 	return
	// }

	// STEP-4: WRITE BACK THE BODY BUFFER TO THE RESPONSE
	// resp.Body = io.NopCloser(bytes.NewReader(bodyBuffer.Bytes()))

	// STEP-5: SEND THE RESPONSE BACK TO CONNECTION
	// err = resp.Write(connection)
	// if err != nil {
	// 	log.Println("Error forwarding response to client:", err)
	// 	return
	// }
}
