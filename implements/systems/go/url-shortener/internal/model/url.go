package model

type URL struct {
	ID       int64  `json:"id"`
	Original string `json:"original"`
	Short    string `json:"short"`
}
