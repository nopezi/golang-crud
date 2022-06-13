package user

type Login struct {
	ClientID     string `json:"clientid"`
	ClientSecret string `json:"clientsecret"`
}

type UserSession struct {
	PERNR      string `json:"PERNR"`
	NIP        string `json:"NIP"`
	SNAME      string `json:"SNAME"`
	WERKS      string `json:"WERKS"`
	BTRTL      string `json:"BTRTL"`
	KOSTL      string `json:"KOSTL"`
	ORGEH      string `json:"ORGEH"`
	STELL      string `json:"STELL"`
	WERKSTX    string `json:"WERKS_TX"`
	BTRTLTX    string `json:"BTRTL_TX"`
	KOSTLTX    string `json:"KOSTL_TX"`
	ORGEHTX    string `json:"ORGEH_TX"`
	STELLTX    string `json:"STELL_TX"`
	PLANSTX    string `json:"PLANS_TX"`
	JGPG       string `json:"JGPG"`
	ORGEHPGS   string `json:"ORGEH_PGS"`
	PLANSPGS   string `json:"PLANS_PGS"`
	ORGEHPGSTX string `json:"ORGEH_PGS_TX"`
	PLANSPGSTX string `json:"PLANS_PGS_TX"`
	SISACT     string `json:"SISA_CT"`
	SISACB     string `json:"SISA_CB"`
	AGAMA      string `json:"AGAMA"`
	TIPEUKER   string `json:"TIPE_UKER"`
	ADDAREA    string `json:"ADD_AREA"`
	PERSG      string `json:"PERSG"`
	PERSK      string `json:"PERSK"`
	STATUS     string `json:"STATUS"`
	BRANCH     string `json:"BRANCH"`
	HILFM      string `json:"HILFM"`
	HTEXT      string `json:"HTEXT"`
	HILFMPGS   string `json:"HILFM_PGS"`
	HTEXTPGS   string `json:"HTEXT_PGS"`
	KAWIN      string `json:"KAWIN"`
	WERKSPGS   string `json:"WERKS_PGS"`
	BTRTLPGS   string `json:"BTRTL_PGS"`
	KOSTLPGS   string `json:"KOSTL_PGS"`
}
