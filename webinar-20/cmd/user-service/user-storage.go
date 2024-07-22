package main

import (
	"errors"
	"fmt"
)

type inMemUserStorage struct {
	users map[int]user // TODO: make it safe with concurrency
}

func newInMemUserStorage() *inMemUserStorage {
	return &inMemUserStorage{
		users: map[int]user{
			1: {
				ID:         1,
				OrdersPaid: 99,
			},
			2: {
				ID:         2,
				OrdersPaid: 180,
				Level:      1,
			},
		},
	}
}

var errUserNotFound = errors.New("user not found")

func (s *inMemUserStorage) GetUserByID(userID int) (*user, error) {
	u, ok := s.users[userID]
	if !ok {
		return nil, fmt.Errorf("%w: %v", errUserNotFound, userID)
	}

	return &u, nil
}

func (s *inMemUserStorage) UpsertUser(u user) error {
	s.users[u.ID] = u
	return nil
}

func (s *inMemUserStorage) GetAllUsers() []user {
	var users []user

	for _, v := range s.users {
		users = append(users, v)
	}

	return users
}
