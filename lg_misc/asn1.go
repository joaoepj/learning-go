package lg_misc

import (
	"encoding/asn1"
	"fmt"
	"time"
)

func Asn1() {
	//t1, _ := time.Parse("20220904131020", "20220904131020")
	t1 := time.Now()
	var t2 time.Time
	b := []byte("20220904131020")
	mdata, _ := asn1.Marshal(t1)
	fmt.Println(mdata, b)
	_, err := asn1.Unmarshal(mdata, &t2)
	fmt.Println(t2)
	fmt.Println(err)
}
