package models

type SubActivityRequest struct {
	ID         int64   `json:"id"`
	ActivityID int64   `json:"activity_id"`
	Name       string  `json:"name"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

type SubActivityResponse struct {
	ID         int64   `json:"id"`
	ActivityID int64   `json:"activity_id"`
	Name       string  `json:"name"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
}

func (p SubActivityRequest) ParseRequest() SubActivity {
	return SubActivity{
		ID:         p.ID,
		ActivityID: p.ActivityID,
		Name:       p.Name,
		// CreatedAt:  p.CreatedAt,
		// UpdatedAt:  p.UpdatedAt,
	}
}

func (p SubActivityResponse) ParseResponse() SubActivity {
	return SubActivity{
		ID:         p.ID,
		ActivityID: p.ActivityID,
		Name:       p.Name,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func (ar SubActivityRequest) TableName() string {
	return "sub_activity"
}

func (ar SubActivityResponse) TableName() string {
	return "sub_activity"
}
