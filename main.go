//gologger displays the port on which the server is running on path "/" and logs a random string on path "/logger/"
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

//constants of the server
const (
	defaultPath = "/"
	loggerPath  = "/logger"
	method      = "GET"
)

//vars of the app
var (
	address    *string
	port       *int
	codes      = []int{101, 105, 108, 140, 155, 1205, 666, 510, 419, 440}
	messages   = []string{"Unauthorized", "Service Not Found", "No Response From Server", "System Error", "Bad Request", "Forbidden", "Authentication Required", "Service Unavailable", "Timeout", "Not Supported"}
	operations = []string{"buyProduct", "getProducts", "queryBalance", "changeUser", "updateUser", "deleteUser", "cancelProduct", "getBalance", "getUser", "addUser"}
	users      = []string{"cthulhu", "yog-sothoth", "dagon", "hastur", "abhoth", "ubbo-sathla", "nyarlathotep", "shub-niggurath", "ghatanothoa", "shoggoth"}
)

func init() {
	address = flag.String("a", "0.0.0.0", "Hostname/IP address")
	port = flag.Int("p", 8080, "Port")
	flag.Parse()
}

func main() {
	err := run()
	if err != nil {
		log.Fatalf("%v\n", err.Error())
	}
}

func run() error {
	mux := http.NewServeMux()
	bind := fmt.Sprintf("%v:%v", *address, *port)

	mux.Handle(defaultPath, isRequestOk(runApp))
	mux.Handle(loggerPath, isRequestOk(loggerApp))

	log.Println("Server started...")
	return http.ListenAndServe(bind, mux)
}

//runApp displays port of server
func runApp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Running App On Port", port)
}

//loggerApp logs transactional data
func loggerApp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", Trans())
}

//isRequestOk is a middleware to validate a request
func isRequestOk(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path != defaultPath && r.URL.Path != loggerPath:
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		case r.Method != method:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		default:
			w.Header().Set("Content-Type", "text/plain")
			endpoint(w, r)
			return
		}
	})
}

//Trans returns a string like a transaction log with some fields separated by pipes
func Trans() string {
	rand.Seed(time.Now().UTC().UnixNano())
	status := "SUCCESS"
	code := 0
	msg := "SUCCESS"

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
	ro := rand.Intn(len(operations))
	operation := operations[ro]

	ri := rand.Intn(len(users))
	user := users[ri]

	timestamp := (time.Now().Local().Format("2006-01-02 15:04:05.000"))
	tid := time.Now().UnixNano()
	s := fmt.Sprintf("%v|%v|%v|%v|%v|%v|%v|%v", timestamp, tid, user, operation, duration, status, code, msg)
	return s
}
