# Encapsulation

> หากมีการใช้คำว่า member นั้นหมยถึง type ทุดชนิดที่อยู่ใต้ package

เป็นการห่อหุ่ม member ไว้ใน package ซึ่งใช้กลไกล public และ private เป็นตัวควบคุมว่า member ไหนที่จะให้ภายนอกเรียกใช้งานได้และ member ไหนที่ต้องซ้อนไว้


## Exported 
export คือ การยินยอมให้ภายนอกเรียกใช้งาน member ใน package 

ซึ่งในตัวอย่างนี้ จะทำการ export ตัวแปรที่ชื่อ Value 

coinPub file 

```go
    package coinPub

    var Value int = 5
```

main file 
```go
    package main

    func main() {
    	fmt.Println("before value change : ", coinPub.Value)
    	coin.Value = 6
	    fmt.Println("after value changed : ", coinPub.Value)
    }
    // before value change :  5
    // after value changed :  6
```

สังเกตุว่าค่า package main สามารถเรียก A ที่อยู่ภายใต้ package a ได้เพราะว่า a ทำการ export  A เป็น public ทำให้คนอื่นสามารถเรียกใช้ได้
แต่ก็มีความเสี่ยงหากคนที่เรียกมีการ assign ค่าเข้าไปใหม่จะทำให้ state ก่อนหน้าที่เป็น 5 โดนเปลียนเป็น 6 

การทำให้ member นั้นให้กลายเป็น public ควรพิจรณาก่อนว่าทำไปเพื่ออะไรทำไมต้อง public ถ้าทุก member เป็น public หมดจะทำให้ state หรือ process ทำงานผิดเพี้ยนไปจากสิ่งที่ควรจะเป็น และเป็นช่องว่างในการโจมตีโปรแกรมของเราได้
  
## Unexported
unexport เป็นการปิดการเข้าถึง member ใน package 

ซึ่งในตัวอย่างนี้ จะทำการ unexport ตัวแปรที่ชื่อ value 

coinPrivate file

```go
    package coinPrivate

    var value int = 5
    var valueVersion int 

    func GetValue() int {
      return value
    }

    func SetValue(n int) {
      value = n
      valueVersion = valueVersion + 1
    }

    func Version() int {
      return valueVersion
    }
    // before  value change :  5
    // after value change :  81
    // version had change :  2
```

value โดนกำหนดให้เป็น private ทำให้ภ่ายนอกไม่สามารถ access ตัวแปร value ได้ ดังนั้นจำเป็นต้องสร้าง method ที่ทำหน้าที่ GetValue พวกเราเรียกสิ่วนี้ว่า  getter

หากต้องการที่จะเปลียนแปลงค่าก็ควรทำผ่าน SetValue พวกเราเรียกว่า Setter

ส่วนการที่จะต้องทำ getter setter ต้องมีเหตุผลว่าทำไมต้องทำไม่อย่างงั้นมันก็ไม่อย่างอะไรกับการที่ ประกาศ variable แบบ public 

จากตัวอย่างด้านบน ไม่ใช่แค่ change state value เฉยๆ แต่มีการ increse  valueVersion ที่ละ 1 เมื่อ value มีการเปลียนแปลงเพื่อตรวจสอบว่า value มีการ change ไปกี่ครั้งแล้ว

main file 

```go
    package main 

    func main() {
        fmt.Println("before  value change : ", coinPrivate.GetValue())
        coinPrivate.SetValue(8)
        coinPrivate.SetValue(81)
        fmt.Println("after value change : ", coinPrivate.GetValue())
        fmt.Println("version had change : ", coinPrivate.Version())
    }

```