package common

type CommonResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SetupTestResponse struct {
	Title   string `json:"title"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type NullableString struct {
	value *string
}

func NewNullableString(s *string, clear bool) *NullableString {
	if s == nil && !clear {
		return nil
	}

	return &NullableString{
		value: s,
	}
}
