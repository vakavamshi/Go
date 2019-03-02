package main
import(
	"encoding/json"
        "fmt"
)

type Person struct{
 First string
 Last  string
 Items []string
}
func main() {
      p1 := Person{
             First: "vamshi",
             Last: "vaka",
             Items: []string{
                        "Items1",
                        "Items2",
                        "Items3",
             },
            }
           
      bs, err := json.Marshal(p1) 
      if err != nil{
       fmt.Println(err)
} 
  fmt.Println(p1)
  fmt.Println("------")
  fmt.Println("json",string(bs))
}
