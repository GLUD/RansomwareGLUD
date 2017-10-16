/* Idea original tomada de: https://gist.github.com/jpillora/cb46d183eca0710d909a, modificada el 16 de octubre de 2017. */
package main

import (
	"log"
	"net/smtp"
)

func main() {
	send("Hello ^^' '")
}

func send(body string) {
	from := "Leiidy.152011@gmail.com"
	pass := "L123456A"
	to := "andres.taamap@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Correo desde Go\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	//log.Print("sent, visit http://foobarbazz.mailinator.com")
}
