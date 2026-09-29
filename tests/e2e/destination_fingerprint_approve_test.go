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
		Hash("mYHp22TZYjGxOzvJZouuU9OG5v2iciu3mGc9gIQtU0o").
		PublicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCxYiAoxFhLmEPfzUSOToeDH+C5A33zzLT9CX5HHGZyvngXLEm6p4RcQn4xsenz73OQMCgOpbOAYFVD70YWNa1gMQr30Zm5M8tH5EPVAraT3tSyMKEDEULql74sBdIGlvYyOOjD7bw4b/PUw0hpqoZ6ZvZoa4bS/9GEv0mHELJCtcxpXEBYaThJZfq/KaXEjSNQZ6g30TPx7ZoFQBIDDpVHWoH5JyVE2pzxaqn30CX+LqZmlF97ys1fI71w/a0A2wJ2NRWSLq2oLoFERa+CbZMDAFoIoyIMFPDyBnQgPFo4BgQbcnzE+BUnnYbL5+ONCPf3gVtoIx+X8Ex/IfjieWxZ test@example.com").
		Do(context.Background())

	if err != nil {
		t.Logf("%+v\n", response)
		t.Error(err)
	}

	testutils.AssertEqual(t, response.Code, "Success")
	testutils.AssertNotEmpty(t, response.Message)
	testutils.AssertEqual(t, response.Data.ValidatedBy, testutils.PredefinedUserGivenName+" "+testutils.PredefinedUserFamilyName)
}
