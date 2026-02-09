package fivetran

import (
    "context"
    "testing"

    testutils "github.com/fivetran/go-fivetran/test_utils"
)

func TestNewConnectCardE2E(t *testing.T) {

    accountInfo, err := testutils.Client.AccountInfo().
        Do(context.Background())

    if err != nil {
        t.Logf("%+v\n", accountInfo)
        t.Error(err)
    }

    testutils.AssertEqual(t, accountInfo.Code, "Success")
    testutils.AssertEqual(t, accountInfo.Message, "Account information retrieved successfully")
    testutils.AssertNotEmpty(t, accountInfo.Data.AccountId)
    testutils.AssertNotEmpty(t, accountInfo.Data.AccountName)
    testutils.AssertEqual(t, accountInfo.Data.UserId, testutils.PredefinedUserId)
	testutils.AssertEmpty(t, accountInfo.Data.SystemKeyId)
}