package model

type (
	SetOption struct {
		Add    *[]string `db:"add" json:"add,omitempty"`
		Remove *[]string `db:"remove" json:"remove,omitempty"`
	}
)
