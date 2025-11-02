package main

import (
	"fmt"
	"time"
)

type customer struct {
	 // creating it for undestanding struct emmbedding
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer  // struct embedding
}

func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here...

	myorder := order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myorder

}

// reciver type
func (o *order) changeStatus(status string) {
	o.status = status

}

// we can remove start from here
func (o *order) getAmmount() float32 {
	return o.amount + 1 // 50-- to 51-

}

func main() {

  newCustomer :=customer{
	name:"john",
	phone:"987654321",
  }
	testEM := order{
		id: "1",
		amount:30,
		status:"recived",
		customer: newCustomer,
	}

  fmt.Println("embeeding struct testing ",testEM)
	// inline struct if need instant only
	language := struct {
		name   string
		isGood bool
	}{"golang", true}
	fmt.Println("language struct", language)

	myorder := newOrder("1", 30.50, "recived")
	fmt.Println("myorder cnstrct", myorder)
	// myorder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "recived",
	// }

	fmt.Println("changed amount ", myorder.getAmmount())

	myorder.changeStatus("confirmed")

	fmt.Println("without time", myorder)
	myorder.createdAt = time.Now()

	fmt.Println("with time", myorder)

	fmt.Println("id is:-", myorder.id)

	/// 2nd instance
	myorder2 := order{
		id:     "2",
		amount: 100.00,
		status: "pending",
	}
	myorder2.createdAt = time.Now()
	myorder.status = "paid"
	fmt.Println("my order 2 is is:\n", myorder2)
	fmt.Println("my order  is is:\n", myorder)

}
