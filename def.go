package zero

const (
	// STUnknown Unknown
	STUnknown = iota
	// STInited Inited
	STInited
	// STRunning Running
	STRunning
	// STStop Stop
	STStop
)

const (
	// MsgHeartbeat heartbeat
	MsgHeartbeat = iota
	// MsgTaskStart 启动任务命令字
	MsgTaskStart
	// MsgTaskStop 停止任务命令字
	MsgTaskStop
)
