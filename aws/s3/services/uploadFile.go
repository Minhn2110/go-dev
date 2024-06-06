package services

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func GetFileFromS3() {
	bucket := "grit-smb-cart-editing"
	item := "dev.txt"

	file, err := os.Create(item)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}

	defer file.Close()

	sess, _ := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
		},
	)

	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		exitErrorf("Unable to download item %q, %v", item, err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func UploadToS3() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Failed to open file", err)
		return
	}
	defer file.Close()

	sess, _ := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
		},
	)

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("grit-smb-cart-editing"),
		Key:    aws.String("dev.txt"),
		Body:   file,
	})
	if err != nil {
		fmt.Println("Failed to upload", err)
		return
	}
	fmt.Println("Successfully uploaded to", "grit-smb-cart-editing")
}
