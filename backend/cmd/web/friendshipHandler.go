package main

import (
	"encoding/json"
	"github.com/isa-mrs-tim6/Projekat/pkg/models"
	"io/ioutil"
	"net/http"
)

func (app *Application) GetFriends(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	var friends []models.User

	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	friends, err = app.Store.GetFriends(user.ID)

	err = json.NewEncoder(w).Encode(friends)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode friends into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) GetPendingRequests(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	var requests []models.User

	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	requests, err = app.Store.GetPendingRequests(user.ID)

	err = json.NewEncoder(w).Encode(requests)
	if err != nil {
		app.ErrorLog.Printf("Cannot encode pending requests into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) CreateFriendRequest(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	var friendshipDto models.FriendshipDto
	err := json.NewDecoder(r.Body).Decode(&friendshipDto)

	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var friendship = models.Friendship{
		User1ID: user.ID,
		User2ID: friendshipDto.User2ID,
		Status: "PENDING",
	}
	user1, err := app.Store.GetUserWithId(friendship.User1ID)
	friendship.User1 = &user1
	user2, err := app.Store.GetUserWithId(friendship.User2ID)
	friendship.User2 = &user2
	err = app.Store.CreateRequest(friendship)
	if err != nil {
		app.ErrorLog.Printf("Cannot create a friend request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) UpdateFriendRequest(w http.ResponseWriter, r *http.Request) {
	var friendshipDto models.FriendshipDto
	email := getEmail(r)
	data, err := ioutil.ReadAll(r.Body)

	if err == nil && data != nil {
		err = json.Unmarshal(data, &friendshipDto)
	} else {
		app.ErrorLog.Println("Could not decode JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	friendshipDto.User1ID = user.ID

	err = app.Store.UpdateFriendRequest(friendshipDto)
	if err != nil {
		app.ErrorLog.Printf("Cannot update friend request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (app *Application) RemoveAFriend(w http.ResponseWriter, r *http.Request) {
	email := getEmail(r)
	var friendshipDto models.FriendshipDto
	err := json.NewDecoder(r.Body).Decode(&friendshipDto)

	user, err := app.Store.GetUser(email)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	friendshipDto.User1ID = user.ID

	err = app.Store.RemoveAFriend(friendshipDto)
	if err != nil {
		app.ErrorLog.Printf("Cannot remove a friend")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
