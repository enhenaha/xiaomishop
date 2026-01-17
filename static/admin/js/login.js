$(function() {
  loginApp.init()
})

var loginApp = {
  init() {
    this.getCaptcha()
    this.captchaImgChange()
  },
  getCaptcha() {
    $.get("/admin/captcha?t=" + Math.random(), (res) => {
      console.log("res: ", res)
      $("#captchaId").val(res.captchaId)
      $("#captchaImg").attr("src", res.captchaImage)
    })
  },
  captchaImgChange() {
    const that = this
    $("#captchaImg").click(() => {
      that.getCaptcha()
    })
  } 
}
/* 
  ?t= + Math.random() 用于防止出现缓存问题
*/