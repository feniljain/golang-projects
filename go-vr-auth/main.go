package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/gddo/httputil/header"
	"io"
	"log"
	"net/http"
	"strings"

	//"time"

	//"github.com/auth0/go-jwt-middleware"
	//"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}

type product struct {
	ID          int
	Name        string
	Slug        string
	Description string
}

//User represents a user
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
	Phone:    "1234567890",
}

var products = []product{
	{ID: 1, Name: "a", Slug: "S1", Description: "desc1"},
	{ID: 2, Name: "b", Slug: "S2", Description: "desc2"},
	{ID: 3, Name: "c", Slug: "S3", Description: "desc3"},
	{ID: 4, Name: "d", Slug: "S4", Description: "desc4"},
}

func main() {

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.Handle("/status", statusHandler).Methods("GET")
	r.Handle("/products", (productsHandler)).Methods("GET")
	r.Handle("/products/{slug}/feedback", (addFeedbackHandler)).Methods("POST")
	//r.Handle("/login", loginHandler).Methods("POST")

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	serv := http.Server{
		Handler: corsWrapper.Handler(r),
		Addr:    ":8080",
		//WriteTimeout: 15 * time.Second,
		//ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Listening on port 8080")
	serv.ListenAndServe()
}

var notImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

var statusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

var productsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	payload, _ := json.Marshal(products)
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(payload))
})

var addFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var matchedProduct product

	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			matchedProduct = p
			break
		}
	}

	bs := make([]byte, 99999)
	_, err := r.Body.Read(bs)
	//fmt.Println(string(bs))
	if err != nil {
		//log.Print(err)
		//http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		//return
	}
	str := make([]string, len(bs))
	//fmt.Println(string(bs))
	for i, b := range bs {
		str[i] = strings.Trim(string(b), "")
	}
	//fmt.Println(str)
	defer r.Body.Close()

	var user User
	//fmt.Println(string(bs))
	err = json.Unmarshal(bs, &user)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusUnprocessableEntity), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-type", "application/json")
	if matchedProduct.Slug != "" {
		payload, _ := json.Marshal(matchedProduct)
		w.Write([]byte(payload))
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
})

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return &malformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &malformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		return &malformedRequest{status: http.StatusBadRequest, msg: msg}
	}

	return nil
}
