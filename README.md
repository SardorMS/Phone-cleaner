# phone-cleaner
Implementing character replacement for formatting phone numbers using command line arguments in Go (Golang).

# Functions

**strings.ReplaceAll:**

This function is used to replace all the old string with a new string. If the given old string is empty, then it matches at the starting of the string and after each UTF-8 sequence it is yielding up to M+1 replacement for M-rune string.

**Syntax:**
```
func ReplaceAll(str, oldstr, newstr string) string
```
Here, str is the original string, oldstr is the string which you wants to replace, and newstr is the new string which replaces the oldstr. Let us discuss this concept with the help of an example:



**strings.NewReplacer():**

Function in Golang returns a new Replacer from a list of previous, new string sets. Substitutions are performed within the order they show up within the target string, without overlapping matches. The old string comparisons are done in argument order. The old string comparisons are done in argument order.

**Syntax**
```
func NewReplacer(oldnew ...string) *Replacer
```
Remember NewReplacer panics if given an odd number of arguments.



# Installation

1. Clone this code into a directory:
 ```
 git clone https://github.com/SardorMS/phone-cleaner
 ```

2. Create a module:

For the replaceAll function use:
 ```sh
 go mod init replaceAll
 ```
For the newReplacer function use: 
 ```sh
 go mod init newReplacer
 ``` 
 
3. Create a generated executable file:
```sh
go build -o cmd/cleaner/main.go
```

4. To RUN use the terminal and one argument of command line:
```sh
cleaner.exe "ANY_EXAMPLE_PHONE_NUMBER"
```
![alt tag](https://github.com/SardorMS/Phone-cleaner/blob/replacement/src/Example_replacement.png)
