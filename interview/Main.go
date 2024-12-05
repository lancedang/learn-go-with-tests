package interview

import "fmt"

func main() {
	arr := [2]int{1, 2}
	arrnew := []*int{}
	for _, v := range arr {
		arrnew = append(arrnew, &v)
	}

	fmt.Println(*arrnew[0], *arrnew[1])
}
