package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

const configPath = "%s/.toptracker"

type configInfo struct {
	Email    string
	Password string
	Token    string
}

func isFileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		return false
	}
}

func getConfigPath() string {
	return fmt.Sprintf(configPath, os.Getenv("HOME"))
}

func readConfig() (configInfo, error) {
	path := getConfigPath()
	config := configInfo{}
	if !isFileExists(path) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("TopTracker login (email): ")
		email, _ := reader.ReadString('\n')
		config.Email = strings.TrimSpace(email)
		fmt.Print("TopTracker password: ")
		password, _ := reader.ReadString('\n')
		config.Password = strings.TrimSpace(password)
		if err := saveConfig(config); err != nil {
			return config, err
		}
		log.Printf("Config saved: %s", path)
	} else {
		file, err := ioutil.ReadFile(path)
		if err != nil {
			return config, err
		}
		err = yaml.Unmarshal(file, &config)
		if err != nil {
			return config, err
		}
	}
	return config, nil
}

func saveConfig(config configInfo) error {
	path := getConfigPath()
	buf, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, buf, 0600)
}
