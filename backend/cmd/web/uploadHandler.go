package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func (app *Application) UpdateProfilePicture(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		app.ErrorLog.Println("Uploaded file too large")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		app.ErrorLog.Println("Could not get uploaded file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	if _, err := fmt.Fprintf(w, "%v", handler.Header); err != nil {
		app.ErrorLog.Println("Could not format the file header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	type Result struct {
		Picture string
	}
	accountType := getAccountType(r)
	email := getEmail(r)
	var fileName string

	switch accountType {
	case "SystemAdmin":
		account, err := app.Store.GetSystemAdmin(email)
		if err != nil {
			app.ErrorLog.Println("No such system admin")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fileName = accountType + strconv.Itoa(int(account.ID))

		f, err := os.OpenFile("./ui/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			app.ErrorLog.Println("Could not open file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer f.Close()
		if _, err := io.Copy(f, file); err != nil {
			app.ErrorLog.Println("Could not copy file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		account.Picture = fileName
		app.Store.Save(&account)

		err = json.NewEncoder(w).Encode(Result{fileName})
		if err != nil {
			app.ErrorLog.Printf("Cannot encode picture path into JSON object")
			w.WriteHeader(http.StatusInternalServerError)

		}
	case "AirlineAdmin":
		account, err := app.Store.GetAirlineAdmin(email)
		if err != nil {
			app.ErrorLog.Println("No such airline admin")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fileName = accountType + strconv.Itoa(int(account.ID))

		f, err := os.OpenFile("./ui/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			app.ErrorLog.Println("Could not open file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer f.Close()
		if _, err := io.Copy(f, file); err != nil {
			app.ErrorLog.Println("Could not copy file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		account.Picture = fileName
		app.Store.Save(&account)

		err = json.NewEncoder(w).Encode(Result{fileName})
		if err != nil {
			app.ErrorLog.Printf("Cannot encode picture path into JSON object")
			w.WriteHeader(http.StatusInternalServerError)

		}
	case "HotelAdmin":
		account, err := app.Store.GetHotelAdmin(email)
		if err != nil {
			app.ErrorLog.Println("No such hotel admin")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fileName = accountType + strconv.Itoa(int(account.ID))

		f, err := os.OpenFile("./ui/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			app.ErrorLog.Println("Could not open file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer f.Close()
		if _, err := io.Copy(f, file); err != nil {
			app.ErrorLog.Println("Could not copy file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		account.Picture = fileName
		app.Store.Save(&account)

		err = json.NewEncoder(w).Encode(Result{fileName})
		if err != nil {
			app.ErrorLog.Printf("Cannot encode picture path into JSON object")
			w.WriteHeader(http.StatusInternalServerError)

		}
	case "Rent-A-CarAdmin":
		account, err := app.Store.GetRACAdmin(email)
		if err != nil {
			app.ErrorLog.Println("No such rent-a-car admin")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fileName = accountType + strconv.Itoa(int(account.ID))

		f, err := os.OpenFile("./ui/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			app.ErrorLog.Println("Could not open file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer f.Close()
		if _, err := io.Copy(f, file); err != nil {
			app.ErrorLog.Println("Could not copy file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		account.Picture = fileName
		app.Store.Save(&account)

		err = json.NewEncoder(w).Encode(Result{fileName})
		if err != nil {
			app.ErrorLog.Printf("Cannot encode picture path into JSON object")
			w.WriteHeader(http.StatusInternalServerError)

		}
	case "User":
		account, err := app.Store.GetUser(email)
		if err != nil {
			app.ErrorLog.Println("No such user")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fileName = accountType + strconv.Itoa(int(account.ID))

		f, err := os.OpenFile("./ui/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			app.ErrorLog.Println("Could not open file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer f.Close()
		if _, err := io.Copy(f, file); err != nil {
			app.ErrorLog.Println("Could not copy file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		account.Picture = fileName
		app.Store.Save(&account)

		err = json.NewEncoder(w).Encode(Result{fileName})
		if err != nil {
			app.ErrorLog.Printf("Cannot encode picture path into JSON object")
			w.WriteHeader(http.StatusInternalServerError)

		}
	}
}

func (app *Application) UpdateHotelProfilePicture(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		app.ErrorLog.Println("Uploaded file too large")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		app.ErrorLog.Println("Could not get uploaded file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	if _, err := fmt.Fprintf(w, "%v", handler.Header); err != nil {
		app.ErrorLog.Println("Could not format the file header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	type Result struct {
		Picture string
	}
	email := getEmail(r)

	account, err := app.Store.GetHotelAdmin(email)
	if err != nil {
		app.ErrorLog.Println("No such hotel admin")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileName := "Hotel" + strconv.Itoa(int(account.HotelID))

	f, err := os.OpenFile("./ui/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		app.ErrorLog.Println("Could not open file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer f.Close()
	if _, err := io.Copy(f, file); err != nil {
		app.ErrorLog.Println("Could not copy file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hotelProfile, err := app.Store.GetHotelProfile(account.HotelID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive hotel profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	hotelProfile.Picture = fileName
	app.Store.Save(&account)

	err = json.NewEncoder(w).Encode(Result{fileName})
	if err != nil {
		app.ErrorLog.Printf("Cannot encode picture path into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) UpdateAirlineProfilePicture(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		app.ErrorLog.Println("Uploaded file too large")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		app.ErrorLog.Println("Could not get uploaded file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	if _, err := fmt.Fprintf(w, "%v", handler.Header); err != nil {
		app.ErrorLog.Println("Could not format the file header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	type Result struct {
		Picture string
	}
	email := getEmail(r)

	account, err := app.Store.GetAirlineAdmin(email)
	if err != nil {
		app.ErrorLog.Println("No such hotel admin")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileName := "Airline" + strconv.Itoa(int(account.AirlineID))

	f, err := os.OpenFile("./ui/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		app.ErrorLog.Println("Could not open file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer f.Close()
	if _, err := io.Copy(f, file); err != nil {
		app.ErrorLog.Println("Could not copy file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	airlineProfile, err := app.Store.GetAirlineProfile(account.AirlineID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive airline profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	airlineProfile.Picture = fileName
	app.Store.Save(&account)

	err = json.NewEncoder(w).Encode(Result{fileName})
	if err != nil {
		app.ErrorLog.Printf("Cannot encode picture path into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) UpdateRACProfilePicture(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		app.ErrorLog.Println("Uploaded file too large")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		app.ErrorLog.Println("Could not get uploaded file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	if _, err := fmt.Fprintf(w, "%v", handler.Header); err != nil {
		app.ErrorLog.Println("Could not format the file header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	type Result struct {
		Picture string
	}
	email := getEmail(r)

	account, err := app.Store.GetRACAdmin(email)
	if err != nil {
		app.ErrorLog.Println("No such hotel admin")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileName := "RAC" + strconv.Itoa(int(account.RentACarCompanyID))

	f, err := os.OpenFile("./ui/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		app.ErrorLog.Println("Could not open file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer f.Close()
	if _, err := io.Copy(f, file); err != nil {
		app.ErrorLog.Println("Could not copy file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	racProfile, err := app.Store.GetRentACarCompanyProfile(account.RentACarCompanyID)
	if err != nil {
		app.ErrorLog.Printf("Could not retrive RAC profile")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	racProfile.Picture = fileName
	app.Store.Save(&account)

	err = json.NewEncoder(w).Encode(Result{fileName})
	if err != nil {
		app.ErrorLog.Printf("Cannot encode picture path into JSON object")
		w.WriteHeader(http.StatusInternalServerError)
	}
}
