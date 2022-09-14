package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func main() {
	newMsg := gomail.NewMessage()
	newMsg.SetHeader("From", "thebigyovadiaz@outlook.com")
	newMsg.SetHeader("To", "thebigyovadiaz@gmail.com")
	newMsg.SetHeader("Subject", "Test Sended Email to Go and GOMAIL")
	newMsg.SetBody("text/html", "<b>This the text in the body email</b>")

	a := gomail.NewDialer("smtp-mail.outlook.com", 587, "example@outlook.com", "ExamplePass")
	if err := a.DialAndSend(newMsg); err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Email Sended!")
}
