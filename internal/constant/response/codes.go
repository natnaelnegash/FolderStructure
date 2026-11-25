package response

type Code string

var (
	UNAUTHORIZED                   Code = "UNAUTHORIZED"
	FORBIDDEN                      Code = "FORBIDDEN"
	CONCURRENT_SESSION_NOT_ALLOWED Code = "CONCURRENT_SESSION_NOT_ALLOWED"
	USER_CREATED                   Code = "USER_CREATED"
)
