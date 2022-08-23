package models

type ActivityRequest struct {
	ID           int64   `json:"id"`
	KodeActivity string  `json:"kode_activity"`
	Name         string  `json:"name"`
	CreateAt     *string `json:"create_at"`
	UpdateAt     *string `json:"update_at"`
}

type ActivityResponse struct {
	ID           int64   `json:"id"`
	KodeActivity string  `json:"kode_activity"`
	Name         string  `json:"name"`
	CreateAt     *string `json:"create_at"`
	UpdateAt     *string `json:"update_at"`
}

func (p ActivityRequest) ParseRequest() Activity {
	return Activity{
		ID:           p.ID,
		KodeActivity: p.KodeActivity,
		Name:         p.Name,
	}
}

func (p ActivityResponse) ParseResponse() Activity {
	return Activity{
		ID:           p.ID,
		KodeActivity: p.KodeActivity,
		Name:         p.Name,
		CreateAt:     p.CreateAt,
		UpdateAt:     p.UpdateAt,
	}
}

func (ar ActivityRequest) TableName() string {
	return "activity"
}

func (ar ActivityResponse) TableName() string {
	return "activity"
}
