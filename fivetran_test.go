package fivetran_test

import (
	"log"
	"os"

	"github.com/fivetran/go-fivetran"
	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func init() {
	var apiUrl string
	var apiKey string
	var apiSecret string

	valuesToLoad := map[string]*string{
		"FIVETRAN_API_URL":               &apiUrl,
		"FIVETRAN_APIKEY":                &apiKey,
		"FIVETRAN_APISECRET":             &apiSecret,
		"FIVETRAN_TEST_CERTIFICATE_HASH": &testutils.CertificateHash,
		"FIVETRAN_TEST_CERTIFICATE":      &testutils.EncodedCertificate,
	}

	for name, value := range valuesToLoad {
		*value = os.Getenv(name)
		if *value == "" {
			log.Fatalf("Environment variable %s is not set!\n", name)
		}
	}

	testutils.Client = fivetran.New(apiKey, apiSecret)
	testutils.Client.BaseURL(apiUrl)
	if testutils.IsPredefinedUserExist() && testutils.IsPredefinedGroupExist() {
		testutils.CleanupAccount()
	} else {
		log.Fatalln("The predefined user doesn't belong to the Testing account. Make sure that credentials are using in the tests belong to the Testing account.")
	}
}
