package types

const UserContextKey = "user"

type AuthenticatedUser struct {
	Email    string
	LoggedIn bool
}
