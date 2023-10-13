package log

/*
   全局默认提供一个Log对外句柄，可以直接使用API系列调用
*/

// 默认日志
var DefaultLog = NewLogger("", BitDefault)

// 获取Log 标记位
func Flags() int {
	return DefaultLog.Flags()
}

// 设置Log标记位
func ResetFlags(flag int) {
	DefaultLog.ResetFlags(flag)
}

// 添加flag标记
func AddFlag(flag int) {
	DefaultLog.AddFlag(flag)
}

// 设置Log 日志头前缀
func SetPrefix(prefix string) {
	DefaultLog.SetPrefix(prefix)
}

// 设置Log绑定的日志文件 ,绑定后不能在控制台输出
func SetLogFile(fileDir string, fileName string) {
	DefaultLog.SetLogFile(fileDir, fileName)
}

// 设置关闭debug
func CloseDebug() {
	DefaultLog.CloseDebug()
}

// 设置打开debug
func OpenDebug() {
	DefaultLog.OpenDebug()
}

// Trace 不做任何输出，占位
func Trace(format string, v ...interface{}) {

}

// ====> Debug <====
func Debug(format string, v ...interface{}) {
	DefaultLog.Debug(format, v...)
}

// ====> Info <====
func Info(format string, v ...interface{}) {
	DefaultLog.Info(format, v...)
}

// ====> Warn <====
func Warn(format string, v ...interface{}) {
	DefaultLog.Warn(format, v...)
}

// ====> Error <====
func Error(format string, v ...interface{}) {
	DefaultLog.Error(format, v...)
}

// ====> Fatal 需要终止程序 <====
func Fatal(format string, v ...interface{}) {
	DefaultLog.Fatal(format, v...)
}

// ====> Panic  <====
func Panicf(format string, v ...interface{}) {
	DefaultLog.Panic(format, v...)
}

// ====> Stack  <====
func Stack(v ...interface{}) {
	DefaultLog.Stack(v...)
}

func init() {
	//因为StdZinxLog对象 对所有输出方法做了一层包裹，所以在打印调用函数的时候，比正常的logger对象多一层调用
	//一般的zinxLogger对象 callDepth=2, StdZinxLog的calldDepth=3
	DefaultLog.callDepth = 3
}
