resource "aws_s3_bucket" "bucket-test-bucket" {
  bucket = "test-bucket"
}

resource "aws_s3_bucket" "bucket-lambda-dump" {
  bucket = "test-bucket-lambda-dump"
}

resource "aws_s3_bucket_notification" "bucket_notification" {
  bucket = aws_s3_bucket.bucket-test-bucket.id

  lambda_function {
    lambda_function_arn = aws_lambda_function.lambda-accept.arn
    events              = ["s3:ObjectCreated:*"]
    filter_prefix       = "AWSLogs/"
    filter_suffix       = ".log"
  }

  depends_on = [aws_lambda_permission.allow_bucket]
}


