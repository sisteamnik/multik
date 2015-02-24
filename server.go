package multik

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"reflect"
)

type Server struct {
	config      Config
	controllers map[string]interface{}
	sites       map[string]*Site
}

func NewServer(config Config) (*Server, error) {
	s := &Server{
		config:      config,
		controllers: map[string]interface{}{},
		sites:       map[string]*Site{},
	}
	sites, err := filepath.Glob(config.Sites)
	if err != nil {
		return s, err
	}
	for _, v := range sites {
		site, err := NewSite(v)
		if err != nil {
			return s, err
		}
		for _, v := range domainsArr(site.Domains) {
			s.sites[v] = site
		}
	}
	return s, nil
}

func (s *Server) Run() {
	http.HandleFunc("/", s.Handler)
	log.Printf("Listen %d", s.config.Port)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(s.config.Port), nil))
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	controller := s.NewController(w, r)
	controller.Apply()
}

func (s *Server) BindControllers(is ...interface{}) {
	for _, v := range is {
		rt := reflect.TypeOf(v)
		s.controllers[rt.Name()] = v
		fmt.Printf("Registered controller %s\n", rt.Name())
	}
}

func (s *Server) NewController(w http.ResponseWriter, r *http.Request) *Controller {
	req := NewRequest(r)
	resp := &Response{Out: w}
	controller := NewController(req, resp)
	controller.Server = s
	return controller
}

func (s *Server) Call(method, action string) {
	tmp := s.controllers[method]
	t := reflect.TypeOf(tmp)
	ptr := reflect.New(t)
	fn := ptr.MethodByName(action)
	if !fn.IsValid() {
		fmt.Println("no")
	} else {
		fn.Call([]reflect.Value{})
	}
}
