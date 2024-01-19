package model

type User struct {
	ID       int    `json:"id"`       // Identifiant unique de l'utilisateur
	Username string `json:"username"` // Nom d'utilisateur
	Password string `json:"password"` // Mot de passe (stocké sous forme hachée)
	Email    string `json:"email"`    // Adresse e-mail de l'utilisateur
}
