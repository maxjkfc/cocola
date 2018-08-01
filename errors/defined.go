package errors

// Default Error Code

var (
	NotError       = New(0, "Success")
	ErrorConfigSet = New(2, "Config Seting Failed")

	// JWT
	ErrorJwtCreateFailedForType = New(102, "Create JWT Token Failed (WrongType)")
	ErrorJwtCreateTokenFailed   = New(103, "Create the Token Failed")
	ErrorJwtStoreTokenFailed    = New(104, "Store the UserToken Failed")
	ErrorJwtSigningMethod       = New(106, "Error SigningMethod")
	ErrorJwtWrongIssuer         = New(107, "Wrong Jwt Issuer")
	ErrorJwtValidFailed         = New(108, "Jwt Validate Failed")
	ErrorJwtTokenNotFound       = New(109, "Jwt Token Not Found")

	// Login
	ErrorAuthLoginFailed  = New(110, "User account or password is incorrect")
	ErrorAuthLogoutFailed = New(111, "User Logout Failed")
	// Database Connect
	ErrorDataBaseType  = New(200, "Error Database Connect Type")
	ErrorConnectFailed = New(201, "Connect Failed.")
	// Database
	ErrorPoolNotFound = New(210, "Not Found the DB in the Pool With the Tag Name")
	ErrorCreateFailed = New(211, "Create Failed")
	ErrorUpdateFailed = New(212, "Update Data Failed")
	ErrorDeleteFailed = New(214, "Delete Data Failed")
	//
	ErrorNotFound = New(220, "Not found this ID")
)
