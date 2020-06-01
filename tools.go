package tools

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GetEnvDefault(key string, defVal interface{}) interface{} {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

func StringToInt(string string) int {
	//if s, err := strconv.Atoi(string); err == nil {
	//	fmt.Printf("%T, %v", s, s)
	//}
	s, _ := strconv.Atoi(string)
	return s
}

func GetMysqlLimitOffset(page string, size string ) (int, int) {
	intPage := StringToInt(page)
	limit := StringToInt(size)
	offset := (intPage - 1) * limit
	return limit, offset
}


// 生成32位MD5
func MD5(text string) string{
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// return len=8  salt
func GetRandomSalt() string {
	return GetRandomString(8)
}

//生成随机字符串
func GetRandomString(lenght int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	bytesLen := len(bytes)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lenght; i++ {
		result = append(result, bytes[r.Intn(bytesLen)])
	}
	return string(result)
}

// bool 转 string
func BoolToString(arg bool) string {
	return  strconv.FormatBool(arg)
}

// string to bool
func StringToBool(string string) bool {
	b, err := strconv.ParseBool(string)
	if err != nil {
		return false
	}
	return b
}

// string to float64

func StringToFloat64(arg string) float64 {
	float, err := strconv.ParseFloat(arg,64)
	if err != nil {
		panic(err)
	}
	return float
}
func StringToInt64(arg string) int64 {
	int64, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		panic(err)
	}
	return int64
}

// slice 去重
func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{}  // 存放不重复主键
	for _, e := range slc{
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l{  // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}