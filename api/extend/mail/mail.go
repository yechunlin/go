package mail

import "gopkg.in/gomail.v2"

func SendMail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "admin@912zufang.com")
	m.SetHeader("To", "2648300401@qq.com")
	m.SetHeader("Subject", "您的会员权益即将到期，请及时续费")
	m.SetBody("text/html", "<html><header><title>会员信息提醒</title></header><body><h2>你™快点续费，快来不及了！</h2></body></html>")

	d := gomail.NewDialer("smtp.exmail.qq.com", 465, "admin@912zufang.com", "Abc12345666")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
