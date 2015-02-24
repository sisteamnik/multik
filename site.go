package multik

import (
	"github.com/itkinside/itkconfig"
	"path"
)

type Site struct {
	Routes  string
	Domains string
	router  *Router
}

func NewSite(confpath string) (*Site, error) {
	cnf := &Site{}
	router := &Router{}
	cnf.router = router
	err := itkconfig.LoadConfig(confpath, cnf)
	if err != nil {
		panic(err)
	}
	err = cnf.router.LoadFromFile(path.Dir(confpath)+"/"+
		cnf.Routes, cnf.Domains)
	if err != nil {
		panic(err)
	}
	return cnf, nil
}
