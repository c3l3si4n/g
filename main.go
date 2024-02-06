package g

import "fmt"


func Exported(arg string){
	fmt.Println(arg)
	return
}

func main(){
	fmt.Println("!")
}
