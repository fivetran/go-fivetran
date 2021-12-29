package fivetran_test

import (
	"context"
	"testing"
)

func TestNewDestinationSetupTestsIntegration(t *testing.T) {
	for version, c := range Clients {
		//todo: remove it after v2 port fixes
		if version == "v2" {
			continue
		}
		t.Run(version, func(t *testing.T) {
			destinationId := CreateTempDestination(t)

			response, err := c.NewDestinationSetupTests().DestinationID(destinationId).
				TrustCertificates(true).
				TrustFingerprints(true).
				Do(context.Background())

			if err != nil {
				t.Logf("%+v\n", response)
				t.Error(err)
			}

			AssertEqual(t, response.Code, "Success")
			AssertEqual(t, response.Message, "Setup tests have been completed")
			AssertEqual(t, response.Data.ID, destinationId)
			AssertEqual(t, response.Data.GroupID, destinationId)
			AssertEqual(t, response.Data.Service, "snowflake")
			AssertEqual(t, response.Data.Region, "US")
			AssertEqual(t, response.Data.TimeZoneOffset, "+10")
			AssertEqual(t, response.Data.SetupStatus, "incomplete")
		})
	}
}
