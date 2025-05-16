package services

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/models"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// TokenClaims представляет содержимое JWT-токена
type TokenClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// AuthService предоставляет методы для работы с аутентификацией и авторизацией пользователей
type AuthService struct {
	userRepo *repository.UserRepository
	jwtSecret string
	jwtExpiration time.Duration
}

// NewAuthService создает новый экземпляр AuthService
func NewAuthService(userRepo *repository.UserRepository, jwtSecret string, jwtExpiration time.Duration) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtSecret: jwtSecret,
		jwtExpiration: jwtExpiration,
	}
}

// Register регистрирует нового пользователя
func (s *AuthService) Register(req models.UserRegistration) (*models.User, error) {
	// Проверка существования пользователя с таким же username
	exists, err := s.userRepo.CheckUsernameExists(req.Username)
	if err != nil {
		log.Printf("Error checking username existence: %v", err)
		return nil, err
	}
	if exists {
		return nil, repository.ErrUserAlreadyExists
	}

	// Проверка существования пользователя с таким же email
	exists, err = s.userRepo.CheckEmailExists(req.Email)
	if err != nil {
		log.Printf("Error checking email existence: %v", err)
		return nil, err
	}
	if exists {
		return nil, repository.ErrUserAlreadyExists
	}

	// Хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, err
	}

	// Создание пользователя
	user := &models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		MiddleName:   req.MiddleName,
		Role:         req.Role,
		Group:        req.Group,
		Faculty:      req.Faculty,
		Department:   req.Department,
	}

	err = s.userRepo.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	return user, nil
}

// Login выполняет аутентификацию пользователя и возвращает JWT-токен
func (s *AuthService) Login(req models.UserLogin) (*models.TokenResponse, error) {
	// Получаем пользователя по username
	user, err := s.userRepo.GetUserByUsername(req.Username)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, repository.ErrInvalidCredentials
		}
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}

	// Проверяем пароль
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		log.Printf("Invalid password for user %s: %v", req.Username, err)
		return nil, repository.ErrInvalidCredentials
	}

	// Генерируем JWT-токен
	tokenString, expiresAt, err := s.GenerateJWT(user)
	if err != nil {
		log.Printf("Error generating JWT: %v", err)
		return nil, err
	}

	// Формируем ответ
	userResponse := models.UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Role:       user.Role,
		Group:      user.Group,
		Faculty:    user.Faculty,
		Department: user.Department,
		CreatedAt:  user.CreatedAt,
	}

	return &models.TokenResponse{
		Token:     tokenString,
		ExpiresAt: expiresAt,
		User:      userResponse,
	}, nil
}

// GenerateJWT генерирует JWT-токен для пользователя
func (s *AuthService) GenerateJWT(user *models.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(s.jwtExpiration)

	claims := &TokenClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "virtual-decanat",
			Subject:   user.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}

// VerifyToken проверяет JWT-токен и возвращает содержащиеся в нем данные
func (s *AuthService) VerifyToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// GetUserByID получает пользователя по ID
func (s *AuthService) GetUserByID(userID int) (*models.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		return nil, err
	}

	return &models.UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Role:       user.Role,
		Group:      user.Group,
		Faculty:    user.Faculty,
		Department: user.Department,
		CreatedAt:  user.CreatedAt,
	}, nil
}

// UpdateUser обновляет данные пользователя
func (s *AuthService) UpdateUser(userID int, updateData map[string]interface{}) (*models.UserResponse, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		return nil, err
	}

	// Обновляем только предоставленные поля
	if email, ok := updateData["email"].(string); ok && email != "" {
		user.Email = email
	}
	if firstName, ok := updateData["first_name"].(string); ok && firstName != "" {
		user.FirstName = firstName
	}
	if lastName, ok := updateData["last_name"].(string); ok && lastName != "" {
		user.LastName = lastName
	}
	if middleName, ok := updateData["middle_name"].(string); ok {
		user.MiddleName = middleName
	}
	if group, ok := updateData["group"].(string); ok {
		user.Group = group
	}
	if faculty, ok := updateData["faculty"].(string); ok {
		user.Faculty = faculty
	}
	if department, ok := updateData["department"].(string); ok {
		user.Department = department
	}

	err = s.userRepo.UpdateUser(user)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return nil, err
	}

	return &models.UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Role:       user.Role,
		Group:      user.Group,
		Faculty:    user.Faculty,
		Department: user.Department,
		CreatedAt:  user.CreatedAt,
	}, nil
}

// ChangePassword изменяет пароль пользователя
func (s *AuthService) ChangePassword(userID int, oldPassword, newPassword string) error {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		return err
	}

	// Проверка старого пароля
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword))
	if err != nil {
		log.Printf("Invalid old password for user %d: %v", userID, err)
		return repository.ErrInvalidCredentials
	}

	// Хэширование нового пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}

	// Обновление пароля в БД
	return s.userRepo.UpdateUserPassword(userID, string(hashedPassword))
} 