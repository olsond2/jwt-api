package main


import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"os"
	"github.com/dgrijalva/jwt-go"
	"time"
	"encoding/json"
	"github.com/auth0/go-jwt-middleware"
)

func main() {
	myRouter := mux.NewRouter()


	myRouter.Handle("/", http.FileServer(http.Dir("./views/")))
	//myRouter.Handle("/status", NotImplemented).Methods("GET")
	//myRouter.Handle("/products", NotImplemented).Methods("GET")
	myRouter.Handle("/get-token", GetTokenHandler).Methods("GET")
	myRouter.Handle("/secure", jwtMiddleware.Handler(ProductsHandler)).Methods("GET")


	myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":7000",handlers.LoggingHandler(os.Stdout, myRouter))


}


var mySigningKey = []byte("secret")


/*	HANDLERS	*/

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)


	
	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = true
	claims["name"] = "John Doe"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()


	tokenString, _ := token.SignedString(mySigningKey)


	w.Write([]byte(tokenString))

})

var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	
	var products struct{}
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))

})


var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
