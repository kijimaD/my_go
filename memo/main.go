package main

import(
	"fmt"
	"os/exec"
	"time"
)

func main() {
	type output struct {
		out []byte
		err error
	}

	ch := make(chan output)
	// start := time.Now()
	go func() {
		cmd := exec.Command("sleep", "3")
		out, err := cmd.CombinedOutput()
		ch <- output{out, err, }
	}()

progress:
	for {
		select {
		case <-ch:
			fmt.Println("finish")
			break progress
		default:
			timer1 := time.NewTimer(1 * time.Second)
			<-timer1.C
			// end := time.Now();
			// fmt.Printf("%f秒",(end.Sub(start)).Seconds())
			fmt.Print(".")
			// println("Waiting for data")
			// time.Sleep(time.Duration(math.MaxInt64))
		}
	}
}

// func main() {
//     type output struct {
// 	    out []byte
// 	    err error
// 	    done bool
//     }

//     ch := make(chan output)

//     go func() {
//         // cmd := exec.Command("sleep", "1")
//         cmd := exec.Command("sleep", "1")
//         // cmd := exec.Command("false")
//         out, err := cmd.CombinedOutput()
//         ch <- output{out, err, true}
//     }()

// 	// x := <-ch
// 	i := 1
// 	for  {
// 		fmt.Println(i)
// 		i++
// 		fmt.Println("a")
// 		if ch != nil  {
// 			// chは自動で待ってくれてる...
// 			fmt.Println(ch)
// 			break
// 		}
// 	}

//     select {
//     case <-time.After(2 * time.Second):
//         fmt.Println("timed out")
//     case x := <-ch:
//         fmt.Printf("program done; out: %q\n", string(x.out))
//         if x.err != nil {
//             fmt.Printf("program errored: %s\n", x.err)
//         }
//     }
// }
