package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate your flag file does not get committed with local flags",
	Long:  "Validate your flag file does not get committed with local flags",
	Run: func(cmd *cobra.Command, args []string) {
		validateFlagFile()
	},
}

var (
	flagFile string
	clean    bool
	git      bool
)

const (
	localBegin = "LOCAL_LCG_FLAGS_BEGIN"
	localEnd   = "LOCAL_LCG_FLAGS_END"
)

func init() {
	validateCmd.PersistentFlags().StringVar(&flagFile, "flagFile", "", "File that is validated")
	if err := viper.BindPFlag("flagFile", validateCmd.PersistentFlags().Lookup("flagFile")); err != nil {
		check(err)
	}
	validateCmd.PersistentFlags().BoolVar(&clean, "clean", false, "Removes any local flags in the flag file")
	if err := viper.BindPFlag("clean", validateCmd.PersistentFlags().Lookup("clean")); err != nil {
		check(err)
	}
	validateCmd.PersistentFlags().BoolVar(&git, "git", false, "Check if the flag file is staged for git commit with local flags present")
	if err := viper.BindPFlag("git", validateCmd.PersistentFlags().Lookup("git")); err != nil {
		check(err)
	}
	rootCmd.AddCommand(validateCmd)
}

func validateFlagFile() {
	if viper.GetString("flagFile") == "" {
		fmt.Println("A flagFile must be provided.")
		os.Exit(1)
	}
	readFile, err := ioutil.ReadFile(viper.GetString("flagFile"))
	check(err)
	outFile := removeLocalFlags(string(readFile), localBegin, localEnd)
	if viper.GetBool("clean") {
		if err := ioutil.WriteFile(viper.GetString("flagFile"), []byte(outFile), 0644); err != nil {
			check(err)
		}
	}
}

func removeLocalFlags(str string, start string, end string) (result string) {
	lines := strings.Split(str, "\n")
	var startLine int
	var endLine int
	var outFile []string
	for idx, line := range lines {
		s := strings.Index(line, start)
		if s > 0 {
			if viper.GetBool("git") && !(viper.GetBool("clean")) {
				cmd := exec.Command("git", "diff", "--cached", "--exit-code", "--", viper.GetString("flagFile"))
				err := cmd.Run()
				var (
					ee *exec.ExitError
				)
				if errors.As(err, &ee) {
					fmt.Println("Local flags found staged!") // ran, but non-zero exit code
					os.Exit(ee.ExitCode())

				} else if err != nil {
					fmt.Printf("%v", err)
					os.Exit(ee.ExitCode())

				}
			} else if !(viper.GetBool("clean")) {
				fmt.Println("Local Flags found!")
				os.Exit(1)
			}
		}
		if s == -1 {
			continue
		}
		startLine = idx
		break
	}
	for idx, line := range lines {
		s := strings.Index(line, end)
		if s == -1 {
			continue
		}
		endLine = idx + 1
		break
	}
	if startLine == -1 && endLine == -1 {
		return str
	}
	outFile = append(outFile, lines[0:startLine]...)
	outFile = append(outFile, lines[endLine:]...)
	return strings.Join(outFile, "\n")
}
