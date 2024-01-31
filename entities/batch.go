package entities

type Batch struct {
	ID        string `json:"id" db:"id"`
	Name      string `json:"name" db:"name" validate:"required"`
	UserID    string `json:"user_id" db:"user_id"`
	AmountQrs int    `json:"amount_qrs" db:"amount_qrs"`
}
