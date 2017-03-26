package datastore

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/lanzafame/glossy/conf"
)

var storeFactories StoreFactoryRegistry

func init() {
	storeFactories = NewRegistry()
	storeFactories.Register("toml", NewTomlStore)
}

type Store interface {
	Write([]byte) error
	Read() ([]byte, error)
}

type StoreFactory func(conf conf.Config) (Store, error)

type StoreFactoryRegistry map[string]StoreFactory

func (sfr StoreFactoryRegistry) Register(name string, factory StoreFactory) {
	if factory == nil {
		log.Fatalf("store factory %s does not exist", name)
	}

	_, registered := sfr[name]
	if !registered {
		log.Printf("store factory %s already registered. ignoring", name)
	}
	sfr[name] = factory
}

func NewRegistry() StoreFactoryRegistry {
	m := make(map[string]StoreFactory)
	return StoreFactoryRegistry(m)
}

func CreateStore(conf conf.Config) (Store, error) {
	// default config "shouldn't" be aware of datastore implementations IMHO
	storeName := conf.Get("STORE", "toml")

	factory, ok := storeFactories[storeName]
	if !ok {
		stores := make([]string, len(storeFactories))
		for i := range storeFactories {
			stores = append(stores, i)
		}
		return nil, errors.New(fmt.Sprintf("Store name %s is unknown. Please use one of the following: %s", storeName, strings.Join(stores, ", ")))
	}

	return factory(conf)
}
