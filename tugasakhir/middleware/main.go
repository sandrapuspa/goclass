package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

var tpl *template.Template
var encryptionKey = "something-very-secret"
var loggedUserSession = sessions.NewCookieStore([]byte(encryptionKey))

type Employee struct {
	Id       int
	Name     string
	Username string
	Email    string
	Role     string
}
type User struct {
	Id       int
	Name     string
	Username string
	Email    string
	Role     string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gtpl"))
}

type Response struct {
	status  int    `json:"status"`
	message string `json:"message"`
	//data    []users
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "test"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.Handle("/notika/", http.StripPrefix("/notika/", http.FileServer(http.Dir("notika"))))
	http.HandleFunc("/login", login)                    // set router
	http.HandleFunc("/register", register)              // set router
	http.HandleFunc("/logout", LogoutHandler)           // set router
	http.HandleFunc("/dashboard", DashBoardPageHandler) // set router
	http.HandleFunc("/edit", Edit)                      // set router

	http.HandleFunc("/contact", contact)     // set router
	err := http.ListenAndServe(":3000", nil) // set listen port
	if err != nil {
		log.Fatal("Error running service: ", err)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./templates/login.html")
		t.Execute(w, nil)
	} else {
		conditionsMap := map[string]interface{}{}
		db, err := sql.Open("mysql", "root:@/test?charset=utf8")
		checkErr(err)
		username := r.FormValue("username")
		password := r.FormValue("password")

		var databaseUsername string
		var databasePassword string

		err = db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

		if err != nil {
			http.Redirect(w, r, "/login", 301)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
		if err != nil {
			http.Redirect(w, r, "/login", 301)
			return
		} else {
			log.Println("Logged in :", username)
			conditionsMap["Username"] = username
			conditionsMap["LoginError"] = false

			// create a new session and redirect to dashboard
			session, _ := loggedUserSession.New(r, "authenticated-user-session")

			session.Values["username"] = username
			err := session.Save(r, w)

			if err != nil {
				log.Println(err)
			}

			http.Redirect(w, r, "/dashboard", http.StatusFound)
			//w.Write([]byte("Hello" + databaseUsername))
		}

	}
}
func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./templates/register.html")
		t.Execute(w, nil)
	} else {
		db, err := sql.Open("mysql", "root:@/test?charset=utf8")
		checkErr(err)

		// insert
		//stmt, err := db.Prepare("insert into `users` values (null, ?, ?, ?, ?)")
		///checkErr(err)

		//res, err := stmt.Exec(r.FormValue("username"), r.FormValue("firstname"), r.FormValue("lastname"), r.FormValue("password"))
		//checkErr(err)
		//affect, err := res.RowsAffected()
		//checkErr(err)

		//fmt.Println(affect)

		//db.Close()

		username := r.FormValue("username")
		password := r.FormValue("password")
		name := r.FormValue("name")
		email := r.FormValue("email")

		var user string

		err = db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

		switch {
		case err == sql.ErrNoRows:
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, "Server error, unable to create your account.", 500)
				return
			}

			_, err = db.Exec("INSERT INTO users(username, password,name,email,role) VALUES(?, ?,?,?,'user')", username, hashedPassword, name, email)
			if err != nil {
				http.Error(w, "Server error, unable to create your account.", 500)
				return
			}

			w.Write([]byte("User created!"))
			http.Redirect(w, r, "/login", 301)
			return
		case err != nil:
			http.Error(w, "Server error, unable to create your account.", 500)
			return
		default:
			http.Redirect(w, r, "/register", 301)
		}
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//read from session
	session, _ := loggedUserSession.Get(r, "authenticated-user-session")

	// remove the username
	session.Values["username"] = ""
	err := session.Save(r, w)

	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/login", 301)
	w.Write([]byte("Logged out!"))
}
func DashBoardPageHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./templates/dashboard.html")
	conditionsMap := map[string]interface{}{}
	//read from session
	session, err := loggedUserSession.Get(r, "authenticated-user-session")

	if err != nil {
		log.Println("Unable to retrieve session data!", err)
	}

	log.Println("Session name : ", session.Name())

	log.Println("Username : ", session.Values["username"])

	conditionsMap["Username"] = session.Values["username"]

	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)
	selDB, err := db.Query("SELECT id,name,username,email,role FROM users ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := User{}
	res := []User{}
	for selDB.Next() {
		var id int
		var name, username, email, role string
		err = selDB.Scan(&id, &name, &username, &email, &role)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Username = username
		emp.Email = email
		emp.Role = role
		res = append(res, emp)
	}

	if err := t.Execute(w, res); err != nil {
		log.Println(err)
	}
	//tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}
func getUsersFromDB() []byte {

	var (
		user  User
		users []User
	)
	rows, err := db.Query("SELECT id,username,name,email,role FROM users")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Name, &user.Email, &user.Role)
		users = append(users, user)
	}
	defer rows.Close()
	jsonResponse, jsonError := json.Marshal(users)
	if jsonError != nil {
		fmt.Println(jsonError)
		return nil
	}
	return jsonResponse
}
func Edit(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./templates/edit.html")
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT id,name,username,email,role FROM users WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := User{}
	for selDB.Next() {
		var id int
		var name, username, email, role string
		err = selDB.Scan(&id, &name, &username, &email, &role)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Username = username
		emp.Email = email
		emp.Role = role
	}
	if err := t.Execute(w, emp); err != nil {
		log.Println(err)
	}
	//tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./templates/login.html")
		t.Execute(w, nil)
	} else {
		//var response Response
		//conditionsMap := map[string]interface{}{}

		name := r.FormValue("name")
		phone := r.FormValue("phone")
		email := r.FormValue("email")
		message := r.FormValue("message")
		db, err := sql.Open("mysql", "root:@/tugasakhir?charset=utf8")
		checkErr(err)

		//var databaseUsername string
		//var databasePassword string

		stmt, err := db.Prepare("INSERT INTO ta_contact VALUES(NULL,?,?,?,?,NOW())")
		checkErr(err)

		_, err = stmt.Exec(name, phone, email, message)

		db.Close()
		if err == nil {
			resp := struct {
				Status  int
				Message string
			}{1, "Sukses"}

			jsonInBytes, err := json.Marshal(resp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonInBytes)
		} else {
			log.Println(err)
		}

	}
}
