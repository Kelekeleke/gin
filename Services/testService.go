package services

import (
	Models "gin/models"
)

type Test struct {
	Id      int    `json:"id"`
	Testcol string `json:"testcol"`
}

func (this *Test) Insert() (id int, err error) {
	var testModel Models.Config
	testModel.Id = this.Id
	id, err = testModel.Insert()
	return
}

func (this *Test) Select() []Models.Config {
	var testModel Models.Config
	// res := make(map[strings]interface{})
	data := testModel.Select()
	// value := reflect.ValueOf(data)

	// for i := 0; i < value.Len(); i++ {
	// 	item := value.Index(i).FieldByName("Id")
	// }
	return data
}
