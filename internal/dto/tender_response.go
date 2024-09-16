package dto

type TenderResponse struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Status      string `json:"status"`
    ServiceType string `json:"serviceType"`
    Version     int    `json:"version"`
    CreatedAt   string `json:"createdAt"`
}
