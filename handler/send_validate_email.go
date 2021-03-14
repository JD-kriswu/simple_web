package handler

import (
	"fmt"
	"net/http"
	logger "onbio/logger"
	"onbio/model"
	"onbio/redis"
	"onbio/utils/email_html"
	"onbio/utils/errcode"
	"onbio/utils/ratelimiter"

	"onbio/utils/mailsender"

	"github.com/gin-gonic/gin"
	redigo "github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	uuid "github.com/satori/go.uuid"
)

const (
	USER_SEND_VALID_EMAIL_LIMIT = "user_send_valid_email_limit:%s"
	USER_EMAIL_VALID_URL        = "http://onb.io/api/user/validate_email?code=%s"
	USER_VALID_CONTENT_PRE      = "user_valid:%s"
)

type EmailParam struct {
	UserEmail string `json:"user_email" binding:"required"`
}

func HandleSendValidateEmailRequest(c *gin.Context) {

	var params EmailParam
	err := c.Bind(&params)
	if err != nil {
		logger.Error("params err ")
		c.Error(errcode.ErrParam)
		return
	}

	//判断是哪个用户的email 。限制次数，每天只能发几次这样子
	err, user := model.GetUserInfo(params.UserEmail, "", 0)

	if err != nil {
		logger.Info("get user info by email failed ", zap.String("user email", params.UserEmail))
		c.Error(errcode.ErrEmail)
		return
	}

	//检查是否已经通过邮箱认证
	if user.IsConfirmed == 1 {
		logger.Info("you are confirmed before ", zap.String("email", params.UserEmail))
		c.Error(errcode.ErrEmailAlReadyValid)
		return
	}
	//检查次数是否超过限制
	key := fmt.Sprintf(USER_SEND_VALID_EMAIL_LIMIT, user.UserName)

	err = ratelimiter.IsRateLimiterExisted(key)

	if err != nil {
		logger.Info("IsRateLimiterExisted ,new one", zap.String("key", key))

		err = ratelimiter.NewRateLimiter(key, 86400, 5) //1天5次

		if err != nil {
			logger.Error("new ratelimit failed ", zap.Error(err))
			c.Error(errcode.ErrInternal)
			return
		}
	}

	//检查是否超过限制
	isExceedLimit := ratelimiter.RateLimitAllow(key)

	if !isExceedLimit {
		logger.Warn("exceed limit ", zap.String("email", params.UserEmail))
		c.Error(errcode.ErrEmailValidLimit)
		return
	}

	//到这里，就可以发邮件了
	err, code := GenValidCode(user.ID, params.UserEmail)
	if err != nil {
		logger.Error("gen valid code failed ,", zap.Error(err))
		c.Error(errcode.ErrInternal)
		return
	}

	sendUrl := fmt.Sprintf(USER_EMAIL_VALID_URL, code)

	//没有接口，先打个log
	logger.Info("valid url ", zap.String("url", sendUrl))

	lan, _ := c.Cookie("i18n_redirected")

	if lan == "" {
		lan = "en-US"
	}
	//高级一点，这里用html
	emailBody, err := email_html.GenerateHtml(user.UserName, sendUrl, lan, 1)

	if err != nil {
		logger.Error("generate email body failed ", zap.String("user", user.UserName))
		c.Error(errcode.ErrInternal)
		return
	}

	var ms mailsender.MailSender = &mailsender.Mail{
		Sender:    "OnBio<welcome@onb.io>", // 可以自定义
		Recipient: params.UserEmail,        // 如果处于Sandbox只能发送已验证过的邮箱
		Subject:   GetTitleWithMultiLan(lan),
		HTMLBody:  emailBody,
		TextBody:  emailBody, // 不支持HTML的话会返回这个
		CharSet:   "UTF-8",   // 固定字符码
	}
	sendRet := ms.SendMail()
	if !sendRet {
		logger.Error("send mail failed ")
		c.Error(errcode.ErrInternal)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": gin.H{},
	})

}

type ValidContent struct {
	UserID    uint64 `json:"user_id"`
	UserEmail string `json:"user_email"`
}

func GenValidCode(userId uint64, userEmail string) (err error, emailValidCode string) {

	var (
		jsons = jsoniter.ConfigCompatibleWithStandardLibrary
	)

	//先生成一个uuid，随机code
	emailValidCode = fmt.Sprintf("%s", uuid.NewV4())

	var validContent ValidContent
	validContent.UserEmail = userEmail
	validContent.UserID = userId

	//序列化
	validContentStr, err := jsons.Marshal(validContent)
	if err != nil {
		logger.Error("err json Marshal", zap.Any("valid content", validContent))
		return
	}
	conn := redis.GetConn("onbio")
	defer conn.Close()

	key := fmt.Sprintf(USER_VALID_CONTENT_PRE, emailValidCode)
	_, err = conn.Do("SET", key, string(validContentStr), "EX", 6*3600, "NX")
	if err != nil && err != redigo.ErrNil {
		logger.Error("err set redis ", zap.String("key", key), zap.Error(err))
		return
	}
	return
}

func GetTitleWithMultiLan(lan string) (title string) {
	switch lan {
	case "zh-CN":
		return "OnBio | 邮箱验证"
	case "en-US":
		return "OnBio | Email verification"
	case "de-DE":
		return "OnBio | E-Mail-Verifizierung"
	case "es-ES":
		return "OnBio | Verificacion de email"
	case "fr-FR":
		return "OnBio | Vérification de l'E-mail"
	case "hi-IN":
		return "OnBio | ई - मेल सत्यापन"
	case "id-ID":
		return "OnBio | Verifikasi email"
	case "it-IT":
		return "OnBio | Verifica email"
	case "jp-JP":
		return "OnBio | メールによる確認"
	case "ko-KR":
		return "OnBio | 이메일 확인"
	case "pt-PT":
		return "OnBio | Verificação de e-mail"
	case "ru-RU":
		return "OnBio | Подтверждение по элетронной почте"
	case "th-TH":
		return "OnBio | การยืนยันอีเมล"
	case "vi-VN":
		return "OnBio | Email xác thực"
	case "zh-TW":
		return "OnBio | 郵箱驗證"
	}
	return "OnBio | Email verification"
}
