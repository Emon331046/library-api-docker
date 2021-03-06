package main

import (
	"github.com/Emon331046/library-management-api/pkg"
	"github.com/Emon331046/library-management-api/pkg/book"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Middleware(next http.Handler) http.Handler {
	println("hello")
	return next
}
func init() {

	pkg.Users = append(pkg.Users, pkg.User{
		ID:       pkg.UserIdCount,
		Name:     "demo user",
		Mail:     "demouser@mail.com",
		Password: "123456",
		PhoneNo:  "019999999",
		UserType: "user",
	})
	pkg.UserIdCount++
	pkg.Users = append(pkg.Users, pkg.User{
		ID:       pkg.UserIdCount,
		Name:     "admin",
		Mail:     "admin@mail.com",
		Password: "123456",
		PhoneNo:  "019999999",
		UserType: "admin",
	})
	pkg.UserIdCount++

	book.BookIDCount++
	book.Books[book.BookIDCount] = book.Book{
		Id:        book.BookIDCount,
		BookName:  "himu",
		Author:    "humayon ahmed",
		Available: true,
	}

	book.BookIDCount++
	book.Books[book.BookIDCount] = book.Book{
		Id:        book.BookIDCount,
		BookName:  "misir ali",
		Author:    "humayon ahmed",
		Available: true,
	}

}
func main() {
	//tokenString, err := pkg.GenerateJWT("admin@mail.com", "admin", 1)
	//if err == nil {
	//	fmt.Println(tokenString)
	//} else {
	//	fmt.Println(err)
	//}
	r := mux.NewRouter()
	r.Use(pkg.JwtMiddleWare)

	r.HandleFunc("/login", pkg.Login).
		Methods("GET")
	r.HandleFunc("/register", pkg.Register).
		Methods("POST")
	r.HandleFunc("/user/{user_id}", pkg.UserProfile).
		Methods("Get")
	r.HandleFunc("/edit-profile/{user_id}", pkg.EditUserProfile).
		Methods("PATCH")
	//r.Handle("/edit-profile/{user_id}", AuthMiddleware(pkg.EditUserProfile(), "user")).
	//	Methods("PATCH")
	r.HandleFunc("/book", book.AddNewBook).
		Methods("POST")
	r.HandleFunc("/book", book.ShowAllBooks).
		Methods("Get")
	r.HandleFunc("/book/{book_id}", book.ShowBook).
		Methods("Get")
	//http.HandleFunc("/df",fu)

	r.HandleFunc("/purchase-book", book.AddNewPurchase).
		Methods("POST")
	r.HandleFunc("/return-book", book.ReturnBook).
		Methods("PUT")
	r.HandleFunc("/delete-book/{book_id}", book.DeleteBook)
	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
