# Variable and type

## Type

type ของ ภาษา GO มีค่อนข้างเยอะให้เลือกใช้มากมาย ซึ่งตามด้านล่างเป็นเพียงแค่บางส่วนที่นิยมใช้งานกัน

```go 
type bool
type byte
type complex128
type complex64
type error
type float32
type float64
type int
type int16
type int32
type int64
type int8
type rune
type string
type uint
type uint16
type uint32
type uint64
type uint8
type uintptr
```


## Variable


```go
func main() {


  //static assign 
  var age int 
  age = 15
  fmt.Println("age : ",age)

  //dynamic assign
  myName := "soyod"
  fmt.Println("my name is : ", myName)
}

```

## Static assign 
คือ variable ที่มีการประกาศ data type ที่มีความชัดเจนตั้งแต่การ define ตัวแปร
อย่างเช่น age มีการประกาศ data type ชนิด int นั้นหมายความว่า ตัวแปร age จะรับค่าได้เฉพาะตัวเลขที่เป็นจำนวนเต็มเท่านั้น

## Dynamic assign 
คือ variable ที่ไม่ต้องประกาศ data type เมื่อมีการ assign ค่าเกิดขึ้น ตัวภาษาจะสร้าง data type ให้เอง 