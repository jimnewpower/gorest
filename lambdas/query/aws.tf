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
    }
  }
}

data "archive_file" "main" {
  type        = "zip"
  source_file = "${path.module}/main"
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