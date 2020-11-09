package utils

import (
	"math/rand"
	"os"
	"time"
)

func GetRoundName(size int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
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