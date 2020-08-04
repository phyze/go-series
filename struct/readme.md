# Struct 

หากใครเคยเขียน java จะมี class ให้ซึ่งทาง go ก็มีคล้าย ๆ กันที่เรียกว่า struct และรองรับแนวคิด oop ถึงแม้ว่าจะรองรับ oop แต่ก็ไม่ใช่ oop แท้ซะทีเดียว

> NOTE :  ในภาษา go  มีแต่ pass by value และ pass by pointer เท่านั้น


##  Struct initialization

การ initialze ใน go มี 3 แบบดังนี้

1. define variable และตามด้วย Type
```go
  var cat Cat
  cat.Run()
```
2. assign Cat type แบบ dynamic typing

```go
  cat1 := Cat{}
  cat1.Run()
```

3. ใช้ keyword new ที่ build-in มากับภาษาเพื่อนำมาใช้ initialize 
  
  ข้อแต่งต่างของการใช้ new  คือเป็นการ initialize Cat  พร้อมกับ return pointer ของ struct Cat ซึ่งแตกต่างกับวิธีที่ 1 และ 2 ที่ return แค่ value เท่านั้น
```go   
  cat2 := new(Cat)
  cat2.Run()
```

  

## Access modifier
ในภาษา java มี 4 ชนิด private , protectd , public , default แต่ในภาษา go นั้นมีแค่ public และ private เท่านั้นเพื่อลดความซับซ้อน 

java
```java
    class Box {
      public int width;
      public int hight;
      private int x;
      private int y;
    }
```
go 
```go
    type Box struct {
      Width int
      Hight int
      x int
      y int
    }
```
สังเกตว่าภาษา Go ไม่มี keywords private public แต่ใช้ตัวเล็กตัวใหญ่ขึ้นต้น 

• การขึ้นต้นด้วยตัวเล็กเป็น private 

• การขึ้นต้นด้วยตัวใหญ่เป็น public

## Behavior of Object

method ในภาษา Go นั้นจะไม่อยู่ใน  struct เหมือนที่ method อยู่ใน class ของ java 

### non-pointer receiver

การที่ go จะรู้ว่า function run นั้นเป็นของ Cat ได้นั้นจำเป็นต้องมี  (c Cat) เราเรียกสิ่งนี้ว่า receiver แบบไม่มี pointer 
```go
    type Cat struct {
        Name string
     }

    func (c Cat) Run() { 
      fmt.Println(c.Name,"is runing.") 
      c.Name = "alter"
    }
```
การที่เป็น non-pointer receiver นั้นจะทำให้ properties ของ struct ไม่ถูกเปลียนแปลง (Immutable) หาก function ต่าง ๆ มีการเรียกใช้ properties ของ struct ขึ้นมาและได้แก้ไข้ state เกิดขึ้นก่อน function จะจบการทำงานลง state ของ properties จะไม่โดนเปลียน ทำให้ไม่กระทบกับการทำงานของ function อื่น (free side-effect) ซึ่งทำให้
function เหล่านั้นกลายเป็น pure function นั้นเอง

ตัวอย่าง 
```go
    func main() {
      cat := Cat{
        Name: "nano"
      }
      cat.Run() //nano...run
      cat.Run() //nano...run
    }
```

ภายใน function Run ได้มีการ set ชื่อใหม่ให้กับ cat ก่อนจะจบการทำงาน จากนั้นเรียก function Run อีกรอบ สิ่งที่ได้คือชื่อไม่ถูกเปลียนแต่ออกมาเป็นชื่อเดิม

### pointer receiver

หากต้องการเปลียนแปลง state ของ propterties ภายใน struct จำเป็นต้องใช้ pointer เพื่อเข้าถึง address ของค่าเหล่านั้น การที่เป็น pointer receriver เมื่อมีการเรียกใช้งาน function  ของ struct นั้น ๆ จะชี้ไปที่ address เดียวกันเสมอทำให้สามารถเปลียนแปลงค่าภายใน struct ได้นั้นเองและถ้า struct มีขนาดใหญ่การใช้ pointer receiver จะทำให้ประหยัด memory ได้มาก
```go
    type Cat struct {
        Name string
     }

    func (c *Cat) Run() { 
      fmt.Println(c.Name,"is runing.") 
      c.Name = "alter"
    }
```
ตัวอย่าง
```go
    func main() {
        cat := Cat{
          Name: "nano"
        }
        cat.Run() //nano ...run
        cat.Run() //alter ...run
    }
```
