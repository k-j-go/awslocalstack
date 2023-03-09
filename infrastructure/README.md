## Lambda to bucket

lambda named "my_lambda" -> "test-bucket-lambda-dump" bucket and defined in the lambda deploy env
#### invoke it
```shell
awslocal lambda invoke --function-name my_lambda --payload '{"id": "000012", "name": "Koba Systems"}' output.json
```

#### check file
```shell
awslocal s3 ls s3://test-bucket-lambda-dump  --recursive --human-readable --summarize
```


## S3 named "test-bucket" trigger lambda named "my_lambda_trigger_accept"
"test-bucket" -> trigger to lambda named "my_lambda_trigger_accept"  -> write to "test-bucket"
#### invoke it
```shell
awslocal lambda invoke --function-name my_lambda_trigger_accept --payload '{"id": "000041", "name": "Koba Systems"}' output.json
```

#### cp to bucket test-bicket to trigger to lambda "my_lambda_trigger_accept"
```shell
aws s3 cp README.md s3://test-bucket/samplefile.txt --endpoint-url http://localhost:4566 --region us-east-1

```
