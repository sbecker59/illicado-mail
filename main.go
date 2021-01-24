package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/gomail.v2"
)

var (
	domain = os.Getenv("ILLICADO_MAIL_DOMAIN")
)

func main() {

	files, err := ioutil.ReadDir("./images")
	if err != nil {
		log.Fatal(err)
	}

	var mails []string
	for _, file := range files {
		mail := strings.Split(strings.ReplaceAll(file.Name(), ".png", ""), "-")[0] + "@" + domain
		if !contains(mails, mail) {
			mails = append(mails, mail)
		}

	}

	for _, mail := range mails {

		fmt.Println(mail)
		m := gomail.NewMessage()
		m.SetHeader("From", os.Getenv("ILLICADO_MAIL_FROM"))
		m.SetHeader("To", mail)
		m.SetHeader("Subject", "CSE Neosoft Lille - Carte Illica")
		m.SetBody("text/html", "")

		for i := 0; i < 10; i++ {
			file := fmt.Sprintf("./images/%s-%d.png", strings.ReplaceAll(mail, "@"+domain, ""), i)
			if fileExists(file) {
				m.Attach(file)
			}
		}

		d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("ILLICADO_MAIL_USER"), os.Getenv("ILLICADO_MAIL_PASSWORD"))

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}

}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		*files = append(*files, path)
		return nil
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
