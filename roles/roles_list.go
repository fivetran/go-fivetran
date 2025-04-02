package roles

import (
	"context"
	"fmt"

	httputils "github.com/fivetran/go-fivetran/http_utils"
)

type RolesListService struct {
	httputils.HttpService
	limit  *int
	cursor *string
}

type RolesListResponse struct {
	Code string `json:"code"`
	Data struct {
		Items []struct {
			Name        		string   `json:"name"`
			Description 		string   `json:"description"`
			IsCustom    		*bool    `json:"is_custom"`
			Scope       		[]string `json:"scope"`
			IsDeprecated    	*bool    `json:"is_deprecated"`
			ReplacementRoleName string   `json:"replacement_role_name"`
		} `json:"items"`
		NextCursor string `json:"next_cursor"`
	} `json:"data"`
}

func (s *RolesListService) Limit(value int) *RolesListService {
	s.limit = &value
	return s
}

func (s *RolesListService) Cursor(value string) *RolesListService {
	s.cursor = &value
	return s
}

func (s *RolesListService) Do(ctx context.Context) (RolesListResponse, error) {
	var response RolesListResponse
	var queries map[string]string = nil
	if s.cursor != nil || s.limit != nil {
		queries = make(map[string]string)
		if s.cursor != nil {
			queries["cursor"] = *s.cursor
		}
		if s.limit != nil {
			queries["limit"] = fmt.Sprintf("%v", *s.limit)
		}
	}
	err := s.HttpService.Do(ctx, "GET", "/roles", nil, queries, 200, &response)
	return response, err
}
