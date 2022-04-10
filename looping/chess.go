package main

import "fmt"

func main() {
	var arr [8][8]int
	col := 8
	row := 8

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf(" %d ",arr[i][j])
			if(j==col-1){
				fmt.Println("")
			}
		}
	}
	

}
