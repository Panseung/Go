# Go with Nico

### Index

### Base 

- [1. import & package](#import--package)
- [2. const let](#const-let)
- [3. defer 함수](#defer-함수)
- [4. 반복문](#반복문)
- [5. 조건문](#조건문)
- [6. switch](#switch)
- [7. memory](#memory)
- [8. Arrays and Slices](#arrays-and-slices)
- [9. map](#map)
- [10. struct](#struct)

### 1st Project: Bank

- [1. pointer★](#pointer)
- [2. method](#method)
- [3. function vs method](#function-vs-method)
- [4. error](#error)
- [5. struct의 특정값 불러오는 내부 method](#struct의-특정값-불러오는-내부-method)
- [6.Bank project main.go & accounts.go](#bank-project-maingo--accountsgo) 

### 2nd Project: Bank

- [Map type](#map-type)
- [Map & pointer](#map--pointer)
- [var 한 번에 선언](#var-한번에-선언)
- [정리: Add, Update, Delete (main.go, mydict.go)](#add-update-delete-maingo-mydictgo)

### 3rd Project: URL CHECKER & GO ROUTINES

- [hitURL & URLChecker](#hitURL--URLChecker)

- [goroutines](#goroutines)
- [channels](#channels)
- [URLChecker with Goroutine](#urlchecker-with-goroutine)

### 4th csv parser

- Fatalln
- 

### Base

#### import & package

main. go

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

something.go

```go
package something

import "fmt"

func sayBye() {
	fmt.Println("Bye")
}

func SayHelo() {
	fmt.Println("Hello")
}
```

sayBye와 같이 소문자로 적으면 export되지 않음
SayHello처럼 대문자로 시작해야 export됨



#### const let

```go
func main() {
    name := "panseung"    /// 이건 축약형 아래 코드와 같음, 함수 안에서 첫번째 값을 기준으로
       					  /// Go가 자동으로 type을 찾아줌
    var name string = "panseung"    /// 함수 안에서만 작동함
}
```



```go
package main

import "fmt"

func multiply(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))
}
```

argument와 return값을 TS처럼 타입을 지정해주어야함



```go
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, _ := lenAndUpper("panseung")
	fmt.Println(totalLength)
}
```

Go는 return 값을 여러개 가질 수 있음
두 개의 return을 반환하는 함수에서 하나의 값만 받고 싶을때는 필요 없는 값은 '_' 언더스코어를 사용

```go
func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
```

위 코드처럼 return값을 아래서 말고 위에서 미리 지정 가능



#### defer 함수

함수 내에 작성하며, 해당 함수가 실행이 종료될 때 실행되는 함수



### ...

```go
func main(numbers ...int) int{
    
}
```

...을 찍으면 arguments로 여러개를 받을 수 있음



#### 반복문

```go
func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}
	return 1
}
```

range는 기본적으로 첫 return값으로 index를 주기 때문에
위와 같이 변수 두개(index, number)를 넣으면
각각 index와 실제 값을 return해줌



#### 조건문

```go
func canIDrink(age int) bool {
	if age < 18{
		return false
	} 
    return true
}
```

기본적인 조건문

```go
func canIDrink(age int) bool {
	if koreanAge := age+2; koreanAge < 18{
		return false
	} 
		return true
}
```

위 코드처럼 if문 안에 새 변수 생성 가능
물론 if문 위에(밖에)서 koreanAge라는 변수 선언 가능하지만
둘의 차이점은 해당 변수가 조건문에서만 쓰이기 위한건지에 대한 차이!
위 코드처럼 짜면 조건문(if문) 안에서만 koreanAge가 사용될 거라는 의도를 
다른 사람들이 쉽게 알 수 있음

이렇게 선언하는걸 `variable expression`이라고 함!



#### switch

```go
func canIDrink(age int) bool {
	switch koreanAge := age+2; koreanAge  {
		case 10:
			return false
		case 18:
			return true
			} 
		return false
}
```

switch는 else if의 난무를 막을 수 있는 문법코드
if문과 마찬가지로 `variable expression` 사용 가능



#### memory

```go
func main() {
	a := 2
	b := &a
	fmt.Println(a, b)
}
```

하면 a의 값과, a의 메모리 주소 찍힘

반대로
```go
func main() {
	a := 2
	b := &a
    a = 5
	fmt.Println(*b)
}
```

5가 출력됨
b는 a의 주소 자체를 copy하고 있고 이 메모리 주소를 '*'를 통해 역참조가 가능(다시 값을 불러올 수 있음)

`*b= 20` 코드를 추가하고
`fmt.Println(a)`를 실행하면 20이 출력됨



#### Arrays and Slices

```go
func main() {
	names := [5]string{"T", "J", "M", "A", "S"}
	names[3] = "AA"
	fmt.Println(names)
}
```

Array는 크기를 지정해주고 시작해야함

Slice는 크기제한이 없는 array! (python list랑 비슷한것같기도?)

```go
func main() {
	names := []string{"T", "J", "M"}
    names = append(names, "AA")
	fmt.Println(names)
}
```

append는 names에 바로 원소를 추가하는 것이 아니라
기존 array에 원소를 추가한 새로운 배열을 return함

#### map

```go
func main() {
	panseung := map[string]string{"name": "panseung", "age": "28"}
	fmt.Println(panseung)
	for _, value := range panseung {
		fmt.Println(value)
	}
}
```

```
map[age:28 name:panseung]
panseung
28
```



#### struct

```go
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"apple", "banana", "melon"}
	// panseung := person{"panseung", 27, favFood}
	panseung := person{name: "panseung", age: 27, favFood: favFood}
	fmt.Println(panseung)
}
```

주석과 같은 코드지만 아래가 좀 더 타인이 봤을 때 명확하게 보이기 때문에
clean한 코드라고 할 수 있음





### 1st Project: Bank

#### pointer ★★★★★★★★★★★★★★★

1. 포인터란 메모리 주소를 값으로 가진 변수(의 타입을)를 포인터라고 한다
2. `*`는 포인터 변수를 "역참조" 하는 데도 사용된다!!!!!!

## ★★★★★★★

```go
package accounts

type BankAccount struct {
	owner   string
	balance int
}

func NewAccount(owner string) *BankAccount {
	account := BankAccount{owner: owner, balance: 0}
	return &account
}
```

복사본을 return하는게 아니라 실제 메모리를 return함으로써
불필요한 메모리를 줄이는 중요한 최적화 작업!!!!!!!!!!!!!!!!!!!!!!!!!

(&account를 return하는데 이는 메모리 주소를 반환하는 것이므로
메모리 주소를 타입으로 가진 변수를 return하기 때문에 
*BankAccount로 타입을 작성)

### Method

receiver의 첫 글자는 struct의 첫 글자로 시작할 것 (단, 소문자로!)
```go
func (a Account) Deposit(amount int) {
	a.balance += amount
}
```

func와 거의 동일한 구조
함수 이름 앞에 괄호를 치고 struct의 첫글자를 소문자로 써주고 그 뒤에 struct를 써주면 끝

★★★★★★★★★★★★★

#### 위코드와 아래코드의 차이점

위 코드는 복사본 account를 가리키지만
아래 코드는 실제 method를 호출한 account를 가리킴
따라서 위 코드는 실제 호출한 account의 balance에 변화가 없지만
아래 코드는 호출한 account의 balance에 변화가 생김
즉 아래코드는 복사본을 만들지 않아 메모리를 절약 

```go
func (a *Account) Deposit(amount int) {
	a.balance += amount
}
```



##### function vs method

- function: 객체(object)와 관계없이 특정 작업을 수행하는 코드
- method: 클래스(struct)내에 포함되어 있는 함수
                  즉 객체와 관련되어 있는 함수



#### error

```go
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw you are poor")
	}
	a.balance -= amount
	return nil
}
```

boolean이 true / false를 return하듯이
errors와 nil 두가지를 return
nil은 python의 none과 같이 아무 일도 일어나지 않음
그렇다고 errors가 error를 발생시키는것도 아님!!!!!!!! ★
Go에는 exception이 없기 때문에 따로 python처럼
에러창에 경고문을 띄우거나 하지 않고 그냥 error를 인식하고
아무 처리를 하지 않음

따라서 경고문(Can't with~~)을 띄우고 싶으면 main.go에 다음과 같이 작성하기
```go
func main() {
	account := accounts.NewAccount("panseung")
	account.Deposit(10)
	err := account.Withdraw(13)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(account.Balance())
	}
}
```



#### struct의 특정값 불러오는 내부 method

```go
func (a *Account) String() string {
    return "whatever you want"
}
```

를 적고 main.go에서 
```go
fmt.Println(account)
```

를 실행하면 "whatever you wnat"가 나오게됨
이와 같이 String()은 해당 struct를 출력할 때 
어떤 문구를 출력할지 정해주는 method
(python에서 `__str__`과 같은 역할)



### Bank project main.go & accounts.go

main.go
```go
package main

import (
	"fmt"

	"github.com/panseung/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("panseung")
	account.Deposit(10)
	fmt.Println(account)
	
	err := account.Withdraw(7)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(account.Balance())
	}
}
```

accounts.go
```go
package accounts

import "errors"

type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (a Account) String() string {
    return "whatever you want"
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw you are poor")
	}
	a.balance -= amount
	return nil
}
```





### 2nd Project: Bank

#### Map Type

Map은 키(Key)에 대응하는 값(Value)을 신속히 찾는 해시테이블(Hash table)을 구현한 자료구조이다.
"map[Key타입]Value타입"

```go
type Dictionary map[string]string
```

Map type은 두가지를 return함
키에 상응하는 value, 키가 존재하는지에 대한 boolean

```go
var errNotFound = errors.New("Not Fount")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}
```



#### Map & pointer

Question

```
Why "Update" method doesn't include point receiver(*) beside Dictionary? 
Is this reason equal to why "Delete" method doesn't include point receiver?
(reason: Dictionary is hash map)

So, if I use struct, when I make a Update, Delete method, must I use point receiver? Plz reply! Thanks! :)
```

answer
```
Correct!
Good question, we don't need the * on the receiver because maps on Go are automatically using *
```



#### var 한번에 선언

```go
var (
	errNotFound = errors.New("Not Fount")
  	errWordExists = errors.New("That word already exists")
 	errCandUpdate = errors.New("Cant update non-existing word")
)
```

이런식으로 var 한 번에 여러개 선언 가능



#### Add, Update, Delete (main.go, mydict.go)

mydict.go
```go
package mydict

import "errors"

type Dictionary map[string]string

var (
	errNotFound = errors.New("Not Fount")
  errWordExists = errors.New("That word already exists")
  errCantUpdate = errors.New("Cant update non-existing word")
	errCantDelete = errors.New("Cant Delete non-existing word")
)


func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists  
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error{
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) error{
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
```

main.go
```go
package main

import (
	"fmt"
	"github.com/panseung/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseword := "hello"
	dictionary.Add(baseword, "First")
	err := dictionary.Update(baseword, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseword)
	fmt.Println(word)
}
 
```





## 3rd Project: URL CHECKER and Goroutines



#### hitURL & URLChecker

main.go

```go
package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {

	var results = map[string]string{}

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

var errRequestFailed = errors.New("Request Failed")

func hitURL(url string) error {
	fmt.Println("Checking: ", url)
	resp, err := http.Get((url))
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
```

http.Get(url) 코드는
해당 주소가 잘 접속이 되는지 확인하는 함수!!



#### Goroutines

함수 앞에 go를 붙여주면 main함수고 실행하는 동안 해당 함수를 비동기로 실행
```go
func main() {
	go sexyCount("nico")
	sexyCount("flynn")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
```

동기로 처리했으면 nico 0~9 출력된 후 flynn 0~9가 출력됨

아래처럼 둘 다 go를 붙이면 아무 일도 일어나지 않음
```go
func main() {
	go sexyCount("nico")
	go sexyCount("flynn")
}
```

왜냐하면 main함수에서 실행하는 동안 go 함수가 진행되는데
둘 다 go가 붙어서 동작시킬 다른 함수가 없기 때문에 main이 바로 종료!



#### Channels

Go 에서는
```go
result := function(a, b)
```

이런식으로 함수의 return값을 변수에 저장하지 못한다.
따라서 이러한 기능이 필요할 때 channel을 사용해야함

```go
func main() {
	 c := make(chan bool)

	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}		
	
	result := <-c
	fmt.Println(result)
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}
```

main() 에서
반복문 코드까지는 순식간에 돌아감 (Goroutine)
그러나 그 밑에 코드에서 reuslt에 c라는 채널을 통해 값을 받고
이후에 print를 해야하는 코드 때문에
main함수를 바로 종료하지 않고 기다리게 되는 것

좀 더 발전된 코드를 보면
main.go

```go
func main() {
	c := make(chan bool)

	people := [2]string{"nico", "flynn"}

	for _, person := range people {
		go isSexy(person, c)
	}		

	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isSexy(person string, c chan bool) {
	switch person {
	case "nico":
		c<-true
	case "flynn":
		c <- false
	}
}
```

위 코드에서 `fmt.Println(<-c)` 코드는 c채널에 메세지가 가는 즉시 프린트 하라는 코드
Go는 smart해서 두개의 Goroutine이 돌고 있는걸 알고 있음
따라서 fmt.Println(<-c) 코드를 하나 더 적으면 받아 올 메세지가 없기 때문에 error 발생

최종 업그레이드 코드
```go
func main() {
	c := make(chan string)

	people := [5]string{"nico", "flynn", "T", "S", "R"}

	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	c <- person + " is sexy"
}
```



#### URLChecker with Goroutine

main.go
```go
package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url string
	status string
}

func main() {
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i:=0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan<- result) {
	resp, err := http.Get((url))
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	} 
	c <- result{url:url, status: status}
}
```





### 4th Project: Job Scrapper

#### Fatalln

Println() 다음에 os.Exit(1)를 호출하는 것과 동일

