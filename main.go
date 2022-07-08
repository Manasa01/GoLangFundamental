package main

import( 
	 "net/http"
	 "github.com/pluralsight/webservice/controllers" 

)

func main(){
	// u := models.User{
	// 	ID: 2,
	// 	FirstName: "Tricia",
	// 	LastName: "McMillan",
	// }
	// fmt.Println("hello from module", u)
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil) //Front and Back controller pattern,
	// nil => default ServeMux - Server multiplexer which acts like a front controller 
	//and sends /redirects the requests to route (back controller) defined with Handle method
}