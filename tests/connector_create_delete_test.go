package tests

import (
	"context"
	"testing"

	"github.com/fivetran/go-fivetran"
	"github.com/corbym/gocrest/has"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestConnectorCreateDelete(t *testing.T) {
	for _, c := range GetClients() {
		created, err := c.NewConnectorCreate().GroupID("_moonbeam").Service("itunes_connect").RunSetupTests(false).
		Config(fivetran.NewConnectorConfig().
			Schema("itunes_e2e_connect").
			Username("fivetran").
			Password("fivetran-api-e2e")).
		Do(context.Background())

		if err != nil {
			t.Logf("%+v\n", created)
			t.Error(err)
		}

		then.AssertThat(t, created.Code, is.EqualTo("Success"))
		then.AssertThat(t, created.Message, is.EqualTo("Connector has been created"))
		then.AssertThat(t, created.Data.ID, is.Not(is.Empty()))
		then.AssertThat(t, created.Data.GroupID, is.EqualTo("_moonbeam"))
		then.AssertThat(t, created.Data.Service, is.EqualTo("itunes_connect"))
		then.AssertThat(t, *created.Data.ServiceVersion, is.EqualTo(1))
		then.AssertThat(t, created.Data.Schema, is.EqualTo("itunes_e2e_connect"))
		then.AssertThat(t, created.Data.ConnectedBy, is.EqualTo("_accountworthy"))
		then.AssertThat(t, created.Data.CreatedAt.IsZero(), is.False())
		then.AssertThat(t, created.Data.SucceededAt.IsZero(), is.True())
		then.AssertThat(t, created.Data.FailedAt.IsZero(), is.True())
		then.AssertThat(t, *created.Data.Paused, is.False())
		then.AssertThat(t, *created.Data.PauseAfterTrial, is.False())
		then.AssertThat(t, *created.Data.SyncFrequency, is.EqualTo(360))
		then.AssertThat(t, created.Data.ScheduleType, is.EqualTo("auto"))

		then.AssertThat(t, created.Data.Status.SetupState, is.EqualTo("incomplete"))
		then.AssertThat(t, created.Data.Status.SyncState, is.EqualTo("scheduled"))
		then.AssertThat(t, created.Data.Status.UpdateState, is.EqualTo("on_schedule"))
		then.AssertThat(t, *created.Data.Status.IsHistoricalSync, is.True())
		then.AssertThat(t, created.Data.Status.Tasks, has.Length(0))
		then.AssertThat(t, created.Data.Status.Warnings, has.Length(0))

		then.AssertThat(t, created.Data.Config.Password, is.EqualTo("******"))
		then.AssertThat(t, created.Data.Config.TimeframeMonths, is.EqualTo("TWELVE"))
		then.AssertThat(t, created.Data.Config.AppSyncMode, is.EqualTo("AllApps"))
		then.AssertThat(t, created.Data.Config.SalesAccountSyncMode, is.EqualTo("AllSalesAccounts"))
		then.AssertThat(t, created.Data.Config.FinanceAccountSyncMode, is.EqualTo("AllFinanceAccounts"))
		then.AssertThat(t, created.Data.Config.Username, is.EqualTo("fivetran"))
		//todo:uncomment after fix T-168279
		//then.AssertThat(t, created.Data.Config.AccountSyncMode, is.EqualTo("AllAccounts"))

		deleted, err := c.NewConnectorDelete().ConnectorID(created.Data.ID).Do(context.Background())

		if err != nil {
			t.Logf("%+v\n", deleted)
			t.Error(err)
		}

		then.AssertThat(t, deleted.Code, is.EqualTo("Success"))
		then.AssertThat(t, deleted.Message, is.EqualTo("Connector with id '" + created.Data.ID + "' has been deleted"))
	}
}