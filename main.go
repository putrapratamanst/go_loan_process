package main


import(
	"bufio"
	"os"
	"fmt"
)

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
	return nil
}
