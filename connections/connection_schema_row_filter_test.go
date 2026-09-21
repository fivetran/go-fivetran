package connections

import (
	"encoding/json"
	"testing"
)

func TestSchemaTableRowFilterRead(t *testing.T) {
	var table ConnectionSchemaConfigTableResponse
	raw := `{"row_filter":{"name":"test","description":"test","column_clauses":[]}}`
	if err := json.Unmarshal([]byte(raw), &table); err != nil {
		t.Fatal(err)
	}
	// Assert read support without depending on a new field at compile time.
	encoded, err := json.Marshal(table)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err = json.Unmarshal(encoded, &fields); err != nil {
		t.Fatal(err)
	}
	if string(fields["row_filter"]) != `{"name":"test","description":"test","column_clauses":[]}` {
		t.Fatalf("row_filter lost: %s", encoded)
	}
}

func TestSchemaTableRowFilterRequest(t *testing.T) {
	for _, tc := range []struct {
		name   string
		filter json.RawMessage
		want   string
	}{
		{"omitted", nil, `{}`},
		{"delete", json.RawMessage(`null`), `{"row_filter":null}`},
		{"object", json.RawMessage(`{"name":"test"}`), `{"row_filter":{"name":"test"}}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			table := new(ConnectionSchemaConfigTable).RowFilter(tc.filter)
			encoded, err := json.Marshal(table.Request())
			if err != nil {
				t.Fatal(err)
			}
			if string(encoded) != tc.want {
				t.Fatalf("got %s, want %s", encoded, tc.want)
			}
		})
	}
}
