# Interface 

interface ในภาษา go ถือว่าเป็น type ชนิดหนึ่งที่เป็นได้ทั้ง  type อะไรก็ได้และยังเป็น collection ของ signature method ซึ่งจะไม่เหมือนกับ OOP ของภาษาอื่น


## Signature method 

เป็น method ที่ถูก defind ว่าใครก็ตามอยากจะได้การกระทำแบบนี้ต้อง implement ตามนี้เท่านั้นนะ คล้ายกับภาษา OOP อื่นๆ เลยซึ่ง OOP บางภาษายินยอมให้มี body แต่ของ GO ไม่

```go
    type Walking interface{
      Walk()  // <-- signature method
    }
```

## Implement Implicitly

การ implement interface ในภาษา Go จะไม่เหมือน OOP ภาษาอื่นเพราะไม่มี keyword implement (implicit implement)  บางคนอาจจะรู้สึกขัดใจอยู่ไม่น้อย การที่ Go เป็น implicit interface ทำให้ ผู้ที่ต้องการ implement นั้น ไม่จำเป็นต้องประกาศ type ของ interface ให้วุ่นวาย (ไม่น่าจะใช่น่ะงงกว่าเดิม)


ตัวอย่างการ implement Walking interface
```go
    type Duck struct {}

    func (d *Duck) Walk() {
      fmt.Println("quack quack walk walk")
    }
```

ตัวอย่าง multiple implement 

```go
    type Running interface {
      Run() 
    }

    type Duck struct {}

    func (d *Duck) Walk() {
      fmt.Println("quack quack walk walk")
    }

    func (d *Duck) Run() {
      fmt.Println("quack quack run run")
    }
```
>  Best Practices : การที่ go เป็น implicit impliment การแตก specifies ออกเป็นส่วน ๆ ช่วยให้ง่ายต่อการ implement และลดการ implement แบบไม่จำเป็นออกไปซะ

การเขียน specification ที่ดีในภาษา go interface ควรจะเล็กเท่าที่เป็นไปได้แยกให้มันเป็นแต่ละหน้าที่ของมันเองจากนั้นให้ใช้การ composition รวมแต่ละ interface เข้ามาเป็น interface เดียวกัน

ยกตัวอย่าง 
```go
    type Duck interface {
      Run()
      Walk()
    }
```
แยกเป็น

```go
    type IDuck interface {
      Running
      Walking
    }

    type Running interface {
      Run()
    }

    type Walking interface {
      Walk()
    }
```


## Trick การตรวจสอบว่า impliment ครบตาม specifies หรือไม่

ความยากของ implicit implement คือเราไม่รู้ว่า impliment ถูกหรือป่าวจะรู้ได้ก็ต่อเมื่อ runtime ทีนี้ทำไง​ ?

มีวิธีดังนี้คือ

1. declare variable เป็น  duck มี type เป็น IDuck
   
```go
    var duck IDuck
```

2. initialize Duck struct จากนั้น assign ให้กับ duck  หากเรา impliment ไม่ครบมันจะด่าทันทีก่อนที่จะ compile อีกเพราะ linter ของ go มันรู้ ฉลาดดี  และ type interface รับเป็น pointer

```go 
  duck := &Duck{}
```

ตัวอย่าง
```go
    type IDuck interface {
      Running
      Walking
    }

    type Running interface {
      Run()
    }

    type Walking interface {
      Walk()
    }

    type Duck struct {}

    func (d *Duck) Run() {}

    function main() {
      var duck IDuck
      duck = &Duck{} // linter จะ error ทันทีว่า คุณ ขาด method Walk นะ impliment ด้วย
    }
```
