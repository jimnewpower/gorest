#!/bin/bash
terraform init
terraform plan

echo ""
echo ""
echo "Run 'terraform apply' to deploy the lambda function."