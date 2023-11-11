package main
import "fmt"
func main(){
	div,err_msg:=divide(10,0)
}
func divide(dividend,divisor int) (int,error){
	if divisor ==0{
		return 0,error.New("Can't divided by zero")
	}
	return dividend/divisor,nil
}   