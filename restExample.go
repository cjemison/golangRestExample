package main
import (
    "net/http"
    "encoding/json"
    "log"
    "github.com/gorilla/mux"
)

type Message struct {
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
}

type Messages[]Message

func addHeaders(response http.ResponseWriter,  request *http.Request){
    response.Header().Set("Accept", "application/json; charset=utf-8")
    response.Header().Set("Content-Type", "application/json; charset=utf-8")
    if origin := request.Header.Get("Origin"); origin != "" {
          response.Header().Set("Access-Control-Allow-Origin", origin)
    } else {
        response.Header().Set("Access-Control-Allow-Origin", "*")
    }
    response.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
    response.Header().Set("Access-Control-Allow-Credentials", "true")
}

func handler(response http.ResponseWriter, request *http.Request) {
    log.Printf("%s %s %s", request.RemoteAddr, request.Method, request.URL)

    vars := mux.Vars(request)
    log.Printf("%s", vars)
    firstName := vars["firstName"]
    lastName := vars["lastName"]

    messages := Messages{
       Message {FirstName: firstName, LastName: lastName},
    }
    log.Printf("Message: %s", messages)

    addHeaders(response, request)
    if err := json.NewEncoder(response).Encode(messages); err != nil {
        panic(err)
    }
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", handler).Methods("GET")
    router.HandleFunc("/{firstName}/{lastName}", handler).Methods("GET")
    log.Fatal(http.ListenAndServe(":8001", router))
}
