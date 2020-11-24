package e

const (
	SUCCESS = 0

	INVALID_PARAMS                 = 401
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 403
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 404
	ERROR_AUTH_TOKEN               = 405
	ERROR_AUTH                     = 406
	FAILE_TO_GET_OPENID            = 407

	ERROR          = 500
	DATABASE_ERROR = 501
	CACHE_ERROR    = 502
	WS_ERROR       = 503
	// login fail
	ERROR_LOGIN_FAIL   = 504
	FAILE_TO_CREATE_OP = 505
	NOT_FOUND_RECORD   = 506
)
