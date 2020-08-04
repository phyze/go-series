# Testing

testing เป็น package ของ go ที่ built-in มาให้พร้อมกับภาษา การใช้งานค่อนข้างง่าย

## Using testing.T

testing คือชื่อ package ที่ built-in กับภาษา ส่วน T หมายถึงเรียกใช้ Test ส่วนการตั้งชื่อ function test ก็แล้วแต่วัฒนธรรมของทีม แต่จำไวว่า ชื่อจำเป็นต้องขึ้นด้วย Test เสมอ และการตั้งชื่อควรสื่อถึงผลลัพธ์ที่ต้องการให้ชัดเจน

ตัวอย่าง การตั้งชื่อแบบ snake case และ camelcase

    func plus(a,b int) int { return a+b }

**snake case**

    func Test_plus_func_must_eq_expected_must_pass(t *testing.T) { 
      expected := 5
      actual := plus(3,2)
      if actual != expected {
		    t.Errorf("expected  %d but actual  %d", expected, actual)
	    }
    }

**camel case**

    func TestPlusFuncMustEqExpectedMustPass(t *testing.T) { 
      expected := 5
      actual := plus(3,2)
      if actual != expected {
		    t.Errorf("expected  %d but actual  %d", expected, actual)
	    }
    }

## Using testing.B



> NOTE : การเขียน test ควรจะอ่านและเข้าใจทันทีว่า test case นี้ต้องการอะไรและเห็นปุ่มต้องอ่านเข้าใจทันที 

## การใช้งาน go test

 go test  แบบ ระบุ file

    go run filename.go

go test แบบ recursive 

    go run ./...