package models

import "fmt"

type FilterOpr string

const (
	OprEq       FilterOpr = "="
	OprNe                 = "<>"
	OprIn                 = "IN"
	OprNin                = "NOT IN"
	OprLt                 = "<"
	OprLte                = "<="
	OprGt                 = ">"
	OprGte                = ">="
	OprContains           = "LIKE"
)

type FilterModel struct {
	Field    string      `json:"field"`
	Operator FilterOpr   `json:"operator"`
	Value    interface{} `json:"value"`
}

func NewFilterModel(field string, operator FilterOpr, value interface{}) *FilterModel {
	m := &FilterModel{
		Field:    field,
		Operator: operator,
		Value:    value,
	}

	return m
}

func (m *FilterModel) BuildClause() (clause string, value interface{}) {
	clause = fmt.Sprintf("%s %s ?", m.Field, m.Operator)

	if m.Operator == OprContains {
		return clause, fmt.Sprintf("%s%v%s", "%", m.Value, "%")
	}

	return clause, m.Value
}
