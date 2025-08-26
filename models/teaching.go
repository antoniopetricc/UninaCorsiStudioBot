package models

type Teaching struct {
	ID                  string  `json:"id" gorm:"primaryKey"`
	AnnoAccademico      int     `json:"annopri"`
	CodInsegnamento     string  `json:"codInsegnamento"`
	DescInsegnamento    string  `json:"descInsegnamento"`
	CodIntegrato        *string `json:"codIntegrato"`
	CodCorso            string  `json:"codCorso"`
	DescInsegnamentoEng string  `json:"descInsegnamentoEng"`
}

type TeachingsResponse struct {
	TotalItems  int        `json:"totalItems"`
	TotalPages  int        `json:"totalPages"`
	CurrentPage int        `json:"currentPage"`
	Items       []Teaching `json:"items"`
}
