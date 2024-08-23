package main

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Person struct {
	Name string
	Age  string
}

var PersonColumn = []string{"Name", "Age"}
var test = []string{"John Dow", "30"}

func sendCSVWithSendGrid(apiKey, fromEmail, toEmail, subject, csvFilePath string) error {
	client := sendgrid.NewSendClient(apiKey)

	from := mail.NewEmail("Example User", fromEmail)
	to := mail.NewEmail("Example User", toEmail)
	message := mail.NewSingleEmail(from, subject, to, "Please find the attached CSV file.", "Please find the attached CSV file.")

	//Read the CSV file
	//csvData, err := ioutil.ReadFile(csvFilePath)
	//if err != nil {
	//	return fmt.Errorf("failed to read CSV file: %w", err)
	//}
	//
	//buffer := bytes.NewBuffer(csvData)

	var records = [][]string{PersonColumn}

	fmt.Println("records", len(records))

	records = append(records, test)

	//arr := []Person{
	//	{
	//		Name: "Name",
	//		Age:  "Age",
	//	},
	//	{
	//		Name: "John Doe",
	//		Age:  "30",
	//	}}

	var b []byte
	buf := bytes.NewBuffer(b)
	w := csv.NewWriter(buf)
	if err := w.WriteAll(records); err != nil {
		fmt.Println("Error writing csv:", err)
	}

	base64Data := base64.StdEncoding.EncodeToString(buf.Bytes())

	fmt.Println("Buffer content:", buf.String())

	// Attach the CSV file
	attachment := mail.NewAttachment()
	attachment.SetContent(base64Data)
	attachment.SetType("text/csv")
	attachment.SetFilename("data.csv")
	attachment.SetDisposition("attachment")
	message.AddAttachment(attachment)

	//Send the email
	response, err := client.Send(message)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	log.Printf("Email sent with status code: %d", response.StatusCode)
	log.Printf("Email sent with body: %s", response.Body)
	return nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("SENDGRID_API_KEY")
	fromEmail := "sender@example.com"
	toEmail := "minhnha@smartosc.com"
	subject := "CSV File Attachment"
	csvFilePath := "test1.csv"

	fmt.Println("apiKey", apiKey)

	_ = sendCSVWithSendGrid(apiKey, fromEmail, toEmail, subject, csvFilePath)
	if err != nil {
		log.Fatalf("Error sending email: %v", err)
	}
}
