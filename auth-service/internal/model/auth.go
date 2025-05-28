package model

// LoginRequest представляет запрос на вход в систему
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterRequest представляет запрос на регистрацию
type RegisterRequest struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	MiddleName string `json:"middle_name"`
	Role       string `json:"role" binding:"required"`
}

// CompleteRegistrationRequest представляет запрос на завершение регистрации
type CompleteRegistrationRequest struct {
	UserID       uint   `json:"user_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	MiddleName   string `json:"middle_name"`
	Role         string `json:"role" binding:"required"`
	BirthDate    string `json:"birth_date"`
	Phone        string `json:"phone"`
	
	// Для студента
	Group        string `json:"group"`
	StudentID    string `json:"student_id"`
	
	// Для преподавателя
	Department   string `json:"department"`
	Position     string `json:"position"`
	Degree       string `json:"degree"`
	
	// Для сотрудника деканата
	InternalPhone string `json:"internal_phone"`
	Gender        string `json:"gender"`
}

// TokenResponse представляет ответ с токенами
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

// RefreshRequest представляет запрос на обновление токена
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// UserResponse представляет информацию о пользователе
type UserResponse struct {
	ID         uint   `json:"id"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name,omitempty"`
	Role       string `json:"role"`
	IsActive   bool   `json:"is_active"`
}

// UpdateProfileRequest представляет запрос на обновление профиля
type UpdateProfileRequest struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
}

// ChangePasswordRequest представляет запрос на изменение пароля
type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

// ErrorResponse представляет ответ с ошибкой
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse представляет успешный ответ
type SuccessResponse struct {
	Message string `json:"message"`
} 