package s3

import (
	"context"
	"fmt"
	"github.com/k-j-go/awslocalstack/internal/aws-clients/s3/ecss3client"
	"io"
	"strings"
	"testing"
)

func TestS3_localstack_builder(t *testing.T) {
	s3svr := ecss3client.Builder().LocalStack().Host().Region("us-east-1").Build()
	err := s3svr.PutObject(context.TODO(), "test-bucket", "test4", strings.NewReader("test"))
	if err != nil {
		return
	}

	data, err := s3svr.GetObject(context.TODO(), "test-bucket", "test4")
	if err != nil {
		return
	}

	all, err := io.ReadAll(data)
	if err != nil {
		return
	}

	fmt.Printf("%s \n", string(all[:]))
}

func TestS3_aws_builder(t *testing.T) {
	ecss3client.Builder().AWS().Region("us-east-1").Build()
}

func TestS3_play(t *testing.T) {
	ecss3client.Builder().LocalStack().
		Region("us-east-1").Build().
		PutObject(context.TODO(), "test-bucket", "some", strings.NewReader("some"))
}
