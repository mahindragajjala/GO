package main

/*
 Group          Key Functions                 Output/Input          Memory Type       
 Print          Print, Printf, Println  Terminal              Temporary strings 
 String Format  Sprint, Sprintf, etc.     Returns strings       Heap              
 File Output    Fprint, Fprintf, etc.     Files, Writers        Efficient         
 Input          Scan, Scanf, etc.         From user             Variable memory   
 Parsing        Sscan, Sscanf, etc.       From string           Heap+GC managed   
 Byte Append    Appendf, Append           To byte slice         Efficient append  
 Error Format   Errorf                      Returns error       Heap              
 Custom Format  Stringer, Formatter       Interfaces for fmt  Advanced usage    
*/


func main(){
  PRINTLN/PRINTF
  
  fmt.Print("without new line Print the text...")
  fmt.Printf("Uses format specifiers (%s, %d, etc.)")
  fmt.Println("with new line.. ")

  SPRINT/SPRINTF/SPRINTLN

	var name string = "stringdata"
	var id int = 123456
	stringdata := fmt.Sprint("username:", "", name, "id:", id)
	fmt.Println(stringdata)

	//user for access specifiers...
	Sprintfdata := fmt.Sprintf("%T", name)
	fmt.Println(Sprintfdata)

	//create the string and create the new line
	dataSprintln := fmt.Sprintln(name)
	fmt.Println(dataSprintln)

  FPRINT/FPRINTLN

  	/*
	   Fprint
	   It prints normal text (no formatting).
	   It sends that text to a writer (like screen, file, or network).
	   You just give it the writer and what to print. */

	fmt.Fprint(os.Stdout, "printing....")

	fmt.Println()

	/*
	   Fprintf
	   It prints formatted text (like Hello, Name: Mahindra).
	   You can use format placeholders like %s, %d, etc.
	   It also sends the result to a writer.
	*/

	name := "Mahindra"
	age := 25
	fmt.Fprintf(os.Stdout, "Name: %s, Age: %d\n", name, age)

	fmt.Println()

	file, _ := os.Create("file1.txt")
	fmt.Fprint(file, "writing in the file....?")

SCANF /SCANLN / SCAN
/*
INPUTS FROM WHERE - IT WILL COME
- STANDARD INPUT [KEYBOARD / TERMINAL]
- COMMAND LINE ARUGMENTS
- FILES 
- ENV
- HTTP REQUESTS
- DB'S

OUTPUTS FROM WHERE - IT WILL COME
- STANDARD OUTPUT[TERMINAL/CONSOLE]
- FILES
- LOGS
- CMD OUTPUTS[RETURN VALUES/EXIT CODES]
- HTTP RESPONSE[WEB APPLICATION]
- JSON/XML/YAML
- NETWORK SOCKETS 
- DATABASES
- EXTERNAL APIS
- GUI OR FRONTEND
*/

  	/*
		Input from the user in Go
	*/
	//fmt.Scan - func Scan(a ...any) (n int, err error)
	/*
			   Reads space-separated input from standard input (stdin).

			   Stores each value into the variables passed by reference (with &).
			   Memory Management:

			   Variables are stored in stack or heap, depending on scope.

		Scan() stores values directly into those memory addresses.
	*/

	fmt.Print("Enter your name and age: ")
	fmt.Scan(&name, &age) // Input: Mahindra 27 */

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	//fmt.Scanf
	//func Scanf(format string, a ...any) (n int, err error)
	/*
			Like Scan, but expects a format string (%d, %s, etc.).

		Gives more controlled parsing.
	*/
	fmt.Scanf("%T", &name)

	/*
		Reads input from stdin, like Scan, but stops at a newline.

		Does not allow more inputs after a newline (\n).

		It throws an error if there are extra tokens after the last expected value.
	*/
	fmt.Print("Enter your name and age: ")
	fmt.Scanln(&name, &age) // Input: Mahindra 27 [Enter]

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)


  SSCAN / SSCANF / SSCANLN

  
	//fmt.Sscan
	//func Sscan(str string, a ...any) (n int, err error)
	/*
			EXTRACTING FROM THE INPUT
		   Parses space-separated values from a string.

		   Assigns parsed values to the provided variables (must use &).
	*/
	var name string
	var age int

	input := "Mahindra 27"
	fmt.Sscan(input, &name, &age)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)


  
	// fmt.Sscanf
	/* func Sscanf(str string, format string, a ...any) (n int, err error) */
	var name string
	var age int

	inputdata := "Name: Mahindra Age: 27"
	fmt.Sscanf(inputdata, "Name: %s Age: %d", &name, &age)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

  
	//fmt.Sscanln
	//func Sscanln(str string, a ...any) (n int, err error)
	/*
	   Similar to Sscan, but stops parsing at newline.

	   Returns an error if there is extra data after the expected variables.
	*/
		var firstName, lastName string

	input := "John Doe"
	fmt.Sscanln(input, &firstName, &lastName)

	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)

ERRORF
  
file, err := os.Open("config.json")
if err != nil {
    return nil, fmt.Errorf("failed to open config.json: %v", err)
}

}
