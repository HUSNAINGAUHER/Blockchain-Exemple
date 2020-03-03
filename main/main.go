package main
import "os"
import "strconv"


import a1 "github.com/HUSNAINGAUHER/assignment02IBC"

func main() {


if len(os.Args) > 2 {
	i1,err := strconv.Atoi(os.Args[2])
	if err != nil {

	}
	a1.Server(os.Args[1],i1)
}


if(len(os.Args) ==2 ) {


	go a1.Client(os.Args[1])
	a1.Server(os.Args[1],-1)
	
	//normal case
}




}