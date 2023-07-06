package db

import "template/service/models"

func RegisterUser(user models.User) error {
	database := Connect()
	defer database.Close()
	query := `INSERT INTO users (username, password) VALUES (:username, :password)`
	_, err := database.NamedExec(query, user)
	if err != nil {
		return err
	}
	return nil
}
