package utils

import "flag"

const (
	Log_debug = iota
	Log_info
	Log_warning
	Log_error
	//Log_fatal	不支持不打印致命输出。既然致命我们一定要尸检然后查看病因啊
	//Log_off
)

//LogLevel 值越小越唠叨, 废话越多，值越大打印的越少，见log_开头的常量;默认是warning级别
var LogLevel int

func init() {
	flag.IntVar(&LogLevel, "ll", Log_warning, "log level,0=debug, 1=info, 2=warning, 3=error, 4=fatal, 5=off")

}

//return LogLevel <= l
func CanLogWithLevel(l int) bool {
	return LogLevel <= l

}

func CanLogErr() bool {
	return LogLevel <= Log_error

}

func CanLogInfo() bool {
	return LogLevel <= Log_info

}
