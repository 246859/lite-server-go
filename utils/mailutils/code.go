package mailutils

const AuthMail = "authMail"
const PasswordMail = "passwordMail"

func RedisMailKey(emil string, key string) string {
	return emil + "|mail|" + key
}
