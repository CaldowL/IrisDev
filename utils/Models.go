package utils

type ResponseBasicBody struct {
	Ret      int    `json:"ret"`
	Data     any    `json:"data,omitempty"`
	ErrorMsg string `json:"errorMsg,omitempty"`
}

type IDPUser struct {
	Id     int    `db:"id"`
	Openid string `db:"openid"`
	Avatar string `db:"avatar"`
	Nick   string `db:"nick"`
}

type ShortLink struct {
	Id    int     `db:"id"`
	Short string  `db:"short"`
	Url   string  `db:"url"`
	T     []uint8 `db:"t"`
}
