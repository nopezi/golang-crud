package models

type VerifikasiRiskControl struct {
	ID            int64
	VerifikasiId  int64
	RiskControlID int64
}

func (vc VerifikasiRiskControl) TableName() string {
	return "verifikasi_risk_control"
}
