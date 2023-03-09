package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/k-j-go/awslocalstack/internal/aws-clients/s3/ecss3client"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// CustomEvent for lambda
type CustomEvent struct {
	ID     string
	Name   string
	Number int32
}

var (
	awsRegion      string
	awsEndpoint    string
	bucketName     string
	localstackHost string

	s3svc *s3.Client
)

func init() {
	awsRegion = os.Getenv("AWS_REGION")
	awsEndpoint = os.Getenv("AWS_ENDPOINT")
	bucketName = os.Getenv("S3_BUCKET")

	localstackHost = os.Getenv("LOCALSTACK_HOSTNAME")
	awsEndpoint = fmt.Sprintf("http://%s:4566", localstackHost)

	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if awsEndpoint != "" {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           awsEndpoint,
				SigningRegion: awsRegion,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithEndpointResolver(customResolver),
	)
	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}

	s3svc = s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
}

func handler(ctx context.Context, event CustomEvent) error {
	s3Key := fmt.Sprintf("%s.txt", event.ID)
	_ = fmt.Sprintf("%s.txt", event.Number)
	body := fmt.Sprintf("Hello, %s", event.Name)

	ecs := ecss3client.
		Builder().
		LocalStack().
		Region("us-east-1").
		Host().
		Build()

	err := ecs.PutObject(ctx, bucketName, s3Key, strings.NewReader(body))
	if err != nil {
		return err
	}
	r, err := ecs.GetObject(ctx, bucketName, s3Key)
	if err != nil {
		return err
	}

	//Converter io.read to string
	//https://dev.to/dayvonjersen/how-to-use-io-reader-and-io-writer-ao0
	readToString := func(r io.Reader) string {
		contents := new(bytes.Buffer)
		_, err = io.Copy(contents, r)
		return contents.String()
	}

	readToString(r)

	return nil
}

func main() {
	lambda.Start(handler)
}
