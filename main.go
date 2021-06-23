//gologger displays the port on which the server is running on path "/" and logs a random string on path "/logger/"
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//constants of the server
const (
	defaultPath = "/"
	loggerPath  = "/logger"
	method      = "GET"
)

//flags of the app
var (
	address, path *string
	port          *int
	version       *bool
)

func init() {
	address = flag.String("a", "localhost", "Hostname/IP address")
	port = flag.Int("p", 8080, "Port")
	version = flag.Bool("v", false, "Display version information")
	flag.Parse()
}

func main() {
	if *version {
		ShowVersion()
	}

	bind := fmt.Sprintf("%v:%v", *address, *port)
	http.HandleFunc(defaultPath, runApp)
	http.HandleFunc(loggerPath, loggerApp)
	log.Println("Server started...")

	err := http.ListenAndServe(bind, nil)
	checkError("ListenAndServe:", err)
}

//runApp displays port of server
func runApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	if r.URL.Path != defaultPath {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method != method {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "Running App On Port", *port)
}

//loggerApp starts logging data to file
func loggerApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	if r.URL.Path != loggerPath {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method != method {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "%v\n", Trans())
}

//ShowVersion shows app version
func ShowVersion() {
	fmt.Printf("Version: 1.0")
	os.Exit(0)
}

//checkError takes a string and an error to log an error message
func checkError(m string, err error) {
	if err != nil {
		log.Fatalf("%v: %v\n", m, err.Error())
	}
}

//Trans returns a string like a transaction log with some fields separated by pipes
func Trans() string {
	rand.Seed(time.Now().UTC().UnixNano())
	status := "SUCCESS"
	code := 0
	msg := "SUCCESS"
	codes := []int{101, 105, 108, 140, 155, 1205, 666, 510, 419, 440}
	messages := []string{"Unauthorized", "Service Not Found", "No Response From Server", "System Error", "Bad Request", "Forbidden", "Authentication Required", "Service Unavailable", "Timeout", "Not Supported"}

	//Set Success response ratio
	ratio := rand.Intn(10)
	p := rand.Intn(10)

	if p > ratio {
		rc := rand.Intn(len(codes))
		status = "FAILED"
		code = codes[rc]
		msg = messages[rc]
	}

	duration := rand.Intn(10000)
	operations := []string{"buyProduct", "getProducts", "queryBalance", "changeUser", "updateUser", "deleteUser", "cancelProduct", "getBalance", "getUser", "addUser"}
	ro := rand.Intn(len(operations))
	operation := operations[ro]

	users := []string{"cthulhu", "yog-sothoth", "dagon", "hastur", "abhoth", "ubbo-sathla", "nyarlathotep", "shub-niggurath", "ghatanothoa", "shoggoth"}
	ri := rand.Intn(len(users))
	user := users[ri]

	timestamp := (time.Now().Local().Format("2006-01-02 15:04:05.000"))
	tid := time.Now().UnixNano()
	s := fmt.Sprintf("%v|%v|%v|%v|%v|%v|%v|%v", timestamp, tid, user, operation, duration, status, code, msg)
	return s
}
