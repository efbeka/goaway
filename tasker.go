    
package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {
    
    dateCmd := exec.Command("date")
    
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    grepCmd := exec.Command("grep", "tasks")

    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("Tasker - Task1 \n Tasker - Task2 "))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

    fmt.Println("> grep 1")
    fmt.Println(string(grepBytes))

    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }

//TODO
//Remove the grep simulation, and write the real tasker!
}