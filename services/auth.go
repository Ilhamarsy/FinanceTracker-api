package services

import (
	"errors"
	"fmt"
	"time"

	"finance-tracker/config"
	"finance-tracker/models"
	"finance-tracker/repositories"
	"finance-tracker/utils"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	jwtSecret string
	repo      *repositories.AuthRepository
}

func NewAuthService(cfg *config.Config, repo *repositories.AuthRepository) *AuthService {
	return &AuthService{
		jwtSecret: cfg.JWTSecret,
		repo:      repo,
	}
}

func (a *AuthService) CreateUser(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}

	return a.repo.CreateUser(&newUser)
}

// GenerateToken generates a new JWT token for a given user ID
func (a *AuthService) GenerateToken(user *models.User) (string, error) {
	availUser, err := a.repo.FindByEmail(user.Email)
	if err == nil && user == nil {
		return "", errors.New("email or password is incorrect")
	}

	fmt.Println(availUser.Password)
	fmt.Println(user.Password)
	err = utils.CheckPassword(availUser.Password, user.Password)
	if err != nil {
		return "", errors.New("email or password is incorrect")
	}

	// Create token claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": availUser.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the JWT token and returns the claims
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		// Return the secret key
		return []byte(config.LoadConfig().JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Extract claims from token
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, errors.New("unable to parse claims")
}
