package auth

type AuthRequest struct {
	Phone string `json:"phone"`
}

type AuthResponse struct {
	SessionId string `json:"sessionId"`
}

type ConfirmRequest struct {
	SessionId string `json:"sessionId"`
	Code      int    `json:"code"`
}
