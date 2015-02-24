package multik

import (
	"errors"
	"io/ioutil"
	"strings"
	"unicode"
)

const (
	ErrRouteLineMustBeConsistOfFourElements = ""
	ErrRouteMustBeOneLine                   = "Must be single line"
	ErrUncknownRoute                        = "Uncknown route"
	ErrRouteMustBeTwoFields                 = "Route line with defined method" +
		" must contains of two fields"
	ErrRouteMustBeThreeFields = "Route line must be" +
		" contains of three fields"

	RouteLineTypeSection = "section"
	RouteLineTypeComment = "comment"
	RouteLineTypeRoute   = "route"
	RouteLineTypeNewLine = "newline"
	RouteLineTypeError   = "error"
)

type Router struct {
	routes []Route
}

type Route struct {
	method string
	domain string
	path   string
	action string

	source *string
}

func (r Route) String() string {
	return r.method + " " + r.domain + r.path + " " + r.action
}

func (r *Router) LoadFromFile(path string, domain string) error {
	var err error
	r.routes, err = loadRoutesFromFile(path, domain, true)
	return err
}

func loadRoutesFromFile(path string, domains string, overrwrite bool) ([]Route,
	error) {
	bts, err := ioutil.ReadFile(path)
	if err != nil {
		return []Route{}, err
	}

	var routes []Route
	arr := domainsArr(domains)
	for _, v := range arr {
		routes, err = string2routes(v, string(bts))
	}
	return routes, err
}

func string2routes(domain, in string) ([]Route, error) {
	var res = []Route{}
	arr := strings.Split(in, "\n")
	method, prefix := "", ""
	for _, v := range arr {
		v = strings.TrimSpace(v)
		v = removeComment(v)
		rtype, err := routeType(method, prefix, v)
		if err != nil {
			return []Route{}, nil
		}
		switch rtype {
		case RouteLineTypeRoute:
			r, err := line2route(method, prefix, v)
			if err != nil {
				return []Route{}, err
			}
			res = append(res, r)
		case RouteLineTypeNewLine:
			method, prefix = "", ""
		case RouteLineTypeSection:
			arr := strings.Fields(v)
			switch len(arr) {
			case 1:
				//todo check length
				switch arr[0][1:2] {
				case "/":
					prefix = prefix + arr[0][1:len(arr[0])-1]
				default:
					method = arr[0]
				}
			case 2:
				method = arr[0][1:]
				prefix = prefix + arr[1][:len(arr[1])-2]
			}

		}

	}
	return res, nil
}

func line2route(method, prefix, line string) (Route, error) {
	r := Route{}
	if method != "" {
		//TODO check allowed methods
		r.method = method

		arr := strings.Fields(line)
		if len(arr) != 2 {
			return Route{}, errors.New(ErrRouteMustBeTwoFields)
		}
		r.path = prefix + arr[0]
		r.action = arr[1]
		return r, nil
	} else {
		arr := strings.Fields(line)
		if len(arr) != 3 {
			return Route{}, errors.New(ErrRouteMustBeThreeFields)
		}
		r.method = arr[0]
		r.path = prefix + arr[1]
		r.action = arr[2]
		return r, nil
	}
	/*flds := strings.Fields(line)
	if len(flds) != 4 {

	}*/

	//fmt.Printf("Route found: %s,%s,%s,%s\n", method, domain, path, action)
	return r, nil
}

func routeType(method, prefix, in string) (string, error) {
	if strings.Contains(in, "\n") {
		return RouteLineTypeError, errors.New(ErrRouteMustBeOneLine)
	}
	if in == "" {
		return RouteLineTypeNewLine, nil
	}
	if in[:1] == "#" {
		return RouteLineTypeComment, nil
	}
	if in[:1] == "[" {
		arr := strings.Fields(in)
		if len(arr) == 1 {
			return RouteLineTypeSection, nil
		}
		if len(arr) > 2 {
			//TODO change err type
			return RouteLineTypeError, errors.New(ErrUncknownRoute)
		}
		return RouteLineTypeSection, nil
	}
	flds := strings.Fields(in)
	if len(flds) == 2 || len(flds) == 3 {
		return RouteLineTypeRoute, nil
	}
	return RouteLineTypeError, errors.New(ErrUncknownRoute)
}

func removeComment(in string) (out string) {
	arr := strings.Split(in, "#")
	return arr[0]
}

func domainsArr(in string) []string {
	return strings.FieldsFunc(in, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '.' &&
			c != '*' && c != '-'
	})
}
