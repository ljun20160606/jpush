package model

type (
	Tags struct {
		Tags []string `db:"tags" json:"tags"`
	}

	TagOption struct {
		SetOption `db:"registration_ids" json:"registration_ids"` // add/remove最多各支持1000个
	}
)
