package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	line := 0
	wc := 0
	ch := 0
	files := os.Args[1:]
	if len(files) == 0 {
		line, wc, ch = countLines(os.Stdin)
	}else{
		for _, arg := range files{
			f,err := os.Open(arg)
			if err != nil{
				fmt.Fprint(os.Stderr, "dup2: ", err)
				return 
			}
			line, wc, ch = countLines(f)
			f.Close();
		}
	}
	fmt.Printf("line wc   ch\n")
	fmt.Printf("%d   %d   %d \n",line, wc, ch)
	
}
func countLines(f *os.File)( int,int, int){
	line :=0
	ch := 0 
	wc := 0
	input := bufio.NewScanner(f)
	for input.Scan(){
		line++
		ch+=len(input.Bytes())
		i := 0
		for i < len(input.Bytes()) {
			if input.Bytes()[i] == 32 && i != 0 && input.Bytes()[i-1] != 32{
				wc++
			}else if i == len(input.Bytes())-1 && input.Bytes()[i] != 32 {
				wc++
			}
			i++
		}
	}
	return line, wc, ch

}
