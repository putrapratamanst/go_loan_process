package controllers

import(
	"strconv"
	"encoding/json"
	"io/ioutil"
	"./models"
	"fmt"
)


func CreateDayMax(tempSave map[string]int, numbers int) string  {
	tempSave["create_day_max"] = numbers
	num := strconv.Itoa(numbers)

	file, _ := json.MarshalIndent(tempSave, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)

	callback := "Created max request with " +num+ " requests" 
	
	return callback
}

func AddDataBorrower(tempSave map[string]int, data string) string {
		var p = models.RequestModel{1, "Singh", 26}
		file, _ := json.MarshalIndent(p, "", " ")
		_ = ioutil.WriteFile("test.json", file, 0644)



	fmt.Println("sdf")

	return ""
}
