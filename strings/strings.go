/*
Strings are represented as a sequence of bytes.
They are immutable, meaning that once a string is created, it cannot be changed. 
To create a new string, simply use double quotes (“”) or backticks (`).
*/

/*
Strings are immutable in Go — functions return a new string instead of modifying the original.
*/
func Stringspackage() {
	//Concatination
	//using + symbol
	A := "a"
	B := "b"
	Concatination := A + B
	fmt.Println(Concatination)

	//using strings package
	sliceOfStrings := []string{"a", "b", "c", "d"}
	afterconcatination := strings.Join(sliceOfStrings, "")
	fmt.Println(afterconcatination)
}

//find the substring
/*
To find a substring within a string, you can use the strings.Contains() and strings.Index() functions.
*/
func Substring_Checking_index() {
	str := "Golang is amazing!"
	fmt.Println(strings.Contains(str, "amazing"))
	fmt.Println(strings.Index(str, "amazing"))
}

//Replace substring
/*
To replace a substring within a string, use the strings.Replace() function.

s string – The original string where replacements will happen.
old string – The substring you want to replace.
new string – The substring you want to insert instead of old.
n int – The number of replacements to perform. important!!!!
*/
func Replace_String() {
	str := "Golang is amazing! Golang"
	newstring := strings.Replace(str, "Golang", "Java", 5)
	replaceAllOccurence := strings.Replace(str, "Golang", "Java", -1)
	fmt.Println(newstring)
	fmt.Println(replaceAllOccurence)
}

//You can split a string into a slice of substrings using the strings.Split() function.

func Splicting_string() {
	data := "substrings"
	slicedata := strings.Split(data, "")
	fmt.Println(slicedata)

	slice := "a,b,c"
	strings.Split(slice ,",")
}


- SEARCHING
- CONTAINS
- HASPREFFIX / HASSUFFIX
- INDEX / LASTINDEX

- COUNTING
- COUNT OCCURENCES

- MODIFYING
- REPLACE / REPLACEALL
- REPEAT 
- TOUPPER / TOLOWER

- TRIMMING
- TRIMSPACE - REMOVE LEADING / TRAILING SPACES
- TRIM / TRIM LEFT / TRIM RIGHT

- SPLITTING AND JOINING
- SPLIT / JOIN / FIELDS

- COMPARSIONS
- EQUAL FOLD - CHECKS CASE SENSITIVE OF THE WORDS
- COMPARE
