package fidelitylink

const (
	packageName         string = "fidelitylink"
	gatewayCategory     string = "fidelidade"
	gatewayType         string = "gateway"
	PointsParameterName string = "FidelidadeEmbaixadorPontosPorIndicado"
)

type AmbassadorReferralLog struct {
	AmbassadorReferralID int64  `db:"FidelidadeEmbaixadorIndicado_Id"`
	RequestType          string `db:"TipoRequest"`
	Request              string `db:"Request"`
	Response             string `db:"Response"`
}

type Gateway struct {
	ID          int64  `json:"id" db:"gateway_id"`
	Category    string `json:"category,omitempty" db:"gateway_category" validate:"required{POST|PUT}" errmsg:"Informe a categoria do gateway"`
	Description string `json:"description,omitempty" db:"gateway_description" validate:"required{POST|PUT}" errmsg:"Informe a descrição do gateway"`
	Host        string `json:"host,omitempty" db:"gateway_host" validate:"required{POST|PUT}" errmsg:"Informe o host do gateway"`
	User        string `json:"user,omitempty" db:"gateway_user" validate:"({{token}}=='')required{POST|PUT}" errmsg:"Informe o usuário do gateway"`
	Password    string `json:"password,omitempty" db:"gateway_password" validate:"({{token}}=='')required{POST|PUT}" errmsg:"Informe a senha do gateway" `
	Token       string `json:"token,omitempty" db:"gateway_token"`
	Type        string `json:"type,omitempty" db:"gateway_type" validate:"required{POST|PUT}" errmsg:"Informe o tipo do gateway"`
	Version     string `json:"version,omitempty" db:"gateway_version" validate:"required{POST|PUT}" errmsg:"Informe a versão do gateway"`
	Active      bool   `json:"active" db:"gateway_active" validate:"required{POST|PUT}" errmsg:"Informe o status do gateway"`
}
