# Function 

function ทั่วไปคงรู้กันอยู่แล้วแต่ผมจะพาทัว anonymous function กับ 1st class และ higher order function ซึ่งใน go ทำได้

## 1. Anonymous function 

```go    
    package main

    func main {
        func() {
          fmt.Println("i'm anonymous function")
        }

    }
```

เป็นการประกาศ function แบบไม่มีชื่อหรือความหมายในการเรียกใช้งาน เป็นการสร้าง function ขึ้นมาแบบใช้แล้วทิ้ง function ลักษณะแบบนี้สามารถเข้าถึงหรือเปลียนแปลง state หรือ ค่า ที่อยู่นอก function ได้ 

## 2. 1st class function
first class function คือการประกาศ function ในรูปแบบ anonymous function แต่สามารถ assign function ให้กับตัวแปรได้

### 2.1.   1st class function
```go

func main(){
    f := func()  {
      fmt.Println("i'm 1st class function")
    }
    
    f()
}
```


### 2.2. pass arg to 1st class function and return resutl

```go

func main(){
    result := func(a,b int) int {
      return a+b
    }(2,3)
    
    //or 

    plus := func(a,b int) int {
      return a+b
    }

    fmt.Println(result)
    fmt.Println(plus(3,2))

    // 5
}
```



## 3. higher order function 

Higher order function หรือ  HOF เป็น function ที่มีการทำงานร่วมกับ function อื่นการที่จะเป็น HOF ได้นั้นจำเป็นต้องส่ง function เข้ามาหรือ return เป็น function 

``` go
  package main
  
  func  pick(alphabets []string) func(pick string) []string { ... } 

  func main() {
      alphabets := []string{"a","a","b","r","t","o","f"}
      picked := pick(alphabets)("a")
      fmt.Println(picked)
  
      // ["a","a"]
  }

```