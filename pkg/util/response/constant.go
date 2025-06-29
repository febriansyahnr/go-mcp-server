package response

const httpCodeService string = "SNP-CR-"

const (
	HttpStatusOK      string = "00"
	HttpStatusCreated string = "01"

	HttpStatusErrorInternal        string = "99"
	HttpStatusErrorDatabase        string = "98"
	HttpStatusErrorRequest         string = "40"
	HttpStatusErrorUnauthorized    string = "41"
	HttpStatusErrorForbidden       string = "43"
	HttpStatusErrorNotFound        string = "44"
	HttpStatusErrorLimiter         string = "48"
	HttpStatusErrorDuplicatedCheck string = "49"
	HttpStatusErrorThirdParty      string = "50"
	HttpStatusErrorBadGateway      string = "51"
)

const (
	HttpErrInternal     string = "ERROR_INTERNAL"
	HttpErrDatabase     string = "ERROR_DATABASE"
	HttpErrThirdParty   string = "ERROR_THIRD_PARTY"
	HttpErrRequest      string = "ERROR_REQUEST"
	HttpErrUnauthorized string = "ERROR_UNAUTHORIZED"
	HttpErrNotFound     string = "ERROR_NOT_FOUND"
	HttpErrDupCheck     string = "ERROR_DUPLICATE_CHECK"
)

func GetHttpCodeService(code string) string {
	return httpCodeService + code
}

func GetHttpCodeServiceError(code string, errCode string) string {
	if errCode != "" {
		return "SNP_" + errCode
	}
	return "SNP_" + code
}
