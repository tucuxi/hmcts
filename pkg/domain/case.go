package domain

type Case struct {
	Id          int     `json:"id"`
	CaseNumber  string  `json:"caseNumber"`
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	Status      string  `json:"status"`
	CreatedDate string  `json:"createdDate"`
}
