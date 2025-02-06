package users

type Configuration struct {
	BaseURL string
	Timeout int
}

var config *Configuration

func SetupConfig(baseUrl string, timeout int) {
	config = &Configuration{
		BaseURL: baseUrl,
		Timeout: timeout,
	}
}
