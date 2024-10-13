package event

type EventReq struct {
	Name     string         `json:"name"`
	Member   int            `json:"member"`
	Date     string         `json:"date"`
	Customer *EventCustomer `json:"customer"`
}

type EventCustomer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EventResp struct {
	Message string `json:"message"`
}
