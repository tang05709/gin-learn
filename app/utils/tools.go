package utils

import (
	"math/rand"
	"os"
)

func GetRoundName(size int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	str := make([]rune, size) 
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	return string(str)
}

func GetSaveDir(savePath string) string {
	// os.stat返回err==nil，说明存在
	_, err := os.Stat(savePath)
	if err != nil {
		// os.IsNotExist(err)为true，说明不存在；否则不确定是否存在
		if os.IsNotExist(err){
			os.Mkdir(savePath, 0777) 
		}
	} 
	return savePath
}