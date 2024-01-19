package model

type Todo struct {
	ID          int    `json:"id"`          // Identifiant unique de la tâche
	UserID      int    `json:"user_id"`     // Identifiant de l'utilisateur associé
	Title       string `json:"title"`       // Titre de la tâche
	Description string `json:"description"` // Description de la tâche
	Completed   bool   `json:"completed"`   // Indique si la tâche est terminée
}
