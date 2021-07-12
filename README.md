# UUID--雪花算法
格式： 42 Bit Timestamp | 10 Bit WorkID | 12 Bit Sequence ID 

```
type SnowFlake struct {
	lastTimestamp uint64  //时间戳
	sequence      uint32  
	workerId      uint32  //区别是哪一个服务器申请的
	lock          sync.Mutex
}

```
