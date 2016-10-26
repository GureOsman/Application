package encoders

import (
	"io"
	"github.com/gureosman/Application/app/models"
	"io/ioutil"
	"encoding/json"
	"log"
)

func EncodeBlood(body io.ReadCloser)(bloodtype models.Blood)  {
	var blood, _ = ioutil.ReadAll(body)
	if err :=json.Unmarshal(blood,&bloodtype); err !=nil{
		log.Println(err)
		return
	}
	return
}
