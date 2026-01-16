package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	
)

// en person i kön
type Nod struct{

	sync.Mutex 
	värde string 
	föreg *Nod
	nästa *Nod

	
}

// dubbelkön = deque
type Dubbelkö struct{

	head *Nod // den pekra på den "falska" starteen
	tail *Nod // slutet
	antal atomic.Int64



	
}

//NyKö skapar en tom ny kö med sentinelnoder som head och tail och detta görs så att man slipper ändra k.head or tail senare
func NyKö() *Dubbelkö{

	h := &Nod{värde: "HEAD_SENTINEL"}
	t := &Nod{värde: "TAIL_SENTINEL"}
	h.nästa =t
	t.föreg =h

	return &Dubbelkö{

		head: h,
		tail: t,
		antal: atomic.Int64{},
	}
	
}

// LäggFram lägger till perosn lägst fram ikön, där k = kö(deque) och v = värde(strängen man lägger in i könn) och låser bara head och den första riktiga  
func (k *Dubbelkö) LäggFram(v string){

	nyNod := &Nod{värde: v}

	

	k.head.Lock()
	// lås nästa nod, allså den som var först innan för att säkra
	first := k.head.nästa
	first.Lock()
	nyNod.nästa =first 
	nyNod.föreg = k.head
	first.föreg =nyNod 
	k.head.nästa = nyNod
	first.Unlock() // lås i ordning
	k.head.Unlock()
	k.antal.Add(1)








	
}

// LäggBak lägger till person längst bak i kön
func (k *Dubbelkö) LäggBak(v string){

	nyNod := &Nod{värde:v}
	k.tail.Lock()
	//lås den som var sist innan
	last := k.tail.föreg
	last.Lock()
	//mellan last och tail
	nyNod.nästa = k.tail
	nyNod.föreg = last 
	last.nästa =nyNod
	k.tail.föreg = nyNod
	last.Unlock()
	k.tail.Unlock()
	k.antal.Add(1)
	
}

//TaFram tar bort första personn i kön
func (k *Dubbelkö) TaFram() (string, bool){

	//hand over hand lås
	k.head.Lock()
	first := k.head.nästa
	first.Lock()
	if first == k.tail{

		first.Unlock() // låss up tail
		k.head.Unlock() // head
		return "",false
	}

	next := first.nästa
	next.Lock()
	värde :=first.värde
	k.head.nästa =next
	next.föreg = k.head
	//släpp låsen
	next.Unlock()
	first.Unlock()
	k.head.Unlock()

	k.antal.Add(-1)
	return värde,true


	
}

//TaBak tar bort siste personen i klön
func (k *Dubbelkö) TaBak() (string, bool){

	k.tail.Lock()
	last := k.tail.föreg
	last.Lock()
	if last ==k.head{

		last.Unlock()
		k.tail.Unlock()
		return "",false
	}

	//för att länka om måste noden låsas innan den susta
	prev := last.föreg
	prev.Lock()
	värde := last.värde
	k.tail.föreg = prev
	prev.nästa= k.tail
	prev.Unlock()
	last.Unlock()
	k.tail.Unlock()
	k.antal.Add(-1)
	return värde, true

	
}

// Antal -> antalet personener i kön
func (k *Dubbelkö) Antal() int{

	
	return int(k.antal.Load()) 
}

//ÄrTom -> kollar om kön om tom eller inte
func (k *Dubbelkö) ÄrTom() bool{

	return k.antal.Load() == 0 // 5 == 0 -> false, 0 == 0 -> very true
}

func (k *Dubbelkö) PrintAll(){

	fmt.Print("KÖN -> ")
	//lås head för att hitta första
	k.head.Lock()
	curr := k.head.nästa
	k.head.Unlock()
	//hand over hand
	for{

		curr.Lock()
		if curr ==k.tail{ // stanna vid tail

			curr.Unlock()
			break
			
		}
		fmt.Printf("%s ", curr.värde)
		next := curr.nästa
		curr.Unlock()
		curr =next
		
	}
	fmt.Println()


	
}

func main(){

	kö := NyKö()
	var wg sync.WaitGroup // typ "jag startar några goroutines, vänta tills alla är klara"

	// första testet: lägg till personer i kön
	fmt.Println("test 1 -> lägg till personer")
	for i := 0; i < 5; i++{

		wg.Add(2) // jag tänker starta x mängd goroutines
		go func(id int){

			defer wg.Done() // "jag är klar" måste köras i varje goroutine
			kö.LäggFram(fmt.Sprintf("Fram-%d", id))
			
		}(i)
		go func(id int){
			defer wg.Done()
			kö.LäggBak(fmt.Sprintf("Bak-%d",id))
		}(i)
	}

	wg.Wait()
	fmt.Printf("Antal i kön: %d\n", kö.Antal())
	kö.PrintAll()

	// andra testet -> ta bort personer från kön
	fmt.Println("\nTest 2-> Tar bort personer")

	for i := 0;i < 3; i++{

		wg.Add(2)
		go func(id int){ // startar goroutine, samt en där nere så 2

			defer wg.Done()
			if v, ok := kö.TaFram(); ok{

				fmt.Printf("Tog fram -> %s\n", v)
			}
		}(i)
		go func(id int) {


			defer wg.Done()
			if v, ok := kö.TaBak(); ok{
				fmt.Printf("Tog bak -> %s\n", v)
			}
			
		}(i)


		
	}

	wg.Wait()
	fmt.Printf("Antal i kön: %d\n", kö.Antal())
	kö.PrintAll()
}
