package url

type (
	urls struct {
		HOME_PATH     string
		ADD_PATH      string
		POST_ADD_PATH string
		SHOW_PATH     string
		SHOWALL_PATH  string
		STATIC_PATH   string
	}
)

func GetURLS() urls {
	var temp urls
	temp.HOME_PATH = "/"
	temp.ADD_PATH = "/add"
	temp.POST_ADD_PATH = "/add/form"
	temp.SHOW_PATH = "/show"
	temp.SHOWALL_PATH = "/showAll"
	temp.STATIC_PATH = "/src/"
	return temp
}
