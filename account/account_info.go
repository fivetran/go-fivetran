package account

import (
	"context"

	"github.com/fivetran/go-fivetran/common"
	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type AccountInfoService struct {
	httputils.HttpService
}

type AccountInfoResponse struct {
	common.CommonResponse
	Data struct {
		AccountId   string `json:"account_id"`
		AccountName string `json:"account_name"`
		UserId      string `json:"user_id"`
		SystemKeyId string `json:"system_key_id"`
	}
}

func (s *AccountInfoService) Do(ctx context.Context) (AccountInfoResponse, error) {
	var response AccountInfoResponse
	url := "/account/info"
	err := s.HttpService.Do(ctx, "GET", url, nil, nil, 200, &response)
	return response, err
}
