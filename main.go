package main

import (
	//"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/mariobris/pecolify"
	"gopkg.in/ini.v1"
	"strings"
)

var profiles []string

func main() {
	// Instanciate pecolify
	pf := pecolify.New()

	//credName := config.DefaultSharedCredentialsFilename()
	confName := config.DefaultSharedConfigFilename()

	f, err := ini.Load(confName)
	if err != nil {
		fmt.Println("Cannot load default shared credentials file! %s", err)
	}
	for _, p := range f.Sections() {
		// We get also DEFAULT section and is empty by default. Let's skip all empty sections
		if len(p.Keys()) != 0 {
			// Some profile names having prefix 'profile' so we will remove it. Might be related to aws cli version.
			profileName := strings.Replace(p.Name(), "profile ", "", 1)
			profiles = append(profiles, profileName)
		}

	}
	// pecolify!
	selected, err := pf.Transform(profiles)
	if err != nil {
		fmt.Printf("Error was occured: %v\n", err)
		return
	}

	fmt.Printf("export AWS_PROFILE=%s\n", selected)

}
