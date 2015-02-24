package multik

import (
	"github.com/itkinside/itkconfig"
)

type Config struct {
	Port  int
	Sites string
}

func ConfigFromFile(path string) (Config, error) {
	cnf := Config{}
	err := itkconfig.LoadConfig(path, &cnf)
	return cnf, err
}
