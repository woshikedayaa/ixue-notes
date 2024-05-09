package models

type UserDetailedInfo struct {
	Status int `json:"status"`
	User   struct {
		UserId            string        `json:"user_id"`
		UserIdentifier    string        `json:"user_identifier"`
		UserAccount       string        `json:"user_account"`
		UserType          string        `json:"user_type"`
		UserMail          string        `json:"user_mail"`
		UserPn            string        `json:"user_pn"`
		UserNickname      string        `json:"user_nickname"`
		UserSex           string        `json:"user_sex"`
		UserBirthday      string        `json:"user_birthday"`
		UserProvinceId    string        `json:"user_province_id"`
		UserRegisterTime  string        `json:"user_register_time"`
		UserLogoUrl       string        `json:"user_logo_url"`
		UserLoginToken    string        `json:"user_login_token"`
		UserStudentNumber string        `json:"user_student_number"`
		Balance           string        `json:"balance"`
		WxUserType        string        `json:"wx_user_type"`
		WxUserPn          string        `json:"wx_user_pn"`
		UserRealName      string        `json:"user_real_name"`
		UserInfo          []interface{} `json:"user_info"`
		UserGrade         interface{}   `json:"user_grade"`
		VipType           interface{}   `json:"vip_type"`
		VipStartTime      interface{}   `json:"vip_start_time"`
		VipEndTime        interface{}   `json:"vip_end_time"`
		IsCanvasser       bool          `json:"isCanvasser"`
		ForceLogin        int           `json:"forceLogin"`
		Role              []interface{} `json:"role"`
	} `json:"user"`
	ErrorMessage []interface{} `json:"errorMessage"`
	Time         int           `json:"time"`
}
