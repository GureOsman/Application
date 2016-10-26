package encoders

import (
	"io"
	"github.com/gureosman/Application/app/models"
	"io/ioutil"
	"encoding/json"
	"log"
)

func EncodeCity(body io.ReadCloser)(cities models.City)  {
	var city, _ = ioutil.ReadAll(body)
	if err :=json.Unmarshal(city,&cities); err !=nil{
		log.Println("donors: ", cities)
		log.Println(err)
		return
	}
	return
}
