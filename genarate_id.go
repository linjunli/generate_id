package generate_id

import (
	"fmt"
	"math/rand"
)


var (

)


func GetOneId()  {
	//timeStamp := time.Time{}.Unix()
	//machine := os.Getgid()


}

func GetUuid() string {
	return fmt.Sprintf(
		"%04x%04x-%04x-%04x-%04x-%04x%04x%04x",
		rand.Int31n(0xffff),
		rand.Int31n(0xffff),
		rand.Int31n(0xffff),
		rand.Int31n(0x0fff) | 0x4000,
		rand.Int31n(0x3fff) | 0x8000,
		rand.Int31n(0xffff),
		rand.Int31n(0xffff),
		rand.Int31n(0xffff),
	)
}