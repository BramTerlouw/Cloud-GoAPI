// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Class struct {
	ID          string        `json:"id"`
	ModuleID    string        `json:"module_Id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Difficulty  LanguageLevel `json:"difficulty"`
	MadeBy      string        `json:"made_by"`
	OpenaiKey   string        `json:"openai_key"`
	CreatedAt   *string       `json:"created_at,omitempty"`
	UpdatedAt   *string       `json:"updated_at,omitempty"`
	SoftDeleted *bool         `json:"soft_deleted,omitempty"`
}

type ClassInfo struct {
	ID          string        `json:"id"`
	ModuleID    string        `json:"module_Id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Difficulty  LanguageLevel `json:"difficulty"`
	MadeBy      string        `json:"made_by"`
}

type ClassInput struct {
	ModuleID    string        `json:"module_Id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Difficulty  LanguageLevel `json:"difficulty"`
}

type ListClassFilter struct {
	SoftDelete *bool          `json:"softDelete,omitempty"`
	ModuleID   *string        `json:"module_id,omitempty"`
	Name       *NameFilter    `json:"name,omitempty"`
	Difficulty *LanguageLevel `json:"difficulty,omitempty"`
	MadeBy     *string        `json:"made_by,omitempty"`
}

type NameFilter struct {
	Input []*string       `json:"input"`
	Type  NameFilterTypes `json:"type"`
}

type Paginator struct {
	Amount int `json:"amount"`
	Step   int `json:"Step"`
}

type LanguageLevel string

const (
	LanguageLevelA1 LanguageLevel = "A1"
	LanguageLevelA2 LanguageLevel = "A2"
	LanguageLevelB1 LanguageLevel = "B1"
	LanguageLevelB2 LanguageLevel = "B2"
	LanguageLevelC1 LanguageLevel = "C1"
	LanguageLevelC2 LanguageLevel = "C2"
)

var AllLanguageLevel = []LanguageLevel{
	LanguageLevelA1,
	LanguageLevelA2,
	LanguageLevelB1,
	LanguageLevelB2,
	LanguageLevelC1,
	LanguageLevelC2,
}

func (e LanguageLevel) IsValid() bool {
	switch e {
	case LanguageLevelA1, LanguageLevelA2, LanguageLevelB1, LanguageLevelB2, LanguageLevelC1, LanguageLevelC2:
		return true
	}
	return false
}

func (e LanguageLevel) String() string {
	return string(e)
}

func (e *LanguageLevel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LanguageLevel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LanguageLevel", str)
	}
	return nil
}

func (e LanguageLevel) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
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
