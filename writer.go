package toolPack

import (
"os"
"log"
	"io"
	"sync"
)

const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)


type WriterFile struct {
	filePath string
	LogWriter *os.File
	err error
	buf    []byte
	out    io.Writer
	mu     sync.Mutex
}


func (a *WriterFile) SetOutput(w io.Writer) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.out = w
}

func (a *WriterFile) Println(text string) {
	a.LogWriter, a.err = os.OpenFile(a.filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if a.err != nil {
		log.Fatal(a.err)
	}
	defer a.LogWriter.Close()
	a.mu.Lock()
	defer a.mu.Unlock()
	buf := []byte(text+"\n")
	_, err := a.LogWriter.Write(buf)
	if err != nil {
		log.Fatal(a.err)
	}
}

func (a *WriterFile) SetPathNewFile(path string)  {
		status, err := CheckExistsFile(path)
     	if err != nil {
			log.Fatal(err)
	    } else if status {
			err := DeleteFile(path)
			if err != nil {
				log.Fatal(err)
			}
		}

		a.filePath = path

}


