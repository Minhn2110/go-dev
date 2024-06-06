package main

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env")  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Mail

	from := mail.NewEmail("Example User", "test@example.com")
	subject := "Hello World from the SendGrid Go Library"
	to := mail.NewEmail("Example User", "minhnha@smartosc.com")
	plainTextContent := "Hello, Email!"
	htmlContent := "<strong>Hello, Email!</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	// Read the file
	//data, err := ioutil.ReadFile("/path/to/your/file.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Create a new attachment
	attachment := mail.NewAttachment()
	attachment.SetContent(string("YXNkYXNkc2Fkc2FkYXNkYXNkYXNk"))
	attachment.SetType("text/plain")
	attachment.SetFilename("file.txt")
	attachment.SetDisposition("attachment")
	attachment.SetContentID("Example Content ID")

	// Add the attachment to the message
	message.AddAttachment(attachment)

	client := sendgrid.NewSendClient(viper.GetString("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
}
