# zero
A Lightweight Socket Service with heartbeat, Can be easily used in TCP server development.

## Requirements

Go version: 1.9.x or later

## Usage

```
go get -u github.com/rocymp/zero
```
### Server Case
```go
import "github.com/rocymp/zero"

func main() {
 	host := "127.0.0.1:18787"

 	ss, err := zero.NewSocketService(host)
	if err != nil {
		return
	}

	// set Heartbeat
	ss.SetHeartBeat(5*time.Second, 30*time.Second)

	// net event
	ss.RegMessageHandler(HandleMessage)
	ss.RegConnectHandler(HandleConnect)
	ss.RegDisconnectHandler(HandleDisconnect)

	ss.Serv()
}


```
### Client Case
```go
package main

import (
	"github.com/rocymp/zero"
)

func main() {

	cs := zero.NewSocketClient("127.0.0.1:18888", 3)
	if cs == nil {
		log.Printf("connect failed\n")
		return 
	}

	// handler server message
	cs.RegMessageHandler(HandleMessage)

	// client online and heartbeat
	cs.Online()

	// client send message
	cs.SendMessage(23, []byte("hello world!"))

	// client stop
	cs.Stop()
	
}


```
