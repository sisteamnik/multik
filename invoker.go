package multik

import (
	"fmt"
	"reflect"
)

func InvokerFilter(c *Controller, _ []Filter) {

	proc, ok := c.Server.controllers[c.Method]
	if !ok {
		fmt.Printf("proccessor not found %s\n", c.Method)
	}

	c.AppController = proc

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

	/*methodValue := reflect.ValueOf(c.AppController).MethodByName(c.Method)

	var resultValue reflect.Value
	if methodValue.Type().IsVariadic() {
		resultValue = methodValue.CallSlice([]reflect.Value{})[0]
	} else {
		resultValue = methodValue.Call([]reflect.Value{})[0]
	}
	if resultValue.Kind() == reflect.Interface && !resultValue.IsNil() {
		c.Result = resultValue.Interface().(Result)
	}*/
}
