package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)
const conferenceTickets int = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

		greetUsers()

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email )

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames();
			fmt.Printf("The first name of bookings are: %v\n", firstNames);

			if remainingTickets == 0 {
				fmt.Println("Sorry, we are sold out. Please try again later.");
				//break;
			}
		} else {
			if !isValidName {
				fmt.Println("Sorry, we are unable to process your booking. Please check your name and try again.");
			}
			if !isValidEmail {
				fmt.Println("Sorry, we are unable to process your booking. Please check your email and try again.");
			}
			if !isValidTicketNumber {
				fmt.Println("Sorry, we are unable to process your booking. Please check your tickets and try again.");
			}
		}
		wg.Wait()
	city := "Vietnam"

	switch city {
		case "Vietnam", "Japan":
			fmt.Println("Welcome to Vietnam and Japan");
		case "Thailand", "Singapore":
			fmt.Println("Welcome to Thailand and Singapore");
		default:
			fmt.Println("No such city");
	}
}
func greetUsers() {
		fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}
func getFirstNames () []string {
		firstNames := []string{};
			for _, booking := range bookings {
				firstNames = append(firstNames, booking.firstName);
			}
			return firstNames
}

func getUserInput() (string, string, string, uint) {
		var firstName string;
		var lastName string;
		var email string;
		var userTickets uint;
		fmt.Println("Enter your first name: ");
		fmt.Scan(&firstName);

		fmt.Println("Enter your last name: ");
		fmt.Scan(&lastName);

		fmt.Println("Enter your email: ");
		fmt.Scan(&email);

		fmt.Println("Enter your tickets: ");
		fmt.Scan(&userTickets);

		return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint , firstName string, lastName string, email string) {
		remainingTickets -= userTickets

		var userData = UserData{
			firstName: firstName,
			lastName: lastName,
			email: email,
			numberOfTickets: userTickets,
		}

		bookings = append(bookings, userData);
		fmt.Printf("List of bookings is %v\n", bookings);

		fmt.Printf("Thank you %v %v for booking %v tickets. Your email is %v\n", firstName, lastName, userTickets, email);
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets,conferenceName)
}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var	ticket = fmt.Sprintf("%v tickets for %v %v\n",userTickets, firstName, lastName)
	fmt.Println("-----------------")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("------------------")
	wg.Done()
}
