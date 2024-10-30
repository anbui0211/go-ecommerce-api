package vo

type UserRegistrationRequest struct {
	Email   string `json:"email" binding:"required,email"`
	Purpose string `json:"Purpose" binding:"required"` // TEST_USER, TRADER, ADMIN etc.
}