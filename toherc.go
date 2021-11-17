/*
    This program is designed to run in the local directory containing a job,
    jcl or a virtual tape and creates the proper devinit to be pasted into
    the Hercules console. It is a rewrite of my Python 3 program.

    By: Bill Blasingim
    On: 09/28/2021
*/
package main

import (
	"fmt"
    "os"
    "path/filepath"
	"strings"
	"github.com/atotto/clipboard"
)

func main() {
	var args []string
	//fmt.Printf("T: %T", os.Args)
	args = os.Args

	//programName := args[0]
	ln := len(args)
    if ln!=2 {
        fmt.Println("Format: toherc filename\n")
        fmt.Println("This program creates the devinit command in your clipboard")
        fmt.Println("ready to paste in to the Hercules console.")
        fmt.Println("This makes it easy to send files to Hercules!")
    }

    fil:=args[1]
    //fmt.Println("fil ",fil)

    path, err := os.Getwd()	// Get working directory
    if err != nil {
        fmt.Println(err)
    }
    path=path+"/"

    ext := strings.ToLower(filepath.Ext(fil))

    cmd:=""
    if ext==".aws" {
        //dirpath = "tapes/"
        cmd="devinit 480 "+path+fil
    } else {
        cmd="devinit 00c "+path+fil}
	clipboard.WriteAll(cmd)
    fmt.Println("'devinit' command ready to be pasted into Hercules console")
}
