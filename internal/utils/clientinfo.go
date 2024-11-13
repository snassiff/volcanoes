package clientinfo

type ClientInfo struct {
	Referer   string `header:"Referer"`
	Origin    string `header:"Origin"`
	UserAgent string `header:"User-Agent"`
	ClientIP  string //Ip del cliente, no sale del header
}
