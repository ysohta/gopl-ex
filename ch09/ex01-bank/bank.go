package bank

type transaction struct {
	amount  int
	success chan bool
}

var deposits = make(chan int)
var withdraws = make(chan *transaction)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Withdraw(amount int) bool {
	success := make(chan bool)
	withdraws <- &transaction{amount, success}
	return <-success
}

func Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case tr := <-withdraws:
			if tr.amount > balance {
				tr.success <- false
				continue
			}
			balance -= tr.amount
			tr.success <- true
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}
