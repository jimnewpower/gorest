
Get Conjur API key:
```bash
curl -k -s -X GET -u admin:"CyberArk_2023!" ${CONJUR_URL}/authn/prima/login
```

Set Conjur Environment Variables:
```bash
export CONJUR_APPLIANCE_URL=https://ec2-34-204-42-151.compute-1.amazonaws.com
export CONJUR_ACCOUNT=prima
export CONJUR_AUTHN_LOGIN=admin
export CONJUR_AUTHN_API_KEY=18wv7sck9a66015fzsv3252qfvp23anzs81qkn4f916fbs3t228p4nb
export CONJUR_CERT_FILE=/home/jim/develop/gorest/lambdas/query/conjur-dev.pem
```

These are the Conjur Environment Variables in the Go API client configuration:
-		ApplianceURL:      os.Getenv("CONJUR_APPLIANCE_URL"),
-		SSLCert:           os.Getenv("CONJUR_SSL_CERTIFICATE"),
-		SSLCertPath:       os.Getenv("CONJUR_CERT_FILE"),
-		Account:           os.Getenv("CONJUR_ACCOUNT"),
-		NetRCPath:         os.Getenv("CONJUR_NETRC_PATH"),
-		CredentialStorage: os.Getenv("CONJUR_CREDENTIAL_STORAGE"),
-		AuthnType:         os.Getenv("CONJUR_AUTHN_TYPE"),
-		ServiceID:         os.Getenv("CONJUR_SERVICE_ID"),


Get dependencies and run:
```bash
go get github.com/cyberark/conjur-api-go/conjurapi
go get github.com/cyberark/conjur-api-go/conjurapi/authn

go run main.go
```


# Build
Build and zip:
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
zip main.zip main conjur-dev.pem
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

# Conjur Go API
[GitHub](https://github.com/cyberark/conjur-api-go)

