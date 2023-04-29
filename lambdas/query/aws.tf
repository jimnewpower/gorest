variable "env_name" {
  description = "Environment name"
}

locals {
  function_name               = "GoGetLogistics"
  function_handler            = "main"
  function_runtime            = "go1.x"
  function_timeout_in_seconds = 5

  function_source_dir = "${path.module}"
}

resource "aws_lambda_function" "function" {
  function_name = "${local.function_name}-${var.env_name}"
  handler       = local.function_handler
  runtime       = local.function_runtime
  timeout       = local.function_timeout_in_seconds

  filename         = "main.zip"
  source_code_hash = data.archive_file.main.output_base64sha256

  role = aws_iam_role.function_role.arn

  environment {
    variables = {
      ENVIRONMENT = var.env_name
      CONJUR_ACCOUNT = "prima"
      CONJUR_APPLIANCE_URL = "https://ec2-34-204-42-151.compute-1.amazonaws.com"
      CONJUR_CERT_FILE = "./conjur-dev.pem"
      CONJUR_AUTHN_LOGIN = "admin"
	    CONJUR_AUTHN_API_KEY = "18wv7sck9a66015fzsv3252qfvp23anzs81qkn4f916fbs3t228p4nb"
      CONJUR_AUTHENTICATOR = "authn-iam"
      HOST = "prima.cvrj95nytzmd.us-west-2.rds.amazonaws.com"
      PORT = "5432"
    }
  }
}

data "archive_file" "main" {
  type        = "zip"
  source_dir = "${path.module}/bin"
  output_path = "${path.module}/main.zip"
}

resource "aws_iam_role" "function_role" {
  name = "${local.function_name}-${var.env_name}"

  assume_role_policy = jsonencode({
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      },
    ]
  })
}