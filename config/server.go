package config

const (
	Tes = "test"
	Dev = "dev"
	Pro = "product"
)

type ServerConfig struct {
	port string `yaml:"port"`
	mode string `yaml:"mode"`
}
