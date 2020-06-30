package main


import(
	"bufio"
	"os"
	"fmt"
	"strings"
	"errors"
	"strconv"
	"os/exec"
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

		fmt.Fprintln(os.Stdout, createDayMax(num))
		return nil
	
	default:
		return errors.New("Command Not Found")
	}
	
	cmd := exec.Command(arrCmdStr[0], arrCmdStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func createDayMax(numbers int) string {

	tempSave["create_day_max"] = numbers
	fmt.Println(tempSave)

	callback := "Created max request with requests" 
	return callback
}
