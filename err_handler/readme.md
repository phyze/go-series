# error handler 

ในภาษ GO นั้นการที่จะดัก error ได้จำเป็นต้องใช้ if ในการเช็ค ซึ่งค่อนข้างจะไม่ดูแปลกๆ
ในหลาย ๆ คนที่พึ่งจะเรียนรู้ภาษา GO 

ในภาษา อื่น ๆ เช่น java, C#, js, python ภาษาเหล่านี้มีการดัก error ด้วยวิธีการที่เรียกว่า Try - catch
แต่ของ GO นั้นไม่มี มาดูกันว่า GO ดัก Error กันยังไง

```go

// error handler
func validatePhoneNumber(x string) error {
	if len(x) != 10 {
		return errors.New("Phone number invalid.")
	}
	return nil
}


func main() {

	//error handler
	phoneNumber := "082231111"
	if err := validatePhoneNumber(phoneNumber); err != nil {
		fmt.Println(err.Error())
	}

	err := validatePhoneNumber(phoneNumber)
	if err != nil {
		fmt.Println(err.Error())
	}
}
```


จะเห็นว่า มีการ return error ออกมาจาก function validatePhoneNumber ซึ่งนี้เป็นวิธีการดัก error ของ GO ในหลักการออกแบบของภาษา ผู้พัฒนาภาษา GO ต้องการให้ นักพัฒนาดักจับ error ก่อนที่จะมี error เกิดขึ้น เพื่อลดความ complex ในการพัฒนา 

NOTE : สิ่งที่ไม่ควรทำคือการ ignore error จะทำให้ program ที่เขียนเกิด bugs ขึ้น และจะกลายเป็น technical dept ในที่สุด