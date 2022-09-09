package main

func main() {
	/*format := "15:04:05"
	  parse, err := time.Parse(format, "16:04:05")
	  if err != nil {
	      println(err.Error())
	  } else {
	      println(parse.Format(time.RFC3339))
	  }

	  now := time.Now()
	  println(now.Format(format))*/

	a := []int{0}
	changeSlice(a)
	println(a[0])
}

func changeSlice(arr []int) {
	arr[0] = 1
}
