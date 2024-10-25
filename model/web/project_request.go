package web

type ProjectCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
type ProjectUpdateRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	CreatedBy   string `json:"created_by"`
}
type ProjectFindByIdRequest struct {
	ID int `json:"id" uri:"id"`
}
