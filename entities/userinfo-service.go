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

func (*LoginAtomicService) SoldierSave(s *Soldiers) error {
    engine.Insert(s)

    return nil
}

func (*LoginAtomicService) GetSoldierByOfficeId(officd_id string) []Soldiers {
    sql := "select * form soldiers where slodiers.serve_office_id = "+officd_id
    var soldiers []Soldiers
    err := engine.SQL(sql).Find(soldiers) 
    checkErr(err)
    return soldiers
}

func (*LoginAtomicService) GetAllBM() []BroadcastMessages {
    var BMs []BroadcastMessages
    err := engine.Find(&BMs)
    checkErr(err)
    return BMs
}


func (*LoginAtomicService) GetBMById(bm_id int) *BroadcastMessages {
    var BM BroadcastMessages
    BM.Bm_id = bm_id
    has, err := engine.Id(bm_id).Get(&BM)
    checkErr(err)
    if has {
        return &BM
    } else {
        return nil
    }
}

func (*LoginAtomicService) AddBM(b *BroadcastMessages) *BroadcastMessages {
    session := engine.NewSession()
    defer session.Close()

    err := session.Begin()
    checkErr(err)

    _, err = session.Insert(b)
    checkErr(err)
    if err == nil {
        session.Commit()
        return b
    } else {
        session.Rollback()
    }
    return nil
}

func (*LoginAtomicService) AddBMO(b *BcMsgOffices) *BcMsgOffices {
    session := engine.NewSession()
    defer session.Close()

    err := session.Begin()
    checkErr(err)

    _, err = session.Insert(b)
    checkErr(err)
    if err == nil {
        session.Commit()
        return b
    } else {
        session.Rollback()
    }
    return nil
}

func (*LoginAtomicService) AddOffice(o *Offices) *Offices {
    session := engine.NewSession()
    defer session.Close()

    err := session.Begin()
    checkErr(err)

    _, err = session.Insert(o)
    checkErr(err)
    if err == nil {
        session.Commit()
        return o
    } else {
        session.Rollback()
    }
    return nil
}

func (*LoginAtomicService) AddBMOrg(b *BcMsgOrgs) *BcMsgOrgs {
    session := engine.NewSession()
    defer session.Close()

    err := session.Begin()
    checkErr(err)

    _, err = session.Insert(b)
    checkErr(err)
    if err == nil {
        session.Commit()
        return b
    } else {
        session.Rollback()
    }
    return nil
}

func (*LoginAtomicService) AddOrg(o *Organizations) *Organizations {
    session := engine.NewSession()
    defer session.Close()

    err := session.Begin()
    checkErr(err)

    _, err = session.Insert(o)
    checkErr(err)
    if err == nil {
        session.Commit()
        return o
    } else {
        session.Rollback()
    }
    return nil
}

func (*LoginAtomicService) AddCN(c *CommonNotifications) *CommonNotifications {
    session := engine.NewSession()
    defer session.Close()

    err := session.Begin()
    checkErr(err)

    _, err = session.Insert(c)
    checkErr(err)
    if err == nil {
        session.Commit()
        return c
    } else {
        session.Rollback()
    }
    return nil
}

func (*LoginAtomicService) AddCNR(c *CmNtReceipts) *CmNtReceipts {
    session := engine.NewSession()
    defer session.Close()

    err := session.Begin()
    checkErr(err)

    _, err = session.Insert(c)
    checkErr(err)
    if err == nil {
        session.Commit()
        return c
    } else {
        session.Rollback()
    }
    return nil
}