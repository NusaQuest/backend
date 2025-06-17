package constant

type Error string

const (

	FailedToParseData					Error = "Oops! We encountered an issue while processing your data. Please try again!"
	FailedToInsertData      			Error = "Sorry, we couldn’t add the new data. Please try again!"
	FailedToGenerateTokenAccess  		Error = "We couldn’t generate an access token. Please try again!"
	FailedToDecodeData           		Error = "We couldn’t process the data. Please try again!"
	FailedToRetrieveData         		Error = "There was an issue retrieving your data. Please try again!"
	InvalidAccountError          		Error = "The email or password entered is incorrect. Please try again!"
	InvalidTokenError            		Error = "The token provided is invalid. Please try again!"
	PermissionDeniedError        		Error = "Access denied! You don’t have permission to access this data."
	ValidationError              		Error = "Your input doesn’t meet the requirements. Please check and try again!"
	DatabaseNotInitialized				Error = "MongoDB hasn’t been initialized or connected properly. Please check the database setup."
	ErrorWhileConnectToDatabase 		Error = "Failed to connect to the MongoDB database. Please check your connection settings."
	ErrorWhilePingToDatabase    		Error = "Unable to reach the MongoDB database. Please ensure it is running and accessible."
	ErrorWhileDisconnectFromDatabase 	Error = "Failed to disconnect from the MongoDB database. Please check the connection state."

)