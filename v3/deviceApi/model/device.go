package model

type (
	SetDeviceOption struct {
		Tags   SetOption `db:"tags" json:"tags"`
		Alias  *string   `db:"alias" json:"alias,omitempty"`
		Mobile *string   `db:"mobile" json:"mobile,omitempty"`
	}

	DeviceInfo struct {
		Tags   []string `db:"tags" json:"tags"`
		Alias  *string  `db:"alias" json:"alias"`
		Mobile *string  `db:"mobile" json:"mobile"`
	}
)
