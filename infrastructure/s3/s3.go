package s3

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"Sharykhin/rent-car/domain"
	"Sharykhin/rent-car/logger"
)

type (
	// Client provides api to upload file on S3 to already specified bucket
	Client struct {
		s3     *s3.S3
		bucket string
	}
)

// NewClient creates a new instance of client. Such params as endpoint and forcePath are used mostly
// for local development to store files in minio
func NewClient(endpoint string, forcePath bool, bucket string, region string) *Client {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	var s3Endpoint *string
	if endpoint != "" {
		s3Endpoint = aws.String(endpoint)
	}

	s3Client := s3.New(sess, &aws.Config{
		Endpoint:         s3Endpoint,
		S3ForcePathStyle: aws.Bool(forcePath),
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
		Region: aws.String(region),
	})

	return &Client{
		s3:     s3Client,
		bucket: bucket,
	}
}

// Upload uploads file on S3 with public-read
func (c *Client) Upload(ctx context.Context, path string, data []byte) error {
	logger.Log.Debugf("[infrastructure][s3][Client][Upload] uploading file to %s/%s", c.bucket, path)
	_, err := c.s3.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(c.bucket),
		Key:           aws.String(path),
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(data),
		ContentLength: aws.Int64(int64(len(data))),
		ContentType:   aws.String(http.DetectContentType(data)),
	})

	if err != nil {
		return domain.NewInternalError(
			fmt.Errorf("failed to upload file on s3: %v", err),
			"[infrastructure][s3][Client][Upload]",
		)
	}

	return nil
}
