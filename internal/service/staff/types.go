package staff

type StaffInfoInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	SessionID string `json:"sessionId"`
}

type StaffInfoResponse struct {
	Status    string
	Comment   string
	SessionId string
}
