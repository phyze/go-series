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

ปัญหามีอยู่ว่า หากโปรแกรมของเรามีการใช้ goroutine หลาย ๆ goroutine แต่ละตัวก็จะมีการทำงานที่แตกต่างกันและ goroutine เหล่านั้นต้องการที่จะแลกเปลียน  data กันแล้วจะมีวิธีทำให้ goroutines แลกเปลียนข้อมูลกันได้ไหม ? 

ปัญหานี้แก้ด้วยการใช้ chan ที่ย่อมาจาก channel

channel คือที่ ๆ หนึ่งที่ให้ goroutines สามารถแลกเปลียนข้อมูลกัน นึกถึงนาย A ใช้มือถือโทรหา  นาย B เมื่อนาย B รับสายทั้งนาย A และ B สามารถพูดคุยสื่อสารกันได้ทั้งสองฝ่ายสามารถได้ยินเสียงของกันและกัน ซึ่งเหมือนกับความสามารถของ Channel ในภาษา go 

โดย go มี slogan ว่า

> Do not communicate by sharing memory; instead, share memory by communicating.

หมายความว่า อย่าติดต่อสือสารกันจากการใช้ sharing memory แต่ให้ใช้การ share memory จากการสือสารแทน

เรามาแกะประโยคด้านบนกันดีกว่าว่ามันหมายถึงอะไร  

ในภาษา java, c++, python และหลาย ๆ ภาษาที่มีกลไกในการใช้ share memory เพื่อแลกเปลียนข้อมูลกันระหว่าง processes สิ่งที่น่ากลัวที่สุดของกลไกนี้คือ deadlock จากการ read/write resource ที่ ๆ เดียวกันจึงเกิดขึ้นง่ายมากผู้ใช้งานจำเป็นต้องมีความเชี่ยวชาญเพราะการ manage process ค่อนข้างยากส่วนมากจะพบการ deadlock ใน concurrency programming ทีนี้ผู้พัฒนาภาษา go จึงใช้กลไก channel เพื่อใช้ในการแลกเปลียนข้อมูลกันผ่านช่องทางการสือสาร (share memory by communicating) แทนการ communicate by sharing memory 

![](github.com/../sharingMemo.png)


channel แบ่งออกเป็น 2 ชนิด คือ
1. UnBuffered 
2. Buffered
   
Unbuffered 


คุณสมบัติของการสื่อสารแบ่งได้ออกเป็น 2 ชนิด 

1. Two way communication
2. One way communication


**Two way communication**    
ชนิดนี้สามารถส่งและรับใน channel เดียวกัน   

**example**   
```go 
package main

import (
	"fmt"
	"sync"
)

type twoWay struct {
	hello string
	world string
}

func hello(pipe chan *twoWay, wg *sync.WaitGroup) {
	defer wg.Done()
	v := <- pipe
	v.hello = "hello"
	pipe <- v
}

func world(pipe chan *twoWay, wg *sync.WaitGroup){
	defer wg.Done()
	v := <- pipe
	v.world = "world"
	pipe <- v
}

func main() {
	wg := &sync.WaitGroup{}
	pipe := make(chan *twoWay,1)

	//send twoWay for init
	pipe <- &twoWay{
		hello: "",
		world: "",
	}

	wg.Add(2)
	go hello(pipe,wg)
	go world(pipe,wg)
	wg.Wait()
	fmt.Println(<- pipe)
}

```

จากตัวอย่างด้านบน เห็นว่ามีการแบ่ง hello และ world ที่เป็น function ทำหน้าที่ assign "hello" และ "world" เข้าไปใน chan *twoWay   
กลับมาที่ main  


One way communication  สามารถใช้ได้แค่ รับ หรือ ส่ง เท่านั้นและต้องมีการ define ที่ชัดเจน โดยใช้เครื่องหมาย <- 