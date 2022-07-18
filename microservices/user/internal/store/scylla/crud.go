package scylla

import (
	"context"
	"time"

	"github.com/gocql/gocql"

	"github.com/MlPablo/gRPCWebSocket/microservices/user/internal/models"
)

type scylla struct {
	session *gocql.Session
}

func New(session *gocql.Session) *scylla {
	return &scylla{session: session}
}

func (s *scylla) Create(ctx context.Context, user models.User) error {
	if err := s.session.Query("INSERT INTO myapp.users (name, password, register_time) VALUES (?, ?, ?)", user.User, user.Password, time.Now()).Exec(); err != nil {
		return err
	}
	return nil
}

func (s *scylla) Read(ctx context.Context, id string) (string, error) {
	var pass string
	if err := s.session.Query("SELECT password FROM myapp.users WHERE name = ?", id).Scan(&pass); err != nil {
		return "", err
	}
	return pass, nil
}

func (s *scylla) Update(ctx context.Context, user models.User) error {
	if err := s.session.Query("INSERT INTO myapp.users (name, password, register_time) VALUES (?, ?, ?)", user.User, user.Password, time.Now()).Exec(); err != nil {
		//if err := s.session.Query("UPDATE users SET password = ? WHERE user = ?;", user.Password, user.User).Exec(); err != nil {
		return err
	}
	return nil
}

func (s *scylla) Delete(ctx context.Context, id string) error {
	if err := s.session.Query("DELETE FROM myapp.users WHERE name = ?", id).Exec(); err != nil {
		return err
	}
	return nil
}
