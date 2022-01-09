package main

import (
    "fmt"
    "sync"
    "time"
)

/*
   Implement the dining philosopher’s problem with the following constraints/modifications.

   There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

   Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

   The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

   In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

   The host allows no more than 2 philosophers to eat concurrently.

   Each philosopher is numbered, 1 through 5.

   When a philosopher starts eating (after it has obtained necessary locks) it prints
   “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

   When a philosopher finishes eating (before it has released its locks) it prints
   “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

type ChopS struct {
    sync.Mutex
}

type Philo struct {
    leftCS, rightCS *ChopS
    id int
}

func (p *Philo) eat(c chan bool, wait *sync.WaitGroup) {
    for i := 0; i < 3; i++ {
        c <- true

        p.leftCS.Lock()
        p.rightCS.Lock()

        fmt.Printf("starting to eat %d\n", p.id)
        time.Sleep(500 * time.Millisecond)
        fmt.Printf("finishing eating %d\n", p.id)

        p.leftCS.Unlock()
        p.rightCS.Unlock()
    }

    fmt.Printf("Philosopher %d has finished eating.\n", p.id)
    wait.Done()
}

func host(c chan bool) {
    for {
        if len(c) == 2 {
            <- c
            <- c

            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    CSticks := make([]*ChopS, 5)
    philos := make([]*Philo, 5)
    c := make(chan bool, 2)
    var wait sync.WaitGroup

    for i := 0; i < 5; i++ {
        CSticks[i] = new(ChopS)
    }


    for i := 0; i < 5; i++ {
        philos[i] = &Philo{CSticks[i],CSticks[(i+1) % 5], i + 1 }
    }

    go host(c)
    for i := 0; i < 5; i++ {
        wait.Add(1)
        go philos[i].eat(c, &wait)
    }

    wait.Wait()
}