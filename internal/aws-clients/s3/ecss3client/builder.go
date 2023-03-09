package ecss3client

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"log"
	"os"
	"time"
)

const LOCALSTACK_HOSTNAME = "LOCALSTACK_HOSTNAME"
const LOCALSTACK_PORT = "4566"

type ECSS3Client interface {
	PutObject(ctx context.Context, bucketName string, s3Key string, data io.Reader) error
	GetObject(ctx context.Context, bucketName string, s3Key string) (io.Reader, error)
}

type ecsS3Client struct {
	s3svc      *s3.Client
	region     string
	localStack *localStack
	timeout    time.Duration
}

func (C *ecsS3Client) GetObject(ctx context.Context, bucketName string, s3Key string) (io.Reader, error) {
	ctx, cancel := context.WithTimeout(ctx, C.timeout)
	defer cancel()

	resp, err := C.s3svc.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &s3Key,
	})

	readToString := func(r io.Reader) string {
		contents := new(bytes.Buffer)
		_, err = io.Copy(contents, r)
		return contents.String()
	}

	fmt.Printf("**** %s \n", readToString(resp.Body))

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("S3 GetObject response: %s", buf.String())
	if err != nil {
		fmt.Printf("{%v}", err)
		return nil, err
	}
	return resp.Body, nil
}

type localStack struct {
	host   string
	region string
}

func (C *ecsS3Client) create() {
	customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		if C.localStack != nil {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           C.localStack.host,
				SigningRegion: C.region,
			}, nil
		}

		// returning EndpointNotFoundError will allow the service to fallback to it's default resolution
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	//Get aws.config based on the localstack
	awsCfg, err := func(stack *localStack) (aws.Config, error) {
		if stack != nil {
			provider := credentials.NewStaticCredentialsProvider("LOCAL", "LOCAL", "")
			return config.LoadDefaultConfig(context.TODO(),
				config.WithRegion(C.region),
				config.WithEndpointResolver(customResolver),
				config.WithCredentialsProvider(provider),
			)
		} else {
			return config.LoadDefaultConfig(context.TODO(),
				config.WithRegion(C.region),
				config.WithEndpointResolver(customResolver),
			)
		}
	}(C.localStack)

	if err != nil {
		log.Fatalf("Cannot load the AWS configs: %s", err)
	}

	s3svc := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	C.s3svc = s3svc
	C.timeout = time.Duration(time.Second) * 1
}

func (C *ecsS3Client) PutObject(ctx context.Context, bucketName string, s3Key string, data io.Reader) error {
	ctx, cancel := context.WithTimeout(ctx, C.timeout)
	defer cancel()
	body, err := io.ReadAll(data)
	if err != nil {
		return err
	}
	resp, err := C.s3svc.PutObject(ctx, &s3.PutObjectInput{
		Bucket:             aws.String(bucketName),
		Key:                aws.String(s3Key),
		Body:               bytes.NewReader(body),
		ContentLength:      int64(len(body)),
		ContentType:        aws.String("application/text"),
		ContentDisposition: aws.String("attachment"),
	})

	log.Printf("S3 PutObject response: %+v", resp)
	if err != nil {
		fmt.Printf("{%v}", err)
		return err
	}
	return nil
}

type ecsS3ClientBuilder struct {
	client *ecsS3Client
}

func (C *ecsS3ClientBuilder) LocalStack() *localStackBuilder {
	return &localStackBuilder{*C}
}

func (C *ecsS3ClientBuilder) AWS() *awsBuilder {
	return &awsBuilder{*C}
}

func (C *ecsS3ClientBuilder) create() {
}

type localStackBuilder struct {
	ecsS3ClientBuilder
}

func (C *localStackBuilder) Host() *localStackBuilder {
	localstackHost := os.Getenv(LOCALSTACK_HOSTNAME)
	if localstackHost != "" {
		localStackEndPoint := fmt.Sprintf("http://%s:%s", localstackHost, LOCALSTACK_PORT)
		localStack := localStack{
			host: localStackEndPoint,
		}
		C.client.localStack = &localStack
	} else {
		localStackEndPoint := fmt.Sprintf("http://%s:%s", "127.0.0.1", LOCALSTACK_PORT)
		localStack := localStack{
			host: localStackEndPoint,
		}
		C.client.localStack = &localStack
	}
	return C
}

func (C *localStackBuilder) Region(region string) *localStackBuilder {
	C.client.region = region
	return C
}

type awsBuilder struct {
	ecsS3ClientBuilder
}

func (C *awsBuilder) Region(region string) *awsBuilder {
	C.client.region = region
	return C
}

func Builder() *ecsS3ClientBuilder {
	return &ecsS3ClientBuilder{client: &ecsS3Client{}}
}

func (C *ecsS3ClientBuilder) Build() ECSS3Client {
	C.client.create()
	return C.client
}
