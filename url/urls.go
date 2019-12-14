package url

type (
	urls struct {
		Home_PATH     string
		ADD_PATH      string
		POST_ADD_PATH string
		SHOW_PATH     string
	}
)

func GetURLS() urls {
	var temp urls
	temp.Home_PATH = "/"
	temp.ADD_PATH = "/add"
	temp.POST_ADD_PATH = "/add/form"
	temp.SHOW_PATH = "show"
	return temp
}
