// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewCategory struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type NewCourse struct {
	CategoryID  string  `json:"categoryId"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
