package main

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"os"
)

type Test struct {
	Name string
}

func SendEmail() {

	data := Test{
		Name: "Minh Nha",
	}

	content, err := os.ReadFile("test.html")
	if err != nil {
		log.Fatal(err)
	}

	// Convert the byte slice to a string
	htmlContent := string(content)

	fmt.Println("aaaaaaaaaaa", htmlContent)

	//tmpl, err := template.New("test").Parse(htmlContent)
	tmpl, err := template.ParseFiles("test.html")

	if err != nil {
		fmt.Println("Error parsing template", err)
	}

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, data)
	if err != nil {
		fmt.Println("Error parsing template", err)
	}

	//err = tmpl.Execute(&buf, data1)

	fmt.Println("aaaaaaaaa", buf.String())

	err = godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//from := mail.NewEmail("Example User", "test@example.com")
	//subject := "Sending with SendGrid is Fun"
	//to := mail.NewEmail("minhnha", "minhnha@smartosc.com")
	//plainTextContent := "and easy to do anywhere, even with Go"
	//
	//message := mail.NewSingleEmail(from, subject, to, plainTextContent, buf.String())
	//
	//key := os.Getenv("SENDGRID_API_KEY")
	//fmt.Println(key)
	//
	//client := sendgrid.NewSendClient(key)
	//response, err := client.Send(message)
	//if err != nil {
	//	log.Println(err)
	//} else {
	//	fmt.Println(response.StatusCode)
	//	fmt.Println(response.Body)
	//	fmt.Println(response.Headers)
	//}
}

func main() {
	SendEmail()
}
