package fidelitylink

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func SetCreditedPointsByReferralIDDBTx(referralID, points int64, tx *sqlx.Tx) error {
	query := `UPDATE sge.fidelidadeembaixadorindicado SET Pontos = ? WHERE Id = ?`

	result, err := tx.Exec(query, points, referralID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("nenhum registro de indicação com o ID %d foi alterado", referralID)
	}

	return nil
}

func InsertErrorLogDBTx(errorLogInfo AmbassadorReferralLog, tx *sqlx.Tx) error {
	query := `
		INSERT INTO sge.fidelidadeembaixadorindicadolog (
			FidelidadeEmbaixadorIndicado_Id,
			TipoRequest,
			Request,
			Response
		)
		VALUES (
			?,
			?,
			?,
			?
		)
	`

	args := []interface{}{
		errorLogInfo.AmbassadorReferralID,
		errorLogInfo.RequestType,
		errorLogInfo.Request,
		errorLogInfo.Response,
	}

	_, err := tx.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func findGatewayByFilterDBTx(filterCategory, filterType string, tx *sqlx.Tx) (gateways []Gateway, err error) {
	query := `
		SELECT
			gateway.id 								gateway_id,
			gateway.category 						gateway_category,
			IFNULL(gateway.description, "") 		gateway_description,
			gateway.host 							gateway_host,
			IFNULL(gateway.user,"") 				gateway_user,
			IFNULL(gateway.password,"") 			gateway_password,
			IFNULL(gateway.token,"") 				gateway_token,
			IFNULL(gateway.type,"") 				gateway_type,
			IFNULL(gateway.version,"") 				gateway_version,
			gateway.active 							gateway_active,
			gateway.created_at 						gateway_created_at,
			gateway.updated_at 						gateway_updated_at
		FROM workspace.gateway 
		LEFT JOIN workspace.gateway_param ON gateway_param.gateway_id = gateway.id
		WHERE TRUE AND gateway.category = ? AND gateway.type = ?
	`

	err = tx.Select(&gateways, query, filterCategory, filterType)
	if err != nil {
		return
	}

	return
}

func findPersonIDByReferralIDDBTx(referralID int64, tx *sqlx.Tx) (personID int64, err error) {
	query := `
		SELECT
			 e.pessoa_id
		FROM sge.fidelidadeembaixadorindicado i
		inner join sge.fidelidadeembaixador e on e.id = i.fidelidadeembaixador_id
		WHERE i.id = ?
	`

	err = tx.Get(&personID, query, referralID)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	err = nil
	return
}
