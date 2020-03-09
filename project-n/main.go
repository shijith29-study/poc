// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	// "fmt" has methods for formatted I/O operations (like printing to the console)
	"fmt"
	// The "net/http" library has methods to implement HTTP clients and servers
	"net/http"
	"html/template"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Tag struct {

    ID string `json:"id"`
    Name string `json:"name"`
}

var tpl = template.Must(template.ParseFiles("index.html"))

func main() {
	// The "HandleFunc" method accepts a path and a function as arguments
	// (Yes, we can pass functions as arguments, and even trat them like variables in Go)
	// However, the handler function has to have the appropriate signature (as described by the "handler" function below)
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/int05")
	//http.HandleFunc("/", htmlOutput)
    //http.HandleFunc("/search", searchOutput)
	// After defining our server, we finally "listen and serve" on port 8080
	// The second argument is the handler, which we will come to later on, but for now it is left as nil,
	// and the handler defined above (in "HandleFunc") is used

    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

    // perform a db.Query insert
    selectQ, err := db.Query("SELECT id, name FROM Accounts WHERE id = '0002d4bc-c8f6-11e7-88a4-0273927739c9'")

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    //defer selectQ.Close()
    //fmt.Println("i am here to check db")

    for selectQ.Next() {
        var tag Tag
        // for each row, scan the result into our tag composite object
        err = selectQ.Scan(&tag.ID, &tag.Name)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        fmt.Printf(tag.Name)
    }
    http.HandleFunc("/", htmlOutput)
	http.ListenAndServe(":8080", nil)
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.

func searchOutput(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hello i am in search world")
    fmt.Println("i am here to check db")
}

func htmlOutput(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	//fmt.Fprintf(w, "Hello World oooo - shijith - output!")
	tpl.Execute(w, nil)
}
