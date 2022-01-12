package TradeLog

import (
	"fmt"
	"runtime"
	"time"
)

//======================================================================================================================
func (f *FileLog) Trace(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if f.level <= TRACE {
		f.logChan <- fmt.Sprintf("[%v]"+time.Now().Format("2006-01-02 15:04:05")+"[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[TRACE] "+format, v...)
	}
}

func (f *FileLog) T(format string, v ...interface{}) {
	f.Trace(format, v...)
}

func (f *FileLog) Info(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if f.level <= INFO {
		f.logChan <- fmt.Sprintf("[%v]"+time.Now().Format("2006-01-02 15:04:05")+"[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[INFO] "+format, v...)
	}
}

func (f *FileLog) I(format string, v ...interface{}) {
	f.Info(format, v...)
}

func (f *FileLog) Warn(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if f.level <= WARN {
		f.logChan <- fmt.Sprintf("[%v]"+time.Now().Format("2006-01-02 15:04:05")+"[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[WARN] "+format, v...)
	}
}

func (f *FileLog) W(format string, v ...interface{}) {
	f.Warn(format, v...)
}

func (f *FileLog) Error(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if f.level <= ERROR {
		f.logChan <- fmt.Sprintf("[%v]"+time.Now().Format("2006-01-02 15:04:05")+"[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[ERROR] "+format, v...)
	}
}

func (f *FileLog) E(format string, v ...interface{}) {
	f.Error(format, v...)
}

//======================================================================================================================
func (f *FileLog) Close() {
	close(f.closeChan)
	if f.buf != "" {
		f.file.WriteString(f.buf)
	}
	if f.file != nil {
		f.file.Close()
	}

}
