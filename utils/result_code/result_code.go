package resultcode

const (
	ErrorURLIsRequired = "1000"

	//db
	ErrorDBConnection  = "2000"
	ErrorDBKeyNotExist = "2001"
)

var ResultCodeMap = map[string]string{
	ErrorURLIsRequired: "URL is required",
	ErrorDBConnection:  "DB connection Error",
	ErrorDBKeyNotExist: "DB Key Not Exist",
}
