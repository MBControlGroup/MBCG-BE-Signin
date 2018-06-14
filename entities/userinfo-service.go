package entities

//LoginAtomicService .
type LoginAtomicService struct{}

//UserInfoService .
var LoginService = LoginAtomicService{}


func (*LoginAtomicService) AdminFindByAccount(account string) *Admins {
    //dao := userInfoDao{mydb}
    sql := "SELECT * FROM Admins WHERE Admins.admin_account=?;"
    var admin Admins
    find, err := engine.SQL(sql, account).Get(&admin)
    checkErr(err)
    if find {
        return &admin
    } else {
        return nil
    }
}

func (*LoginAtomicService) AdminFindById(id int) *Admins {
    //dao := userInfoDao{mydb}
    var admin Admins
    admin.Admin_id = id
    _, err := engine.Get(&admin)
    checkErr(err)
    return &admin//dao.FindByID(id)
}

func (*LoginAtomicService) AddIMUser() error {
    session := engine.NewSession()
    defer session.Close()

    err := session.Begin()
    checkErr(err)

    sql := "INSERT INTO IMUsers (user_id) VALUEs (?)"

    _, err = engine.Exec(sql,0) 
    checkErr(err)
    
    if err == nil {
        session.Commit()
    } else {
        session.Rollback()
    }
    return nil
}

func (*LoginAtomicService) AdminSave(a *Admins) *Admins {
    session := engine.NewSession()
    defer session.Close()

    err := session.Begin()
    checkErr(err)

    sql := "INSERT INTO Admins (admin_account, admin_passwd, admin_type, im_user_id) VALUES (?,?,?,?);"

    _, err = engine.Exec(sql, a.Admin_account, a.Admin_passwd, a.Admin_type, a.Im_user_id) 
    checkErr(err)
    
    if err == nil {
        session.Commit()
        return a
    } else {
        session.Rollback()
    }
    return nil
}



