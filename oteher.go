package toolPack

import (
	"log"
	"os"
	"fmt"
	"time"
	"math"
)

func Fck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func FckText(text string,err error) {
	if err != nil {
		log.Fatal(text,err)
	}
}
func DeleteFile(path string)error {
	// delete file
	var err = os.Remove(path)
	 if err != nil {
		return  err
	}
	return nil
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func Getime()string  {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}