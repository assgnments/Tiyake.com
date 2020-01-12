package util

import "time"

var AvailableFuncMaps=map[string] interface{}{
	"readableDate": readableTime,

}

func readableTime(t time.Time)string{
	return  t.Format(time.RFC822)
}