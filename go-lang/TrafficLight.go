package main

import "fmt"
import "bufio"
import "os"
import "strings"

func main() {
	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Please input the light(red, yellow, and green)");

	light, _ := reader.ReadString('\n')
	//In this part is a bit tricky because at the end of the input contain "\n" 
	//so we need to trim it
	if strings.TrimRight(light, "\n") == "red" {
		fmt.Println("Stop");
	} else if strings.TrimRight(light, "\n") == "yellow" {
		fmt.Println("Stand By");
	} else if strings.TrimRight(light, "\n") == "green" {
		fmt.Println("Run");
	}
}