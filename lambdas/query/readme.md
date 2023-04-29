# Build
Build and zip:
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
zip main.zip main
```

# Deploy
- Manual deploy: Upload `main.zip` to AWS Lambda.
- Automated deploy with Terraform (type `dev` for environment, when prompted):
```bash
aws configure
terraform init
terraform plan
terraform apply
```

# Test
```bash
aws lambda invoke --function-name query --payload '{"id": "1"}' output.txt
curl https://n3ipnwcww3hoxcidozmtiudzjy0ziddl.lambda-url.us-west-2.on.aws/
```

# URL
- [Dedicated URL](https://n3ipnwcww3hoxcidozmtiudzjy0ziddl.lambda-url.us-west-2.on.aws/)
