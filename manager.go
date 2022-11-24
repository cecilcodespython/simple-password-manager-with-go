package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const pwd_db = "password.db"

func checkErr(err error){
	if err != nil {
		log.Fatal(err)
		return
	}
}

func store(platform,username,password string){
	entry := platform +","+username+","+password+"\n"
	f,err := os.OpenFile(pwd_db,os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	l,err := f.WriteString(entry)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(l,"bytes written")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func collect(platform string){
	f,err := os.Open(pwd_db)
	checkErr(err)
	input := bufio.NewScanner(f)
	for input.Scan() {
		entry := strings.Split(input.Text(), ",")
		if entry[0] == platform{
			fmt.Println(entry[1],entry[2])
			return
		}
	}
	fmt.Print("Platform %s not know \n", platform)
}




func main()  {
	var args []string = os.Args
	if args[1] == "add"{
		store(args[2],args[3],args[4])
	}else if args[1] == "get"{
		collect(args[2])
	}else {
		fmt.Println("Invalid Operation",args[1])
	}
	
	
}