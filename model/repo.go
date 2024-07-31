package model

type RepoService interface {
	Read() (map[string]string, error)
	ReadOne(email string) (string, error)
	Delete(email string) (int64, error)
	Create(cred Credential) error
}
