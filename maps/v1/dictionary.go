package main

// Dictionary store definitions to words.
type Dictionary map[string]string

// Search find a word in the dictionary.
func (d Dictionary) Search(word string) string {
	return d[word]
}

//func (d Dictionary) Search(word string) string: Defines a method named Search on the Dictionary type.
//d[word]: Returns the value associated with word in the Dictionary map.
//The Search method looks up and returns the value for a given word from the Dictionary map.

// func Search(dictionary map[string]string, word string) string {
// 	return dictionary[word]
// }:

//"dictionary map[string]string" is the first parameter of the function

//"dictionary" is the name of the parameter

//map[string]string is the type of the parameter

//"word string" is the second parameter -
//"word" is the name of the parameter -
//"string" is the type of the parameter, which is a string

//"string" is the return type of the function -
//the func returns a value of type "string"

//putting it all together 'func Search(dictionary map[string]string, word string)'
//string defines a function named 'Search' that takes two parameters: a map ('dictionary')
//with string keys and string values, and a string ('word').The function returns a string.
