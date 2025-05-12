package store

import (
	"context"
	"fmt"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedOn string `json:"created_on"`
}

func (s *SQLiteStore) CreateUser(ctx context.Context, user *User) error {
	// Prepare the SQL statement
	stmt, err := s.DB.PrepareContext(ctx, "INSERT INTO users (username, password) VALUES ($1, $2)")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}

	return nil
}

func (s *SQLiteStore) GetUserByUsername(ctx context.Context, user *User) error {
	// Prepare the SQL statement
	stmt, err := s.DB.PrepareContext(ctx, "SELECT id, username, password FROM users WHERE username = $1")
	if err != nil {
		return fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	// Execute the SQL statement
	err = stmt.QueryRowContext(ctx, user.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return fmt.Errorf("error executing statement: %w", err)
	}

	return nil
}
