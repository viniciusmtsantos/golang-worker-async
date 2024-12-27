package fidelitylink_test

import (
	"testing"

	fidelitylink "worker-demo/fidelity_link"

	"github.com/stretchr/testify/assert"
)

const (
	invalidID       int64  = 0
	validReferralID int64  = 1
	points          int64  = 10
	validUserID     int64  = 99999
	validDocument   string = "35555818807"
	invalidDocument string = ""
)

func TestFidelidadeEmbaixadorIndicadoLog(t *testing.T) {
	errorLog := fidelitylink.AmbassadorReferralLog{
		AmbassadorReferralID: 1,
		RequestType:          "PUT",
		Request:              "payload",
		Response:             "error",
	}

	assert.Equal(t, int64(1), errorLog.AmbassadorReferralID, "FidelidadeEmbaixadorIndicadoID deveria ser 1")
	assert.Equal(t, "PUT", errorLog.RequestType, "TipoRequest deveria ser PUT")
	assert.Equal(t, "payload", errorLog.Request, "Request deveria ser payload")
	assert.Equal(t, "error", errorLog.Response, "Response deveria ser error")

	assert.NotEqual(t, int64(2), errorLog.AmbassadorReferralID, "FidelidadeEmbaixadorIndicadoID deveria ser 2")
	assert.NotEmpty(t, errorLog.RequestType, "TipoRequest não deveria ser vazio")
}
