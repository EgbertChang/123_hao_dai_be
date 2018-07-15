package server

type CreateResponse struct {
	Msg string `json:"msg"`
}

type RetrieveResponse struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type DeleteResponse struct {
	Msg string `json:"msg"`
}

type UploadResponse struct {
	Msg  string `json:"msg"`
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
}

type A struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	BNum int    `json:"bNum"`
}

type addAParams struct {
	Name string `json:"name"`
}

type listAParams struct {
	Name      string `json:"name"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}

type B struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	PartyAUrl  string `json:"partyAUrl"`
	PartyBUrl  string `json:"partyBUrl"`
	ClickCount int    `json:"clickCount"`
}

type addBParams struct {
	PartyAId  int    `json:"partyAId"`
	Name      string `json:"name"`
	PartyAUrl string `json:"partyAUrl"`
	PartyBUrl string `json:"partyBUrl"`
}

type listBParams struct {
	PartyAId  int `json:"partyAId"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type product struct {
	Name                  string      `json:"name"`
	Url                   string      `json:"url"`
	Type                  string      `json:"type"`
	PersonalQualification []string    `json:"personalQualification"`
	LimitMin              string      `json:"limitMin"`
	LimitMax              string      `json:"limitMax"`
	LogoUrl               string      `json:"logoUrl"`
	Slogan                string      `json:"slogan"`
	ApplyNumber           int         `json:"applyNumber"`
	Term                  []string    `json:"term"`
	Interest              interface{} `json:"interest"`
	LendingRate           interface{} `json:"lendingRate"`
	Credit                string      `json:"credit"`
	AuditType             string      `json:"auditType"`
	AccountInType         string      `json:"accountInType"`
	ApplyStrategy         interface{} `json:"applyStrategy"`
}
