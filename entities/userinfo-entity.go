package entities

import (
    "time"
)

const (
    OF = iota
    OR
)

// UserInfo .
type UserInfo struct {
    Username    string  `json:"username"`
    Password    string  `json:"password"`
}

type Admins struct {
    Admin_id        int    
    Admin_account   string    
    Admin_passwd    string
    Admin_type      string
    Im_user_id      int
}

type Soldiers struct {
    Soldier_id      int   
    Rank            string
    Id_num          string
    Name            string
    Phone_num       string
    Wechat_openid   string
    Commander_id    int
    Serve_office_id int
}

type Task struct {
    Task_id         int   
    Title           string
    Mem_count       int
    Launch_admin_id int 
    Launch_datetime *time.Time   
    Gather_datetime *time.Time
    Detail          string
    Gather_place_id int
    Finish_datetime *time.Time    
}

type BroadcastMessages struct {
    Bm_id           int   
    Title           string
    Detail          string
    Bm_type         string
    Wechat_notice   bool
    Sms_notice      bool
    Voice_notice    bool
}

type BcMsgOffices struct {
    Bmo_id          int   
    Msg_id          int
    Msg_office_id   int
}

type BcMsgOrgs struct {
    Bmo_id          int   
    Msg_id          int
    Msg_org_id      int
}

type Organizations struct {
    Org_id          int   
    Serve_office_id int
    Leader_sid      int
    Name            string
}

type Offices struct {
    Office_id           int   
    Office_level        int
    Higher_office_id    int
    Name                string
}

type CommonNotifications struct {
    Cn_id           int   
    Cn_bm_id        int
    Recv_soldier_id int
}

type CmNtReceipts struct {
    Cnr_id          int   
    Cn_id           int
    Rec_content     string
}


