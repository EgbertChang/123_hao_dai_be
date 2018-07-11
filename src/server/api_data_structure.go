package server

type CreateResult struct {
	Msg string `json:"msg"`
}

type RetrieveResult struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type DeleteResult struct {
	Msg string `json:"msg"`
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
