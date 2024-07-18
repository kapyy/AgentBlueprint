package logger

import (
	"runtime"
	"strings"
)

const (
	maxDepth = 10
)

func Recover(data interface{}) {
	log := GetLogger().WithField("func", "Recover")
	if err := recover(); err != nil {
		log.Errorf("Panic|%s|%v", err, data)
		for i := 0; i < maxDepth; i++ {
			pc, file, line, ok := runtime.Caller(i)
			if ok {
				p := strings.Index(file, "/src/")
				if p != -1 {
					file = file[p+len("/src/"):]
				}
				log.Errorf("frame %d:[func:%s,file:%s,line:%d]",
					i, runtime.FuncForPC(pc).Name(), file, line)
			} else {
				break
			}
		}
	}
}
