package security

type TokenGenerator interface {
	Generate(userID string) (string, error)
	Validate(token string) (string, error)
}
