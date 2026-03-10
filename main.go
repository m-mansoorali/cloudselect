package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"runtime"

	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v3"
)

type CloudsConfig struct {
	Clouds map[string]Cloud `yaml:"clouds"`
}

type Cloud struct {
	// Auth       Auth   `yaml:"auth"`
	// RegionName string `yaml:"region_name"`
}

// type Auth struct {
// AuthUrl string `yaml:"auth_url"`
// }
// type CloudsConfigTwo struct {
// 	CloudsTwo map[string]interface{} `yaml:"clouds"`
// }

func main() {
	osType := runtime.GOOS
	var clouds CloudsConfig
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	readFile, err := os.ReadFile(homeDir + "/.config/openstack/clouds.yaml")
	// readFile, err := os.ReadFile("./clouds.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(readFile, &clouds)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, clouds.Clouds)

	// Collect all cloud names into a slice
	var cloudNames []string
	for name := range clouds.Clouds {
		cloudNames = append(cloudNames, name)
	}

	// Create and run the prompt
	prompt := promptui.Select{
		Label: "Select Cloud Name",
		Items: cloudNames,
	}

	_, result, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	// Output export command to stdout
	// Note: promptui may write escape codes to stdout during interaction
	// To use with eval, run: eval $(go run main.go 2>/dev/null | grep "^export OS_CLOUD")
	fmt.Printf("export OS_CLOUD=%s\n", result)
	var zshrcPath = homeDir + "/.zshrc"
	zshrcContent, err := os.ReadFile(zshrcPath)
	content := string(zshrcContent)
	// Pattern to match OS_CLOUD= with any value (with or without export, with or without quotes)
	pattern := regexp.MustCompile(`(?m)^\s*(export\s+)?OS_CLOUD=.*$`)
	newLine := fmt.Sprintf("export OS_CLOUD=%q", result)

	if pattern.MatchString(content) {
		// Replace existing line
		content = pattern.ReplaceAllString(content, newLine)
	} else {
		// Add new line at the end
		content = strings.TrimSpace(content) + "\n" + newLine + "\n"
	}

	// Write back to .zshrc
	err = os.WriteFile(zshrcPath, []byte(content), 0644)
	if err != nil {
		panic(fmt.Errorf("failed to write .zshrc: %v", err))
	}

	fmt.Fprintf(os.Stderr, "Updated .zshrc: export OS_CLOUD=%q\n", result)
	fmt.Fprintf(os.Stderr, "Run 'source ~/.zshrc' or open a new terminal to apply changes.\n")
	fmt.Println("OS is = ", osType)
}
