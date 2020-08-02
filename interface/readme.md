# Interface 

interface ในภาษา go ถือว่าเป็น type ชนิดหนึ่งที่เป็นได้ทั้ง  type อะไรก็ได้และยังเป็น collection ของ signature method ซึ่งจะไม่เหมือนกับ OOP ของภาษาอื่น


## Signature method 

เป็น method ที่ถูก defind ว่าใครก็ตามอยากจะได้การกระทำแบบนี้ต้อง implement ตามนี้เท่านั้นนะ คล้ายกับภาษา OOP อื่นๆ เลยซึ่ง OOP บางภาษายินยอมให้มี body แต่ของ GO ไม่

    type Walking interface{
      Walk()  // <-- signature method
    }


## Implement Implicitly

การ implement interface ในภาษา Go จะไม่เหมือน OOP ภาษาอื่นเพราะไม่มี keyword implement (implicit implement)  บางคนอาจจะรู้สึกขัดใจอยู่ไม่น้อย การที่ Go เป็น implicit interface ทำให้ ผู้ที่ต้องการ implement นั้น ไม่จำเป็นต้องประกาศ type ของ interface ให้วุ่นวาย (ไม่น่าจะใช่น่ะงงกว่าเดิม)


ตัวอย่างการ implement Walking interface

    type Duck struct {}

    func (d *Duck) Walk() {
      fmt.Println("quack quack walk walk")
    }


ตัวอย่าง multiple implement 


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

>  Best Practices : การที่ go เป็น implicit impliment การแตก specifies ออกเป็นส่วน ๆ ทำให้ง่ายต่อการ implement และลดการ implement แบบไม่จำเป็นออกไปซะ

การเขียน specification ที่ดีในภาษา go interface ควรจะเล็กเท่าที่เป็นไปได้แยกให้มันเป็นแต่ละหน้าที่ของมันเองจากนั้นให้ใช้การ composition รวมแต่ละ interface เข้ามาเป็น interface เดียวกัน

ยกตัวอย่าง 

    type FileSystem interface {
      Read() ([]byte,error)
      Write([]byte) error
      Close() error
    }

แยกเป็น

    type Reader interface {
      Read() ([]byte,error)
    }

    type Writer interface {
      Write([]byte) error 
    }

    type Closer interface {
      Close() error
    }

    type FileSystemComposition interface {
      Reader
      Writer
      Closer
    }