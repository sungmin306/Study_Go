package listing46go

import "fmt"

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("기간 : %d", *d)
}

func main() {
	//duration(42).pretty() // 에러 -> 메서드 집합 위배
}
