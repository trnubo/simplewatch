package main

import "flag"
import "time"
import "os"
import "os/exec"
import "strings"
import "log"
import "syscall"
import "bytes"

func main() {

	firstWait := flag.Int("firstwait", 30, "wait in seconds before first check")
	wait := flag.Int("wait", 30, "wait in seconds between checks")

	flag.Parse()

	checkCmd := flag.Args()

	if len(checkCmd) == 0 {
		log.Fatalln("No command defined, try ./simplewatch -h")
	}

	log.Println("Starting simplewatch. First Wait: ", *firstWait, " Wait: ", *wait, " CMD: ", checkCmd)

	time.Sleep(time.Duration(*firstWait) * time.Second)

	exitStatus := 0
	for exitStatus == 0 {

		cmd := exec.Command(checkCmd[0], checkCmd[1:]...)

		var buffStdout bytes.Buffer
		var buffStderr bytes.Buffer
		cmd.Stdout = &buffStdout
		cmd.Stderr = &buffStderr

		if err := cmd.Start(); err != nil {
			log.Fatalf("cmd.Start: %v")
		}

		err := cmd.Wait()

		if len(buffStderr.Bytes()) > 0 {
			log.Println("==> Error: ", strings.TrimSpace(buffStderr.String()))
		}
		log.Println("==> Output: ", strings.TrimSpace(buffStdout.String()))

		if err != nil {

			if exiterr, ok := err.(*exec.ExitError); ok {
				// The program has exited with an exit code != 0

				// This works on both Unix and Windows. Although package
				// syscall is generally platform dependent, WaitStatus is
				// defined for both Unix and Windows and in both cases has
				// an ExitStatus() method with the same signature.
				if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
					exitStatus = status.ExitStatus()
					log.Printf("Exit Status: %d", status.ExitStatus())
					continue
				}
			} else {
				log.Fatalf("cmd.Wait: %v", err)
			}
		}

		time.Sleep(time.Duration(*wait) * time.Second)

	}

	log.Println("Done. Exit: ", exitStatus)
	os.Exit(exitStatus)

}
