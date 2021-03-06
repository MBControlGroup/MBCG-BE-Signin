package service

import (
	//"time"
	"encoding/json"
	//"fmt"
    //"time"
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
        //headMap := make(map[string][]string)
        //headMap["Access-Control-Allow-Origin"] = append(headMap["Access-Control-Allow-Origin"],"*")
        //headMap["Access-Control-Allow-Headers"] = append(headMap["Access-Control-Allow-Headers"],"Origin, X-Requested-With, Content-Type, Accept")
        //headMap["Access-Control-Allow-Methods"] = append(headMap["Access-Control-Allow-Methods"],"GET, POST, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	    w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Credentials")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") 
        w.Header().Set("Access-Control-Allow-Credentials", "true") 

        var user entities.UserInfo
        
        err := json.NewDecoder(req.Body).Decode(&user)
        fmt.Println(user)
        checkErr(err)

        u := entities.LoginService.AdminFindByAccount(user.Username)
        //u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["username"][0]})
        
        if u==nil || u.Admin_passwd != user.Password {
            formatter.JSON(w, http.StatusBadRequest, struct{ Code int `json:"code"`;Enmsg string `json:"enmsg"`;Cnmsg string `json:"cnmsg"`; Data interface{} `json:"data"`}{400, "fail", "失败", nil})
        } else {
            fmt.Println(u.Admin_id)
            tokenString, err := token.Generate(u.Admin_id)
            //cookie := http.Cookie{Name:"token", Value:tokenString, MaxAge:86400}
            //sid=kcyt333ns5i5xiejizlz0fjpxfd6d59f; expires=Fri, 28-Jun-2019 15:06:09 GMT; Max-Age=31536000; path=/; secure; HttpOnly
            cookie := http.Cookie{Name:"token", Value:tokenString, MaxAge:86400, }
            http.SetCookie(w, &cookie)
            //headMap["set-cookie"] = append(headMap["set-cookie"],cookie.String())
            //w.Header = headMap
            //w.Header().Add("Set-Cookie", cookie.String())
            //w.Header().Set("set-cookie", cookie.String())
	        //w.Header().Set("set-cookie", "token="+tokenString+";expires="+time.Now().Format("2006-01-02 15:04:05")+";Max-Age=31536000;path=/;secure;HttpOnly") 
            checkErr(err)
            formatter.JSON(w, http.StatusOK, struct{ Code int `json:"code"`;Enmsg string `json:"enmsg"`;Cnmsg string `json:"cnmsg"`; Data interface{} `json:"data"`}{200, "ok", "成功", nil})
        }
    }
} 

func signoutHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	    w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Credentials")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") 
        w.Header().Set("Access-Control-Allow-Credentials", "true") 
        
        cookie, err := req.Cookie("token")

        if err != nil || cookie.Value == "" {
            formatter.JSON(w, http.StatusBadRequest, struct{ Code int `json:"code"`;Enmsg string `json:"enmsg"`;Cnmsg string `json:"cnmsg"`; Data interface{} `json:"data"`}{302, "fail", "失败", nil})
	        return;
        }

	    _, err = token.Valid(cookie.Value)

	    if err != nil {
	        formatter.JSON(w, http.StatusBadRequest, struct{ Code int `json:"code"`;Enmsg string `json:"enmsg"`;Cnmsg string `json:"cnmsg"`; Data interface{} `json:"data"`}{302, "fail", "失败", nil})
	        return;
        }

        cookie = &http.Cookie{Name: "token", Path: "/", MaxAge: -1}
        http.SetCookie(w, cookie)
        formatter.JSON(w, http.StatusOK, struct{ Code int `json:"code"`;Enmsg string `json:"enmsg"`;Cnmsg string `json:"cnmsg"`; Data interface{} `json:"data"`}{200, "ok", "成功", nil})
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

func preOptionHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	    w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Credentials")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") 
        w.Header().Set("Access-Control-Allow-Credentials", "true") 
        formatter.JSON(w, http.StatusOK, "")
    }
}

func testToken(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {

        cookie, err := req.Cookie("token")
	    if err != nil || cookie.Value == ""{
	        formatter.JSON(w, 403, struct{Error string `json:"error"`}{"token not found."})
	        return;
	    }

	    user_id, err := token.Valid(cookie.Value)

	    if err != nil {
	        formatter.JSON(w, 403, struct{Error string `json:"error"`}{"bad token"})
	        return;
        }

        id, err := strconv.Atoi(user_id)
        admin := entities.LoginService.AdminFindById(id)
        //fmt.Println(user_id)

        formatter.JSON(w, http.StatusOK, 
            struct{ Success bool `json:"success"`;
                    Content string `json:"content"`;
                    AdminInfo entities.Admins `json:"admin_info"`}{
                    true, 
                    "The token is valid.",
                    *admin})
    }
}

func tokenValid(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        var tt struct{Token string `json:"token"`}
        tt.Token = ""
        err := json.NewDecoder(req.Body).Decode(&tt)
        checkErr(err)

        if len(tt.Token) == 0 {
            formatter.JSON(w, http.StatusBadRequest, 
                struct{ Success bool `json:"success"`;
                        Detail  string `json:"detail"`;
                        Id      int `json:"id"`}{
                        false, "Notfound token", -1,
                        })
            return
        }

        tokenString := tt.Token

	    user_id, err := token.Valid(tokenString)

	    if err != nil {
            formatter.JSON(w, 403, 
                struct{ Success bool `json:"success"`;
                        Detail  string `json:"detail"`;
                        Id      int `json:"id"`}{
                        false, "Invalid token", -1,
                        })
            return
        }

        id, _ := strconv.Atoi(user_id)
        formatter.JSON(w, http.StatusOK, 
            struct{ Success bool `json:"success"`;
                    Detail  string `json:"detail"`;
                    Id      int `json:"id"`}{
                    true, 
                    "The token is valid.",
                    id})
    }
}
