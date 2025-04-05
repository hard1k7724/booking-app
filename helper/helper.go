package helper

var MyVar string = "someVar" // eg of global vars in `Go`

// NOTE: Unlike in solidity or other langs Go can can return any no. of values from a func not just one.
// NOTE: Capitalizating the name of the func will export it explicitly from the main to the helper package
// Eg- exported- `ValidateUserInputs` not exported - `validateUserInputs`
func ValidateUserInputs(firstName string, lastName string, userTickets uint8, remainingTickets uint8) (bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets
	return isValidName, isValidTicketNumber
}
