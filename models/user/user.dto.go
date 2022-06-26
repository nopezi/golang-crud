package user

type LoginRequest struct {
	Pernr    string `json:"pernr"`
	Password string `json:"password"`
}

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

type UserSessionIncognito struct {
	PERNR      string `json:"PERNR"`
	WERKS      string `json:"WERKS"`
	BTRTL      string `json:"BTRTL"`
	KOSTL      string `json:"KOSTL"`
	ORGEH      string `json:"ORGEH"`
	ORGEHPGS   string `json:"ORGEH_PGS"`
	STELL      string `json:"STELL"`
	SNAME      string `json:"SNAME"`
	WERKSTX    string `json:"WERKS_TX"`
	BTRTLTX    string `json:"BTRTL_TX"`
	KOSTLTX    string `json:"KOSTL_TX"`
	ORGEHTX    string `json:"ORGEH_TX"`
	ORGEHPGSTX string `json:"ORGEH_PGS_TX"`
	STELLTX    string `json:"STELL_TX"`
	BRANCH     string `json:"BRANCH"`
	TIPEUKER   string `json:"TIPE_UKER"`
	HILFM      string `json:"HILFM"`
	HILFMPGS   string `json:"HILFM_PGS"`
	HTEXT      string `json:"HTEXT"`
	HTEXTPGS   string `json:"HTEXT_PGS"`
	CORPTITLE  string `json:"CORP_TITLE"`
}

type Menu struct {
	MenuID     int64  `json:"menu_id"`
	Title      string `json:"title"`
	Url        string `json:"url"`
	Deskripsi  string `json:"deskripsi"`
	Icon       string `json:"icon"`
	Atribut    string `json:"atribut"`
	Badge      int64  `json:"badge"`
	ParentID   int64  `json:"parent_id"`
	Target     string `json:"target"`
	Urutan     int64  `json:"urutan"`
	RoleAccess int64  `json:"role_access"`
	KanpusOnly int64  `json:"kanpus_only"`
	Jenis      int64  `json:"jenis"`
	Posisi     int64  `json:"posisi"`
	Status     int64  `json:"status"`
}

type Menus []Menu

type MenuResponse struct {
	Title string              `json:"title"`
	Url   string              `json:"url"`
	Icon  string              `json:"icon"`
	Child []ChildMenuResponse `json:"child"`
}
type ChildMenuResponse struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Icon  string `json:"icon"`
}
