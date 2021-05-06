package TradeLog

import (
	"sync"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	trace := NewFileLog("", "TRACE.log", TRACE)
	defer trace.Close()
	trace.Start()

	info := NewFileLog("", "INFO.log", INFO)
	defer info.Close()
	info.Start()

	warn := NewFileLog("", "WARN.log", WARN)
	defer warn.Close()
	warn.Start()

	error := NewFileLog("", "ERROR.log", ERROR)
	defer error.Close()
	error.Start()

	w := sync.WaitGroup{}

	w.Add(1)
	go func() {
		for i := 1; i <= 1001; i++ {
			trace.T("This is the No[%v] TRACE log using fileLogger.", i)
		}
		w.Done()
	}()

	w.Add(1)
	go func() {
		for i := 1; i <= 1001; i++ {
			info.I("This is the No[%v] INFO log using fileLogger.", i)
		}
		w.Done()
	}()

	w.Add(1)
	go func() {
		for i := 1; i <= 1001; i++ {
			warn.W("This is the No[%v] WARN log using fileLogger.", i)
		}
		w.Done()
	}()

	w.Add(1)
	go func() {
		for i := 1; i <= 1001; i++ {
			error.E("This is the No[%v] ERROR log using fileLogger.", i)
		}
		w.Done()
	}()

	w.Wait()
	time.Sleep(time.Second * 10)


}
