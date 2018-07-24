package errors

// Default Error Code

var (
	ErrorConfigSet = New(102, "Config Seting Failed")

	ErrorConnectFailed = New(201, "Connect Failed.")

	// Database
	ErrorPoolNotFound = New(210, "Not Find the DB in the Pool With the Tag Name")
)
