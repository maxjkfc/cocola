package errors

// Default Error Code

var (
	NotError = New(0, "Success")

	ErrorConfigSet = New(102, "Config Seting Failed")

	ErrorDataBaseType = New(200, "Error Database Connect Type")

	ErrorConnectFailed = New(201, "Connect Failed.")

	// Database

	ErrorPoolNotFound = New(210, "Not Find the DB in the Pool With the Tag Name")

	ErrorCreateFailed = New(211, "Create Failed")
)
