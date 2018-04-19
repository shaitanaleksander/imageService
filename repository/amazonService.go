package repository

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"bytes"
)

const S3_BUCKET = "shaitanbucket"


var ses *session.Session
var uploader *s3manager.Uploader
var downloader *s3manager.Downloader
var svc *s3.S3

func init() {
	ses = session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(endpoints.UsEast1RegionID),
		Credentials: credentials.NewSharedCredentials("awsConfig", ""),
	}))
	uploader = s3manager.NewUploader(ses)
	downloader = s3manager.NewDownloader(ses)
	svc = s3.New(ses)
}

func UploadFile(name string, buffer *bytes.Buffer) error {

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(name),
		Body:   buffer,
	})
	return err
}

func DownloadFile(name string, w io.Writer) error {

	buff := &aws.WriteAtBuffer{}
	_, err := downloader.Download(buff, &s3.GetObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(name),
	})
	if err != nil {
		return err
	}
	w.Write(buff.Bytes())
	return nil
}

func DeleteImage(name string) error{
	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(name)})
	return err
}
