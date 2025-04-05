package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

// NOTE: Package level vars: These are defined at the top outside of all the funcs (these can be called package level vars)
// Global vars in `Go` can be declared with capital first letters. See helper package for ref.
var conferenceName string = "Go Conference"

const conferenceTickets uint8 = 50

var remainingTickets uint8 = 50
var bookingSlice = make([]UserData, 0)

// `type is for creating custom types. for eg we creating a custom type called `UserData`
type UserData struct {
	firstName       string
	lastName        string
	amountOfTickets uint8
}

// Note: main" function is the entrypoint for a Go program.
// 2. A program can only have one "main" function coz u can only have 1 entrypoint
func main() {

	greetUsers()

	for remainingTickets > 0 || len(bookingSlice) < 50 {

		firstName, lastName, userTickets := getUserInput()

		isValidName, isValidTicketNumber := helper.ValidateUserInputs(firstName, lastName, userTickets, remainingTickets)

		if isValidName && isValidTicketNumber {
			bookTickets(userTickets, firstName, lastName)
			go sendTicket(userTickets, firstName, lastName) //`go` keyword starts a new go routine. goRoutine is light weight thread managed by the go runtime

			firstNames := getFirstNames()
			fmt.Printf("These first names of bookings are: %v\n", firstNames)

			var noTicketsLeft bool = remainingTickets == 0
			if noTicketsLeft {
				// end the program
				fmt.Println("Tickets sold out.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Please enter a valid name")
			}
			if !isValidTicketNumber {
				fmt.Println("Please enter a valid ticket number")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Hello! Welcome to our %v booking app.\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "and the remaining amount of tickets are currently", remainingTickets)
	fmt.Println("To attend the", conferenceName, "get your tickets here!!")
}

func getFirstNames() []string {
	firstNames := []string{}
	// When ietrating through a list we get 2 values for each iteration i.e index and we get the element itself.
	//after the syntax `:=` we define what list are we looping through which in this case is the `bookingSlice`.
	// To iterate/loop through a slice we need a `range` operation. Range iterates over elements for diff data structs (not only array and slices)

	// Strings.Fields- splits the string with white space as operator
	// In Go `_` are used to identify unused variables
	for _, booking := range bookingSlice {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// NOTE: Unlike in solidity or other langs Go can can return any no. of values from a func not just one.
// func validateUserInputs(firstName string, lastName string, userTickets uint8) (bool, bool) {
// 	isValidName := len(firstName) >= 2 && len(lastName) >= 2
// 	isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets
// 	return isValidName, isValidTicketNumber
// } //PASTED IN THE HELPER>GO FILE

func getUserInput() (string, string, uint8) {
	// This is basically can be called an array type
	//var bookings [50]string //Note: Arrays in golang are fixed size. //For dynamic Golang has something called slices similar to an array

	var firstName string
	var lastName string
	var userTickets uint8

	fmt.Println("Enter First Name: ")
	fmt.Scan(&firstName) // This `&` here is called as pointer

	fmt.Println("Enter Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("No. of tickets to buy: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userTickets
}

func bookTickets(userTickets uint8, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets
	//	bookings[0] = userName
	// Append- used to add elements to the end of the slice. And, grows the slice if greater capacity is needed and returns the updated slice val

	// Here we'll create a map for the users. Note- its like a mapping in solidity
	// var userDataMap = make(map[string]string) //`make` is for creating an empty map. Also this is like string => string(in solidity)
	var userDataMap = UserData{
		firstName:       firstName,
		lastName:        lastName,
		amountOfTickets: userTickets,
	}
	// userDataMap["firstName"] = firstName      // This is the first key value pair that we're savin in the user map
	// userDataMap["lastName"] = lastName
	// //The above is like mapping from firstName to firstName(string => string)
	// userDataMap["amount of tickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookingSlice = append(bookingSlice, userDataMap)
	fmt.Printf("List of bookings: %v\n", bookingSlice)
	//Slice
	fmt.Printf("The whole slice: %v\n", bookingSlice)
	fmt.Printf("Slice Length: %v\n", len(bookingSlice))
	//Array
	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first index: %v\n", bookings[0])
	// fmt.Printf("The array type: %T\n", bookings)
	// fmt.Printf("Array Length: %v\n", len(bookings))

	fmt.Printf("Tickets Left: %v\n", remainingTickets)
}

func sendTicket(userTickets uint8, firstName string, lastName string) {
	time.Sleep(10 * time.Second) // This is like the vm.warp(kind of) it stucks for the specified amount of time
	var ticket = fmt.Sprintf("%v tickets for user %v %v\n", userTickets, firstName, lastName)
	fmt.Println("########")
	fmt.Printf("Sending ticket:\n %v \nto name- %v\n", ticket, firstName)
	fmt.Println("###########")
}
