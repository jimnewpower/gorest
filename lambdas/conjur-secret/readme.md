
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
export CONJUR_CERT_FILE=/home/jim/develop/gorest/lambdas/conjur-secret/conjur-dev.pem
```

Get dependencies and run:
```bash
go get github.com/cyberark/conjur-api-go/conjurapi
go get github.com/cyberark/conjur-api-go/conjurapi/authn

go run main.go
```