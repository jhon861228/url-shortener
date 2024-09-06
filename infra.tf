provider "aws" {
  region = var.region
  profile = "personal"
}

variable "region" {
  description = "The AWS region to deploy the infrastructure"
  type        = string
  default     = "us-east-1"
}

# Tabla DynamoDB
resource "aws_dynamodb_table" "shorturl_table" {
  name           = "shorturl-table"
  billing_mode    = "PAY_PER_REQUEST"
  hash_key        = "id"
  
  attribute {
    name = "id"
    type = "S"
  }
}

# Rol de ejecución de Lambda
resource "aws_iam_role" "lambda_execution_role" {
  name = "lambda_url_shortener_execution_role"
  
  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

# Política de permisos al rol de Lambda
resource "aws_iam_role_policy_attachment" "lambda_policy_attachment" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role     = aws_iam_role.lambda_execution_role.name
}

# Política IAM para acceder a DynamoDB
resource "aws_iam_policy" "lambda_dynamodb_policy" {
  name        = "lambda-dynamodb-access"
  description = "Allow Lambda function to access DynamoDB table"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect   = "Allow",
        Action   = [
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "dynamodb:UpdateItem",
          "dynamodb:DeleteItem",
          "dynamodb:Scan",
          "dynamodb:Query"
        ],
        Resource = aws_dynamodb_table.shorturl_table.arn
      }
    ]
  })
}

# Política de DynamoDB al rol Lambda
resource "aws_iam_role_policy_attachment" "lambda_dynamodb_attachment" {
  policy_arn = aws_iam_policy.lambda_dynamodb_policy.arn
  role     = aws_iam_role.lambda_execution_role.name
}

# Función Lambda
resource "aws_lambda_function" "my_lambda_function" {
  function_name = "url-shortener"
  role          = aws_iam_role.lambda_execution_role.arn
  handler       = "main"
  runtime       = "provided.al2023"
  
  filename    = "./backend/url-shortener.zip"

  source_code_hash = filebase64sha256("./backend/url-shortener.zip")

  memory_size = 1024
  
  timeout = 30

  environment {
    variables = {
      TABLE_NAME = aws_dynamodb_table.shorturl_table.name
    }
  }
}

# API Gateway
resource "aws_api_gateway_rest_api" "api" {
  name        = "api-url-shortener"
  description = "API for my Lambda function"

  body = file("${path.module}/openapi/shorturl_api.json")
}

# API Gateway invoque la función Lambda
resource "aws_lambda_permission" "api_gateway" {

  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.my_lambda_function.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_api_gateway_rest_api.api.execution_arn}/*/*"
}

# Etapa de despliegue para la API
resource "aws_api_gateway_deployment" "deployment" {
  depends_on   = [aws_api_gateway_rest_api.api]
  rest_api_id  = aws_api_gateway_rest_api.api.id
  stage_name   = "prod"
  triggers = {
    redeploy = sha256(file("${path.module}/openapi/shorturl_api.json"))
  }
}

# Api key
resource "aws_api_gateway_api_key" "api_key" {
  name        = "api-key-frontend"
  description = "API Key for frontend consumer"
  enabled     = true
}

# Plan de uso
resource "aws_api_gateway_usage_plan" "usage_plan" {
  name = "FrontendUsagePlan"

  api_stages {
    api_id = aws_api_gateway_rest_api.api.id
    stage  = aws_api_gateway_deployment.deployment.stage_name
  }

  throttle_settings {
    burst_limit = 5
    rate_limit  = 10
  }

  quota_settings {
    limit  = 10000
    period = "MONTH"
  }
  
}

# Asociación de API Key con el Usage Plan
resource "aws_api_gateway_usage_plan_key" "my_usage_plan_key" {
  key_id        = aws_api_gateway_api_key.api_key.id
  key_type      = "API_KEY"
  usage_plan_id = aws_api_gateway_usage_plan.usage_plan.id
}

output "api_gateway_url" {
  description = "URL del API Gateway"
  value       = "https://${aws_api_gateway_rest_api.api.id}.execute-api.${var.region}.amazonaws.com/${aws_api_gateway_deployment.deployment.stage_name}"
}