package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

// +-------------------------------------------------------+
// | 42 Bit Timestamp | 10 Bit WorkID | 12 Bit Sequence ID |
// +-------------------------------------------------------+

const (
	epoch           int64 = 1526285084373
	numWorkerBits         = 10
	numSequenceBits       = 12
	MaxWorkId             = -1 ^ (-1 << numWorkerBits) //取低numWorkerBits位
	MaxSequence           = -1 ^ (-1 << numSequenceBits)
)

type SnowFlake struct {
	lastTimestamp uint64
	sequence      uint32
	workerId      uint32
	lock          sync.Mutex
}

func (sf *SnowFlake) pack() uint64 {
	uuid:=(sf.lastTimestamp<<(numSequenceBits+numWorkerBits))|(uint64(sf.workerId)<<numSequenceBits)|(uint64(sf.sequence))
	return uuid
}

func timestamp() uint64 {
	return uint64(time.Now().UnixNano()/int64(1000000)-epoch)
}

func New(workerId uint32) (*SnowFlake,error) {
	if workerId<0||workerId>MaxWorkId {
		return nil,errors.New("invaild worker id")
	}
	return &SnowFlake{workerId: workerId},nil
}

func (sf *SnowFlake) Generate() (uint64,error) {
	sf.lock.Lock()
	defer sf.lock.Unlock()

	ts:=timestamp()
	if ts==sf.lastTimestamp {
		sf.sequence=(sf.sequence+1)&MaxSequence
		if sf.sequence==0 {
			ts=sf.waitNextMilli(ts)
		}
	}else{
		sf.sequence=0
	}

	if ts<sf.lastTimestamp {
		return 0,errors.New("invaild system clock")
	}

	sf.lastTimestamp=ts
	return sf.pack(),nil
}

func (sf *SnowFlake) waitNextMilli(ts uint64) uint64 {
	for ts == sf.lastTimestamp {
		time.Sleep(100*time.Millisecond)
		ts=timestamp()
	}
	return ts
}

func main() {
	sn,err:=New(1)
	if err!=nil {
		log.Fatal("雪花算法初始化出错")
		return
	}
	uuid,err:=sn.Generate()
	if err!=nil {
		fmt.Println("uuid 生成出错")
		return
	}
	fmt.Printf("uuid=%d",uuid)
}
