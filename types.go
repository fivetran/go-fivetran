package fivetran

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
