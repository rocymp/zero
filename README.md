# zero
A Lightweight Socket Service with heartbeat, Can be easily used in TCP server development.

## Requirements

Go version: 1.9.x or later

## Usage

```
go get -u github.com/rocymp/zero
```

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
