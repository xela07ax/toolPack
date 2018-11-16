package toolPack

import (
	"log"
	"github.com/go-ini/ini"
)

func OpenINIfile(path_file string, cfg_struct interface{}) {
    //example OpenConf("file.ini", &cfg_struct)
	//Читаем параметры из схемы файла
	cfg, err := ini.Load(path_file)
	if err != nil {
		log.Fatal(err)
	}
	//t := ds.Config{}
	err = cfg.MapTo(cfg_struct)
	if err != nil {
		log.Fatal(err)
	}

}
