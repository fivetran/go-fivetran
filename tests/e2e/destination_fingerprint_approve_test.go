package fivetran_test

import (
	"context"
	"testing"

	testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewCertificateDestinationFingerprintApproveE2E(t *testing.T) {
	destinationId := testutils.CreateTempDestination(t)
	response, err := testutils.Client.NewCertificateDestinationFingerprintApprove().
		DestinationID(destinationId).
		Hash("test_hash").
		PublicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDK3K7vPLVzJ7zE2YvZ5K8vK9m9L2vM3nO4pP5qR6sT7uV8wX9yZ0aB1cD2eF3gH4iJ5kL6mN7oP8qR9sT0uV1wX2yZ3aB4cD5eF6gH7iJ8kL9mN0oP1qR2sT3uV4wX5yZ6aB7cD8eF9gH0iJ1kL2mN3oP4qR5sT6uV7wX8yZ7aB0cD1eF2gH3iJ4kL5mN6oP7qR8sT9uV8wX9yZ8 test@example.com").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertEqual(t, response.Data.ValidatedBy, testutils.PredefinedUserGivenName+" "+testutils.PredefinedUserFamilyName)
}
