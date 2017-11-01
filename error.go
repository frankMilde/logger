// logger provide some wrappers around go's logging tool
package logger

import "log"

func Err(e error) {
	if e != nil {
		info := getCallerInfo()
		log.Println(info + e.Error())
	}
}

func Panic(e error) {
	if e != nil {
		info := getCallerInfo()
		log.Println(info + e.Error())
		panic(e)
	}
}

func Fatal(e error) {
	if e != nil {
		info := getCallerInfo()
		log.Fatal(info + e.Error())
	}
}

func Error(msg string, args ...interface{}) {
	info := getCallerInfo()
	log.Printf(info+msg, args...)
}

func Info(msg string, args ...interface{}) {
	info := getCallerInfo()
	log.Printf(info+msg, args...)
}
