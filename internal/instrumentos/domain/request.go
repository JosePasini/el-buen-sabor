package domain

type Login struct {
	Usuario string `json:"usuario"`
	Hash    string `json:"hash"`
}
