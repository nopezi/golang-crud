package models

type FaqRequest struct {
	ID       int64  `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type FaqResponse struct {
	ID        int64  `json:"id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (p FaqRequest) ParseRequest() Faqs {
	return Faqs{
		ID:       p.ID,
		Answer:   p.Answer,
		Question: p.Question,
	}
}

func (p FaqResponse) ParseResponse() Faqs {
	return Faqs{
		ID:        p.ID,
		Answer:    p.Answer,
		Question:  p.Question,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
