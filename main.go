/**
 *
 * https://www.youtube.com/watch?v=yyUHQIec83I
 * https://gitlab.com/nanuchi/go-full-course-youtube
 *
 */

package main

import (
	//"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

//var bookings []string
//var bookings = make([]map[string]string, 0)
var bookings = make([]User, 0)

type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

/**
 *
 *
 */
func main() {

	greetUsers()

	//fmt.Printf("Welcome to %v booking application \n", conferenceName)
	//fmt.Println("We have a total of ", conferenceTickets, " tickets and ", remainingTickets, " tickets")
	//fmt.Println("Get your tickets here to attend")
	//fmt.Printf("Valor: %b ", 63)

	// Assk user for their name
	for {
		firstName, lastName, email, userTickets := getUserInput()

		//isValidName, isVaidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		isValidName, isVaidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isVaidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket()

			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings: %v\n", firstNames)

			var noTicketRemaining bool = remainingTickets == 0
			if noTicketRemaining {
				fmt.Println("Sold out!!!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("User's firstname or lastname is too short")
			}
			if !isVaidEmail {
				fmt.Println("Entered email is not valid")
			}
			if !isValidTicketNumber {
				fmt.Printf("The requested tickets of %v exceeds the available ticket of %v ", userTickets, remainingTickets)
			}

			fmt.Printf("Your input data is not valid. Please try again!")
			continue
		}

	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have a total of ", conferenceTickets, " tickets and ", remainingTickets, " tickets")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)
		//firstNames = append(firstNames, names[0])
		//firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name...")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name...")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email...")
	fmt.Scan(&email)

	fmt.Println("Please enter the number of ticket...")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	// First attempt
	//bookings = append(bookings, firstName+" "+lastName)

	// Second attempt: Using maps
	// Create a map for a user
	// map[key_data_type]value_data_type
	/*
		var userData = make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

		bookings = append(bookings, userData)
	*/

	// Third attempt: Using structs
	var user = User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, user)

	/*
		fmt.Printf("Whole array: %v\n", bookings)
		fmt.Printf("First element: %v\n", bookings[0])
		fmt.Printf("Type array: %T\n", bookings)
		fmt.Printf("Length array: %v\n", len(bookings))
	*/

	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation email at %v \n", email)
	fmt.Printf("Remaining tickets: [%v] for the conference [%v]\n", remainingTickets, conferenceName)
}

func sendTicket() {
	time.Sleep(10 * time.Second)
	fmt.Println("===========================")
	fmt.Println("Sending tickets to the user")
	fmt.Println("===========================")
}
