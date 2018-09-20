package model

type Contato struct {
	Id        int    `json:"id"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
