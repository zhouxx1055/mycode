package main

import (
	"log"
	"time"
)

type DevTypeStr struct {
	DevType int
	Info    map[string]string
}

var ch_syncCheckFail chan int
var ch_syncCheckSuccess chan DevTypeStr
var ch_execCheckResult chan string

var devTypeList []int

func main() {

	defer func() {
		if ch_syncCheckFail != nil {
			close(ch_syncCheckFail)
		}
		if ch_syncCheckSuccess != nil {
			close(ch_syncCheckSuccess)
		}
		if ch_execCheckResult != nil {
			close(ch_execCheckResult)
		}
	}()
	ch_syncCheckFail = make(chan int, 100)
	ch_syncCheckSuccess = make(chan DevTypeStr, 100)
	ch_execCheckResult = make(chan string, 100)

	// 读取devtype配置文件
	devTypeList = []int{12, 14, 33}

	// 登录、创建家庭
	// 添加设备，同时添加多个设备
	time.Sleep(1000 * time.Millisecond)

	// 启10个线程 同时检测功能是完整，检查成功的，放到ch_syncCheckSuccess；失败的放到ch_syncCheckFail
	time.Sleep(1000 * time.Millisecond)

	// ch_syncCheckFail 接收完 所有检查失败的devtype，统一告警：fmt.println(dsts告警信息)

}

//只能向chan里写数据
func sendX(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
		log.Println("send===>i:%d, len:%d", i, len(c))
	}
}

//只能取channel中的数据
func recvX(c <-chan int) {
	for i := range c { //fatal error: all goroutines are asleep - deadlock!
		log.Println("recv===>i:%d, len:%d", i, len(c))
		time.Sleep(1 * time.Second)
	}
}
