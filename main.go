package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings:=string
	greetusers(conferenceName,conferenceTickets,remainingTickets)
	fmt.Printf("Conferencetickets", "remainingTicketsis", "conferenceName is%T,Conferencetickets,remainingTicketsis,conferenceName is")

		fmt.Println(conferenceName)
	for {
		var firstName String
		var lastName String

		bookings := string
		
		var userName String
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

func greetusers(confName string,confTickets int,remainingTickets uint){
	fmt.Printf("Welcome to %v booking application.")
	fmt.Println("We have total of %v tickets and %v and this many are still available.", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func printFirstNames() []string {
	firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(bookings)
			firstNames = append(firstNames, names[0])
		}
		return firstNames
}
func isValidName,isValidEmail,isValidTicketNumber=validateUserInput(firstName string, lastName string, email string, userTickets uint,remainingTickets uint) (bool, bool, bool){
	isValidName:=len(firstName)>=2&&len(lastName)>=2
	isValidemail:=strings.Contains(email,"@")
	isValidTicketNumber:=userTickets>0&&userTickets<=remainingTickets
	return isValidName,isValidemail,isValidTicketNumber
}