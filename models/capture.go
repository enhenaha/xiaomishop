package models

import (
  "image/color"

  "github.com/mojocn/base64Captcha"
)

// 创建store, 存储验证码
var store = base64Captcha.DefaultMemStore

// 获取验证码
func MakeCaptcha() (string, string, error) {
  var driver base64Captcha.Driver
  // 创建字符串验证码
  driverString := base64Captcha.DriverString{
    Height:          40, // 验证码图片高度
    Width:           100, // 验证码图片宽度
    NoiseCount:      0, // 在图片上随机生成的小点 (干扰点)
    ShowLineOptions: 2 | 4, // 生成 波浪线 + 正弦曲线的干扰线类型
    Length:          1, // 验证码字符长度
    Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm", // 字符来源池 (随机字符从这里选)
    BgColor: &color.RGBA{ // 背景色
      R: 3,
      G: 102,
      B: 214,
      A: 125,
    },
    Fonts: []string{},
  }

  driver = driverString.ConvertFonts()

  c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := c.Generate()
	return id, b64s, err
  /* 
    id: 验证码的唯一标识(key), 用于校验用户输入
    b64s: 验证码图片的 base64 数据, 用于给前端显示的验证码图片
    answer: 验证码的正确答案
    err: 生成验证码失败的原因
  */
}

// 验证验证码
func VerifyCaptcha(id string, VerifyValue string) bool { 
  return store.Verify(id, VerifyValue, true) // 第三个参数

  /* 
    id: c.Generate()返回的id
    VerifyValue: 用户输入的验证码
    true: 验证成功后删除store中存储的验证码
  */
}