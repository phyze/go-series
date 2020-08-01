# Struct 

หากใครเคยเขียน java จะมี class ให้ซึ่งทาง go ก็มีคล้าย ๆ กันที่เรียกว่า struct และรองรับแนวคิด oop ถึงแม้ว่าจะรองรับ oop แต่ก็ไม่ใช่ oop แท้ซะทีเดียว

> NOTE :  ในภาษา go  มีแต่ pass by value และ pass by pointer เท่านั้น

## Access modifier
ในภาษา java มี 4 ชนิด private , protectd , public , default แต่ในภาษา go นั้นมีแค่ public และ private เท่านั้นเพื่อลดความซับซ้อน 

java

    class Box {
      public int width;
      public int hight;
      private int x;
      private int y;
    }

go 

    type Box struct {
      Width int
      Hight int
      x int
      y int
    }

สังเกตว่าภาษา Go ไม่มี keywords private public แต่ใช้ตัวเล็กตัวใหญ่ขึ้นต้น 

• การขึ้นต้นด้วยตัวเล็กเป็น private 

• การขึ้นต้นด้วยตัวใหญ่เป็น public

## Behavior of Object

method ในภาษา Go นั้นจะไม่อยู่ใน  struct เหมือนที่ method อยู่ใน class ของ java 

java 

    class Cat {
      public void run() { System.out.println("run...") }
    }

go 

    type Cat struct { }

    func (c Cat) Run() { fmt.Println("run...") }


การที่ go จะรู้ว่า function run นั้นเป็นของ Cat ได้นั้นจำเป็นต้องมี  (c Cat) เราเรียกสิ่งนี้ว่า receiver 

##  Struct initialization

การ initialze ใน go มี 3 แบบดังนี้

1. define variable และถามด้วย Type
  
        var cat Cat
        cat.Run()

2. assign Cat type แบบ dynamic typing

        cat1 := Cat{}
        cat1.Run()

3. ใช้ keyword new ที่ build-in มากับภาษาเพื่อทำการ initialize 
   
        cat2 := new(Cat)
        cat2.Run()

    ข้อแต่งต่างของการใช้ new  คือเป็นการ initialize Cat  พร้อมกับ return pointer ของ struct Cat  

  