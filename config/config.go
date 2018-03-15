package config

import (
	"fmt"
	"log"
	"sync"

	"os"

	"github.com/Nitecon/1Password/exit"
	"github.com/Nitecon/1Password/jsont"
	"github.com/Nitecon/1Password/user"
)

var (
	config     *Configuration
	configLock = new(sync.RWMutex)
)

type Configuration struct {
	MyPath         string
	VaultLocations []string
}

// Get function returns the currently set active configuration to be used.
func Get() *Configuration {
	configLock.RLock()
	defer configLock.RUnlock()
	return config

}

func newConfigFile(cf *jsont.JsonFile) *Configuration {
	config := &Configuration{MyPath: cf.FileName}
	err := cf.NewSave(config)
	if err != nil {
		exit.Fatalm(fmt.Sprintf("Could not save initial 1password config file (%s)", cf.FileName), err)
	}
	return config
}

func LoadFromDisk() error {
	config := &Configuration{}
	cp, err := user.GetConfPath()
	if err != nil {
		exit.Fatalm("User does not appear to be valid, exiting:", err)
	}
	cf := &jsont.JsonFile{FileName: cp}

	if err := cf.Load(config); os.IsNotExist(err) {
		log.Println("File doesn't exist so lets make it")
		config = newConfigFile(cf)
	}

	fmt.Printf("Config loaded: %#v\n", config)
	MemSet(config)
	return nil
}

// MemSet sets the global configuration
func MemSet(cfg *Configuration) *Configuration {
	configLock.Lock()
	defer configLock.Unlock()
	config = cfg
	return config
}
