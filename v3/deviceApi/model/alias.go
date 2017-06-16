package model

type AliasDevices struct {
	RegistrationIds []string `db:"registration_ids" json:"registration_ids"`
}
