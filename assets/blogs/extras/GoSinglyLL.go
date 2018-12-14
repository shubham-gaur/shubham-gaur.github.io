package main
import "fmt"

type ele struct {
    val int
    next *ele    
}

type singlyLL struct{
    len int
    head *ele
}

func initList() *singlyLL{
    return &singlyLL{}
}

func (s *singlyLL) AddFront(val int){
    ele := &ele{
        val: val,
        next: nil,
    }
    if s.head == nil {
        s.head = ele
    } else {
        ele.next = s.head
        s.head = ele
    }
    s.len++
}

func (s *singlyLL) RemoveFront() error{
    if s.head == nil {
        return fmt.Errorf("Remove Error: List is empty")
    } else {
        s.head = s.head.next
        s.len--
        return nil
    }
    
}

func (s *singlyLL) AddBack(val int){
    ele := &ele{
        val: val,
        next: nil,
    }
    if s.head == nil {
        s.head = ele
    } else {
        curr := s.head
        for curr != nil {
            if curr.next == nil {
                curr.next = ele
                break       
            }
            curr = curr.next
        }
    }
    s.len++
}

func (s *singlyLL) RemoveBack() error{
    if s.head == nil {
        return fmt.Errorf("Remove Error: List is empty")
    } else {
        prev := s.head
        curr := s.head
        if prev.next == nil {
            s.head = nil
            s.len--
            return nil
        }
        for curr != nil {
            curr = curr.next
            if curr.next == nil {
                prev.next =  nil
                s.len--
                return nil
            }
            prev = curr
        }
        return fmt.Errorf("Remove Error: Something went wrong")
    }   
}

func (s *singlyLL) RemoveDuplicates() error{
    if s.head == nil {
        return fmt.Errorf("Remove Error: List is empty")
    } else {
        prev := s.head
        curr := s.head
        if curr.next == nil {
            return nil
        }
        for prev != nil {
            for curr.next != nil {
                curr = curr.next
                if curr.val == prev.val {
                    prev.next =  curr.next
                    s.len--
                }
            }
            if prev == nil {
                return nil
            }
            return nil
        }
        
        return fmt.Errorf("Remove Error: Something went wrong")
    }   
}

func (s *singlyLL) Length()int{
    return s.len
}

func (s *singlyLL) Traverse() error{
    if s.head == nil {
        return fmt.Errorf("Tranverse Error: List is empty")
    } else {
        curr := s.head
        for curr != nil {
            fmt.Printf("%v ->", curr.val)
            curr = curr.next
        }
        fmt.Println()
        return nil
    }
}

func main(){
    singleList := initList()
 
    fmt.Printf("AddFront: 1\n")
    singleList.AddFront(1)
    fmt.Printf("Size: %v\n", singleList.Length())
 
    fmt.Printf("AddFront: 2\n")
    singleList.AddFront(2)
    fmt.Printf("AddFront: 2\n")
    singleList.AddFront(2)
    fmt.Printf("AddFront: 2\n")
    singleList.AddFront(2)
    fmt.Printf("AddFront: 2\n")
    singleList.AddFront(2)

    fmt.Printf("Size: %v\n", singleList.Length())

    fmt.Printf("AddBack: 3\n")
    singleList.AddBack(3)
    fmt.Printf("Size: %v\n", singleList.Length())
 
    fmt.Printf("AddBack: 4\n")
    singleList.AddBack(4)
    fmt.Printf("Size: %v\n", singleList.Length())
    
    err := singleList.Traverse()
    if err != nil {
        fmt.Println(err.Error())
    }
    
    singleList.RemoveDuplicates()
    fmt.Printf("Duplicates removed \n")
    fmt.Printf("Size: %v\n", singleList.Length())
    err = singleList.Traverse()
    if err != nil {
        fmt.Println(err.Error())
    }
    
    singleList.RemoveBack()
    fmt.Printf("Last element removed \n")
    fmt.Printf("Size: %v\n", singleList.Length())
    err = singleList.Traverse()
    if err != nil {
        fmt.Println(err.Error())
    }
    
    singleList.RemoveBack()
    fmt.Printf("Last element removed \n")
    fmt.Printf("Size: %v\n", singleList.Length())
    err = singleList.Traverse()
    if err != nil {
        fmt.Println(err.Error())
    }
    
    singleList.RemoveBack()
    fmt.Printf("Last element removed \n")
    fmt.Printf("Size: %v\n", singleList.Length())
    err = singleList.Traverse()
    if err != nil {
        fmt.Println(err.Error())
    }
    
    singleList.RemoveBack()
    fmt.Printf("Last element removed \n")
    fmt.Printf("Size: %v\n", singleList.Length())
    err = singleList.Traverse()
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Println()
}
