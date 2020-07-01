package controllers

import(
	"strconv"
	// "encoding/json"
	// "io/ioutil"
	// "./models"
	// "fmt"
)


func CreateDayMax(tempSave map[string]int, numbers int) string  {

	tempSave["create_day_max"] = numbers
	num := strconv.Itoa(numbers)

	// file, _ := json.MarshalIndent(tempSave, "", " ")
	// _ = ioutil.WriteFile("test.json", file, 0644)

	callback := "Created max request with " +num+ " requests" 
	
	return callback
}

func AddDataBorrower(tempSaveArr map[string][]string, data string) string {
	callback := "Success: "+ data
	return callback
}

func CheckStatus(tempSaveArr map[string][]string, loan string) string {

	dataLoan := tempSaveArr[loan]
	if dataLoan == nil{
		return "Data Loan Not Found"
	}

	amount, _ := strconv.Atoi(dataLoan[2])
	if amount % 1000000 != 0{
		return "Loan ID " +loan+ " is Rejected"

	}
	return "Loan ID " +loan+ " is Accepted"
}



func FindByAmountAccepted(tempSaveArr map[string][]string, loan string) string {
	s := " "

	for key, arg := range tempSaveArr {
		if arg[2] == loan{
			s += key + " "

		} 
	}
	return s
}

func FindByAmountRejected(tempSaveArr map[string][]string, loan string) string {
	s := " "
	amount, _ := strconv.Atoi(loan)

	if amount % 1000000 == 0{
		return "Sorry, doesn't found it"
	}
	
	for key, arg := range tempSaveArr {
		if arg[2] == loan{
			s += key + " "

		} 
	}
	return s
}
