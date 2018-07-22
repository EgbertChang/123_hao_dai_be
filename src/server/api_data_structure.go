package server

type CreateResponse struct {
	Msg string `json:"msg"`
}

type DeleteResponse struct {
	Msg string `json:"msg"`
}

type UpdateResponse struct {
	Msg string `json:"msg"`
}

type RetrieveResponse struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
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
	PartyAId  string `json:"partyAId"`
	Name      string `json:"name"`
	PartyAUrl string `json:"partyAUrl"`
	PartyBUrl string `json:"partyBUrl"`
}

type listBParams struct {
	PartyAId  string `json:"partyAId"`
	PageIndex int    `json:"pageIndex"`
	PageSize  int    `json:"pageSize"`
}

type product struct {
	Name                  string      `json:"name"`
	Url                   string      `json:"url"`
	Type                  []string    `json:"type"`
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
	ApplyStrategy         string      `json:"applyStrategy"`
}

type productInfo struct {
	Id                    string      `json:"id"`
	Name                  string      `json:"name"`
	Url                   string      `json:"url"`
	Type                  []string    `json:"type"`
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
	ApplyStrategy         string      `json:"applyStrategy"`
}

type productList struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	LimitMin    string   `json:"limitMin"`
	LimitMax    string   `json:"limitMax"`
	Slogan      string   `json:"slogan"`
	ApplyNumber int      `json:"applyNumber"`
	Interest    interest `json:"interest"`
}

type productListParams struct {
	Name                  string   `json:"name"`
	LimitMin              string   `json:"limitMin"`
	LimitMax              string   `json:"limitMax"`
	Type                  []string `json:"type"`
	PersonalQualification []string `json:"personalQualification"`
	PageIndex             int      `json:"pageIndex"`
	PageSize              int      `json:"pageSize"`
}

type interest struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type lendingRate struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
