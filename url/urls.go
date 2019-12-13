package url

type (
	urls struct {
		Home_PATH string
	}
)

func GetURLS() urls {
	var temp urls
	temp.Home_PATH = "/"
	return temp
}
