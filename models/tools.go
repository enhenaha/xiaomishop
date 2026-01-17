package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

// 时间戳转为日期
func UnixToTime(timestamp int) string {
  t := time.Unix(int64(timestamp), 0)
  return t.Format("2006-01-02 15:04:05")
}

// 日期转换成时间戳
func TimeToUnix(str string) int64 {
  template := "2006-01-02 15:04:05"
  t, err := time.ParseInLocation(template, str, time.Local)
  if err != nil {
    return 0
  }
  return t.Unix()
}

// 获取当前时间戳
func GetUnix() int64 {
  return time.Now().Unix()
}

// 获取当前时间
func GetDate() string {
  template := "2006-01-02 15:04:05"
  return time.Now().Format(template)
}

// 获取年月日
func GetDay() string {
  template := "20060102"
  return time.Now().Format(template)
}

// md5加密
func Md5(str string) string {
  // 1. 创建一个 MD5 哈希计算器
  h := md5.New()
  // 2. 把字符串内容写入 MD5 计算器
	io.WriteString(h, str)
  // 3. 将 MD5 的二进制编码为十六进制字符串返回
	return fmt.Sprintf("%x", h.Sum(nil))
}

// string 转 int
func Int(str string) (int, error) {
  n, err := strconv.Atoi(str)
  return n, err
}

// int 转 string
func String(n int) string {
  str := strconv.Itoa(n)
  return str
}