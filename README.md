# UUID--雪花算法
格式： 42 Bit Timestamp | 10 Bit WorkID | 12 Bit Sequence ID  

1.42位，用来记录时间戳（毫秒）。  

2.10位，用来记录工作机器id。  

3.可以部署在2^10=1024个节点。  

4.12位，用来记录同毫秒内产生的不同id。  

```
type SnowFlake struct {
	lastTimestamp uint64  //时间戳
	sequence      uint32  
	workerId      uint32  //区别是哪一个服务器申请的
	lock          sync.Mutex
}

```
  
  
SnowFlake可以保证：

* 所有生成的id按时间趋势递增
* 整个分布式系统内不会产生重复id（因为有workerId来做区分）

