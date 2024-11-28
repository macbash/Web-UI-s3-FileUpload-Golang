package main

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

)

const (
	// S3 bucket information
	BucketName = "s320241128"
	Region     = "us-east-1" // e.g., us-west-2
)

func uploadFileToS3(file multipart.File, fileName string) error {
	// Initialize a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewEnvCredentials(), // or manually set your access keys
	})
	if err != nil {
		return fmt.Errorf("unable to create AWS session: %v", err)
	}

	// Initialize S3 service client
	s3Svc := s3.New(sess)

	// Upload the file to S3
	_, err = s3Svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("unable to upload file to S3: %v", err)
	}

	return nil
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming form data (multipart form data)
	err := r.ParseMultipartForm(10 << 20) // Limit to 10MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from the form
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get the file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get the file name from the form
	fileName := r.FormValue("filename")

	// Upload the file to S3
	err = uploadFileToS3(file, fileName)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error uploading to S3: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond to the user
	w.Write([]byte("File uploaded successfully!"))
}

func main() {
	// Setup HTTP handler for file upload
	http.HandleFunc("/upload", uploadHandler)

	// Start the web server
	port := ":8080"
	log.Printf("Starting server on %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
