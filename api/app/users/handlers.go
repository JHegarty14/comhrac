package users

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// ResponseResult => typing for http response
type ResponseResult struct {
	status  bool
	message string
	Error   string
}

// LogoutHandler => Handler for logout route
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var token Token
	var res ResponseResult
	keys, ok := r.URL.Query()["username"]
	tokencollection := sqldb.globalDB.Model(&token)

	if !ok || len(keys[0]) < 1 {
		res.Error = "No user found in request!"
		json.NewEncoder(w).Encode(res)
		return
	}

	var err = tokencollection.FindOne(context.TODO(), bson.D{{"username", token.Username}}).Decode(&token)

	if err != nil {
		res.Error = "Invalid username"
		json.NewEncoder(w).Encode(res)
		return
	}

	var resp = map[string]interface{}{"status": false, "message": "logged out"}

	json.NewEncoder(w).Encode(resp)
	return
}

// LoginHandler => Handler for login route
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	usercollection := sqldb.globalDB.Model(&user)

	if err != nil {
		log.Fatal(err)
	}
	var resultToken Token
	var resultUser User
	var res ResponseResult

	err = usercollection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&resultUser)

	if err != nil {
		res.Error = "Invalid username"
		json.NewEncoder(w).Encode(res)
		return
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(resultUser.Password), []byte(user.Password))

	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		var res = map[string]interface{}{"status": false, "message": "Invalid login credential. Please try again"}
		json.NewEncoder(w).Encode(res)
		return
	}
	tk := &Token{
		UserID: user.ID,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		res.Error = "Error while generating token, Try again"
		json.NewEncoder(w).Encode(res)
		return
	}
	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString
	resp["tk"] = tk

	json.NewEncoder(w).Encode(resp)
}

// RegisterHandler => Handler for signup router
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	var res ResponseResult
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	var result User
	err = sqldb.globalDB.Where("user", body.Email).First(&user)

	if err != nil {
		if err.Error() == "postgres: no documents in result" {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

			if err != nil {
				res.Error = "Error While Hashing Password, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}
			user.Password = string(hash)

			_, err = usercollection.InsertOne(context.TODO(), user)
			if err != nil {
				res.Error = "Error While Creating User, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}
			res.Result = "Register Successful"
			json.NewEncoder(w).Encode(res)
			return
		}
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Result = "A user with that email address already exists!"
	json.NewEncoder(w).Encode(res)
	return
}

// UserSettingsHandler => Handler to get user settings
func UserSettingsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userSettings UserSettings
	keys, ok := r.URL.Query()["userId"]

	if !ok || len(keys[0]) < 1 {
		res.Error = "Invalid request: no user id found"
		json.NewEncoder(w).Encode(res)
		return
	}

	var resultUser UserSettings
	var res ResponseResult

	var err = sqldb.globalDB.Where("userId = ?", keys[0]).First(&userSettings)

	if err != nil {
		// TODO: handle error
	}
}
