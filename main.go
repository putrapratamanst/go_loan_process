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

)

var tempSave = make(map[string]int)

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
		fmt.Println(tempSave)

		return nil

	case "exit":
		os.Exit(0)

	case "create_day_max":

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

		if len(arrCmdStr) != 4 {
			return errors.New("Required for 3 arguments")
		} 
		
		var arrNum []string
		for i, arg := range arrCmdStr {
				if i == 0 {
				continue
			}
			arrNum = append(arrNum, arg)
		}


		fmt.Println(arrNum);
		fmt.Fprintln(os.Stdout, controllers.AddDataBorrower(tempSave, "sdf"))
		return nil

	default:
		return errors.New("Command Not Found")
	}
	
	cmd := exec.Command(arrCmdStr[0], arrCmdStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
