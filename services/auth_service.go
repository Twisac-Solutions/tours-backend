package services

import (
	"time"

	"github.com/Twisac-Solutions/tours-backend/blacklist"
	"github.com/Twisac-Solutions/tours-backend/config"
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// AuthResponse represents the response returned after authentication actions.
// swagger:model
type AuthResponse struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Token   string        `json:"token,omitempty"` // Token is optional
	User    *UserResponse `json:"user,omitempty"`  // User is optional
}

// UserResponse represents a user in API responses.
// swagger:model
type UserResponse struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
}

// RegisterInput represents the expected input for user registration.
// swagger:model
type RegisterInput struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user account with auto-generated username
// @Tags Auth
// @Accept json
// @Produce json
// @Param loginInput body LoginInput true "Login Input"
// @Success 201 {object} AuthResponse
// @Failure 400 {object} AuthResponse
// @Failure 500 {object} AuthResponse
// @Router /api/register [post]
func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	username := utils.GenerateUsername(data["name"])

	user := models.User{
		ID:              uuid.New(),
		Name:            data["name"],
		Email:           data["email"],
		Username:        username,
		Password:        string(hash),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		EmailVerifiedAt: time.Now(),
		IsVerified:      true,
	}

	database.DB.Create(&user)
	token, _ := utils.GenerateJWT(user.ID.String())
	return c.JSON(fiber.Map{"token": token})
}

// Login godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param loginInput body RegisterInput true "Login Input"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} AuthResponse
// @Failure 401 {object} AuthResponse
// @Failure 500 {object} AuthResponse
// @Router /api/login [post]
func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(AuthResponse{
			Status:  "error",
			Message: "Invalid input",
		})
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == uuid.Nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(AuthResponse{
			Status:  "error",
			Message: "Invalid credentials",
		})
	}

	token, _ := utils.GenerateJWT(user.ID.String())
	return c.JSON(AuthResponse{
		Status:  "success",
		Message: "Login successful",
		Token:   token,
		User: &UserResponse{
			ID:       user.ID.String(),
			Email:    user.Email,
			Name:     user.Name,
			Username: user.Username,
			// ProfilePicture: user.ProfilePicture, // Uncomment if available in your model
		},
	})
}

func generateJWT(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

func GoogleLogin(c *fiber.Ctx) error {
	url := utils.GetGoogleOAuthURL()
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

// GoogleCallback handles the callback from Google OAuth.
func GoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Code not found"})
	}

	// Exchange the code for an access token and fetch user info.
	userInfo, err := utils.GetGoogleUserInfo(code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if a user with this email exists.
	var user models.User
	if err := database.DB.Where("email = ?", userInfo.Email).First(&user).Error; err != nil {
		// If not, create a new user with auto‑generated username.
		username := utils.GenerateUsername(userInfo.Name)
		user = models.User{
			Email:    userInfo.Email,
			Name:     userInfo.Name,
			Username: username,
		}
		database.DB.Create(&user)
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create token"})
	}

	return c.JSON(fiber.Map{"token": token, "user": user})
}

// Logout godoc
// @Summary Logout a user
// @Description Invalidate the current JWT token by blacklisting it
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} AuthResponse
// @Failure 401 {object} AuthResponse
// @Router /api/logout [post]
func Logout(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(AuthResponse{
			Status:  "error",
			Message: "Missing Authorization header",
		})
	}

	// Extract token from "Bearer <token>"
	var token string
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		token = authHeader[7:]
	} else {
		token = authHeader
	}

	// Optionally, parse the token to get its expiry (for now, use a default expiry)
	expiry := time.Now().Add(24 * time.Hour) // Set to 24h for demonstration

	// Use a package-level blacklist instance (should be initialized in main or as a singleton)
	if globalBlacklist == nil {
		globalBlacklist = blacklist.NewBlacklist()
	}
	globalBlacklist.Add(token, expiry)

	return c.JSON(AuthResponse{
		Status:  "success",
		Message: "Successfully logged out",
	})
}

// globalBlacklist is a package-level variable for demonstration.
// In production, use a better-scoped or persistent solution.
var globalBlacklist *blacklist.Blacklist
