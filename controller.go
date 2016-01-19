package multik

import (
	"fmt"
	"reflect"
)

type Controller struct {
	Request  *Request
	Response *Response
	Server   *Server

	Result Result

	Method string
	Action string

	AppController interface{}
}

func NewController(req *Request, resp *Response) *Controller {
	return &Controller{
		Request:  req,
		Response: resp,
	}
}

func (c *Controller) GetProcessor(method, action string) {
	_, ok := c.Server.controllers[method]
	if !ok {
		panic("can't get actio")
	}
	c.Method = method
	c.Action = action
}

func (c *Controller) Apply() {
	fmt.Fprintf(c.Response.Out, "hello %s", "sobaka")

	//todo router hange it
	c.GetProcessor("Users", "Get")

	proc, ok := c.Server.controllers[c.Method]
	if !ok {
		fmt.Printf("proccessor not found %s\n", c.Method)
	}

	acv := reflect.ValueOf(proc)

	vc := reflect.ValueOf(c)
	el := reflect.New(acv.Type()).Elem()
	mv := el.FieldByName("Controller")
	fmt.Println("da", mv)

	if !mv.CanSet() {
		fmt.Println("cant set")
		//mv.Set(vc)
	} else {
		mv.Set(vc)
	}
	if !el.CanAddr() {
		fmt.Println("cant addr")
	}
	ptr := el.Addr()
	fn := ptr.MethodByName(c.Action)
	if !fn.IsValid() {
		fmt.Println("no")
	} else {
		fn.Call([]reflect.Value{})
	}
}
