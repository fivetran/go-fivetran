package common

import "encoding/json"

type CommonResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SetupTestResponse struct {
	Title   string `json:"title"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Details string `json:"details"`
}

type NullableString struct {
	value *string
}

func (n *NullableString) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.value)
}

func NewNullableString(s *string, clear bool) *NullableString {
	if s == nil && !clear {
		return nil
	}

	return &NullableString{
		value: s,
	}
}
