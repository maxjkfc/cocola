package errors

// Default Error Code

var (
	NotError                    = New(0, "Success")
	ErrorConfigSet              = New(2, "Config seting failed")
	ErrorJwtCreateFailedForType = New(102, "Create jwt token failed (WrongType)")
	ErrorJwtCreateTokenFailed   = New(103, "Create the token failed")
	ErrorJwtStoreTokenFailed    = New(104, "Store the user token failed")
	ErrorJwtSigningMethod       = New(106, "Error signing method")
	ErrorJwtWrongIssuer         = New(107, "Wrong jwt issuer")
	ErrorJwtValidFailed         = New(108, "Jwt validate failed")
	ErrorJwtTokenNotFound       = New(109, "Jwt token not found")
	ErrorDataBaseType           = New(200, "Error database connect type")
	ErrorConnectFailed          = New(201, "Connect failed.")
	ErrorPoolNotFound           = New(210, "Not found the db in the pool with the tag name")
	ErrorCreateFailed           = New(211, "Create failed")
	ErrorUpdateFailed           = New(212, "Update data failed")
	ErrorDeleteFailed           = New(214, "Delete data failed")
	ErrorDataNotFound           = New(220, "Not found the data")
	ErrorDataIDFailed           = New(221, "Wrong data ID type")
	ErrorRequestBindFailed      = New(300, "Request bind failed")
	ErrorRequestValidFailed     = New(301, "Request valid failed")
	ErrorOperationFailed        = New(310, "Operation failed")
	ErrorIllegalOperation       = New(311, "Illegal operation")
)
