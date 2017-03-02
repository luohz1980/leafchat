package conf

import (
	"log"
	"time"
)

var (
	// gate conf
	Encoding               = "json" // "json" or "protobuf"
	PendingWriteNum        = 2000
	LenMsgLen              = 2
	MinMsgLen       uint32 = 2
	MaxMsgLen       uint32 = 4096
	LittleEndian           = false
	HTTPTimeout            = 10 * time.Second

	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	ChanRPCLen         = 10000

	// log conf
	LogFlag = log.LstdFlags
)
