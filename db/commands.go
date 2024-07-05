package db

import "fmt"

func Echo(values []Value) Value {
	echoMsg := fmt.Sprintf("+%s\r\n", values[1].bulk)
	return Value{typ: BULK, bulk: echoMsg}
}

func Set(ys *YoshiStore, values []Value) Value {
	if len(values) != 3 {
		return Value{typ: ERROR, err: "-Err Invalid amount of arguments for Set command (2 required)\r\n"}
	}

	ys.Set(values[1].bulk, values[2].bulk)

	return Value{typ: SIMPLE, simple: "+OK\r\n"}
}

func Get(ys *YoshiStore, values []Value) Value {
	if len(values) != 2 {
		return Value{typ: ERROR, err: "-Err Invalid amount of arguments for Get command (1 required)\r\n"}
	}
	val, ok := ys.Get(values[1].bulk)
	if !ok {
		fmt.Println("invalid get")
		return Value{typ: ERROR, err: "-ERR not a valid key\r\n"}
	}
	return Value{typ: BULK, bulk: "+" + val + "\r\n"}
}
