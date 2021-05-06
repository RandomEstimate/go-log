package TradeLog

import (
	"fmt"
	"runtime"
)

//======================================================================================================================
func (f *fileLog) Trace(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if f.level <= TRACE {
		f.logChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[TRACE] "+format, v...)
	}
}

func (f *fileLog) T(format string, v ...interface{}) {
	f.Trace(format, v...)
}

func (f *fileLog) Info(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if f.level <= INFO {
		f.logChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[INFO] "+format, v...)
	}
}

func (f *fileLog) I(format string, v ...interface{}) {
	f.Info(format, v...)
}

func (f *fileLog) Warn(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if f.level <= WARN {
		f.logChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[WARN] "+format, v...)
	}
}

func (f *fileLog) W(format string, v ...interface{}) {
	f.Warn(format, v...)
}

func (f *fileLog) Error(format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2) //calldepth=3
	if f.level <= ERROR {
		f.logChan <- fmt.Sprintf("[%v:%v]", shortFileName(file), line) + fmt.Sprintf("[ERROR] "+format, v...)
	}
}

func (f *fileLog) E(format string, v ...interface{}) {
	f.Error(format, v...)
}

//======================================================================================================================
func (f *fileLog) Close() {
	close(f.closeChan)
	if f.buf != "" {
		f.file.WriteString(f.buf)
	}
	if f.file != nil {
		f.file.Close()
	}

}
