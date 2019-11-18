package main

import (
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

var tpl *template.Template
var encryptionKey = "something-very-secret"
var loggedUserSession = sessions.NewCookieStore([]byte(encryptionKey))
var store *sessions.CookieStore

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
type Article struct {
	Id        int
	Title     string
	Subtitle  string
	Content   string
	Publisher string
	Published string
}

/*type Sess struct {
	Id       int
	Username string
	Role     string
}*/

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)

	store = sessions.NewCookieStore(
		authKeyOne,
		encryptionKeyOne,
	)

	store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}

	gob.Register(User{})

	tpl = template.Must(template.ParseGlob("templates/*.html"))
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
func GetFoo(f interface{}) string {
	if f != nil {
		if foo, ok := f.(string); ok {
			return foo
		}
	}
	return ""
}
func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.Handle("/notika/", http.StripPrefix("/notika/", http.FileServer(http.Dir("notika"))))

	http.Handle("/blog/", http.StripPrefix("/blog/", http.FileServer(http.Dir("blog"))))
	http.HandleFunc("/login", login)                    // set router
	http.HandleFunc("/register", register)              // set router
	http.HandleFunc("/logout", LogoutHandler)           // set router
	http.HandleFunc("/dashboard", DashBoardPageHandler) // set router
	http.HandleFunc("/edit", Edit)                      // set router
	http.HandleFunc("/forbidden", forbidden)
	http.HandleFunc("/submit_edit", submit_edit)
	http.HandleFunc("/delete", deleteHandler)
	//Blog Routes
	http.HandleFunc("/index", home)
	http.HandleFunc("/post_detail", post_detail)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/submit_contact", submit_contact)
	http.HandleFunc("/article_list", article_list)
	http.HandleFunc("/article_edit", article_edit)
	http.HandleFunc("/article_new", article_new)
	http.HandleFunc("/article_update", article_update)
	http.HandleFunc("/article_create", article_create)
	http.HandleFunc("/article_delete", article_delete)
	err := http.ListenAndServe(":9090", nil) // set listen port
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
		var databaseRole string
		var databaseId string

		err = db.QueryRow("SELECT id,username, password, role FROM users WHERE username=?", username).Scan(&databaseId, &databaseUsername, &databasePassword, &databaseRole)

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

			session.Values["role"] = databaseRole
			session.Values["user_id"] = databaseId
			err := session.Save(r, w)

			if err != nil {
				log.Println(err)
			}

			http.Redirect(w, r, "/index", http.StatusFound)
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
	session.Values["role"] = ""
	session.Values["user_id"] = ""
	session.Options.MaxAge = -1
	err := session.Save(r, w)

	if err != nil {
		log.Println(err)
	}
	w.Write([]byte("Logged out!"))
	//http.Redirect(w, r, "/index", 200)
}

func DashBoardPageHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./templates/dashboard.html")
	conditionsMap := map[string]interface{}{}
	//read from session
	session, err := loggedUserSession.Get(r, "authenticated-user-session")

	if err != nil || session.Values["usernamme"] == "" {
		log.Println("Unable to retrieve session data!", err)
	}

	log.Println("Session name : ", session.Name())

	log.Println("Username : ", session.Values["username"])

	log.Println("Role : ", session.Values["role"])

	conditionsMap["Username"] = session.Values["username"]
	conditionsMap["Id"] = session.Values["user_id"]
	conditionsMap["Role"] = session.Values["role"]
	if session.Values["role"] == "admin" {
		conditionsMap["deletes"] = "Ya"
	} else {
		conditionsMap["deletes"] = "No"
	}
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

	if err := t.Execute(w, map[string]interface{}{
		"listuser": res,
		"sesi":     conditionsMap,
	}); err != nil {
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

	nId := r.URL.Query().Get("id")
	session, err := loggedUserSession.Get(r, "authenticated-user-session")

	if nId != session.Values["user_id"] && session.Values["role"] != "admin" {
		session.AddFlash("You don't have access!")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/forbidden", http.StatusFound)
		return
	} else {

		t, _ := template.ParseFiles("./templates/edit.html")
		db, err := sql.Open("mysql", "root:@/test?charset=utf8")
		checkErr(err)
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
}
func forbidden(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./templates/forbidden.html")
	flashMessages := "Anda Tidak Berhak Mengedit User ini"
	//tpl.ExecuteTemplate(w, "/templates/forbidden.html", flashMessages)
	if err := t.Execute(w, flashMessages); err != nil {
		log.Println(err)
	}
}
func submit_edit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "POST" {
		db, err := sql.Open("mysql", "root:@/test?charset=utf8")
		checkErr(err)

		username := r.FormValue("username")
		name := r.FormValue("name")
		email := r.FormValue("email")
		id := r.FormValue("user_id")

		// update
		stmt, err := db.Prepare("update `users` set username=?,name=?,email=? where id=?")
		checkErr(err)

		_, err = stmt.Exec(username, name, email, id)
		checkErr(err)

		http.Redirect(w, r, "/dashboard", 301)
	}
}
func deleteHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		nId := r.URL.Query().Get("id")
		db, err := sql.Open("mysql", "root:@/test?charset=utf8")
		checkErr(err)

		// update
		stmt, err := db.Prepare("DELETE FROM `users` where id=?")
		checkErr(err)

		_, err = stmt.Exec(nId)
		checkErr(err)

		http.Redirect(w, r, "/dashboard", 301)
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./templates/home.html")
		conditionsMap := map[string]interface{}{}
		session, err := loggedUserSession.Get(r, "authenticated-user-session")

		if err != nil {
			log.Println("Unable to retrieve session data!", err)

		}
		foo := GetFoo(session.Values["username"])
		if foo == "" {

			conditionsMap["logged"] = 0
		} else {

			conditionsMap["logged"] = 1
		}

		log.Println("Session name : ", session.Name())

		log.Println("Username : ", session.Values["username"])

		log.Println("Role : ", session.Values["role"])
		conditionsMap["Username"] = session.Values["username"]
		db, err := sql.Open("mysql", "root:@/test?charset=utf8")
		checkErr(err)
		selDB, err := db.Query("SELECT id,title,subtitle,content,publisher,published FROM article ORDER BY id DESC")
		if err != nil {
			panic(err.Error())
		}
		emp := Article{}
		res := []Article{}
		for selDB.Next() {
			var id int
			var title, subtitle, content, publisher, published string
			err = selDB.Scan(&id, &title, &subtitle, &content, &publisher, &published)
			if err != nil {
				panic(err.Error())
			}
			emp.Id = id
			emp.Title = title
			emp.Subtitle = subtitle
			emp.Content = content
			emp.Publisher = publisher
			emp.Published = published
			res = append(res, emp)
		}

		if err := t.Execute(w, map[string]interface{}{
			"listarticle": res,
			"sesi":        conditionsMap,
		}); err != nil {
			log.Println(err)
		}
		//tmpl.ExecuteTemplate(w, "Index", res)
		defer db.Close()
	}
}
func post_detail(w http.ResponseWriter, r *http.Request) {

	nId := r.URL.Query().Get("id")

	t, _ := template.ParseFiles("./templates/post.html")
	conditionsMap := map[string]interface{}{}
	session, err := loggedUserSession.Get(r, "authenticated-user-session")

	if err != nil {
		log.Println("Unable to retrieve session data!", err)

	}
	foo := GetFoo(session.Values["username"])
	if foo == "" {

		conditionsMap["logged"] = 0
	} else {

		conditionsMap["logged"] = 1
	}

	log.Println("Session name : ", session.Name())

	log.Println("Username : ", session.Values["username"])

	log.Println("Role : ", session.Values["role"])
	conditionsMap["Username"] = session.Values["username"]
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)
	selDB, err := db.Query("SELECT id,title,subtitle,content,publisher,published FROM article WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Article{}
	for selDB.Next() {
		var id int
		var title, subtitle, content, publisher, published string
		err = selDB.Scan(&id, &title, &subtitle, &content, &publisher, &published)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Title = title
		emp.Subtitle = subtitle
		emp.Content = content
		emp.Publisher = publisher
		emp.Published = published
	}
	if err := t.Execute(w, map[string]interface{}{
		"article": emp,
		"sesi":    conditionsMap,
	}); err != nil {
		log.Println(err)
	}
	//tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()

}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./templates/about.html")
		conditionsMap := map[string]interface{}{}
		session, err := loggedUserSession.Get(r, "authenticated-user-session")

		if err != nil {
			log.Println("Unable to retrieve session data!", err)

		}
		foo := GetFoo(session.Values["username"])
		if foo == "" {

			conditionsMap["logged"] = 0
		} else {

			conditionsMap["logged"] = 1
		}
		conditionsMap["Username"] = session.Values["username"]
		t.Execute(w, map[string]interface{}{
			"sesi": conditionsMap,
		})
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./templates/contact.html")
		conditionsMap := map[string]interface{}{}
		session, err := loggedUserSession.Get(r, "authenticated-user-session")

		if err != nil {
			log.Println("Unable to retrieve session data!", err)

		}
		foo := GetFoo(session.Values["username"])
		if foo == "" {

			conditionsMap["logged"] = 0
		} else {

			conditionsMap["logged"] = 1
		}
		conditionsMap["Username"] = session.Values["username"]
		t.Execute(w, map[string]interface{}{
			"sesi": conditionsMap,
		})
	}
}
func submit_contact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "POST" {
		db, err := sql.Open("mysql", "root:@/test?charset=utf8")
		checkErr(err)

		name := r.FormValue("name")
		email := r.FormValue("email")
		phone := r.FormValue("phone")
		message := r.FormValue("message")

		// update
		stmt, err := db.Prepare("insert into `contact` VALUES(NULL,?,?,?,?,NOW())")
		checkErr(err)

		_, err = stmt.Exec(name, email, phone, message)
		checkErr(err)

		http.Redirect(w, r, "/contact", 301)
	}
}
func article_list(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./templates/article_list.html")
	conditionsMap := map[string]interface{}{}
	//read from session
	session, err := loggedUserSession.Get(r, "authenticated-user-session")

	if err != nil {
		log.Println("Unable to retrieve session data!", err)

	}
	foo := GetFoo(session.Values["username"])
	if foo == "" {
		http.Redirect(w, r, "/dashboard", http.StatusFound)
	}

	log.Println("Session name : ", session.Name())

	log.Println("Username : ", session.Values["username"])

	log.Println("Role : ", session.Values["role"])

	conditionsMap["Username"] = session.Values["username"]
	conditionsMap["Id"] = session.Values["user_id"]
	conditionsMap["Role"] = session.Values["role"]
	if session.Values["role"] == "admin" {
		conditionsMap["deletes"] = "Ya"
	} else {
		conditionsMap["deletes"] = "No"
	}
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)
	selDB, err := db.Query("SELECT id,title,subtitle,content,publisher,published FROM article ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Article{}
	res := []Article{}
	for selDB.Next() {
		var id int
		var title, subtitle, content, publisher, published string
		err = selDB.Scan(&id, &title, &subtitle, &content, &publisher, &published)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Title = title
		emp.Subtitle = subtitle
		emp.Content = content
		emp.Publisher = publisher
		emp.Published = published
		res = append(res, emp)
	}

	if err := t.Execute(w, map[string]interface{}{
		"listarticle": res,
		"sesi":        conditionsMap,
	}); err != nil {
		log.Println(err)
	}
	//tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}
func article_edit(w http.ResponseWriter, r *http.Request) {

	nId := r.URL.Query().Get("id")
	session, err := loggedUserSession.Get(r, "authenticated-user-session")

	if nId != session.Values["user_id"] && session.Values["role"] != "admin" {
		session.AddFlash("You don't have access!")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/index", http.StatusFound)
		return
	} else {

		t, _ := template.ParseFiles("./templates/article_edit.html")
		db, err := sql.Open("mysql", "root:@/test?charset=utf8")
		checkErr(err)
		selDB, err := db.Query("SELECT id,title,subtitle,content,publisher,published FROM article WHERE id=?", nId)
		if err != nil {
			panic(err.Error())
		}
		emp := Article{}
		for selDB.Next() {
			var id int
			var title, subtitle, content, publisher, published string
			err = selDB.Scan(&id, &title, &subtitle, &content, &publisher, &published)
			if err != nil {
				panic(err.Error())
			}
			emp.Id = id
			emp.Title = title
			emp.Subtitle = subtitle
			emp.Content = content
			emp.Publisher = publisher
			emp.Published = published
		}
		if err := t.Execute(w, emp); err != nil {
			log.Println(err)
		}
		//tmpl.ExecuteTemplate(w, "Index", res)
		defer db.Close()
	}
}
func article_new(w http.ResponseWriter, r *http.Request) {

	session, err := loggedUserSession.Get(r, "authenticated-user-session")
	foo := GetFoo(session.Values["username"])
	if foo == "" {

		session.AddFlash("You don't have access!")
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/forbidden", http.StatusFound)
		return
	} else {

		t, _ := template.ParseFiles("./templates/article_new.html")

		if err := t.Execute(w, nil); err != nil {
			log.Println(err)
		}
		//tmpl.ExecuteTemplate(w, "Index", res)

	}
}
func article_update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "POST" {
		db, err := sql.Open("mysql", "root:@/test?charset=utf8")
		checkErr(err)

		title := r.FormValue("title")
		subtitle := r.FormValue("subtitle")
		content := r.FormValue("content")
		publisher := r.FormValue("publisher")
		id := r.FormValue("id")

		// update
		stmt, err := db.Prepare("update `article` set title=?,subtitle=?,content=?,publisher=? where id=?")
		checkErr(err)

		_, err = stmt.Exec(title, subtitle, content, publisher, id)
		checkErr(err)

		http.Redirect(w, r, "/article_list", http.StatusFound)
	}
}
func article_create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "POST" {
		db, err := sql.Open("mysql", "root:@/test?charset=utf8")
		checkErr(err)

		title := r.FormValue("title")
		subtitle := r.FormValue("subtitle")
		content := r.FormValue("content")
		publisher := r.FormValue("publisher")

		// update
		stmt, err := db.Prepare("insert into `article` set title=?,subtitle=?,content=?,publisher=?,published=CURDATE(),created_on=NOW()")
		checkErr(err)

		_, err = stmt.Exec(title, subtitle, content, publisher)
		checkErr(err)

		http.Redirect(w, r, "/article_list", http.StatusFound)
	}
}
func article_delete(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		nId := r.URL.Query().Get("id")
		db, err := sql.Open("mysql", "root:@/test?charset=utf8")
		checkErr(err)

		// update
		stmt, err := db.Prepare("DELETE FROM `article` where id=?")
		checkErr(err)

		_, err = stmt.Exec(nId)
		checkErr(err)

		http.Redirect(w, r, "/article_list", http.StatusFound)
	}
}
