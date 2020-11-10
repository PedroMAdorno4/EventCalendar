package rest

import (
	"encoding/json"
	"flag"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/PedroMAdorno4/Desafio/pkg/auth"
	"github.com/PedroMAdorno4/Desafio/pkg/create"
	"github.com/PedroMAdorno4/Desafio/pkg/delete"
	"github.com/PedroMAdorno4/Desafio/pkg/read"
	"github.com/PedroMAdorno4/Desafio/pkg/update"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

func Handler(a auth.Service, c create.Service, r read.Service, u update.Service, d delete.Service) http.Handler {
	var dir string
	flag.StringVar(&dir, "dir", "assets/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router := mux.NewRouter()
	// router.Use(auth.JwtAuth)
	// router.Use(auth.Authorization)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	// router.Path("/player").Queries("game", "{id}").HandlerFunc(playerPage)
	// router.HandleFunc("/API/auth", createToken(a, r)).Methods("POST", "OPTIONS")
	// router.HandleFunc("/API/parser/mediaResources", getMediaResources(p)).Methods("GET", "OPTIONS")
	// router.HandleFunc("/API/user/get", getUser(r)).Methods("GET")
	// router.Path("/API/act/get").Queries("act", "{id}").HandlerFunc(getAct(r)).Methods("GET", "OPTIONS")
	// router.HandleFunc("/API/media/icons/upload", uploadIcons(a)).Methods("POST")

	router.HandleFunc("/API/user/create", createUser(c)).Methods("POST")
	router.HandleFunc("/API/user/get", getUser(r)).Methods("GET")
	router.HandleFunc("/API/user/update", updateUser(u)).Methods("POST")
	router.HandleFunc("/API/user/delete", deleteUser(d)).Methods("POST")

	return router
}

var contextKeyToken = auth.ContextKey("authToken")
var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func setSession(token string, w http.ResponseWriter) {
	value := map[string]string{
		"token": token,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func getSession(r *http.Request) (string, error) {
	var token string
	var err error
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			token = cookieValue["token"]
		}
	}
	return token, err
}

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	clearSession(w)

	http.Redirect(w, r, "/", http.StatusFound)
	return

}

func createToken(a auth.Service, reader read.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userInfo UserLogin
		err := json.NewDecoder(r.Body).Decode(&userInfo)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		logged, err := a.CheckCredentials(userInfo.Username, userInfo.Password)
		if err == auth.ErrCredentialsMissing {
			http.Error(w, "Credentials missing in request", http.StatusBadRequest)
			return
		} else if err == auth.ErrCredentialsInvalid {
			http.Error(w, "Credentials invalid", http.StatusUnauthorized)
			return
		}

		token := ""
		if logged {
			id, _ := reader.GetIDByUsername(userInfo.Username)
			token, err = a.CreateToken(id)
		}

		json.NewEncoder(w).Encode(token)
	}
}

func loginHandler(a auth.Service, reader read.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		logged, err := a.CheckCredentials(username, password)
		if err == auth.ErrCredentialsMissing {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		} else if err == auth.ErrCredentialsInvalid {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		if logged {
			id, _ := reader.GetIDByUsername(username)
			token, _ := a.CreateToken(id)

			setSession(token, w)

		}
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

//User CRUD \\
func createUser(s create.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var userInfo create.User
		err := json.NewDecoder(r.Body).Decode(&userInfo)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		insertedID, err := s.CreateUser(userInfo)
		/*var userGroup create.Group

		userGroup.OwnerID = insertedID

		switch userInfo.Permissions {
		case auth.PermPlayer:
			userGroup.Depth = 2
		case auth.PermManager:
			userGroup.Depth = 1
		}

		s.CreateGroup(userGroup)*/

		json.NewEncoder(w).Encode("Usuario " + userInfo.Username + "(" + insertedID.Hex() + ")" + " inserido com sucesso")
	}
}

func getUser(s read.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var id primitive.ObjectID
		err := json.NewDecoder(r.Body).Decode(&id)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := s.GetUser(id)
		if err == read.ErrUserNotFound {
			http.Error(w, "The user you requested does not exist", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func updateUser(s update.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var user update.User
		err := json.NewDecoder(r.Body).Decode(&user)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.UpdateUser(user)
		if err == update.ErrUserNotFound {
			http.Error(w, "The user you requested does not exist", http.StatusNotFound)
			return
		} else if err == update.ErrUserDuplicate {
			http.Error(w, "The new username already exists", http.StatusNotAcceptable)
			return
		}

		json.NewEncoder(w).Encode("Informacoes do usuario atualizadas com sucesso")
	}
}

func deleteUser(s delete.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var id primitive.ObjectID
		err := json.NewDecoder(r.Body).Decode(&id)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.DeleteUser(id)
		if err == read.ErrUserNotFound {
			http.Error(w, "The user you requested does not exist", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode("Usuario excluido com sucesso")
	}
}
