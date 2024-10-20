package handlers

import (
	"Package-Tracker/database"
	"Package-Tracker/models"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"time"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func NewClaims(ID string, role string, standardClaims jwt.StandardClaims) *Claims {
	return &Claims{ID: ID, Role: role, StandardClaims: standardClaims}
}

func Login(writer http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", creds.Email).First(&user).Error; err != nil {
		http.Error(writer, "User not found", http.StatusUnauthorized)
		return
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		http.Error(writer, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Token creation
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID:   strconv.Itoa(int(user.ID)),
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(writer, "Could not generate token "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(map[string]string{"token": tokenString})
}

func RegisterUser(writer http.ResponseWriter, req *http.Request) {
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Hash the user's password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(writer, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// check if all fields are filled
	if user.Name == "" {
		http.Error(writer, "Name is required", http.StatusBadRequest)
		return
	}
	if user.Email == "" {
		http.Error(writer, "Email is required", http.StatusBadRequest)
		return
	}
	if user.Phone == "" {
		http.Error(writer, "Phone is required", http.StatusBadRequest)
		return
	}
	if user.Password == "" {
		http.Error(writer, "Password is required", http.StatusBadRequest)
		return
	}
	if user.Role == "" {
		http.Error(writer, "Role is required", http.StatusBadRequest)
		return
	}

	// check if user exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		http.Error(writer, "User already exists", http.StatusConflict)
		return
	}

	// Save to DB
	if err := database.DB.Create(&user).Error; err != nil {
		http.Error(writer, "Could not create user", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(user)
}
