package constants

type Sentinel string

func (s Sentinel) Error() string {
	return string(s)
}

const (
	ErrPasswordTooShort      = Sentinel("password is too short")
	ErrPasswordTooLong       = Sentinel("password is too long")
	ErrMustContainDigit      = Sentinel("password must contain at least 1 digit")
	ErrMustContainAlphabetic = Sentinel("password must contain at least 1 alphabetic")
	ErrEmailAddress          = Sentinel("invalid email address. valid e-mail can contain only latin letters, numbers, '@' and '.'")
	ErrInvalidRequestBody    = Sentinel("invalid request body")
	ErrIncorrectNameValue    = Sentinel("name must include minumum 3 and maximum 255 charachters")
	ErrAuthIncorrect         = Sentinel("auth incorrect")
	ErrAuthNotGiven          = Sentinel("auth not given")

	ProdEnvironment    = "prod"
	DevEnvironment     = "dev"
	StagingEnvironment = "staging"

	InvalidToken = Sentinel("invalid jwt token")
)
