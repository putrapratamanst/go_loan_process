package main


import(
	"bufio"
	"os"
	"fmt"
	"strings"
	"errors"
	"strconv"
	"os/exec"
	"./controllers"
	"time"

)

var tempSave = make(map[string]int)
var tempSaveArr = make(map[string][]string)
var tempId int

const MAX_REQUEST_PER_DATE = 50

func main(){
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = runCommand(cmdString)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func runCommand(cmdStr string) error {
	cmdStr = strings.TrimSuffix(cmdStr, "\n")
	arrCmdStr := strings.Fields(cmdStr)
	

	switch arrCmdStr[0] {

	case "check_temp":
		if tempSave == nil {
			fmt.Println("temp is nil")
		} 

		if tempSaveArr == nil {
			fmt.Println("temp Arr is nil")
		} 
		fmt.Println(tempSave)
		fmt.Println(tempSaveArr)
		fmt.Println(tempId)

		return nil

	case "exit":
		os.Exit(0)

	case "create_day_max":

		if _, ok := tempSave["create_day_max"]; ok {
				return errors.New("Max Create Day Already Setup")
		}

		if len(arrCmdStr) != 2 {
			return errors.New("Required for 1 arguments")
		} 
		
		num, _ := strconv.Atoi(arrCmdStr[1])

		if num != MAX_REQUEST_PER_DATE {
			return errors.New("Argument Value Must Be 50")
		} 

		fmt.Fprintln(os.Stdout, controllers.CreateDayMax(tempSave, num))
		return nil
	
	case "add":
		if len(tempSave) == 0{
			return errors.New("Max Create Day Not Setup Yet")
		}
		var newId string

		if len(arrCmdStr) != 4 {
			return errors.New("Required for 3 arguments")
		} 

		if _, err := strconv.Atoi(arrCmdStr[1]); err != nil {
			return errors.New("No KTP Must Be Integer!")
		} 

		if _, err := strconv.Atoi(arrCmdStr[3]); err != nil {
			return errors.New("Loan Must Be Integer!")
		}

		var arrNum []string
		for i, arg := range arrCmdStr {
				if i == 0 {
					continue
				}
				
			arrNum = append(arrNum, arg)
		}

		newId = reformatId(tempId) 
		tempSaveArr[newId] = arrNum
		fmt.Fprintln(os.Stdout, controllers.AddDataBorrower(tempSaveArr, newId))
		return nil

	case "status":
		if len(tempSave) == 0{
			return errors.New("Max Create Day Not Setup Yet")
		}
		if len(arrCmdStr) != 2 {
			return errors.New("Required for 1 arguments")
		} 

		fmt.Fprintln(os.Stdout, controllers.CheckStatus(tempSaveArr, arrCmdStr[1]))

		return nil
	
	case "installment":
		if len(tempSave) == 0{
			return errors.New("Max Create Day Not Setup Yet")
		}

		if len(arrCmdStr) != 3 {
			return errors.New("Required for 2 arguments")
		} 

		if len(tempSaveArr) == 0{
			return errors.New("Data Loan is Empty")
		}

		fmt.Fprintln(os.Stdout, controllers.Installment(tempSaveArr, arrCmdStr[1], arrCmdStr[2]))

		return nil

	case "find_by_amount_accepted":

		if len(tempSave) == 0{
			return errors.New("Max Create Day Not Setup Yet")
		}

		if len(arrCmdStr) != 2 {
			return errors.New("Required for 1 arguments")
		} 

		if len(tempSaveArr) == 0{
			return errors.New("Data Loan is Empty")
		}

		fmt.Fprintln(os.Stdout, controllers.FindByAmountAccepted(tempSaveArr, arrCmdStr[1]))

		return nil
	
	case "find_by_amount_rejected":

		if len(tempSave) == 0{
			return errors.New("Max Create Day Not Setup Yet")
		}
		if len(arrCmdStr) != 2 {
			return errors.New("Required for 1 arguments")
		} 

		if len(tempSaveArr) == 0{
			return errors.New("Data Loan is Empty")
		}
		fmt.Fprintln(os.Stdout, controllers.FindByAmountRejected(tempSaveArr, arrCmdStr[1]))

		return nil

	default:
		return errors.New("Command Not Found")
	}
	
	cmd := exec.Command(arrCmdStr[0], arrCmdStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}


func reformatId(id int)string{
	dt := time.Now()
	newFormatDate := dt.Format("010206")
	newId := strconv.Itoa(id)
	tempId = tempId + 1

	return newFormatDate + newId
}
