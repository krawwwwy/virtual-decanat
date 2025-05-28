package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/model"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/repository"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// AuthService интерфейс для аутентификации
type AuthService interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.UserResponse, error)
	CompleteRegistration(ctx context.Context, req model.CompleteRegistrationRequest) (*model.UserResponse, error)
	Login(ctx context.Context, req model.LoginRequest) (*model.TokenResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*model.TokenResponse, error)
	GetUserByID(ctx context.Context, userID uint) (*model.UserResponse, error)
	UpdateProfile(ctx context.Context, userID uint, req model.UpdateProfileRequest) (*model.UserResponse, error)
	ChangePassword(ctx context.Context, userID uint, req model.ChangePasswordRequest) error
}

// AuthServiceImpl реализация AuthService
type AuthServiceImpl struct {
	userRepo repository.UserRepository
	logger   *zap.Logger
}

// NewAuthService создает новый AuthService
func NewAuthService(userRepo repository.UserRepository, logger *zap.Logger) AuthService {
	return &AuthServiceImpl{
		userRepo: userRepo,
		logger:   logger,
	}
}

// Register регистрирует нового пользователя
func (s *AuthServiceImpl) Register(ctx context.Context, req model.RegisterRequest) (*model.UserResponse, error) {
	// Проверяем, существует ли пользователь с таким email
	existingUser, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		s.logger.Error("Failed to check existing user", zap.Error(err))
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Находим роль
	role, err := s.userRepo.FindRoleByName(ctx, req.Role)
	if err != nil {
		s.logger.Error("Failed to find role", zap.Error(err))
		return nil, fmt.Errorf("failed to find role: %w", err)
	}
	if role == nil {
		return nil, errors.New("invalid role")
	}

	// Создаем пользователя
	user := &model.User{
		Email:      req.Email,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		MiddleName: req.MiddleName,
		RoleID:     role.ID,
		IsActive:   true,
	}

	// Устанавливаем пароль
	if err := user.SetPassword(req.Password); err != nil {
		s.logger.Error("Failed to set password", zap.Error(err))
		return nil, fmt.Errorf("failed to set password: %w", err)
	}

	// Сохраняем пользователя
	if err := s.userRepo.Create(ctx, user); err != nil {
		s.logger.Error("Failed to create user", zap.Error(err))
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Если роль - студент или преподаватель, создаем соответствующую запись
	if role.Name == "student" {
		student := &model.Student{
			UserID:    user.ID,
			GroupID:   1, // Временное значение, должно быть указано в запросе
			StudentID: fmt.Sprintf("S%d", user.ID),
		}
		if err := s.userRepo.CreateStudent(ctx, student); err != nil {
			s.logger.Error("Failed to create student", zap.Error(err))
			return nil, fmt.Errorf("failed to create student: %w", err)
		}
	} else if role.Name == "teacher" {
		teacher := &model.Teacher{
			UserID:     user.ID,
			Department: "Default Department", // Временное значение, должно быть указано в запросе
			Position:   "Default Position",   // Временное значение, должно быть указано в запросе
		}
		if err := s.userRepo.CreateTeacher(ctx, teacher); err != nil {
			s.logger.Error("Failed to create teacher", zap.Error(err))
			return nil, fmt.Errorf("failed to create teacher: %w", err)
		}
	}

	// Возвращаем информацию о пользователе
	return &model.UserResponse{
		ID:         user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Role:       role.Name,
		IsActive:   user.IsActive,
	}, nil
}

// Login выполняет вход пользователя
func (s *AuthServiceImpl) Login(ctx context.Context, req model.LoginRequest) (*model.TokenResponse, error) {
	// Находим пользователя по email
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		s.logger.Error("Failed to find user", zap.Error(err))
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Проверяем пароль
	if !user.CheckPassword(req.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Проверяем активность пользователя
	if !user.IsActive {
		return nil, errors.New("user is not active")
	}

	// Генерируем токены
	accessToken, accessExpires, err := s.generateAccessToken(user)
	if err != nil {
		s.logger.Error("Failed to generate access token", zap.Error(err))
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, _, err := s.generateRefreshToken(user)
	if err != nil {
		s.logger.Error("Failed to generate refresh token", zap.Error(err))
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	// Возвращаем токены
	return &model.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    accessExpires,
	}, nil
}

// RefreshToken обновляет токен доступа
func (s *AuthServiceImpl) RefreshToken(ctx context.Context, refreshToken string) (*model.TokenResponse, error) {
	// Парсим refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Проверяем алгоритм подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt.secret")), nil
	})

	if err != nil {
		s.logger.Error("Failed to parse refresh token", zap.Error(err))
		return nil, errors.New("invalid refresh token")
	}

	if !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	// Получаем claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	// Проверяем тип токена
	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "refresh" {
		return nil, errors.New("invalid token type")
	}

	// Получаем ID пользователя
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("invalid user ID in token")
	}

	// Находим пользователя
	user, err := s.userRepo.FindByID(ctx, uint(userID))
	if err != nil {
		s.logger.Error("Failed to find user", zap.Error(err))
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Проверяем активность пользователя
	if !user.IsActive {
		return nil, errors.New("user is not active")
	}

	// Генерируем новые токены
	accessToken, accessExpires, err := s.generateAccessToken(user)
	if err != nil {
		s.logger.Error("Failed to generate access token", zap.Error(err))
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	newRefreshToken, _, err := s.generateRefreshToken(user)
	if err != nil {
		s.logger.Error("Failed to generate refresh token", zap.Error(err))
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	// Возвращаем токены
	return &model.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    accessExpires,
	}, nil
}

// GetUserByID получает информацию о пользователе по ID
func (s *AuthServiceImpl) GetUserByID(ctx context.Context, userID uint) (*model.UserResponse, error) {
	// Находим пользователя
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to find user", zap.Error(err))
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Возвращаем информацию о пользователе
	return &model.UserResponse{
		ID:         user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Role:       user.Role.Name,
		IsActive:   user.IsActive,
	}, nil
}

// UpdateProfile обновляет профиль пользователя
func (s *AuthServiceImpl) UpdateProfile(ctx context.Context, userID uint, req model.UpdateProfileRequest) (*model.UserResponse, error) {
	// Находим пользователя
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to find user", zap.Error(err))
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Обновляем данные пользователя
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.MiddleName = req.MiddleName

	// Сохраняем изменения
	if err := s.userRepo.Update(ctx, user); err != nil {
		s.logger.Error("Failed to update user", zap.Error(err))
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// Возвращаем обновленную информацию о пользователе
	return &model.UserResponse{
		ID:         user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Role:       user.Role.Name,
		IsActive:   user.IsActive,
	}, nil
}

// ChangePassword изменяет пароль пользователя
func (s *AuthServiceImpl) ChangePassword(ctx context.Context, userID uint, req model.ChangePasswordRequest) error {
	// Находим пользователя
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to find user", zap.Error(err))
		return fmt.Errorf("failed to find user: %w", err)
	}
	if user == nil {
		return errors.New("user not found")
	}

	// Проверяем старый пароль
	if !user.CheckPassword(req.OldPassword) {
		return errors.New("invalid old password")
	}

	// Устанавливаем новый пароль
	if err := user.SetPassword(req.NewPassword); err != nil {
		s.logger.Error("Failed to set new password", zap.Error(err))
		return fmt.Errorf("failed to set new password: %w", err)
	}

	// Сохраняем изменения
	if err := s.userRepo.Update(ctx, user); err != nil {
		s.logger.Error("Failed to update user", zap.Error(err))
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// generateAccessToken генерирует токен доступа
func (s *AuthServiceImpl) generateAccessToken(user *model.User) (string, int64, error) {
	// Получаем время жизни токена из конфигурации
	expirationTime := viper.GetDuration("jwt.access_expiration")
	if expirationTime == 0 {
		expirationTime = 15 * time.Minute // По умолчанию 15 минут
	}

	// Создаем claims для токена
	expiresAt := time.Now().Add(expirationTime)
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role.Name,
		"type":    "access",
		"exp":     expiresAt.Unix(),
		"iat":     time.Now().Unix(),
	}

	// Создаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен
	tokenString, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return "", 0, err
	}

	return tokenString, int64(expirationTime.Seconds()), nil
}

// generateRefreshToken генерирует токен обновления
func (s *AuthServiceImpl) generateRefreshToken(user *model.User) (string, int64, error) {
	// Получаем время жизни токена из конфигурации
	expirationTime := viper.GetDuration("jwt.refresh_expiration")
	if expirationTime == 0 {
		expirationTime = 24 * time.Hour // По умолчанию 24 часа
	}

	// Создаем claims для токена
	expiresAt := time.Now().Add(expirationTime)
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"type":    "refresh",
		"exp":     expiresAt.Unix(),
		"iat":     time.Now().Unix(),
	}

	// Создаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен
	tokenString, err := token.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return "", 0, err
	}

	return tokenString, int64(expirationTime.Seconds()), nil
}

// CompleteRegistration завершает регистрацию пользователя в зависимости от роли
func (s *AuthServiceImpl) CompleteRegistration(ctx context.Context, req model.CompleteRegistrationRequest) (*model.UserResponse, error) {
	// Находим пользователя
	user, err := s.userRepo.FindByID(ctx, req.UserID)
	if err != nil {
		s.logger.Error("Failed to find user", zap.Error(err))
		return nil, fmt.Errorf("failed to find user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Обновляем основные данные пользователя
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.MiddleName = req.MiddleName

	// Находим роль
	role, err := s.userRepo.FindRoleByName(ctx, req.Role)
	if err != nil {
		s.logger.Error("Failed to find role", zap.Error(err))
		return nil, fmt.Errorf("failed to find role: %w", err)
	}
	if role == nil {
		return nil, errors.New("invalid role")
	}

	// Обновляем пользователя
	if err := s.userRepo.Update(ctx, user); err != nil {
		s.logger.Error("Failed to update user", zap.Error(err))
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// В зависимости от роли, создаем или обновляем соответствующие записи
	switch role.Name {
	case "student":
		// Создаем или обновляем данные студента
		student, err := s.userRepo.FindStudentByUserID(ctx, user.ID)
		if err != nil {
			s.logger.Error("Failed to find student", zap.Error(err))
			return nil, fmt.Errorf("failed to find student: %w", err)
		}

		// Парсим дату рождения
		var birthDate time.Time
		if req.BirthDate != "" {
			birthDate, err = time.Parse("2006-01-02", req.BirthDate)
			if err != nil {
				s.logger.Error("Failed to parse birth date", zap.Error(err))
				return nil, fmt.Errorf("invalid birth date format: %w", err)
			}
		}

		// Проверяем, существует ли группа по имени
		group, err := s.userRepo.FindGroupByName(ctx, req.Group)
		if err != nil {
			s.logger.Error("Failed to find group", zap.Error(err))
			return nil, fmt.Errorf("failed to find group: %w", err)
		}

		// Если группа не существует, используем дефолтную группу
		var groupID uint = 1 // По умолчанию
		if group != nil {
			groupID = group.ID
		}

		// Логирование полученных данных для отладки
		s.logger.Info("Received student data",
			zap.String("fio", req.FirstName + " " + req.MiddleName + " " + req.LastName),
			zap.String("birth_date", req.BirthDate),
			zap.String("group", req.Group),
			zap.String("student_id", req.StudentID),
			zap.String("phone", req.Phone))

		if student == nil {
			// Создаем нового студента
			student = &model.Student{
				UserID:    user.ID,
				GroupID:   groupID,
				StudentID: req.StudentID,
				BirthDate: birthDate,
				Phone:     req.Phone,
			}
			if err := s.userRepo.CreateStudent(ctx, student); err != nil {
				s.logger.Error("Failed to create student", zap.Error(err))
				return nil, fmt.Errorf("failed to create student: %w", err)
			}
		} else {
			// Обновляем существующего студента
			student.StudentID = req.StudentID
			student.GroupID = groupID
			student.BirthDate = birthDate
			student.Phone = req.Phone
			
			if err := s.userRepo.UpdateStudent(ctx, student); err != nil {
				s.logger.Error("Failed to update student", zap.Error(err))
				return nil, fmt.Errorf("failed to update student: %w", err)
			}
		}

	case "teacher":
		// Создаем или обновляем данные преподавателя
		teacher, err := s.userRepo.FindTeacherByUserID(ctx, user.ID)
		if err != nil {
			s.logger.Error("Failed to find teacher", zap.Error(err))
			return nil, fmt.Errorf("failed to find teacher: %w", err)
		}

		// Парсим дату рождения
		var birthDate time.Time
		if req.BirthDate != "" {
			birthDate, err = time.Parse("2006-01-02", req.BirthDate)
			if err != nil {
				s.logger.Error("Failed to parse birth date", zap.Error(err))
				return nil, fmt.Errorf("invalid birth date format: %w", err)
			}
		}

		// Логирование полученных данных для отладки
		s.logger.Info("Received teacher data",
			zap.String("fio", req.FirstName + " " + req.MiddleName + " " + req.LastName),
			zap.String("department", req.Department),
			zap.String("position", req.Position),
			zap.String("degree", req.Degree),
			zap.String("birth_date", req.BirthDate),
			zap.String("phone", req.Phone))

		if teacher == nil {
			// Создаем нового преподавателя
			teacher = &model.Teacher{
				UserID:     user.ID,
				Department: req.Department,
				Position:   req.Position,
				Degree:     req.Degree,
				BirthDate:  birthDate,
				Phone:      req.Phone,
			}
			if err := s.userRepo.CreateTeacher(ctx, teacher); err != nil {
				s.logger.Error("Failed to create teacher", zap.Error(err))
				return nil, fmt.Errorf("failed to create teacher: %w", err)
			}
		} else {
			// Обновляем существующего преподавателя
			teacher.Department = req.Department
			teacher.Position = req.Position
			teacher.Degree = req.Degree
			teacher.BirthDate = birthDate
			teacher.Phone = req.Phone
			
			if err := s.userRepo.UpdateTeacher(ctx, teacher); err != nil {
				s.logger.Error("Failed to update teacher", zap.Error(err))
				return nil, fmt.Errorf("failed to update teacher: %w", err)
			}
		}

	case "dean_office":
		// Создаем или обновляем данные сотрудника деканата
		staff, err := s.userRepo.FindStaffByUserID(ctx, user.ID)
		if err != nil {
			s.logger.Error("Failed to find staff", zap.Error(err))
			return nil, fmt.Errorf("failed to find staff: %w", err)
		}

		if staff == nil {
			// Создаем нового сотрудника деканата
			staff = &model.Staff{
				UserID:        user.ID,
				Department:    req.Department,
				Position:      req.Position,
				InternalPhone: req.InternalPhone,
				Gender:        req.Gender,
			}
			if err := s.userRepo.CreateStaff(ctx, staff); err != nil {
				s.logger.Error("Failed to create staff", zap.Error(err))
				return nil, fmt.Errorf("failed to create staff: %w", err)
			}
		} else {
			// Обновляем существующего сотрудника деканата
			staff.Department = req.Department
			staff.Position = req.Position
			staff.InternalPhone = req.InternalPhone
			staff.Gender = req.Gender
			if err := s.userRepo.UpdateStaff(ctx, staff); err != nil {
				s.logger.Error("Failed to update staff", zap.Error(err))
				return nil, fmt.Errorf("failed to update staff: %w", err)
			}
		}
	}

	// Возвращаем информацию о пользователе
	return &model.UserResponse{
		ID:         user.ID,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Role:       role.Name,
		IsActive:   user.IsActive,
	}, nil
}