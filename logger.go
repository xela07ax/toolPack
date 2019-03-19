package toolPack

import (
	"os"
	"log"
	"fmt"
	"strings"
)




type Logger struct {
	filePath string //Имя дирректории для записи логов
	LogWriter *os.File //Дескриптор файла
	err error 
}


func (a *Logger) openWriter()  {
	a.LogWriter, a.err = os.OpenFile(a.filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if a.err != nil {
		log.Fatal(a.err)
	}
}


func (a *Logger) LogPrintln(funct string, text string) {
	a.openWriter()
	defer a.LogWriter.Close() //defer to close when you're done with it, not because you think it's idiomatic!
	if text != "" {
		message := fmt.Sprintf("%s  |  %s",Getime(),text)
		fmt.Println(message)

		res:= strings.Split(text, ":")
	if len(res) == 1 {
		text = fmt.Sprintf("TEXT: %s",text)
	}
		message = fmt.Sprintf("%s  |  FUNC: %s  |  %s",Getime(),funct,text)
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
		//fmt.Fprintln(a.LogWriter,message)
		log.SetOutput(a.LogWriter)
		log.Println(message)
	}
}

func (a *Logger) LogPrintlnFile(funct string, text string) {
	a.openWriter()
	defer a.LogWriter.Close() //defer to close when you're done with it, not because you think it's idiomatic!
	if text != "" {
		res:= strings.Split(text, ":")
		if len(res) == 1 {
			text = fmt.Sprintf("TEXT: %s",text)
		}
		message := fmt.Sprintf("%s  |  FUNC: %s  |  %s",Getime(),funct,text)
		//fmt.Println(message)
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
		log.SetOutput(a.LogWriter)
		log.Println(message)
	}
}

func (a *Logger) LogFatal(text string, err error) {

	a.openWriter()
	defer a.LogWriter.Close() //defer to close when you're done with it, not because you think it's idiomatic!
	if a.LogWriter != nil {

		message := fmt.Errorf("%s  |  FUNC: @ERROR START@  |  STEP: %s  |  TEXT: %s  |  @ERROR END@\n\n",Getime(),text,err)
		fmt.Fprintln(os.Stderr, message)
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
		log.SetOutput(a.LogWriter)
		log.Fatal(message)
	}
}
func (a *Logger) SetFilePath(path string)  {
	if path != "" {
		a.filePath = path
		/*
var std = New(os.Stderr, "", LstdFlags)
&Logger{out: out, prefix: prefix, flag: flag}
*/
	}
}
