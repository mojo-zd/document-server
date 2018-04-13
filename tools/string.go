package tools

import (
	"math/rand"
	"strings"
	"time"

	uuid2 "github.com/pborman/uuid"
)

var (
	RANGE    = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	RANGE_32 = 32
)

type Remove struct {
	Split   string
	Reverse bool
}

func RandomString(length int) string {
	bytes := []byte(RANGE)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func RemoveSuffix(str string, removes []Remove) string {
	for _, remove := range removes {
		index := strings.LastIndex(str, remove.Split)
		if remove.Reverse {
			str = str[0:index]
		} else {
			str = str[index+1:]
		}
	}
	return str
}

func NewUUID() (uuid string) {
	uuid = uuid2.NewUUID().String()
	return
}

func Suffix(str, symbol string) (s string) {
	if !strings.Contains(str, symbol) {
		s = str
		return
	}
	s = str[strings.Index(str, symbol):]
	return
}

func Substring(str, symbol string) (s string) {
	if !strings.Contains(str, symbol) {
		s = str
		return
	}
	s = str[:strings.Index(str, symbol)]
	return
}
