# Polymorphism

ที่มาของ polymorphism มาจากรากศัพท์ของ geek

poly แปลว่า  มาก , จำนวนมาก   
morph แปลว่า เปลียนแปลง , รูปแบบ 

polymorphism แปลว่า สภาวะหนึ่ง ๆ ที่มีได้หลายรูปแบบหากนำมาใช้กับ oop จะมีความหมายว่า  interface ที่มีการ define ให้มี พฤติกรรมทำบางสิ่งบางอย่าง แล้วมีการนำไป implement ด้วย type อื่นโดยที่ interface ไม่รู้ว่าคนที่นำไป implement นั้นเป็น type อะไรแต่สามารถเรียกใช้ พฤติกรรมของ type ที่ implement ได้ตามที่มีการ define ไว้ในรูปแบบ signature method

ตัวอย่าง การวาดรูป

```go

type Shape interface {
  Draw()  //signature method
}
```
define type Shape ที่เป็น interface เพื่อไว้สำหรับการ implement และมี method Draw สำหรับวาดรูป

```go
type Rectangle struct {
  Width int
  Height int
}

func (r *Rectangle) Draw() {
  fmt.Println("Rectangle")
}
```

Rectangle ทำการ implement ตาม specification ของ Shape 

```go
type Circle() {
  Radiant int
}

func (c *Circle) Draw() {
  fmt.Println("Circle")
}
```

Circle ทำการ implement ตาม specification ของ Shape

```go
func main() {
  var rectangle Shape
  rectangle = &Rectangle { 
    Width: 2,
    Height: 4,
  }
  rectangle.Draw() //Rectangle

  var circle Shape 
  circle = &Circle {
    Radiant: 5,
  }
  circle.Draw()  //Circle

}
```

สังเกตว่า  variable rectangle และ circle ประกาศเป็น type Shape ซึ่งถูก assign ด้วย type struct  rectangle ก็รับ type Rectangle และ 
circle ถูก assign ด้วย type Circle ทีนี้สังเกตว่าทั้ง rectangle และ circle เรียก method Draw() ถึงชื่อเหมือนกันแต่การทำงานต่างกัน

