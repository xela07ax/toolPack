package toolPack

import (
	"fmt"
	"os"
	"path/filepath"
	"io/ioutil"
)

func ThisDir()string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}
//find, err := CheckExistsFile(path)
//if err != nil {
//log.Fatal(err)
//} else if !find {
//log.Fatal("Not Found:", path)
//}
func CheckExistsFile(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func BinDir() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "" , err
	}
	return filepath.Dir(ex), nil
}

func FindFile(filename string) (Path string, err error) {
	// Цель этого модуля искать файл
	// C начала модуль просто проверяет путь
	// Потом модуль посставляет к пути полную директорию до бин файла
	//fmt.Printf("Find %s: %s\n",filename, ThisDir())
	flagF, err := CheckExistsFile(filename)
	if err != nil {
		return "" , err
	}

	if flagF {
		return filename, nil

	} else {
		ex, err := os.Executable()
		if err != nil {
			return "" , err
		}
		exPath := filepath.Dir(ex)
		fmt.Printf("Find dir in files: %s\n", exPath)
		flagF, err = CheckExistsFile(fmt.Sprintf("%s/%s",exPath,filename))
		if err != nil {
			return "" , err
		}
		if flagF {
			return fmt.Sprintf("%s/%s",exPath,filename), nil
		}
	}
	return "", fmt.Errorf("Нигде нет этого файла: %s", filename)
}

func ReadfileWorkDirIntel(filename string) (text string, err error) {
	var data []byte
	path, err := FindFile(filename)
	if err != nil {
		return "" , err
	}
	data, err = ReadFile(path)
	if err != nil {
		return "" , err
	}
	return string(data), nil
}

func ReadFile (filename string) (data []byte, err error) {
	///var  fileInfo os.FileInfo

	//fileInfo, err = os.Stat(filename)
	//if err != nil {
		//return nil , err
	//}

	//fmt.Printf("File: %s | Permissions: %s\n",fileInfo.Name(),fileInfo.Mode())
	//fmt.Println("File list name:", fileInfo.Name())
	//fmt.Println("    Size in bytes:", fileInfo.Size())
	//fmt.Println("    Permissions:", fileInfo.Mode())
	//fmt.Println("    Last modified:", fileInfo.ModTime())
	//fmt.Println("\n")
	////fmt.Println("Is Directory: ", fileInfo.IsDir())
	//fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	//fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		return nil , err
	}
	data, err = ioutil.ReadAll(file)
	if err != nil {
		return nil , err
	}
	return data, nil
}
