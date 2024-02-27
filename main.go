package main

import (
	"fmt"
	"time"
	"strings"
	"booking-app/helper"
)

	var conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookTicket()
	var bookings=make([]UserData,0)
	type UserData struct{
		firstName string
		lastName string
		email string
		numberOfTickets uint
	}
	"Neil Raney"
	firstName: Neil
	lastName: Raney
	email: n@r.com
	Tickets: 3

var wg = weight group

func main() {
	greetusers(conferenceName,conferenceTickets,remainingTickets)
	fmt.Printf("Conferencetickets", "remainingTicketsis", "conferenceName is%T,Conferencetickets,remainingTicketsis,conferenceName is")

		fmt.Println(conferenceName)
	
		firstName,lastName,emailaddress,userTickets:=getusertickets()
		isValidName:=len(firstName)>=2&&len(lastName)>=2
		isValidemail:=strings.Contains(email,"@")
		isValidTicketNumber:=userTickets>0&&userTickets<=remainingTickets
		if isValidName&&isValidemail&&isValidTicketNumber{
		remainingTickets = remainingTickets - userTickets
		if userTickets<=remainingTickets{
			noremainingTickets:=remainingTickets==0
		if remainingTickets==0{
			//end program
			fmt.println("Our conference is booked out. Come back next year.")
			break
		}
		else{
		if !isValidName{
			fmt.Println("first name or last name is too short")
		}
		if!isValidEmail{
			fmt.Println("email does not contain an @ sign")
					}
		elseif !isValidTicketNumber
		fmt.println("number of tickets you entered is not valid")
		}
		wg.wait()
	}

		bookings = append(booking, firstName+" "+lastName)

		fmt.Printf("The whole array%v\n", bookings)
		fmt.Printf("The first value%v\n", bookings[0])
		fmt.Printf("Array type%v\n", bookings)
		fmt.Printf("Array length%v\n", len(bookings))

		fmt.Println("Thank you for booking your tickets. You will receive a confirmation e-mail at")
		userName = "Tom"
		userTickets = 2
		fmt.Println(firstName, userTickets)
		fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)
		firstNames := []string{}
		fmt.Printf("These are all our bookings.", bookings)

		
	}
}
city:="London"

}

func greetusers(conferenceName string,confTickets int,remainingTickets uint){
	fmt.Printf("Welcome to %v booking application.")
	fmt.Println("We have total of %v tickets and %v and this many are still available.", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func printFirstNames() []string {
	firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(bookings)
			firstNames = append(firstNames, booking.firstName)
		}
		return firstNames
}
func isValidName,isValidEmail,isValidTicketNumber=helper.ValidateUserInput(firstName string, lastName string, email string, userTickets uint,remainingTickets uint) (bool, bool, bool){
	isValidName:=len(firstName)>=2&&len(lastName)>=2
	isValidemail:=strings.Contains(email,"@")
	isValidTicketNumber:=userTickets>0&&userTickets<=remainingTickets
	return isValidName,isValidemail,isValidTicketNumber
}

func getUserInput(){
	var firstName String
		var lastName String

		bookings := string
		
		var userName string
		var userTickets int
		//ask user for their name
		fmt.Println("Enter your first name.")
		fmt.Scan(&firstname)
		fmt.Println("Enter your last name.")
		fmt.Scan(&lastname)
		fmt.Println("Enter your e-mail address.")
		fmt.Scan(&emailaddress)
		fmt.Println("Enter your tickets.")
		fmt.Scan(&userTickets)
		return firstName,lastName,emailaddress,userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string){
	var myslice []string
	var mymap map[string]string
	//create a map for a user
	var userData=make(map[string]string){
	firstName:firstName,
	lastName:lastName,
	email:email,
	userTickets:numberofTickets,
	}
	userData["firstName"]=firstName
	userData["lastName"]=lastName
	userData["email"]=email
	userData["numberOfTickets"]=strconv.FormatUint(uint64(userTickets),10)
}

func sendTicket(userTickets uint, firstName string, lastName string){
	time.Sleep(10*time.Seconds)
	fmt.Sprintf("%v tickets for %v%v",userTickets, first name, last name)
	fmt.Println("##########")
	fmt.printf("Sending ticket:\n%v to e-mail address%v\n",ticket,email)
	fmt.Println("##########")
	wg.Done()
}