package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

func main() {
	// 创建新的邮件消息
	m := gomail.NewMessage()

	// 设置邮件头部信息
	m.SetHeader("From", "")                 // TODO 发送方
	m.SetHeader("To", "")                   // TODO 接收方
	m.SetHeader("Subject", "邮件标题")          // 邮件主题
	m.SetBody("text/html", "<h2>邮件内容</h2>") // 邮件内容，支持HTML格式

	// 设置邮件服务器信息
	d := gomail.NewDialer("smtp.163.com", // SMTP服务器地址
		25, // 端口号
		"", // TODO 发件人邮箱账号
		"", // TODO 发件人邮箱授权码
	)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("邮件发送失败: %v", err) // 错误处理
	}
	log.Println("邮件发送成功")
}
