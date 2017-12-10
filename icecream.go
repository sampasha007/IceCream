package main
import("fmt"
	"time"
)

c1:= make(chan int)
func clerk(cust_id int){
	
	fmt.Println("cust %d is being served" )
	c1<- cust_id

}

func manager(){
	seed(time.Now().seconds)
	decision:= rand.Intn(2)
	if decision==1{
	id:<-c1
	fmt.Println("Customer %d is selected", id)
	c1<-id
	}
	else{
	go clerk(id)
	}
}

func cashier(){
	cust<-c1
	fmt.Println("Cash collected for cost %d" , cust)
}

func main(){
	var(
	cust_id int=1
	n int
	)
for{
	fmt.Scanf(&n)
	if n==1{
	go clerk(cust_id)
	go manager()
	go wait()
	go cashier()
	cust_id+=1
	}
}
else break
}
