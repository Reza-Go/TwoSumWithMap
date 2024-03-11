package main 


import "fmt"


func main (){
      
     nums := []int { 2 , 3 , 5 , 8}
	 target := 10 
	 
	 result := twoSum (nums , target)
     fmt.Printf("result is:%v\n" , result)


}

func twoSum (nums []int , target int) []int {

	 keys := map[int]int{}
	  for i,item := range nums {
           complement := target - item 
		   
		   if _ , ok := keys[complement]; ok {
                  return []int{keys[complement], i}
		   }
		   keys[item] = i 
		   
	  }
	  return []int{0,0}
}