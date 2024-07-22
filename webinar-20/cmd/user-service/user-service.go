package main

import "fmt"

type userStorage interface {
	GetUserByID(userID int) (*user, error)
	UpsertUser(user) error
	GetAllUsers() []user
}

type user struct {
	ID         int
	Level      int // Level indicates how many orders user paid. One level = 100 orders.
	OrdersPaid int
}

type userService struct {
	storage userStorage
}

type orderPaidEvent struct {
	UserID  int
	OrderID int
}

func (s *userService) ConsumeOrderPaidEvent(e orderPaidEvent) error {
	u, err := s.storage.GetUserByID(e.UserID)
	if err != nil {
		return fmt.Errorf("getting user by id: %w", err)
	}

	u.OrdersPaid++
	u.Level = u.OrdersPaid / 100

	if err := s.storage.UpsertUser(*u); err != nil {
		return fmt.Errorf("updating user: %w", err)
	}

	return nil
}
