package log

/*
	日志类全部方法 及 API
*/

import (
	"bytes"
	"fmt"
	"github.com/jzyong/golib/util"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	LOG_MAX_BUF = 1024 * 1024
)

//日志头部信息标记位，采用bitmap方式，用户可以选择头部需要哪些标记位被打印
const (
	BitDate         = 1 << iota                            //日期标记位  2019/01/23
	BitTime                                                //时间标记位  01:23:12
	BitMicroSeconds                                        //微秒级标记位 01:23:12.111222
	BitLongFile                                            // 完整文件名称 /home/go/src/zinx/server.go
	BitShortFile                                           // 最后文件名   server.go
	BitLevel                                               // 当前日志级别： 0(Debug), 1(Info), 2(Warn), 3(Error), 4(Panic), 5(Fatal)
	BitStdFlag      = BitDate | BitTime                    //标准头部日志格式
	BitDefault      = BitLevel | BitShortFile | BitStdFlag //默认日志头部格式
)

//日志级别
const (
	LogDebug = iota
	LogInfo
	LogWarn
	LogError
	LogPanic
	LogFatal
)

//日志级别对应的显示字符串
var levels = []string{
	"[DEBUG]",
	"[INFO]",
	"[WARN]",
	"[ERROR]",
	"[PANIC]",
	"[FATAL]",
}

// 日志
type Logger struct {
	//确保多协程读写文件，防止文件内容混乱，做到协程安全
	mu sync.Mutex
	//每行log日志的前缀字符串,拥有日志标记
	prefix string
	//日志标记位
	flag int
	//控制台输出的文件描述符
	consoleOutWriter io.Writer
	//输出的缓冲区
	buf bytes.Buffer
	//当前日志绑定的输出文件
	file *os.File
	// 文件创建时间
	fileCreateTime time.Time
	//文件名称
	fileName string
	//文件路径
	filePath string
	//是否打印调试debug信息
	debugClose bool
	//获取日志文件名和代码上述的runtime.Call 的函数调用层数
	callDepth int
}

/*
	创建一个日志
	consoleOutWriter: 标准输出的文件io
	prefix: 日志的前缀
	flag: 当前日志头部信息的标记位
*/
func NewLogger(prefix string, flag int) *Logger {

	//默认 debug打开， calledDepth深度为2,ZinxLogger对象调用日志打印方法最多调用两层到达output函数
	log := &Logger{consoleOutWriter: os.Stderr, prefix: prefix, flag: flag, file: nil, debugClose: false, callDepth: 2}
	//设置log对象 回收资源 析构方法(不设置也可以，go的Gc会自动回收，强迫症没办法)
	runtime.SetFinalizer(log, CleanLog)
	return log
}

/*
   回收日志处理
*/
func CleanLog(log *Logger) {
	log.closeFile()
}

/*
   制作当条日志数据的 格式头信息
*/
func (log *Logger) formatHeader(buf *bytes.Buffer, t time.Time, file string, line int, level int) {
	//如果当前前缀字符串不为空，那么需要先写前缀
	if log.prefix != "" {
		buf.WriteByte('<')
		buf.WriteString(log.prefix)
		buf.WriteByte('>')
	}

	//已经设置了时间相关的标识位,那么需要加时间信息在日志头部
	if log.flag&(BitDate|BitTime|BitMicroSeconds) != 0 {
		//日期位被标记
		if log.flag&BitDate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			buf.WriteByte('/') // "2019/"
			itoa(buf, int(month), 2)
			buf.WriteByte('/') // "2019/04/"
			itoa(buf, day, 2)
			buf.WriteByte(' ') // "2019/04/11 "
		}

		//时钟位被标记
		if log.flag&(BitTime|BitMicroSeconds) != 0 {
			hour, min, sec := t.Clock()
			itoa(buf, hour, 2)
			buf.WriteByte(':') // "11:"
			itoa(buf, min, 2)
			buf.WriteByte(':') // "11:15:"
			itoa(buf, sec, 2)  // "11:15:33"
			//微秒被标记
			if log.flag&BitMicroSeconds != 0 {
				buf.WriteByte('.')
				itoa(buf, t.Nanosecond()/1e3, 6) // "11:15:33.123123
			}
			buf.WriteByte(' ')
		}

		// 日志级别位被标记
		if log.flag&BitLevel != 0 {
			buf.WriteString(levels[level])
		}

		//日志当前代码调用文件名名称位被标记
		if log.flag&(BitShortFile|BitLongFile) != 0 {
			//短文件名称
			if log.flag&BitShortFile != 0 {
				short := file
				for i := len(file) - 1; i > 0; i-- {
					if file[i] == '/' {
						//找到最后一个'/'之后的文件名称  如:/home/go/src/zinx.go 得到 "zinx.go"
						short = file[i+1:]
						break
					}
				}
				file = short
			}
			buf.WriteString(file)
			buf.WriteString("(")
			itoa(buf, line, -1) //行数
			buf.WriteString(")")
			buf.WriteString("--> ")
		}
	}
}

/*
   输出日志文件,原方法
*/
func (log *Logger) OutPut(level int, s string) error {

	now := time.Now() // 得到当前时间
	var file string   //当前调用日志接口的文件名称
	var line int      //当前代码行数
	log.mu.Lock()
	defer log.mu.Unlock()

	if log.flag&(BitShortFile|BitLongFile) != 0 {
		log.mu.Unlock()
		var ok bool
		//得到当前调用者的文件名称和执行到的代码行数
		_, file, line, ok = runtime.Caller(log.callDepth)
		if !ok {
			file = "unknown-file"
			line = 0
		}
		log.mu.Lock()
	}

	//清零buf
	log.buf.Reset()
	//写日志头
	log.formatHeader(&log.buf, now, file, line, level)
	//写日志内容
	log.buf.WriteString(s)
	//补充回车
	if len(s) > 0 && s[len(s)-1] != '\n' {
		log.buf.WriteByte('\n')
	}

	//将填充好的buf 写到IO输出上
	_, err := log.consoleOutWriter.Write(log.buf.Bytes())
	if log.file != nil {
		//检测是否创建新的日志
		if !util.SameDay(now, log.fileCreateTime) {
			SetLogFile(log.filePath, log.fileName)
		}
		_, err = log.file.Write(log.buf.Bytes())
	}
	return err
}

// ====> Debug <====
func (log *Logger) Debug(format string, v ...interface{}) {
	if log.debugClose == true {
		return
	}
	_ = log.OutPut(LogDebug, fmt.Sprintf(format, v...))
}

// ====> Info <====
func (log *Logger) Info(format string, v ...interface{}) {
	_ = log.OutPut(LogInfo, fmt.Sprintf(format, v...))
}

// ====> Warn <====
func (log *Logger) Warn(format string, v ...interface{}) {
	_ = log.OutPut(LogWarn, fmt.Sprintf(format, v...))
}

// ====> Error <====
func (log *Logger) Error(format string, v ...interface{}) {
	_ = log.OutPut(LogError, fmt.Sprintf(format, v...))
}

// ====> Fatal 需要终止程序 <====
func (log *Logger) Fatal(format string, v ...interface{}) {
	_ = log.OutPut(LogFatal, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// ====> Panic  <====
func (log *Logger) Panic(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	_ = log.OutPut(LogPanic, s)
	panic(s)
}

// ====> Stack  <====
func (log *Logger) Stack(v ...interface{}) {
	s := fmt.Sprint(v...)
	s += "\n"
	buf := make([]byte, LOG_MAX_BUF)
	n := runtime.Stack(buf, true) //得到当前堆栈信息
	s += string(buf[:n])
	s += "\n"
	_ = log.OutPut(LogError, s)
}

//获取当前日志bitmap标记
func (log *Logger) Flags() int {
	log.mu.Lock()
	defer log.mu.Unlock()
	return log.flag
}

//重新设置日志Flags bitMap 标记位
func (log *Logger) ResetFlags(flag int) {
	log.mu.Lock()
	defer log.mu.Unlock()
	log.flag = flag
}

//添加flag标记
func (log *Logger) AddFlag(flag int) {
	log.mu.Lock()
	defer log.mu.Unlock()
	log.flag |= flag
}

//设置日志的 用户自定义前缀字符串
func (log *Logger) SetPrefix(prefix string) {
	log.mu.Lock()
	defer log.mu.Unlock()
	log.prefix = prefix
}

//设置日志文件输出
func (log *Logger) SetLogFile(filePath string, fileName string) {
	var file *os.File
	log.fileName = fileName
	log.filePath = filePath
	log.fileCreateTime = time.Now()
	//创建日志文件夹
	_ = mkdirLog(filePath)

	fullPath := filePath + "/" + fileName + "_" + time.Now().Format("2006-01-02") + ".log"
	if log.checkFileExist(fullPath) {
		//文件存在，打开
		file, _ = os.OpenFile(fullPath, os.O_APPEND|os.O_RDWR, 0644)
	} else {
		//文件不存在，创建
		file, _ = os.OpenFile(fullPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	}

	log.mu.Lock()
	defer log.mu.Unlock()

	//关闭之前绑定的文件
	log.closeFile()
	log.file = file
	//log.consoleOutWriter = file
}

//关闭日志绑定的文件
func (log *Logger) closeFile() {
	if log.file != nil {
		_ = log.file.Close()
		log.file = nil
	}
}

func (log *Logger) CloseDebug() {
	log.debugClose = true
}

func (log *Logger) OpenDebug() {
	log.debugClose = false
}

// ================== 以下是一些工具方法 ==========

//判断日志文件是否存在
func (log *Logger) checkFileExist(filename string) bool {
	exist := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//创建日志目录
func mkdirLog(dir string) (e error) {
	_, er := os.Stat(dir)
	b := er == nil || os.IsExist(er)
	if !b {
		if err := os.MkdirAll(dir, 0775); err != nil {
			if os.IsPermission(err) {
				e = err
			}
		}
	}
	return
}

//将一个整形转换成一个固定长度的字符串，字符串宽度应该是大于0的
//要确保buffer是有容量空间的
func itoa(buf *bytes.Buffer, i int, wid int) {
	var u uint = uint(i)
	if u == 0 && wid <= 1 {
		buf.WriteByte('0')
		return
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}

	// avoid slicing b to avoid an allocation.
	for bp < len(b) {
		buf.WriteByte(b[bp])
		bp++
	}
}
