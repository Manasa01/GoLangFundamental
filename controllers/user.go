package controllers

import (
	"net/http"
	"regexp"
	"github.com/pluralsight/webservice/models"
	"encoding/json"
	"strconv"
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
	if r.URL.Path == "/users" {
		 switch r.Method{
		 case http.MethodGet: 
			 uc.getAll(w,r)
		 case http.MethodPost: 
			 uc.post(w,r)
		 default: 
		     w.WriteHeader(http.StatusNotImplemented)
		 }
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
			case http.MethodGet: 
			uc.get(id, w)
		case http.MethodPut: 
			uc.put(id, w,r)
		case http.MethodDelete: 
			uc.delete(id, w)
		default: 
		    w.WriteHeader(http.StatusNotImplemented)
		}

	}
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request){
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc *userController) get(id int, w http.ResponseWriter){
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request){
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	u,err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request){
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	u,err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) delete(id int, w http.ResponseWriter){
	err := models.RemoveUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (uc *userController) parseRequest(r *http.Request) (models.User, error){
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil{
		return models.User{}, err
	}
	return u,nil
}
