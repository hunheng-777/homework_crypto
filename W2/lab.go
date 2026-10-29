package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)
 var A int = 2 
 var B int =4
 //lab1
func lab01() {
	var c ,d int 
	fmt.Println("input num c")
	fmt.Scan(&c)
	fmt.Println("input num d")
	fmt.Scan(&d)
	x := c
	fmt.Println("%v",x)
	x +=d
	fmt.Println("%v",x)
	x -=d
	fmt.Println("%v",x)
	x *=d
	fmt.Println("%v",x)
	if c !=0{
		x /=c
	}else{
		fmt.Println("cint can't devide by 0")
	}
	
	fmt.Println("%v",x)

}
//lab2
func lab02() {
	var a ,b int 
	var y bool = false
	fmt.Println("input num c")
	fmt.Scan(&a)
	fmt.Println("input num d")
	fmt.Scan(&b)
	if a >0 && b > 0{
		y = true
		fmt.Println(y)
	}else{
		y = false 
		fmt.Println(y)
	}
	fmt.Println(a>b || b>a)
	fmt.Println(a != b)
}
//lab3
func myAND()  {a:=2 ;b:=4;result :=  a&b ;fmt.Println( "%b",result)}
func myOR()  {a:=2 ;b:=4;result :=  a|b ;fmt.Println( "%b",result)}
func myXOR()  {a:=2 ;b:=4;result :=  a^b ;fmt.Println( "%b",result)}
func myLEFT_SHIFT()  {a:=2 ;b:=4;result :=  a<<b ;fmt.Println( "%b",result)}
func myRIGHT_SHIFT()  {a:=2 ;b:=4;result :=  a>>b ;fmt.Println( "%b",result)}


func lab03() {
	myAND()
	myOR()
	myXOR()
	myLEFT_SHIFT()
	myRIGHT_SHIFT()
}

//lab4

func lab04(){
	var options int
	var a ,b int 
	
fmt.Println("===== Mini Calculator =====")
fmt.Println(" 1.Add\n 2.Sub\n 3.Mul\n 4.Div\n 5.Mod\n 6.Exit\n PLease chose options")
fmt.Scan( &options)
if options ==6 {
		return
	}
fmt.Println("input num c")
	fmt.Scan(&a)
	fmt.Println("input num d")
	fmt.Scan(&b)
	
switch options{
case 1 : x:=a+b;fmt.Println("result of Add",x)
case 2 : x:=a-b;fmt.Println("result of Sub",x)
case 3 : x:=a*b;fmt.Println("result of Mul",x)
case 4 : if b == 0 {
			fmt.Println("Error: Cannot divide by zero!")
		} else {
			x := a / b
			fmt.Println("result of Div:", x)
		}
case 5 :if b == 0 {
			fmt.Println("Error: Cannot divide by zero!")
		} else {
			x := a % b
			fmt.Println("result of Mod:", x)
		}
}
}
//lab05
func lab05(){
var text string
fmt.Println("please input text:")
fmt.Scan(&text)
for i:=0; i< len(text);i++{
fmt.Printf("convert to binary: %08b",text[i])
}
hexResult := hex.EncodeToString([]byte(text))
fmt.Println("convert to Hexadecimal:", hexResult, "\n")
base64Result := base64.StdEncoding.EncodeToString([]byte(text))
fmt.Println("convert to Base64:", base64Result, "\n")

}

func xorEncrypt(text string, key byte) string {
    result := make([]byte, len(text))

    for i := 0; i < len(text); i++ {
        result[i] = text[i] ^ key
    }

    return string(result)
}

func lab06() {
    var plaintext string
    var keyChar byte

    fmt.Print("Enter plaintext: ")
    fmt.Scanln(&plaintext)

    fmt.Print("Enter a single-character key: ")
    fmt.Scanf("%c", &keyChar)

    // Encrypt the message
    ciphertext := xorEncrypt(plaintext, keyChar)

    fmt.Printf("\nCiphertext (hex): %x\n", []byte(ciphertext))

    // Decrypt the message using the same function and key
    decryptedText := xorEncrypt(ciphertext, keyChar)

    fmt.Println("Decrypted message:", decryptedText)
}
func main (){
	// lab01()
	// lab02()
	// lab03()
	//  lab04()
	//  lab05()
	 lab06()

}
