package zero

import (
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
	conn       *net.TCPConn
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
		conn:       conn,
		status:     STInited,
		startTS:    time.Now().Unix(),
	}

	return
}

func (sc *SocketClient) SendMessage(cmdId int32, dataIn []byte) error {
	msg := NewMessage(cmdId, dataIn)
	data, err := Encode(msg)
	if err != nil {
		return err
	}
	sc.conn.Write(data)

	return nil
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
func (s *SocketClient) Online() {
	s.status = STRunning

	go func() {
		hbData := make([]byte, 0)
		timer := time.NewTicker(s.hbInterval)

		for {
			select {
			case <-s.stopCh:
				return
			case <-timer.C:
				s.SendMessage(MsgHeartbeat, hbData)
			}
		}
	}()
}
