package ammodel

type Br struct {
	Code int
	Data interface{}
	Msg  string
	Ext  string
}

func Ok(d interface{}) Br {
	res := Br{Code: 0, Data: d}
	return res
}

func Error(m string) Br {
	res := Br{Code: -1, Msg: m}
	return res
}
