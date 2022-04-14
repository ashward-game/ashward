package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

var (
	ss                                               *session.Session
	s3Session                                        *s3.S3
	region, bucketName, accessKeyID, secretAccessKey string
)

func Init() {
	region = os.Getenv("AWS_REGION")
	bucketName = os.Getenv("AWS_BUCKET_NAME")
	accessKeyID = os.Getenv("AWS_ACCESS_KEY")
	secretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")

	// Create session
	ss = session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,
			secretAccessKey,
			"",
		),
	}))

	s3Session = s3.New(ss)
}

func UploadFile(path string, filename string) (resp *s3manager.UploadOutput) {
	// Open the file for use
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	uploader := s3manager.NewUploader(ss)
	fmt.Println("Uploading .... ", filename)
	resp, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(path + filename),
		Body:   file,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("Uploaded successfully !")
	return resp
}

func ListFiles() (resp *s3.ListObjectsV2Output) {
	resp, err := s3Session.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		panic(err)
	}
	return resp
}
