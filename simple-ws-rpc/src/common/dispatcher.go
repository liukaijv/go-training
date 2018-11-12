package common

import (
	"fmt"
	"log"
	"reflect"
	"runtime/debug"
	"sync"
	"time"
)

type Func struct {
	Handler  reflect.Value
	IsMethod bool
	Caller   reflect.Value
}

type Dispatcher struct {
	*sync.RWMutex
	funcList map[string]*Func
}

func NewDispatcher() *Dispatcher {

	return &Dispatcher{
		new(sync.RWMutex),
		map[string]*Func{},
	}

}

func (d *Dispatcher) AddFunc(funcName string, caller interface{}, handler interface{}) {

	if d.HasFunc(funcName) {
		log.Printf("Dispatcher.AddFunc already has funcName:[%s]", funcName)
		return
	}

	f := new(Func)
	if caller != nil {
		if reflect.TypeOf(handler) == reflect.TypeOf("") {
			f.Handler = reflect.ValueOf(caller).MethodByName(handler.(string))
			f.IsMethod = true
		} else {
			f.Caller = reflect.ValueOf(caller)
			f.Handler = reflect.ValueOf(handler)
		}
	} else {
		f.Handler = reflect.ValueOf(handler)
	}

	if f.Handler.IsValid() {
		d.Lock()
		d.funcList[funcName] = f
		d.Unlock()
	} else {
		log.Printf("Dispatcher.AddFunc::funcName:[%s] fail.", funcName)
	}

}

func (d *Dispatcher) RemoveFunc(funcName string) {
	d.Lock()
	delete(d.funcList, funcName)
	d.Unlock()
}

func (d *Dispatcher) Run(funcName string, args ...interface{}) (resp []interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmtStr := fmt.Sprintf("Dispatcher.Run::funcName:[%s], args:%v error.", funcName, args)
			log.Println(fmtStr, err, string(debug.Stack()))
		}
	}()
	d.RLock()
	f, exist := d.funcList[funcName]
	d.RUnlock()
	if exist {

		callArgs := make([]reflect.Value, 0)
		if f.Caller.IsValid() && !f.IsMethod {
			callArgs = append(callArgs, f.Caller)
		}
		for _, arg := range args {
			callArgs = append(callArgs, reflect.ValueOf(arg))
		}
		t := time.Now()
		handlerRes := f.Handler.Call(callArgs)
		elapsed := time.Now().Sub(t)

		if elapsed.Nanoseconds() > int64(100*1000*1000) {
			log.Printf("Dispatcher.Run::funcName[%s] elapsed > 100ms, elapsed time: %s, args: %v", funcName, elapsed, args)
		}

		for _, v := range handlerRes {
			resp = append(resp, v.Interface())
		}

	} else {
		log.Printf("Dispatcher.Run::funcName:[%s] not found", funcName)
	}
	return
}

func (d *Dispatcher) HasFunc(funcName string) bool {
	d.RLock()
	_, exist := d.funcList[funcName]
	d.RUnlock()
	return exist
}

func (d *Dispatcher) ClearFunc() {
	d.Lock()
	for key, _ := range d.funcList {
		delete(d.funcList, key)
	}
	d.Unlock()
}
