package fidelitylink_test

import (
	"testing"
	"time"

	fidelitylink "worker-demo/fidelity_link"

	"github.com/stretchr/testify/assert"
)

// Para correto funcionamento do teste, as constantes config.*_FILE devem ser alteradas para o caminho relativo à pasta 'test'
func TestGetDocumentByUserID_Success(t *testing.T) {

	document, err := fidelitylink.GetDocumentByUserID(validUserID)

	assert.NoError(t, err, "O erro deveria ser nil")
	assert.NotEmpty(t, document, "O documento não deveria ser nil")
	assert.NotContains(t, document, []string{".", "-"}, "o documento retornado deveria vir formatado")
}

func TestGetDocumentByUserID_Error(t *testing.T) {

	testCases := []struct {
		name   string
		userID int64
	}{
		{"Error_InvalidUserID", invalidID},
		{"Error_NotFoundUserID", 50},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			document, err := fidelitylink.GetDocumentByUserID(tc.userID)

			assert.Error(t, err, "O erro não deveria ser nil")
			assert.Empty(t, document, "O documento deveria ser vazio")
		})
	}
}

func TestGetPointsByParameterName_Success(t *testing.T) {
	p, err := fidelitylink.GetPointsByParameterName(fidelitylink.PointsParameterName)

	assert.NoError(t, err, "O erro deveria ser nil")
	assert.NotEmpty(t, p, "Points não deveria ser nil")
}

func TestGetPointsByParameterName_Error(t *testing.T) {
	testCases := []struct {
		name        string
		parameterID string
	}{
		{"Error_InvalidParameterID", ""},
		{"Error_NotFoundParameterID", "notFound"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := fidelitylink.GetPointsByParameterName(tc.parameterID)

			assert.Error(t, err, "O erro não deveria ser nil")
			assert.Zero(t, p, "Points deveria ser zero")
		})
	}
}

// SetCreditedPointsByReferralID é executada duas vezes para que a indicação registrada com os pontos possa ser executada com sucesso em TestSetCreditedPointsByReferralIDDBTx_Success
func TestSetCreditedPointsByReferralID_Success(t *testing.T) {
	assert.NoError(t, fidelitylink.SetCreditedPointsByReferralID(validReferralID, points), "O erro deveria ser nil")
	assert.NoError(t, fidelitylink.SetCreditedPointsByReferralID(validReferralID, 0), "O erro deveria ser nil")
}

func TestSetCreditedPointsByReferralID_Error(t *testing.T) {
	assert.Error(t, fidelitylink.SetCreditedPointsByReferralID(invalidID, points), "O erro não deveria ser nil")
}

// HandleErrorLog é um método do package processors e está sendo utilizado para chamar o método InserErrorLog do package fidelitylink
func TestHandleErrorLog_Success(t *testing.T) {
	assert.EqualError(t, fidelitylink.HandleErrorLog(validReferralID, "teste > mensagem de erro"), "teste > mensagem de erro")
}

func TestHandleErrorLog_Error(t *testing.T) {
	testCases := []struct {
		name          string
		referralID    int64
		errorLog      string
		expectedError string
	}{
		{"Error_InvalidReferralID", invalidID, "teste > mensagem de erro", "o id da indicação ou a mensagem de erro não foram definidos"},
		{"Error_NotSetErrMsg", validReferralID, "", "o id da indicação ou a mensagem de erro não foram definidos"},
		{"Error_NotFoundReferralID", 2, "teste > mensagem de erro", "teste > mensagem de erro"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.ErrorContains(t, fidelitylink.HandleErrorLog(tc.referralID, tc.errorLog), tc.expectedError)
		})
	}
}

// TestGetDateOneYearFromNow verifica se a função GetDateOneYearFromNow retorna a data correta no formato esperado.
func TestGetDateOneYearFromNow(t *testing.T) {
	currentTime := time.Now().UTC()
	expectedTime := currentTime.AddDate(1, 0, 0)
	expected := expectedTime.Format(time.RFC3339)

	result := fidelitylink.GetDateOneYearFromNow()

	assert.Equal(t, expected, result, "O resultado de GetDateOneYearFromNow() deveria ser igual à variavel expected")
}
