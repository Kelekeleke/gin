package Services

import (
	Models "gin/models"
)

type Test struct {
	Id      int    `json:"id"`
	Testcol string `json:"testcol"`
}

func (this *Test) Insert() (id int, err error) {
	var testModel Models.Test
	testModel.Id = this.Id
	testModel.Testcol = this.Testcol
	id, err = testModel.Insert()
	return
}

func Select() interface{} {
	var testModel Models.Test
	// res := make(map[strings]interface{})
	data := testModel.Select()
	// value := reflect.ValueOf(data)

	// for i := 0; i < value.Len(); i++ {
	// 	item := value.Index(i).FieldByName("Id")
	// }
	return data
}
