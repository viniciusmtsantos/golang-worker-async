package taskflow

import "github.com/jmoiron/sqlx"

func GetDocumentByUserID(userID int64) (string, error) {
	return "", nil
}

func GetPointsByParameterName(parameterName string) (int64, error) {
	return 0, nil
}

func CreditPointsToReferrer(document string, points int64) error {
	return nil
}

func SetCreditedPointsByReferralID(referralID, points int64) error {
	return nil
}

func GetClientApi() error {
	return nil
}

func GetDateOneYearFromNow() string {
	return ""
}

func HandleErrorLog(referralID int64, errMessage string) error {
	return nil
}

func insertErrorLogTx(errorLogInfo AmbassadorReferralLog, tx *sqlx.Tx) error {
	return nil
}

func findGatewayByFilterTx(filterCategory string, filterType string, tx *sqlx.Tx) ([]Gateway, error) {
	return []Gateway{}, nil
}

func setCreditedPointsByReferralIDTx(referralID, points int64, tx *sqlx.Tx) error {
	return nil
}

func FindPersonIDByReferralIDTx(referralID int64, tx *sqlx.Tx) (personID int64, err error) {
	return 0, nil
}

func FindPersonIDByReferralID(referralID int64) (personID int64, err error) {
	return 0, nil
}
