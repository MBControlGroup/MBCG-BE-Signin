package service

import (
	"encoding/json"
	//"fmt"
    "net/http"
    "strconv"
    "fmt"
    "github.com/MBControlGroup/MBCG-BE-Login/entities"
    "github.com/MBControlGroup/MBCG-BE-Login/token"
    "github.com/unrolled/render"
    //"github.com/dgrijalva/jwt-go/request"
    //"github.com/dgrijalva/jwt-go"

)

func signinHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        var user entities.UserInfo
        
        err := json.NewDecoder(req.Body).Decode(&user)
        fmt.Println(user)
        checkErr(err)

        u := entities.LoginService.AdminFindByAccount(user.Username)
        //u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["username"][0]})
        
        if u==nil || u.Admin_passwd != user.Password {
            formatter.JSON(w, http.StatusBadRequest, struct{ Code int;Enmsg string;Cnmsg string; Data interface{}}{400, "fail", "失败", nil})
        } else {
            //fmt.Println(u.Admin_id)
            tokenString, err := token.Generate(u.Admin_id)
            cookie := http.Cookie{Name:"token", Value:tokenString, Path:"/", MaxAge:86400}
            http.SetCookie(w, &cookie)
            checkErr(err)
            formatter.JSON(w, http.StatusOK, struct{ Code int;Enmsg string;Cnmsg string; Data interface{}}{200, "ok", "成功", nil})
        }
    }
} 

func signoutHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        cookie, err := req.Cookie("token")

        if err != nil || cookie.Value == "" {
            formatter.JSON(w, http.StatusBadRequest, struct{ Code int;Enmsg string;Cnmsg string; Data interface{}}{302, "fail", "失败", nil})
	        return;
        }

	    _, err = token.Valid(cookie.Value)

	    if err != nil {
	        formatter.JSON(w, http.StatusBadRequest, struct{ Code int;Enmsg string;Cnmsg string; Data interface{}}{302, "fail", "失败", nil})
	        return;
        }

        cookie = &http.Cookie{Name: "token", Path: "/", MaxAge: -1}
        http.SetCookie(w, cookie)
        formatter.JSON(w, http.StatusOK, struct{ Code int;Enmsg string;Cnmsg string; Data interface{}}{200, "ok", "成功", nil})
    }
} 

func addAdminHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        fmt.Println(123)
        req.ParseForm()
        var admin entities.Admins
        //fmt.Println(req.Form["admin_password"][0])
        fmt.Println(123)
        admin.Admin_passwd = req.Form["admin_password"][0]
        admin.Admin_account = req.Form["admin_account"][0]
        //admin_type, err := strconv.Atoi(req.Form["admin_type"][0])
        //checkErr(err)
        admin.Admin_type = req.Form["admin_type"][0]
        admin_Im_user_id, err := strconv.Atoi(req.Form["admin_im_user_id"][0])
        checkErr(err)
        admin.Im_user_id = admin_Im_user_id
        fmt.Println(admin)
        entities.LoginService.AdminSave(&admin)
    }
}

func addIMUserHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        entities.LoginService.AddIMUser()
    }
}

func testToken(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {

        cookie, err := req.Cookie("token")
	    if err != nil || cookie.Value == ""{
	        formatter.JSON(w, 403, struct{Error string}{"token not found."})
	        return;
	    }

	    user_id, err := token.Valid(cookie.Value)

	    if err != nil {
	        formatter.JSON(w, 403, struct{Error string}{"bad token"})
	        return;
        }

        id, err := strconv.Atoi(user_id)
        admin := entities.LoginService.AdminFindById(id)
        //fmt.Println(user_id)

        formatter.JSON(w, http.StatusOK, 
            struct{ Success bool;
                    Content string;
                    AdminInfo entities.Admins}{
                    true, 
                    "The token is valid.",
                    *admin})
    }
}

func tokenValid(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()

        if len(req.Form["token"]) == 0 {
            formatter.JSON(w, http.StatusBadRequest, 
                struct{ Success bool;
                        Detail  string;
                        Id      int}{
                        false, "Notfound token", -1,
                        })
            return
        }

        tokenString := req.Form["token"][0]

	    user_id, err := token.Valid(tokenString)

	    if err != nil {
            formatter.JSON(w, 403, 
                struct{ Success bool;
                        Detail  string;
                        Id      int}{
                        false, "Invalid token", -1,
                        })
            return
        }

        id, _ := strconv.Atoi(user_id)
        formatter.JSON(w, http.StatusOK, 
            struct{ Success bool;
                    Detail string;
                    Id      int}{
                    true, 
                    "The token is valid.",
                    id})
    }
}