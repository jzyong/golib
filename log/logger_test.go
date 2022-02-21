package log

import (
	"testing"
	"time"
)

func TestLogPrint(t *testing.T) {

	//测试 默认debug输出
	Debug("debug content1")
	Debug("debug content2")

	Debug(" debug a = %d\n", 10)

	//设置log标记位，加上长文件名称 和 微秒 标记
	ResetFlags(BitDate | BitLongFile | BitLevel)
	Info("info content")

	//设置日志前缀，主要标记当前日志模块
	SetPrefix("MODULE")
	Error("golib error content")

	//添加标记位
	AddFlag(BitShortFile | BitTime)
	Stack("golib Stack! ")

	//设置日志写入文件
	SetLogFile("./log", "golib")
	Debug("===> golib debug content ~~666")
	Debug("===> golib debug content ~~888")
	Error("===> golib Error!!!! ~~~555~~~")

	//关闭debug调试
	CloseDebug()
	Debug("===> 我不应该出现~！")
	Debug("===> 我不应该出现~！")
	Error("===> golib Error  after debug close !!!!")

}

//测试跨天创建文件
func TestLogCreateFile(t *testing.T) {

	//设置日志写入文件
	SetLogFile("./log", "golib")
	go func() {
		for true {
			now := time.Now()
			Info("current second :%v", now.Second())

			time.Sleep(1 * time.Second)
		}
	}()
	select {}
}
