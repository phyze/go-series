# Goroutine & Channel


## Goroutine
ในภาษา go ผู้ใช้งานไม่สามารถเรียกใช้งาน Thread ได้โดยตรงได้ ​ซึ่งภาษา go ได้สร้าง keyword ที่ใช้งานทดแทน thread นั้นคือ
go ที่เป็น keyword built-in มาในภาษา เพื่อลดความซับซ้อนการใช้งาน เพราะการใช้งาน thread ต้องบริหารจัดการเองซึ่งค่อนข้างยากจนหลาย ๆ ครั้งเกิด
thread deadlock  "แต่ไม่ได้หมายความว่า goroutine จะไม่มี deadlock นะแค่โอกาสเกิดน้อยกว่า" 

go keyword จะถูกบริหารจัดการโดย go-runtime และ go-runtime จะเรียกใช้งาน OS threads ให้


  > NOTE : การที่จะใช้งาน go keyword ได้นั้น stement ข้างหลังจำเป็นต้องเป็น function เท่านั้น

**ตัวอย่าง การใช้ go keyword แบบ non-blocking**

```go
package main

import (
    "fmt"
    "time"
)

func hello() {
    fmt.Print("hello ")
}

func main() { 
  go hello()
  time.Sleep(time.Millisecond * 50)
  fmt.Println("world")
}

// hello world
```

ตรง go hello() เป็นการบอก go-runtime ว่าแตก thread ให้หน่อยและ execute hello function ด้วยนะ  
ส่วน time.Sleep เป็นการบอกให้ main thread ทำการ sleep ตัวเอง 50 mlisec เพื่อรอ hello() ทำการ print hello ออกมา

จากตัวอย่างด้านบน หากเจองานที่ไม่สามารถระบุเวลาเสร็จของ function ที่ใช้ร่วมกับ goroutine ได้ เราจะ block มันยังไงเพื่อไม่ให้ function main มันจบการทำงาน ?

**ตัวอย่าง go keyword แบบ blocking**

การที่ program เป็นแบบ non-blocking ทำให้ function main มันจบการทำงานทันที ดังนั้นจำเป็นต้องมีการรอให้ การทำงานที่ใช้ goroutine เสร็จก่อน ซึ่งจะใช้ package ที่ชื่อ sync

```go
package main

import (
    "fmt"
    "time"
    "sync"
)

func hello(wg *sync.WaitGroup) {
    fmt.Print("hello ")
    wg.Done()
}

func main() { 
  wg := sync.WaitGroup{}
  wg.Add(1)
  go hello(&wg)
  wg.Wait()
  fmt.Println("world")
}

// hello world
```

function main  :    

ประกาศ wg และ assign sync.WaitGroup ให้กัับ wg
จากนั้นเรียกใช้ function Add(n int) ของ WaitGroup ซึ่งตรงนี้เป็นการเพิ่มจำนวน count  ที่ต้องการจะรอให้ goroutine ทำงานเสร็จ ตัวอย่างด้านบนมี goroutine แค่ 1 ก็ใส่ลงไปใน Add(1) ซึ่งค่าเริ่มตอนคือ 0 จากนั้นส่ง pointer ของ WaitGroup เข้าไปใน hello(&wg) และสั่งให้โปรแกรม waiting จนกว่า count ของ goroutine จะเท่ากับ 0 ถ้าเท่ากับ 0 แล้วจะทำบรรทัดถัดไป ก็คือ print world ออกมา

function hello :   
 รับ pointer ของ WaitGroup เพื่อทำการลดจำนวน count ของ goroutine โดยการเรีอก Done function ซึ่งการเรียก 1 ครั้งเท่ากับ count ลบทีละ 1 และไม่สามารถติดลบได้ หากติดลบ WaitGroup จะตี panic ทันที 


---

## Channel

ปัญหามีอยู่ว่า หากโปรแกรมของเรามีการใช้ goroutine หลาย ๆ goroutine แต่ละตัวก็จะมีการทำงานที่แตกต่างกันและ goroutine เหล่านั้นต้องการที่จะแลกเปลียน  data กันละจะมีวิธีทำให้ goroutines แลกเปลียนกันได้ไหม ? 

ปัญหานี้แก้ด้วยการใช้ chan ที่ย่อมาจาก channel


channel คือที่ ๆ หนึ่งที่ให้ goroutines สามารถแลกเปลียนข้อมูลกัน    

**ยกตัวอย่าง**   
สถานีปล่อยน้ำปะปา น้ำที่ถูกปล่อยจากสถานีจะไหลไปตามท่อของกรมปะปา บ้านไหนที่ได้ต่อท่อกับทางกรมปะปาจะได้รับน้ำปะปาไป channel ก็คล้ายกัน