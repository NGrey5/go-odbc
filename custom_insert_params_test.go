package odbc

import (
	"strings"
	"testing"
)

func TestCustomInsertParams_NoParams(t *testing.T) {

	sql := `
		SELECT * FROM "Table"
		WHERE "column" = 'value'
	`

	want := sql
	got, _ := customInsertParams(sql, nil)

	if want != got {
		t.Fatalf("Wanted %v, Got %v", want, got)
	}

}

func TestCustomInsertParams_Params(t *testing.T) {
	sql := `
		SELECT * FROM "Table"
		WHERE "column" = ? AND
		"column2" = ?
	`
	params := []QueryParameter{"value", "value2"}

	want := `
		SELECT * FROM "Table"
		WHERE "column" = 'value' AND
		"column2" = 'value2'
	`
	got, err := customInsertParams(sql, params)
	if err != nil {
		t.Fatalf("%s", err)
	}

	if strings.TrimSpace(want) != strings.TrimSpace(got) {
		t.Fatalf("Wanted %v, Got %v", want, got)
	}
}

func TestCustomInsertParams_InvalidParamCount(t *testing.T) {

	sql1 := `SELECT * FROM "Table"`
	params1 := []QueryParameter{"invalidparam"}

	_, err := customInsertParams(sql1, params1)
	if err == nil {
		t.Fatalf("should get invalid params error")
	}

	sql2 := `SELECT * FROM "Table" WHERE "column" = ?`
	_, err = customInsertParams(sql2, nil)
	if err == nil {
		t.Fatalf("should get invalid params error")
	}

}

func TestCustomInsertParams_EscapeSingleQuotes(t *testing.T) {
	sql := `SELECT * FROM "Table" WHERE "column" = ? AND "column2" = ?`
	params := []QueryParameter{"sq'", "sq2'"}

	want := `SELECT * FROM "Table" WHERE "column" = 'sq''' AND "column2" = 'sq2'''`

	got, err := customInsertParams(sql, params)
	if err != nil {
		t.Fatalf("%s", err)
	}

	if want != got {
		t.Fatalf("Want %s, Got %s", want, got)
	}
}
