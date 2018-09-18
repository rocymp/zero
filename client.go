package zero

import (
	"context"
	"log"
	"net"
	"time"
)

// SocketService struct
type SocketClient struct {
	onMessage  func(*Message)
	hbInterval time.Duration
	saddr      string
	status     int
	startTS    int64
	stopCh     chan int
	conn       *Conn
}

func NewSocketClient(addr string, interval int) (sc *SocketClient) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return
	}

	sc = &SocketClient{
		hbInterval: time.Duration(interval) * time.Second,
		stopCh:     make(chan int),
		saddr:      addr,
		status:     STInited,
		startTS:    time.Now().Unix(),
	}

	cconn := NewConn(conn, sc.hbInterval, sc.hbInterval*3)

	sc.conn = cconn
	return
}

func (sc *SocketClient) SendMessage(cmdId int32, dataIn []byte) error {
	msg := NewMessage(cmdId, dataIn)

	return sc.conn.SendMessage(msg)
}

// RegMessageHandler register message handler
func (s *SocketClient) RegMessageHandler(handler func(*Message)) {
	s.onMessage = handler
}

// RegMessageHandler register message handler
func (s *SocketClient) Stop() {
	s.stopCh <- 0
}

// Serv Start socket service
func (s *SocketClient) online() {
	s.status = STRunning
	ctx, cancel := context.WithCancel(context.Background())

	defer func() {
		cancel()
		s.conn.Close()
	}()

	go s.conn.readCoroutine(ctx)
	go s.conn.writeCoroutine(ctx)

	for {
		select {
		case err := <-s.conn.done:
			log.Printf("Conn Error %#v\n", err)
			return

		case msg := <-s.conn.messageCh:
			if s.onMessage != nil {
				s.onMessage(msg)
			}
		case <-s.stopCh:
			return
		}
	}
}

// Serv Start socket service
func (s *SocketClient) Online() {
	go s.online()
	return
}
