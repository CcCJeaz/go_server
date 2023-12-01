package xlog

/*
TODO:
可参考zap的设计
1. 使用对象池, 减少gc压力
2. 异步化
*/

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
	FATAL
)
var logger *xLog

func Init(level int, output string) {
	logger = &xLog{
		level: level,
		output: output,
	}
}

type xLog struct {
	level int
	output string
}


func Debug() {

}

func Info() {

}

func Warning() {

}

func Error() {

}

func Fatal() {

}