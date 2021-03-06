package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation   = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase     = &Errno{Code: 20002, Message: "Database error."}
	ErrToken        = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrUnauthorized = &Errno{Code: 20004, Message: "Error Unauthorized."}

	// user errors
	ErrEncrypt            = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound       = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid       = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect  = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrUsernameExist      = &Errno{Code: 20105, Message: "The username was existed."}
	ErrDBArgumentInvalid  = &Errno{Code: 20106, Message: "The argument is invalid in DB layer."}
	ErrServiceArgInvalid  = &Errno{Code: 20107, Message: "The argument is invalid in Service layer."}
	ErrApiArgumentInvalid = &Errno{Code: 20108, Message: "The argument is invalid in API layer."}
	ErrUserInvalid        = &Errno{Code: 20109, Message: "The user status was invalid."}
	ErrQueryValueInvalid  = &Errno{Code: 20110, Message: "The query value in url was invalid."}
)
