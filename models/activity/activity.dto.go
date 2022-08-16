package models

type ActivityRequest struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	CreateAt *string `json:"create_at"`
	UpdateAt *string `json:"update_at"`
}

type ActivityResponse struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	CreateAt *string `json:"create_at"`
	UpdateAt *string `json:"update_at"`
}

func (p ActivityRequest) ParseRequest() Activity {
	return Activity{
		ID:   p.ID,
		Name: p.Name,
	}
}

func (p ActivityResponse) ParseResponse() Activity {
	return Activity{
		ID:       p.ID,
		Name:     p.Name,
		CreateAt: p.CreateAt,
		UpdateAt: p.UpdateAt,
	}
}

func (ar ActivityRequest) TableName() string {
	return "activity"
}

func (ar ActivityResponse) TableName() string {
	return "activity"
}
