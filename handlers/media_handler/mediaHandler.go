package media_handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var (
	region = "ap-south-1"
	bucket = "bongobasi"
)

func UploadMediaToS3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes := bytes.NewBuffer(nil)
	if _, err := fileBytes.ReadFrom(file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileType := header.Header.Get("Content-Type")
	fileName := header.Filename

	fmt.Println("File Name:", fileName)
	fmt.Println("File Type:", fileType)

	// Upload the file to S3
	objectKey, err := uploadToS3(fileBytes, fileName, fileType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the public image link
	imageURL := getPublicURL(objectKey)
	// send img url in json for
	data := map[string]interface{}{
		"img_url": imageURL,
	}
	json.NewEncoder(w).Encode(data)

}

func uploadToS3(file *bytes.Buffer, fileName, fileType string) (string, error) {
	// get environment variables
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	bucket := os.Getenv("AWS_BUCKET_NAME")
	// Create a bytes.Reader from the file buffer
	fileReader := bytes.NewReader(file.Bytes())

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return "", err
	}

	svc := s3.New(sess)

	// Generate a unique object key
	objectKey := generateObjectKey(fileName)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(objectKey),
		ACL:         aws.String("public-read"),
		ContentType: aws.String(fileType),
		Body:        fileReader, // Use the bytes.Reader
	})
	if err != nil {
		return "", err
	}

	return objectKey, nil
}

func generateObjectKey(fileName string) string {
	fileExt := filepath.Ext(fileName)
	objectKey := fmt.Sprintf("%s%s", generateUUID(fileName), fileExt)
	return objectKey
}

func generateUUID(fileName string) string {
	// Generate a unique ID using your preferred method
	// This is just a placeholder
	return strings.ReplaceAll(fileName, " ", "") + uuid.New().String()
}

func getPublicURL(objectKey string) string {
	return fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", bucket, region, objectKey)
}
