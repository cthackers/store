package auth

type Authenticator interface{}

func New() Authenticator {
	return nil
}
