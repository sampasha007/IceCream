package main
import("fmt"
	   "time"
	   "math/rand"
)

 var c1= make(chan int)
 var c2= make(chan int)
 var c3= make(chan int)

func clerk(cust_id int){
	
	fmt.Printf("\ncust %d is being served",cust_id)
	c1<- cust_id
	time.Sleep(time.Millisecond * 200)

}

func manager(){

	rand.Seed(time.Now().UTC().UnixNano())
	decision:= rand.Intn(2)
	id:= <-c1
	//fmt.Println(id)
	if decision==1{
	
	fmt.Printf("\nCustomer %d is selected", id)
	c2<-id
	} else{
	fmt.Printf("\nCustomer %d is Rejected  \nTrying again.........   ", id)	
	time.Sleep(time.Millisecond * 100)

	go clerk(id)
	go manager()
	go wait()
	go cashier()
	}
}
func wait(){
    id:=<-c2
	time.Sleep(time.Second * 2)
	c3<-id
     
}

func cashier(){
	cust:=<-c3
	fmt.Printf("\nCash collected for cost %d" , cust)
	time.Sleep(time.Millisecond * 100)

	
}

func main(){
	var( 
	cust_id int=1
	n int
	)

	fmt.Println("Welcome TO the ICECREAM SIMULATOR\n=======================\n\n")
for{
	fmt.Scanln(&n)
	if n==1{
	go clerk(cust_id)
	go manager()
	go wait()
	go cashier()
	cust_id+=1
	}else {
		break}
}
}
