package toolPack

import (
	"os"
	"bufio"
)

//tp.WriteFile(file,[]byte(result))
func  WriteFile(filname string, content []byte) (error) {

	var data, _= os.Create(filname)
	writer := bufio.NewWriter(data)
	defer data.Close()
	//detailjson, _ := json.Marshal(a.Forms)
	writer.Write(content) // запись строки
	writer.Flush() // сбрасываем данные из буфера в файл
	return nil
}

