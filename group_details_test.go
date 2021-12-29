package fivetran_test

import (
	"context"
	"testing"
)

func TestNewGroupDetailsIntegration(t *testing.T) {
	for version, c := range Clients {
		t.Run(version, func(t *testing.T) {
			result, err := c.NewGroupDetails().GroupID("climbed_consulted").Do(context.Background())
			if err != nil {
				t.Logf("%+v\n", result)
				t.Error(err)
			}

			AssertEqual(t, result.Code, "Success")
			AssertEqual(t, result.Data.ID, "climbed_consulted")
			AssertEqual(t, result.Data.Name, "Warehouse")
			AssertEqual(t, result.Data.CreatedAt.IsZero(), false)
		})
	}
}
