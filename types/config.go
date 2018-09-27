package types

//Config structure to store configuration datas
//nolint:lll
type Config struct {
	TelegramToken string `envconfig:"CIRILLA_TOKEN"`
	Debug         bool   `envconfig:"CIRILLA_DEBUG" default:"false"`
	Timeout       int    `envconfig:"CIRILLA_TIMEOUT"`
	CommandPrefix string `envconfig:"CIRILLA_CMD_PREFIX" default:"/"`
}
