package main

import(
  "strings"
  "bufio"
  "bytes"
  "flag"
  "fmt"
  "os"
  "os/exec"

  "github.com/manifoldco/promptui"
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
  prompt := promptui.Select{
		Label: "Select the type if change that you're commiting",
		Items: []string{"feat", "fix", "docs", "style", "refactor", "perf"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
  commitMessage += result
 
  promptValidate := promptui.Prompt{
		Label:    "Select the scope of this change",
		Validate: nil,
	}

	result, err = promptValidate.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

  if len(result) == 0 {
    commitMessage += ":"
  } else {
    commitMessage += "(" + result + "):"
  }

  promptValidate = promptui.Prompt{
		Label:    "Write a short imperative description of the change",
		Validate: nil,
	}

	result, err = promptValidate.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

  if len(result) == 0 {
    commitMessage += "No commit message :("
  } else {
    commitMessage += "result"
  }


  promptValidate = promptui.Prompt{
		Label:    "Provide a longer description of the change",
		Validate: nil,
	}

	result, err = promptValidate.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

  commitMessage += result

  promptValidate = promptui.Prompt{
		Label:    "List any breaking changes or issues closed by this change",
		Validate: nil,
	}

	result, err = promptValidate.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

  commitMessage += result
}

func executeCommitMessage() {
  command := ""
  if !*stageAll {
    command += "-a"
  } else {
    command += strings.Join(os.Args[2:], " ")
  }
  command += "-m \"" + commitMessage + "\""
  fmt.Printf("\n\n Executing < git commit \"%s\" >", commitMessage)

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
