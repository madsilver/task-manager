package presenter

type Task struct {
	ID      uint64 `json:"id"`
	UserID  uint64 `json:"user_id"`
	Summary string `json:"summary"`
	Date    string `json:"date"`
}
