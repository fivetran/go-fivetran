package fivetran

import "encoding/json"

type nullableString struct {
	value *string
}

func newNullableString(s *string, clear bool) *nullableString {
	if s == nil && !clear {
		return nil
	}

	return &nullableString{
		value: s,
	}
}

func (n *nullableString) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.value)
}
