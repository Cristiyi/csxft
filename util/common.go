package util

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

type PointRange struct {
	MaxLat float64
	MinLat float64
	MaxLng float64
	MinLng float64
}

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// 去除字符串尾部的所有空格
func DeleteTailBlank(str string) string {
	spaceNum := 0
	for i := len(str)-1; i >= 0; i-- {
		if str[i] == ' ' {
			spaceNum++
		} else {
			break
		}
	}
	return str[:len(str)-spaceNum]
}

//字符串转时间
func StrToTime(dateStr string) time.Time {
	tm, _ := time.Parse("2006/01/02", dateStr)
	return tm
}

//字符串转时间
func StrToTime1(dateStr string) time.Time {
	tm, _ := time.Parse(`"`+time.RFC3339+`"`, dateStr)
	return tm
}

//interface转换为string
func InConvertString(inter interface{}) string {
	b, ok := inter.(string) // 肯定转换失败的，如果是string，则 b 为空
	if ok{
		return b
	} else {
		return ""
	}
}

//interface转换为int
func InConvertInt(inter interface{}) int {
	b, ok := inter.(int) // 肯定转换失败的，如果是string，则 b 为空
	fmt.Println(ok)
	if ok{
		return b
	} else {
		return 0
	}
}

//interface转换为float64
func InConvertFloat64(inter interface{}) float64 {
	b, ok := inter.(float64) // 肯定转换失败的，如果是string，则 b 为空
	if ok{
		return b
	} else {
		return 0
	}
}

//interface转换为time
func InConvertTime(inter interface{}) time.Time {
	b, ok := inter.(time.Time) // 肯定转换失败的，如果是string，则 b 为空
	if ok{
		return b
	} else {
		return time.Time{}
	}
}

//float转string
func Float2String(des float64, byte int) string {
	return strconv.FormatFloat(des, 'e', -1, byte)
}

//string转float64
func String2Float64(str string) float64 {
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	} else {
		return v
	}
}

//string转int
func String2Int(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		return 0
	} else {
		return v
	}
}

//时间转时间戳
func timeToUnix(targetTime time.Time) int64 {
	return targetTime.Unix()
}

//获取当天0点时间戳
func GetTodayUnix() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	return t.Unix()
}


//身份证打码
func HideIdCard(idCard string) string {
	first := idCard[0:6]
	target := idCard[6:14]
	targetRune := []rune(target)
	last := idCard[14:len(idCard)]
	for i, _ := range targetRune {
		targetRune[i] = '*'
	}
	target = string(targetRune)
	return first + target + last
}

//根据经纬度计算距离
func GetDistancePointRange(lat, lng, distant float64) PointRange {
	var calRange float64 = 180 / math.Pi * distant / 6372.797
	var lngR float64 = calRange / math.Cos(lat * math.Pi / 180)
	pointRange := new(PointRange)
	pointRange.maxLat = lat + calRange //最大纬度
	pointRange.minLat= lat - calRange //最小纬度
	pointRange.maxLng = lng + lngR//最大经度
	pointRange.minLng = lng - lngR//最小经度
	return *pointRange
}

