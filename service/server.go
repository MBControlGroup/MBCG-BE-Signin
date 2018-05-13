package service

import (
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

    formatter := render.New(render.Options{
        IndentJSON: true,
    })

    n := negroni.Classic()
    mx := mux.NewRouter()

    initRoutes(mx, formatter)

    n.UseHandler(mx)
    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
    mx.HandleFunc("/testToken", testToken(formatter)).Methods("GET")
    mx.HandleFunc("/tokenValid", tokenValid(formatter)).Methods("POST")
    mx.HandleFunc("/signin", signinHandler(formatter)).Methods("POST")
    mx.HandleFunc("/signout", signoutHandler(formatter)).Methods("POST")
    mx.HandleFunc("/pmanage/admin", addAdminHandler(formatter)).Methods("POST")
    mx.HandleFunc("/pmanage/IMUsers", addIMUserHandler(formatter)).Methods("POST")

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}