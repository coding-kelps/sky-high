package repositories

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/coding-kelps/sky-high/pkg/config"
)

var S3 *s3.S3

func InitS3Client(cfg *config.S3Config) error {
	// Create an AWS session
	awsSession, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(cfg.Endpoint),
		Region:   aws.String(cfg.Region),
		Credentials: credentials.NewStaticCredentials(
			cfg.AccessKeyID, cfg.AccessKeySecret, ""),
	})
	if err != nil {
		return err
	}

	S3 = s3.New(awsSession)

	return nil
}
