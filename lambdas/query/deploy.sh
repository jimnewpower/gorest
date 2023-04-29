#!/bin/bash
terraform init
terraform plan
echo "Run terraform apply to deploy the lambda function"