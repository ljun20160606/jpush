package v3

type (
	Error struct {
		Code    int    `db:"code" json:"code"`
		Message string `db:"message" json:"message"`
	}

	ErrorInfo struct {
		Error Error `db:"error" json:"error"`
	}
)
