package main
import (
	"errors"
	"fmt"
	"os"
)
func main(){
	filename:="myfile.txt"
	file,err := openFile(filename)

	//hanndle the error here
	if err != nil{
		fmt.Println(err)
		return
	}

	defer file.Close()
}
func openFile(fileName string) (*os.File,error){
	file,err:=os.Open(fileName)
	if err != nil {
		return nil,errors.New(fmt.Sprintf("failed to open file %v",fileName,err))
	}
	return file,nil
}