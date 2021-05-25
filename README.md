Backend General Questions & Go Questions
1. What is the function of DNS*
map domain name with ip adress

2. What is TTL in DNS*
time for update DNS

3. Tell us a couple advantages when set default DNS TTL for very long period, eg. 24 hr*
reduce load for dns
redure cost for server

4. Tell us a couple disadvantages when set default DNS TTL for very long period, eg. 24 hr*
increase load for server
if failed over client have to wait untill dns update to new server

5. Tell us a couple advantages when set default DNS TTL for very short period, eg. 10 secs*
most advantage with loadbalance for server
if server update user will update the system more quickly

6. Tell us a couple disadvantages when set default DNS TTL for very short period, eg. 10 secs*
increase load for dns
increase cost for server

7. What is a default HTTP and HTTPS port*
http 80
https 443

8. What is a common meaning of HTTP 2xx response code*
request that send to api get status look like success

9. What is a common meaning of HTTP 4xx response code*
request that send to api get status look like client error

10. What is a common meaning of HTTP 5xx response code*
request that send to api get status look like server error

11. If a web client sends a request to a server using JavaScript and parses the response body using JSON.parse(response.body), is it necessary to set the response’s header as Content-Type: application/json ? and why?*
for tell server that we send content json not for media and protect server for send any that harmful to server

12. Which one is worse in terms of performance and why?: Increment an integer variable vs INCR a key on Redis (both perform 100k times sequentially)*
INCR because INCR store data as string and contains will slow in long term of performance

13. Given endpoints /auth/signin and /me which are used to authenticate and get basic information of authenticated users respectively. There is a requirement to fetch basic information which requires authentication first, can you draw a sequence diagram with 3 actors: Browser, Server and Database to demonstrate the requirement? You can choose any authentication mechanism.*

/auth/singin for authenticate
/me for information of authenticated

                browser                             server                      database
/auth/singin ---> || send request user/pass --------> || authen user ------------> || ---
                  ||                                  ||                           ||   | query user
                  || <------ send authenticated/token || <----- validate user/pass || <-|

/me ------------> || send request with token -------> || validate token ---
                  ||                                  ||                   |
                  ||                                  ||   <---------------|
                  ||                                  || get user information ---> || ---
                  ||                                  ||                           ||   | query information
                  || <----------send user information || <------------------------ || <-|

14. Please write a very short introduction about docker*
software that provide mini server for application
15. Please write a very short introduction about kubernetes*
software for manage container from docker
16. Can you briefly explain how a single pod in kubernetes communicates with another pod on another machine?*
in single pod look like single server and can access to another pod by ip:port

17. Write a couple features that Layer 4 load balancer has*

18. Write a couple features that Layer 7 load balancer have over Layer 4 load balancer*

19. If you want to route requests to different machines by HTTP path prefix, what kind of load balancer is needed and why?*

20. What are the differences between Encryption and Hashing? Give us some algorithms for both.*
encryption aes256
hash sha256
encryption can decryption from cipher text to plain text but hash can only verify data cant reverse to plain text

21. What is different between Symmetric and Asymmetric encryption? Give us some algorithms for both.*
symmetric aes256
sysmetric use the same key to encrypt and decrypt

asymmetric rsa256
asymmetric user public key to encrypt and private key for decrypt

22. If you’re using macOS and you have to build a Go program to run on Linux, how would you execute the go build command?*
env GOOS=linux GOARCH=arm64 go build

23. What’s wrong with this Go program?*
package prog1

func main() { }

go program support main function on package main only, have to change prog1 to main

24. Given the following statement, what is the difference between these 2 variables?*
var v1, V2 int

two variable of v1 and V2 are the same type and value but go convension if the first character is lowercase for only this function but
if the first character is uppercase for export variable or let another package use too

25. Write a statement that create a slice of int with length 0 and capacity 100*
slice := []int{100}

26. Write a statement that create a array of int with size 5*
array := [5]int{}

27. Given an empty slice of int (numbers := []int{}), write a statement that add two numbers 100 and 200 to that slice*
numbers = append(numbers, 100)
numbers = append(numbers, 200)

28. Write a short summary about defer keyword*
defer will do after end function, if we push state of function to defer it will process at last function

29. Is Go routine the same as thread and why?*
go routine can provide tasks to workers 

30. Write a snippet of struct definition that can be render as following JSON*
{ “greeting_message”: “some message” }

type Greeting struct {
    GreetingMessage string `json:"greeting_message"`
}

31. Write a snippet that print value from slice of string given var employees []string*
fmt.Printf("%v", employees)

32. What is the difference between the = and := operator?*
:= operator for assign new variable
= operator for assign allocated variable

33. Write an if statement to show how to check whether a given map (var ages map[string]int) has a key “John”?*
if ages["John"] == 0 {

}

34. Given 2 variables item := “Chocolate” and price := 15.33333, write a statement that create new variable “line” with this format “Chocolate (15.33)”*
line := fmt.Sprintf("%v%v", item, price)

35. What is sync.Mutex and when you use, tell us a use case in practice*
sync.Mutex for goroutine that provide lock/unlock option for prevent another access allocated variable or function until unlock

36. Write a statement to print current system’s time in second since epoch*
fmt.Println(time.Now().Format("2006-01-02 15:04:05.000000000"))

37. Please create a simple go project that satisfy the following requirements*
You're asked to write a simple JSON API to summarize COVID-19 stats using this public API, https://covid19.th-stat.com/api/open/cases.

1. Your project must use Go, Go module, and Gin framework

2. You create a JSON API at this endpoint /covid/summary

3. The JSON API must count number of cases by provinces and age group

4. There are 3 age groups: 0-30, 31-60, and 60+ if the case has no age data, please count as "N/A" group

5. We encourage you to write tests, which we will give you some extra score

6. Please zip the whole project and upload to the form.

--- sample response --

{

"Province": {

"Samut Sakhon": 3613,

"Bangkok": 2774

},

"AgeGroup": {

"0-30": 300,

"31-60": 150,

"61+": 250,

"N/A": 4

}

}