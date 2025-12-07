package main

type Code int

func (code Code)GetMsg() (msg string) {
	switch code {
		case SuccessCode:
			return "Success"
		case ServiceErrCode:
			return "Service Error"
		case NetworkErrCode:
			return "Network Error"
	}
	return ""
}
func (c Code) Ok() (code Code, msg string) {
	return c, c.GetMsg()
}

const (
	SuccessCode    Code = 0
	ServiceErrCode Code = 1001
	NetworkErrCode Code = 1002
)

func webServer(name string) (code Code, msg string) {
	if name == "1" {
		return ServiceErrCode.Ok()
	}
	if name == "2" {
		return NetworkErrCode.Ok()
	}
	return SuccessCode.Ok()
}

func main() {
}

