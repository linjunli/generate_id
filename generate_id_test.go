package generate_id

import (
	"fmt"
	"log"
)

func ExampleGetUuid() {
	uuid := GetUuid()
	if uuid == "" {
		log.Fatal(fmt.Errorf("generate uuid err"))
	}
	log.Printf("generate uuid : "+uuid)
}