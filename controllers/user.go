package controllers

import (
	"net/http"
	"regexp"
) 
// an emptystruct is used sometimes to collect all the related functions/methods together
//though the struct itself maynot have any data value

//custom type
type userController struct {
    userIDPattern *regexp.Regexp
}
//dont have constructors in Go- instead constructorFunc is used -> depends on how a func is used rather than it being a special function
func newUserController() *userController{ //conventional name for constructor
     return &userController{
		 userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	 }
}
//no tight coupling bw methods and struct/object
//func, specify a type to bind the func to => this makes function to a method, method name, params
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from the User Controller!"))
}