package TradeLog

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	DATEFORMAT         = "2006-01-02"
	DEFAULT_BUF_AMOUNT = 10
	DEFAULT_SQL_TIME   = 300 //s
	DEFAULT_SCAN_TIME  = 300 //s
)

type LEVEL int

const (
	TRACE LEVEL = iota
	INFO
	WARN
	ERROR
	OFF
)

type FileLog struct {
	fileName string
	dirName  string
	fileTime string
	level    LEVEL

	//mu *sync.Mutex

	file    *os.File
	logChan chan string

	buf       string // log buf
	bufAmount int
	bufCount  int

	once *sync.Once

	closeChan chan int
}

func NewFileLog(dirName, fileName string, l LEVEL) *FileLog {
	f, err := newFileLog(dirName, fileName, l)
	if err != nil {
		panic(fmt.Sprintf("NewFileLog() err : %s", err))
	}
	return f

}

func newFileLog(dirName, fileName string, l LEVEL) (*FileLog, error) {
	fl := &FileLog{
		fileName: fileName,
		dirName:  dirName,
		level:    l,
		//mu:        new(sync.Mutex),
		logChan:   make(chan string, 100),
		buf:       "",
		bufAmount: DEFAULT_BUF_AMOUNT,
		bufCount:  0,
		once:      new(sync.Once),
		closeChan: make(chan int),
	}

	if !Exist(dirName) {
		os.Mkdir(dirName, 0755)
	}
	t := time.Now().Format(DATEFORMAT)
	logPath := joinFilePath(dirName, t+"_"+fileName)

	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	fl.file = f
	fl.fileTime = t
	return fl, err
}

func (f *FileLog) Start() {

	f.once.Do(func() {
		go f.logGoroutine()
	})
}

func (f *FileLog) logGoroutine() {
	T := time.NewTicker(time.Duration(DEFAULT_SQL_TIME) * time.Second)
	T2 := time.NewTicker(time.Duration(DEFAULT_SCAN_TIME) * time.Second)
	for {
		select {
		case <-T.C:
			f.l()
		case d := <-f.logChan:
			f.p(d)
			//fmt.Println("log")
		case <-T2.C:
			f.scan()
		case <-f.closeChan:
			goto Exit
		}

	}
Exit:
}

func (f *FileLog) l() {
	f.file.WriteString("//===========================================LEN:" + fmt.Sprint(len(f.logChan)))
}

func (f *FileLog) p(d string) {
	if f.bufCount == f.bufAmount {
		_, err := f.file.WriteString(f.buf)
		if err != nil {
			fmt.Println(err)
		}
		f.bufCount = 0
	} else {
		f.buf += d + "\n"
		f.bufCount++
	}

}

func (f *FileLog) scan() {
	t := time.Now().Format(DATEFORMAT)
	if t != f.fileTime {
		f.file.Close()
	}
	logPath := joinFilePath(f.dirName, t+"_"+f.fileName)
	f_, err := os.OpenFile(logPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(fmt.Sprintf("scan() error: %s", err))
	}
	f.file = f_
	f.fileTime = t
}
