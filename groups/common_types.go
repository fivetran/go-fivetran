package groups

import (
	"time"

	"github.com/fivetran/go-fivetran/common"
)

type GroupItem struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type GroupDetailsResponse struct {
	common.CommonResponse
	Data GroupItem `json:"data"`
}
