package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found.")
	}
}

func main() {
	awsRegion := os.Getenv("AWS_REGION")
	s3Bucket := os.Getenv("S3_BUCKET")
	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	if awsRegion == "" || s3Bucket == "" || awsAccessKeyID == "" || awsSecretAccessKey == "" {
		log.Fatal("AWS_REGION, S3_BUCKET, AWS_ACCESS_KEY_ID, and AWS_SECRET_ACCESS_KEY are required.")
	}

	s3Service := NewS3Uploader(awsRegion, s3Bucket, awsAccessKeyID, awsSecretAccessKey)

	r := mux.NewRouter()
	r.HandleFunc("/upload", FileUploadHandler(s3Service)).Methods("POST")

	log.Println("Server is running on port 8001...")
	log.Fatal(http.ListenAndServe(":8001", r))
}
