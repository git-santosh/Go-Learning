package main 
import ( 
	"log"
	"errors"
	"fmt"
)
var MathError = errors.New("Math Error : Sqrt of negitive number")
func main(){ 
	fmt.Printf("%T\n",MathError)
	_,err:=sqrt(-10)
	if err != nil { 
		log.Fatalln(err)
	}
}
func sqrt(f float64)(float64,error){
	if f<0 { 
		// return 0,errors.New("Math Error : Sqrt of negitive number")
		// return 0 , MathError
		return 0, fmt.Errorf("Math Error : Sqrt of negitive number : %v",f)
	}
	return 42,nil
}