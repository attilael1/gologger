//gologger displays the port on which the server is running on path "/" and logs a random string on path "/logger/"
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"
)

//constants of the server
const (
	defaultPath = "/"
	loggerPath  = "/logger"
	method      = "GET"
	//By sets creator
	By     = "attilael1"
	output = `{{.Timestamp }}|{{.Tid}}|{{.User}}|{{.Operation}}|{{.Duration}}|{{.Status}}|{{.Code}}|{{.Msg}}
`
)

//flags of the app
var (
	outputTmpl    *template.Template
	address, path *string
	port, ratio   *int
	version       *bool
	//Version sets app version
	Version string
	//GitCommit sets commit info
	GitCommit string
	//GitBranch sets git branch
	GitBranch string
	//BuildDate sets build date
	BuildDate string
)

type transaction struct {
	Timestamp string
	Tid       int64
	User      string
	Operation string
	Duration  int
	Status    string
	Code      int
	Msg       string
}

func init() {
	outputTmpl = template.Must(template.New("output").Parse(output))
	address = flag.String("a", "localhost", "Hostname/IP address")
	port = flag.Int("p", 8080, "Port")
	ratio = flag.Int("r", 5, "Failure ratio")
	version = flag.Bool("v", false, "Display version information")
	flag.Parse()
}

func main() {
	if *version == true {
		ShowVersion()
	}

	bind := fmt.Sprintf("%v:%v", *address, *port)
	http.HandleFunc(defaultPath, runApp)
	http.HandleFunc(loggerPath, logger)
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

//logger starts logging data to file
func logger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	if r.URL.Path != loggerPath {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method != method {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	t := getTransaction()
	outputTmpl.Execute(w, t)
}

//ShowVersion shows app version
func ShowVersion() {
	fmt.Printf("Version: %v, Commit: %v-%v, Build: %v, By: %v\n", Version, GitCommit, GitBranch, BuildDate, By)
	os.Exit(0)
}

//checkError takes a string and an error to log an error message
func checkError(m string, err error) {
	if err != nil {
		log.Fatalf("%v: %v\n", m, err.Error())
	}
}

func getTransaction() transaction {
	//var t transaction
	rand.Seed(time.Now().UTC().UnixNano())
	status := "SUCCESS"
	code := 0
	msg := "SUCCESS"
	codes := []int{101, 105, 108, 140, 155, 1205, 666, 510, 419, 440}
	messages := []string{"Unauthorized", "Service Not Found", "No Response From Server", "System Error", "Bad Request", "Forbidden", "Authentication Required", "Service Unavailable", "Timeout", "Not Supported"}

	//Set Success response ratio to %
	p := rand.Intn(10)
	if p > *ratio {
		rc := rand.Intn(len(codes))
		code = codes[rc]
		msg = messages[rc]
		status = "FAILED"
	}

	duration := rand.Intn(10000)
	operations := []string{"buyProduct", "getProducts", "queryBalance", "changeUser", "updateUser", "deleteUser", "cancelProduct", "getBalance", "getUser", "addUser"}
	ro := rand.Intn(len(operations))
	operation := operations[ro]

	users := []string{"cthulhu", "yog-sothoth", "dagon", "hastur", "abhoth", "ubbo-sathla", "nyarlathotep", "shub-niggurath", "ghatanothoa", "shoggoth"}
	ri := rand.Intn(len(users))
	user := users[ri]

	timestamp := (time.Now().Local()).Format("2006-01-02 15:04:05.000")
	tid := time.Now().UnixNano()
	return transaction{
		Code:      code,
		Status:    status,
		Msg:       msg,
		Duration:  duration,
		Operation: operation,
		User:      user,
		Timestamp: timestamp,
		Tid:       tid,
	}
}
