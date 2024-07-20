package dto

type CredentialsRegisteredDTO struct {
	Lenght  int      `json:"lenght"`
	Content []string `json:"content"`
	ReadAt  string   `json:"readtAt"`
}
