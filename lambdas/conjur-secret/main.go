package main

import (
    "os"
    "fmt"
    "github.com/cyberark/conjur-api-go/conjurapi"
    "github.com/cyberark/conjur-api-go/conjurapi/authn"
)

func GetConjurClient() (*conjurapi.Client, error) {
	config, err := conjurapi.LoadConfig()
	if err != nil {
		return nil, err
	}

	conjur, err := conjurapi.NewClientFromKey(config,
		authn.LoginPair{
			Login:  os.Getenv("CONJUR_AUTHN_LOGIN"),
			APIKey: os.Getenv("CONJUR_AUTHN_API_KEY"),
		},
	)
	if err != nil {
		return nil, err
	}

	return conjur, nil
}

func RetrieveSecret(variableIdentifier string) ([]byte, error) {
	conjur, err := GetConjurClient()
	if err != nil {
		return nil, err
	}

	// Retrieve a secret into []byte.
	secretValue, err := conjur.RetrieveSecret(variableIdentifier)
	if err != nil {
		return nil, err
	}

    // Retrieve a secret into io.ReadCloser, then read into []byte.
    // Alternatively, you can transfer the secret directly into secure memory,
    // vault, keychain, etc.
    secretResponse, err := conjur.RetrieveSecretReader(variableIdentifier)
    if err != nil {
        panic(err)
    }

	secretValue, err = conjurapi.ReadResponseBody(secretResponse)
    if err != nil {
        panic(err)
    }

	return secretValue, nil
}

func main() {
    variableIdentifier := "postgresDBApp/password"
	secretValue, err := RetrieveSecret(variableIdentifier) // returns []byte, error
    if err != nil {
        panic(err)
    }
	fmt.Println(fmt.Sprintf("%s: %s", variableIdentifier, string(secretValue)))
}