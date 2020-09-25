package util

import "sync"

type WebData struct {
	data map[string]string
	lock sync.RWMutex
}

var clientData  *WebData = &WebData{
	data:make(map[string]string),
}
func Instance()*WebData{
	return clientData
}

func(w *WebData)Put(key,value string){
	w.lock.Lock()
	defer w.lock.Unlock()
	w.data[key] = value
	//将内容保存在文件中

}

func(w *WebData)Get(key string)(string,bool){
	w.lock.Lock()
	defer w.lock.Unlock()
	value,ok := w.data[key]
	return value,ok
}