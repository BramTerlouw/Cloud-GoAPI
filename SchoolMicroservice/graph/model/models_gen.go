// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type ListSchoolFilter struct {
	SoftDelete *bool           `json:"softDelete,omitempty"`
	Name       *NameFilter     `json:"name,omitempty"`
	Location   *LocationFilter `json:"location,omitempty"`
	MadeBy     *string         `json:"made_by,omitempty"`
}

type LocationFilter struct {
	Input []*string       `json:"input"`
	Type  NameFilterTypes `json:"type"`
}

type NameFilter struct {
	Input []*string       `json:"input"`
	Type  NameFilterTypes `json:"type"`
}

type Paginator struct {
	Amount int `json:"amount"`
	Step   int `json:"Step"`
}

type School struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Location    string  `json:"location"`
	MadeBy      string  `json:"made_by"`
	CreatedAt   *string `json:"created_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
	SoftDeleted *bool   `json:"soft_deleted,omitempty"`
}

type SchoolInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	MadeBy   string `json:"made_by"`
}

type SchoolInput struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

type NameFilterTypes string

const (
	NameFilterTypesEq     NameFilterTypes = "eq"
	NameFilterTypesNe     NameFilterTypes = "ne"
	NameFilterTypesStarts NameFilterTypes = "starts"
	NameFilterTypesEnds   NameFilterTypes = "ends"
)

var AllNameFilterTypes = []NameFilterTypes{
	NameFilterTypesEq,
	NameFilterTypesNe,
	NameFilterTypesStarts,
	NameFilterTypesEnds,
}

func (e NameFilterTypes) IsValid() bool {
	switch e {
	case NameFilterTypesEq, NameFilterTypesNe, NameFilterTypesStarts, NameFilterTypesEnds:
		return true
	}
	return false
}

func (e NameFilterTypes) String() string {
	return string(e)
}

func (e *NameFilterTypes) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NameFilterTypes(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NameFilterTypes", str)
	}
	return nil
}

func (e NameFilterTypes) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
