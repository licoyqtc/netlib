package protocol

import (
	"encoding/json"
	"reflect"
	"encoding/base64"
	"fmt"
)



type Protocol struct {
	Reqname  string  `json:"reqname"`
	Data []byte `json:"data"`
}

type ReqRegisterSdp struct {
}


type RegisterSdp struct {
	BoxId string `json:"box_id"`
	Sdp   string `json:"sdp"`
}

type PushAppSdp struct {
	AppSdp string `json:"app_sdp"`
}
type PushRes struct {
	ErrNo  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
}

var instance *ProtManager

type ProtManager struct {
	handler map[string]func(Protocol)
}

func GetProtManagerIns () *ProtManager{
	if instance == nil {
		instance = new(ProtManager)
		instance.handler = make(map[string]func(Protocol))
	}
	return instance
}

func (pm *ProtManager) SetFuncHandler(req interface{} , f func(Protocol)) {
	t := reflect.TypeOf(req)
	pm.handler[t.Name()] = f
}


func (pm *ProtManager)PackData(req interface{}) string {

	t := reflect.TypeOf(req)
	stName := t.Name()

	p := Protocol{}
	p.Reqname = stName
	body , _ := json.Marshal(req)
	p.Data = body

	data , _ := json.Marshal(p)

	return base64.StdEncoding.EncodeToString(data)
}

func (pm *ProtManager)HandleRequest(data string) interface{}{

	bdata , _ := base64.StdEncoding.DecodeString(data)

	p := Protocol{}

	json.Unmarshal(bdata , &p)
	fmt.Printf("HandleRequest data : %+v\n",p)

	f , ok := pm.handler[p.Reqname]
	if ok {
		fmt.Printf("handle %s func...\n",p.Reqname)
		f(p)
	}

	return ""
}

