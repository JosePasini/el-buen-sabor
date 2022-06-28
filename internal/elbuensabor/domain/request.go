package domain

type Login struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}
