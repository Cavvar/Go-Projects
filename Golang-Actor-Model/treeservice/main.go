package main

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/ob-vss-ws19/blatt-3-angelosolo/treeservice/service"

	"flag"
	"fmt"
	"sync"
)

func newRemoteActor() actor.Producer {
	return func() actor.Actor {
		return &service.NodeService{
			Trees:  make(map[int32]service.TreeIdent),
			NextID: 1337,
		}
	}
}

func newRemote(server, name string) {
	remote.Start(server)
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(newRemoteActor())
	pid, err := context.SpawnNamed(props, name)

	remote.Register("hello", props)
	var waitGroup sync.WaitGroup

	if err == nil {
		// Yeah
		fmt.Printf("Successfully created actor with pid %v\n", pid)
		waitGroup.Add(1)
		waitGroup.Wait()
	} else {
		fmt.Printf("Something went wrong with creating the named actor")
	}
}

func main() {
	fmt.Println("Hello Tree-Service!")

	flagHost := flag.String("host", "localhost:1860", "Adresse, wo der Remote-Actor gestartet wird (Sechzig:P)")
	flagServiceName := flag.String("serviceName", "treeservice", "Name des 'Baumservices'")

	flag.Parse()

	// Starten des Remote-Actors an der Adresse localhost:1860
	newRemote(*flagHost, *flagServiceName)
}
