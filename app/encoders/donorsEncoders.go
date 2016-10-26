package encoders

import (
	"github.com/gureosman/Application/app/models"
	"io/ioutil"
	"encoding/json"
	"log"
	"io"
)

func EncodeDonors(body io.ReadCloser)(donors models.Donors)  {
	var donor, _ = ioutil.ReadAll(body)
	if err :=json.Unmarshal(donor,&donors); err !=nil{
		log.Println(err)
		return
	}
	return
}
