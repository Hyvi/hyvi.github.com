package main

import (
	"http"
)

func init() {
	http.HandleFunc("/view", viewRecord)
}

func viewRecord(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {

		http.Error(w, err.Error(), 500)
		return
	}

	if err := viewTemplate.Execute(w, record); err != nil {
		http.Error(w, err.Error(), 500)
	}

}

func main() {
	// add some more HTTP handlers and you quickly end up with many copies of identical error handling code
	// then reduce the repetition,  use own HTTP appHandler type
}

type appHandler func(http.ResponseWriter, *http.Request) error

// change our viewRecord function to return errors

func viewRecord2(w http.ResponseWriter, r *http.Request) error {
	// 省略相同的代码
	if err := datastore.Get(c, key, record); err != nill {
		return err
	}

	return viewTemplate.Execute(w, record)
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

// notes the method's receiver, fn, is a function ,
// the method invoke the function by call the receiver in the expression fn(w, r)

func init2() {
	// convert the function viewRecord2 to type appHandler
	http.Handle("/view", appHandler(viewRecord2))
}

// better to give the user a simple error message with an appropriate HTTP status code
// while logging the full error to app Engine developer console for debugging purposes

type appError struct {
	Error   error
	Message string
	Code    int
}

type appHandler func(http.ResponseWriter, *http.Request) *appError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Reqeust) {
	if e := fn(w, r); e != nil {
		c := appengine.NewContext(r)
		c.ErrorF("%v", e.Error)
		http.Error(w, e.Message, e.Code)
	}
}

func viewRecord3(w http.ResponseWriter, r *http.Request) *appError {
	// 省略
	if err := datastore.Get(c, key, record); err != nil {
		return &appError{err, "Record not found", 404}
	}
	if err := viewTemplate.Execute(w, record); err != nil {
		return &appError{err, "Can't display record", 500}
	}

	return nil
}

// it doesn't end here,

// give the error handler a pretty HTML template

// make debuggging easier by writing the stack trace to the HTTP response when the user is administrator,

// write a constructor function for appError that stores the stack trace for easier debugging,

// recover from panics inside the appHandler, logging the error to console as  "Critical" while telling the user " a serious error has occurred"
