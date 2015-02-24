package multik

import (
	"io/ioutil"
	"testing"
)

var (
	ROUTES_STRING string
	test_router   *Router
	/*[get /api/v1]
	/users/:id Users.Get
	/users/search Users.Search
	#comment

	[/api/v2]
	get /articles/:id Article.Get
	post /articles/new Article.Create*/
	asertRoutes = []Route{
		Route{
			method: "get",
			domain: "example.com",
			path:   "/api/v1/users/:id",
			action: "Users.Get",
		},
	}

	assertmap = map[string]string{
		"[get /api/v1]":                     RouteLineTypeSection,
		"/users/:id Users.Get":              RouteLineTypeRoute,
		"/users/search Users.Search":        RouteLineTypeRoute,
		"#comment":                          RouteLineTypeComment,
		"[/api/v2]":                         RouteLineTypeSection,
		"get /articles/:id Article.Get":     RouteLineTypeRoute,
		"post /articles/new Article.Create": RouteLineTypeRoute,
	}
)

func init() {
	rs, err := ioutil.ReadFile("./assets/conf.routes")
	if err != nil {
		panic(err)
	}
	ROUTES_STRING = string(rs)
	test_router = new(Router)
}

func TestRoutes(t *testing.T) {
	rts, err := string2routes("example.com", ROUTES_STRING)
	if err != nil {
		t.Error(err.Error())
	}
	for _, ar := range asertRoutes {
		for _, r := range rts {
			if ar.action == r.action &&
				ar.domain == r.domain &&
				ar.method == r.method &&
				ar.path == r.path {
			}
		}
	}
	if len(rts) != 4 {
		t.Errorf("Route file contains four routes, but %d found\n", len(rts))
	}
}

func TestRouteLineTypeDetection(t *testing.T) {
	for k, v := range assertmap {
		r, err := routeType("", "", k)
		if err != nil {
			t.Error(err.Error())
		}
		if r != v {
			t.Errorf("%s must be %s", r, v)
		}
	}
}

func TestRoutesFromFile(t *testing.T) {
	err := test_router.LoadFromFile("./assets/conf.routes", "*")
	if err != nil {
		t.Error(err.Error())
	}

}

/*func TestLine2Route(t *testing.T) {
	var assrt = map[string]Route{
		""
	}

}*/
