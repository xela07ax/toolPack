package toolPack

import "os"

func CheckMkdir(workFolder string)error  {

	if _, err := os.Stat(workFolder); os.IsNotExist(err) {
		err = os.Mkdir(workFolder, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

