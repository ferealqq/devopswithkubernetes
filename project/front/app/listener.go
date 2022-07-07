package main


var Listener func()

func Listen(f func()){
	Listener = f
}

func Dispatch(){
	if Listener != nil {
		Listener()
	}else{
		panic("cannot dispatch without a listener")
	}
}