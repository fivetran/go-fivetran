package common

type CommonResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SetupTestResponse struct {
	Title   string `json:"title"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
