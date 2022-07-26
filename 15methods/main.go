package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs in GoLang")
	//no inheritance in go lang //no super or parent concepts

	neeraj := User{"Neeraj","neeraj08@gmail.com",true,22}
	fmt.Println("Neeraj data: ",neeraj)

	fmt.Printf("Neeraj details are :\n%+v\n",neeraj)

	fmt.Printf("Neeraj name is %v and email is %v \n",neeraj.Name,neeraj.Email)
	
	neeraj.GetStatus()
	neeraj.NewMail()

	fmt.Printf("Neeraj name is %v and email is %v \n",neeraj.Name,neeraj.Email)
	
	
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is User active: ",u.Status)
}

func (u User) NewMail(){
	u.Email = "test.go"
	fmt.Println("New main of this person is:",u.Email)
}
