package odbc

import (
	"errors"
	"fmt"
	"strings"
)

const (
	PARAM_IDENTIFIER = "?"
)

func customInsertParams(
	sql string,
	parameters []QueryParameter,
) (string, error) {

	countProvided := len(parameters)
	countExpected := strings.Count(sql, PARAM_IDENTIFIER)

	// If no parameters, return just the sql provided
	if countProvided == 0 && countExpected == 0 {
		return sql, nil
	}

	// Check that the number of provided params equals the number of expected params
	if countExpected != countProvided {
		errStr := fmt.Sprintf("expected %d parameters but got %d", countExpected, countProvided)
		return "", errors.New(errStr)
	}

	newSql := sql

	// Loop through the parameters, escape the param, and then make any replacements to the sql
	for _, p := range parameters {
		modP := modifyParam(
			p,
			escapeSingleQuote,
			wrapInSingleQuotes,
		)

		newSql = strings.Replace(newSql, PARAM_IDENTIFIER, fmt.Sprint(modP), 1)
	}

	return newSql, nil
}

func modifyParam(p QueryParameter, modifiers ...func(p *QueryParameter)) QueryParameter {
	param := p

	for _, mod := range modifiers {
		mod(&param)
	}

	return param
}

func escapeSingleQuote(p *QueryParameter) {
	switch (*p).(type) {
	case string:
		*p = strings.ReplaceAll(fmt.Sprint(*p), "'", "''")
	}
	*p = fmt.Sprint(*p)
}

func wrapInSingleQuotes(p *QueryParameter) {
	switch (*p).(type) {
	case string:
		*p = fmt.Sprintf("'%v'", *p)
	}
	*p = fmt.Sprint(*p)
}
