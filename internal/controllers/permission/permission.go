package permission

const user = "user_token"
const admin = "admin_token"

func IsVadlidToken(token string) bool {
	return token == user || token == admin
}

func IsAdmin(token string) bool {
	return token == admin
}
