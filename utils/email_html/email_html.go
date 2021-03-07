package email_html

import (
	"onbio/logger"

	"github.com/matcornic/hermes/v2"
	"go.uber.org/zap"
)

type EmailType uint32

const (
	Register_Account EmailType = 1
	Reset_Pwd        EmailType = 2
)

/*这个内容不需要了
func (p EmailType) GetContent() string {
	switch p {
	case Register_Account:
		return "Thanks for signing up, great to have you!"
	case Reset_Pwd:
		return "You are trying to reset your password"
	default:
		return "UNKNOWN"
	}
}*/

func (p EmailType) GetMultiLanContentWithRegister(lan string) (pre string, after string) {
	switch lan {
	case "zh-CN":
		pre = "请点击下方按钮或打开此链接（"
		after = "），验证电子邮箱地址"
		return
	case "en-US":
		pre = "Please click on the button below or open this link ("
		after = ") to verify your email address"
		return
	case "de-DE":
		pre = "Bitte klicken Sie auf die Schaltfläche unten oder öffnen Sie diesen Link ("
		after = "), um Ihre E-Mail-Adresse zu bestätigen"
		return
	case "es-ES":
		pre = "Por favor, haga clic en el botón de abajo o abra este enlace ("
		after = ") para verificar su dirección de correo electrónico"
		return
	case "fr-FR":
		pre = "Veuillez cliquer sur le bouton ci-dessous ou ouvrir ce lien ("
		after = ") pour vérifier votre adresse e-mail"
		return
	case "hi-IN":
		pre = "कृपया नीचे दिए गए बटन पर क्लिक करें या अपने ईमेल पते को सत्यापित करने के लिए इस लिंक ("
		after = ") को खोलें"
		return
	case "id-ID":
		pre = "Silakan klik tombol di bawah ini atau buka tautan ini ("
		after = ") untuk memverifikasi alamat email Anda"
		return
	case "it-IT":
		pre = "Si prega di fare clic sul pulsante qui sotto o aprire questo link ("
		after = ") per verificare il tuo indirizzo e-mail"
		return
	case "jp-JP":
		pre = "下のボタンをクリックするか、このリンク("
		after = ")を開いてメールアドレスを確認してください"
		return
	case "ko-KR":
		pre = "아래 버튼을 클릭하거나 이 링크("
		after = ")를 열어 이메일 주소를 확인하십시오."
		return
	case "pt-PT":
		pre = "Clique no botão abaixo ou abra este link ("
		after = ") para verificar o seu endereço de e-mail"
		return
	case "ru-RU":
		pre = "Пожалуйста, нажмите на кнопку ниже или откройте эту ссылку ("
		after = "), чтобы проверить свой адрес электронной почты"
		return
	case "th-TH":
		pre = "กรุณาคลิกที่ปุ่มด้านล่างหรือเปิดลิงค์นี้ ("
		after = "), чтобы проверить свой адрес электронной почты"
		return
	case "vi-VN":
		pre = "Vui lòng nhấp vào nút bên dưới hoặc mở liên kết này ("
		after = ") để xác minh địa chỉ email của bạn"
		return
	case "zh-TW":
		pre = "請點擊下方按鈕或打開此連結（"
		after = "），驗證電子郵箱位址"
		return
	}
	pre = "Please click on the button below or open this link ("
	after = ") to verify your email address"
	return
}

func (p EmailType) GetGreetingWithMultiLan(lan string) (greeting string) {
	switch lan {
	case "zh-CN":
		return "Hi"
	case "en-US":
		return "Hi"
	case "de-DE":
		return "Hallo"
	case "es-ES":
		return "Hola"
	case "fr-FR":
		return "Salut"
	case "hi-IN":
		return "नमस्ते "
	case "id-ID":
		return "Hai"
	case "it-IT":
		return "Ciao"
	case "jp-JP":
		return "こんにちは"
	case "ko-KR":
		return "하이 "
	case "pt-PT":
		return "Oi "
	case "ru-RU":
		return "Привет "
	case "th-TH":
		return "สวัสดี "
	case "vi-VN":
		return "Chào"
	case "zh-TW":
		return "Hi"
	}
	return "Hi"
}
func (p EmailType) GetMultiLanContentWithReset(lan string) (pre string, after string) {
	switch lan {
	case "zh-CN":
		pre = "请点击下方按钮或打开此链接（"
		after = "），设置新密码"
		return
	case "en-US":
		pre = "Please click on the button below or open this link ("
		after = ") to set a new password"
		return
	case "de-DE":
		pre = "Bitte klicken Sie auf die Schaltfläche unten oder öffnen Sie diesen Link ("
		after = "), um ein neues Passwort festzulegen"
		return
	case "es-ES":
		pre = "Por favor, haga clic en el botón de abajo o abra este enlace ("
		after = ") para establecer una nueva contraseña"
		return
	case "fr-FR":
		pre = "Veuillez cliquer sur le bouton ci-dessous ou ouvrir ce lien ("
		after = ") pour vérifier votre adresse e-mail"
		return
	case "hi-IN":
		pre = "कृपया नीचे दिए गए बटन पर क्लिक करें या एक नया पासवर्ड सेट करने के लिए इस लिंक ("
		after = ") को खोलें"
		return
	case "id-ID":
		pre = "Silakan klik tombol di bawah ini atau buka tautan ini ("
		after = ") untuk menyetel sandi baru"
		return
	case "it-IT":
		pre = "Si prega di fare clic sul pulsante qui sotto o aprire questo link ("
		after = ") per impostare una nuova password"
		return
	case "jp-JP":
		pre = "下のボタンをクリックするか、このリンク("
		after = ")を開いて新しいパスワードを設定してください"
		return
	case "ko-KR":
		pre = "아래 버튼을 클릭하거나 이 링크("
		after = ")를 열어 새 비밀번호를 설정하십시오."
		return
	case "pt-PT":
		pre = "Clique no botão abaixo ou abra este link ("
		after = ") para definir uma nova senha"
		return
	case "ru-RU":
		pre = "Пожалуйста, нажмите на кнопку ниже или откройте эту ссылку ("
		after = "), чтобы установить новый пароль"
		return
	case "th-TH":
		pre = "กรุณาคลิกที่ปุ่มด้านล่างหรือเปิดลิงค์นี้ ("
		after = ") เพื่อตั้งรหัสผ่านใหม่"
		return
	case "vi-VN":
		pre = "Vui lòng nhấp vào nút bên dưới hoặc mở liên kết này ("
		after = ") để xác minh địa chỉ email của bạn"
		return
	case "zh-TW":
		pre = "請點擊下方按鈕或打開此連結（"
		after = "），設置新密碼"
		return
	}
	pre = "Please click on the button below or open this link ("
	after = ") to set a new password"
	return
}

func (p EmailType) GetIntroContent(url string, lan string) string {
	switch p {
	case Register_Account:
		pre, after := p.GetMultiLanContentWithRegister(lan)
		return pre + url + after
	case Reset_Pwd:
		pre, after := p.GetMultiLanContentWithReset(lan)
		return pre + url + after
	default:
		return "UNKNOWN"
	}
}

func GenerateHtml(userName, url, lan string, emailType EmailType) (emailBody string, err error) {

	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "Onb.io",
			Link: "https://onb.io/",
			// Optional product logo
			Logo:      "http://onb.io/_nuxt/assets/images/logo.png",
			Copyright: "Onb.io",
		},
	}
	intros := []string{""}
	email := hermes.Email{
		Body: hermes.Body{
			Greeting: emailType.GetGreetingWithMultiLan(lan),
			Name:     userName,
			Intros:   intros,
			Actions: []hermes.Action{
				{
					Instructions: emailType.GetIntroContent(url, lan),
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "Go",
						Link:  url,
					},
				},
			},
			Outros: []string{
				"",
			},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err = h.GenerateHTML(email)
	if err != nil {
		logger.Error("GenerateHTML failed ", zap.Error(err))
		return
	}
	return
}
