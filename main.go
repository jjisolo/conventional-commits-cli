package main

import(
  "strings"
  "bufio"
  "bytes"
  "flag"
  "fmt"
  "os"
  "os/exec"
)

var helpMessage   = flag.Bool("help", false, "Show help message")
var stageAll      = flag.Bool("a", false, "Stage all commits")

var commitMessage string

func getString(reader *bufio.Reader) string {
  input, err := reader.ReadString('\n')
	if err != nil {
	  fmt.Println("An error occured while reading input. Please try again", err)
	}

  return input
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func constructCommitMessage() {
  possibleCommitTypes := []string{"feat", "fix", "docs", "style", "refactor", "perf"}
  reader := bufio.NewReader(os.Stdin)
  
  fmt.Printf("Select the type if change that you're commiting [%s]: ", strings.Join(possibleCommitTypes[:], "/"))
  input := strings.ToLower(strings.TrimSuffix(getString(reader), "\n"))
  if stringInSlice(input, possibleCommitTypes) {
 	  commitMessage += input
  } else {
    fmt.Println("This type is not conventional!")
    os.Exit(1)
  }

  fmt.Print("\nSelect the scope of this change: ")
  input = strings.ToLower(strings.TrimSuffix(getString(reader), "\n"))
  if len(input) == 0 {
    commitMessage += ":"
  } else {
    commitMessage += "(" + strings.TrimSuffix(input, "\n") + "):"
  }

  fmt.Print("\nWrite a short imperative description of the change: ")
  input = strings.TrimSuffix(getString(reader), "\n")
  if len(input) != 0 {
    commitMessage += " " + input
  } else {
    commitMessage += " No commit message :("
  }

  fmt.Print("\nProvide a longer description of the change: ")
  input = strings.TrimSuffix(getString(reader), "\n")
  if len(input) != 0 {
    commitMessage += "\n" + input
  }
 
  fmt.Print("\nList any breaking changes or issues closed by this change: ")
  input = strings.TrimSuffix(getString(reader), "\n")
  if len(input) != 0 {
    commitMessage += "\n" + input
  }
}

func executeCommitMessage() {
  command := ""
  if !*stageAll {
    command += "-a"
  } else {
    command += strings.Join(os.Args[2:], " ")
  }
  command += "-m \"" + commitMessage + "\""

  var stdout bytes.Buffer
  var stderr bytes.Buffer
  cmd := exec.Command("git", "commit", command)
  cmd.Stdout = &stdout
  cmd.Stderr = &stderr
  err := cmd.Run()
  if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + stderr.String()) 
  }

  fmt.Println(stdout.String())

}

func main() {
  flag.Parse()

  if *helpMessage {
    flag.Usage()
    os.Exit(1)
  }

  constructCommitMessage()
  executeCommitMessage()
}
