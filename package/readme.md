# Package

เป็น collection ของ function , structs, variable , interface ที่มีการเปิดให้ภายนอกเรียกใช้งาน แต่จะสามารถเรียกใช้ได้เฉพาะที่ยินยอมเท่านั้น (public)  Package มีไว้เพื่อแบ่ง scope ของงานออกไปเป็นส่วน ๆ ทำให้การนำไป reuse และ maintain นั้นง่ายขึ้น 

ส่วนการนำไปใข้ ชื่อ package ต้องอ้างอิงตามชื่อ directory  ถ้า directory ชื่อ duck ไพล์ที่อยู่ใน directory duck จะต้องใช้ชื่อ package ว่า duck

package ของ go จะโดน load แค่ครั้งเดียวเท่านั้นหากมี package A เรียก package B  go ยินยอมให้มีการ share address ของ package B ไปให้ package A ตามภาพด้านล่าง

![alt text](github.com/../pkg.jpg)

> NOTE : ควรระวังไว้ก็คือ การเข้าไปเปลียน state ของ package อื่นนั้นอันตรายมากอะไรที่สำคัญควรจะทำการ encapsulate ด้วย


## การสร้าง package

> NOTE : การตั้งชื่อ package ต้องขึ้นต้นด้วย lowercase เสมอและหลีกเลี่องการใช้ _ (underscore)


main file

    package main

    import (
      "github.com/go-series/package/a"
    )

    func main() {
      a.Greet()
    }

a file 

    package a

    func Greet() { fmt.Println("hello") }




หากมีการสร้าง directory ซ้อน directory ให้ยึด current path เป็นชื่อ package

      root directory
       - r1
         - r2
           - file <-- currect path is here so i used package name r2


## Alies package name

alies เป็นการตั้งชื่อ package ใหม่หลังจาก import เข้ามาหรือสามา  รถ ignore ก็ได้

### new alies

    package main

    format "fmt"

    func main() {
      format.Println("hello")
    }

### ignore package 

นี้ไม่มีไรแค่เติม _ (Underscore) หน้า package ที่ import การ ignore เป็นบอกให้ go ไม่สนใจ package ที่ import มา

    _ "fmt"
