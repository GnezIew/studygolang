package main

//声明日志写入器(接口)
type LogWriter interface {
	Write(data interface{}) error
}

//日志器
type Logger struct {
	//这个日志器用到的日志写入器
	writeList []LogWriter
}

//注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter){
	l.writeList = append(l.writeList,writer)
}

//将一个data 类型的数据写入数据
func (l *Logger) Log(data interface{}){
	//遍历所有注册的写入器
	for _,writer :=range l.writeList{
		writer.Write(data)
	}
}

//创建日志器实例
func NewLogger() *Logger{
	return &Logger{}
}

