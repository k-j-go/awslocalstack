## Create a lambda and it writes to a s3 bucket 
- using terraform to create lambda, s3
- runing on localstack
- implement the S2 client build pattern

```shell
terraform init
terraform apply --auto-approve
```

#### call lambda 
```shell
awslocal lambda invoke --function-name my_lambda --payload '{"id": "00002", "name": "Koba Systems"}' output.json
```

#### check lambda write to bucker
```shell
awslocal s3 ls s3://test-bucket-lambda-dump --recursive --human-readable --summarize
```

/////////////////////////////////////////////////////////////////////////

## test s3 trigger lambda

#### check lambda works
#### call lambda
```shell
awslocal lambda invoke --function-name my_lambda_trigger_accept --payload '{"id": "00302", "name": "Koba Systems"}' output.json
```
#####  lambda write file to s3 bucket
```shell
awslocal s3 ls s3://test-bucket-lambda-dump --recursive --human-readable --summarize
```
##### cp to s2 test-bucket to trigger lambda
```shell
aws s3 cp main.go s3://test-bucket/samplefile.txt --endpoint-url http://localhost:4566 --region us-east-1
```


[good](https://github.com/jagonzalr/go-lambda-terraform-setup)