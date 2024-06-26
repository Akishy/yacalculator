package auth

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
)

func NewUser(name, password string) *User {
	return &User{
		Name:           name,
		Password:       password,
		IsAdmin:        false,
		AmountOfAgents: 0,
		TimeToCalc:     3 * time.Second,
	}
}

func InsertUser(ctx context.Context, db *sql.DB, user *User) (int64, error) {
	q := `INSERT INTO users (username, password, is_admin, amount_of_agents, time_to_calc) VALUES ($1, $2, $3, $4, $5);`

	result, err := db.ExecContext(ctx, q, user.Name, user.Password, user.IsAdmin, user.AmountOfAgents, user.TimeToCalc)
	if err != nil {
		log.Printf("[ERROR] insertUser: Error inserting new user: %v", err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("[ERROR] insertUser: Error getting integer generated by the database: %v", err)
		return 0, err
	}

	return id, nil
}

func GetUserIdByName(ctx context.Context, db *sql.DB, name string) (int, error) {
	q := `SELECT id FROM users WHERE username = $1;`

	var id int
	err := db.QueryRowContext(ctx, q, name).Scan(&id)
	if err != nil {
		log.Printf("[ERROR] GetUserIdByName: Error checking if user exists: %v", err)
		return -1, err
	}

	return id, nil
}

func GetUserInfoById(ctx context.Context, db *sql.DB, id int) (*User, error) {
	q := `SELECT * FROM users WHERE id = $1;`

	user := &User{}
	err := db.QueryRowContext(ctx, q, id).Scan(&user.ID, &user.Name, &user.Password, &user.IsAdmin, &user.AmountOfAgents, &user.TimeToCalc)
	if err != nil {
		log.Printf("[ERROR] GetUserInfoById: Error finding user: %v", err)
		return nil, err
	}

	return user, nil
}

func GetUserInfoByName(ctx context.Context, db *sql.DB, name string) (*User, error) {
	q := `SELECT * FROM users WHERE username = $1;`

	user := &User{}
	err := db.QueryRowContext(ctx, q, name).Scan(&user.ID, &user.Name, &user.Password, &user.IsAdmin, &user.AmountOfAgents, &user.TimeToCalc)
	if err != nil {
		log.Printf("[ERROR] GetUserInfoByName: Error finding user: %v", err)
		return nil, err
	}

	return user, nil
}

// CheckIfUserExists - returns true, nil when user exists and false, nil, when user not exist
func CheckIfUserExists(ctx context.Context, db *sql.DB, username string) (bool, error) {
	q := `SELECT count(*) FROM users WHERE username = $1`

	var count int
	err := db.QueryRowContext(ctx, q, username).Scan(&count)
	if err != nil {
		log.Printf("[ERROR] CheckIfUserExists: Error checking if user exists: %v", err)
		return false, err
	}

	log.Println("[DEBUG] CheckIfUserExists: Number of users found:", count)

	return count > 0, nil
}

// CheckLoginPassword - check if given username and password correct to identify user's credentials
func CheckLoginPassword(ctx context.Context, db *sql.DB, username string, password string) (*User, error) {
	q := `SELECT username, is_admin, amount_of_agents, time_to_calc FROM users WHERE username = $1 AND password = $2`

	user := new(User)
	err := db.QueryRowContext(ctx, q, username, password).Scan(&user.Name, &user.IsAdmin, &user.AmountOfAgents, &user.TimeToCalc)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		log.Println("[ERROR] CheckLoginPassword: No user found")
		return nil, err
	case err != nil:
		log.Printf("[ERROR] CheckLoginPassword: Error checking if user exists: %v", err)
		return nil, err
	default:
		log.Println("[DEBUG] CheckLoginPassword: user found:", user)
		return user, nil
	}
}
