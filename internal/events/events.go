package events

import (
	"errors"
	"time"
)

type Event struct {
	Id                   int
	Title                string
	Date                 time.Time
	Deadline             time.Time
	Description          string
	User                 *User
	NotificationDeadline time.Time
}

type Notification struct {
	Id    int
	Title string
	Date  time.Time
	User  *User
}

type User struct {
	Id       int
	Username string
	Password string
}

var Events []*Event
var Users []*User
var Notifications []*Notification

func AddUser(username string, password string) *User {
	newUser := User{
		Id:       len(Users),
		Username: username,
		Password: password,
	}
	Users = append(Users, &newUser)
	return &newUser
}

func GetUserById(id int) *User {
	for i, u := range Users {
		if u.Id == id {
			return Users[i]
		}
	}
	return nil
}

func AddEvent(event *Event) {
	Events = append(Events, event)
}

func GetEventById(id int) *Event {
	for i, e := range Events {
		if e.Id == id {
			return Events[i]
		}
	}
	return nil
}

func UpdateEvent(id int, event *Event) error {
	for i, e := range Events {
		if e.Id == id {
			Events[i] = event
			return nil
		}
	}
	return errors.New("event not found")
}

func DeleteEventById(id int) {
	for i, e := range Events {
		if e.Id == id {
			Events[i] = Events[len(Events)-1]
			Events = Events[:len(Events)-1]
		}
	}
}

func GetEventsByDate(date time.Time) []*Event {
	var dateEvents []*Event
	for i, e := range Events {
		if e.Date.Truncate(24 * time.Hour).Equal(date.Truncate(24 * time.Hour)) {
			dateEvents = append(dateEvents, Events[i])
		}
	}
	return dateEvents
}

func GetEventsWeek(dateStart time.Time) []*Event {
	var dateEvents []*Event
	endDate := dateStart.AddDate(0, 0, 7)
	for i, e := range Events {
		if (e.Date.Equal(endDate) || e.Date.Equal(dateStart)) || (e.Date.Before(endDate) && e.Date.After(dateStart)) {
			dateEvents = append(dateEvents, Events[i])
		}
	}
	return dateEvents
}
func GetEventsMonth(dateStart time.Time) []*Event {
	var dateEvents []*Event
	endDate := dateStart.AddDate(0, 1, 0)
	for i, e := range Events {
		if (e.Date.Equal(endDate) || e.Date.Equal(dateStart)) || (e.Date.Before(endDate) && e.Date.After(dateStart)) {
			dateEvents = append(dateEvents, Events[i])
		}
	}
	return dateEvents
}
