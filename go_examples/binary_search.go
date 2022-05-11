package main  
import "fmt" 

func binary_search(list []int, item int) int{
    low := 0
    high := len(list) - 1
    
   // mid, guess := 0, 0
    
    for low <= high {
        mid := (low + high) / 2
        guess := list[mid]
        
        if guess == item {
            return mid
        }
        if guess > item {
            high = mid - 1 
        } else {
            low = mid + 1
        }
    }
    return 0
}

func main() { 
    my_list := []int{1, 3, 5, 7, 9}
    
    fmt.Println("Start") 
    
    finded := binary_search(my_list, 3)
    fmt.Println("Result: ", finded)
    
}