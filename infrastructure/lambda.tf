## Lambda
data "aws_caller_identity" "current" {}

data "aws_region" "current" {}

resource "aws_lambda_permission" "allow_bucket" {
  statement_id  = "AllowExecutionFromS3Bucket"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda.arn
  principal     = "s3.amazonaws.com"
  source_arn    = aws_s3_bucket.bucket-lambda-dump.arn
}

resource "aws_lambda_function" "lambda" {
  filename      = "../deployment.zip"
  function_name = "my_lambda"
  role          = aws_iam_role .lambda_role.arn
  handler       = "main"
  runtime       = "go1.x"

  environment {
    variables = {
      AWS_REGION = "us-east-1"
      AWS_ENDPOINT = "http://localhost:4566"
      S3_BUCKET = "test-bucket-lambda-dump"
    }
  }
}

resource "aws_lambda_function" "lambda-accept" {
  filename      = "../deployment.zip"
  function_name = "my_lambda_trigger_accept"
  role          = aws_iam_role.lambda_role.arn
  handler       = "main"
  runtime       = "go1.x"

  environment {
    variables = {
      AWS_REGION = "us-east-1"
      AWS_ENDPOINT = "http://localhost:4566"
      S3_BUCKET = "test-bucket-lambda-dump"
    }
  }
}