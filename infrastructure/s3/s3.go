package s3

import (
	"Sharykhin/rent-car/logger"
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"Sharykhin/rent-car/domain"
)

type (
	Client struct {
		s3     *s3.S3
		bucket string
	}
)

func NewClient(endpoint string, forcePath bool, bucket string) *Client {
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
	})

	return &Client{
		s3:     s3Client,
		bucket: bucket,
	}
}

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
