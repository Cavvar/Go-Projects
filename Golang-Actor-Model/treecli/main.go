package main

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/ob-vss-ws19/blatt-3-angelosolo/messages"

	"errors"
	"flag"
	"fmt"
	"sync"
	"time"
)

type CLIActor struct {
	waitGroup *sync.WaitGroup
}

func (cli CLIActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.CreateNewTreeResponse:
		fmt.Printf("Created tree with token %v and id %v\n", msg.Token, msg.Id)
	case *messages.TreeTokenOrIDInvalid:
		fmt.Printf("The inputted token or id is invalid!\n")
	case *messages.SuccessFindValue:
		fmt.Printf("Found key %v with value %v\n", msg.Key, msg.Value)
	case *messages.TraverseResponse:
		fmt.Printf("Traversal\n%v\n", msg.KvPair)
	case *messages.ErrorFindingValue:
		fmt.Printf("Key could not be found for find operation!\n")
	case *messages.CouldNotFindKey:
		fmt.Printf("Key could not be found!\n")
	case *messages.SuccessDeleteKey:
		fmt.Printf("Successfully deleted key %v with value %v from tree\n", msg.Key, msg.Value)
	case *messages.ErrorKeyDoesNotExist:
		fmt.Printf("This key %v does not exist!\n", msg.Key)
	case *messages.SuccessfulTreeDelete:
		fmt.Printf("Successfully deleted tree with token %v and id %v\n", msg.Token, msg.Id)
	case *messages.SuccessAdd:
		fmt.Printf("Successfully added key %v with value %v to tree", msg.Key, msg.Value)
	case *messages.SendBackTreeResponse:
		fmt.Printf("%v\n", msg.Trees)
	}
}

func main() {
	fmt.Println("Hello Tree-CLI!")
	flagBind := flag.String("bind", "localhost:1338", "bind address to")
	flagRemote := flag.String("remote", "localhost:1860", "remote address")

	flagCreateTree := flag.Bool("newtree", false, "create tree. default not creating tree")
	flagInsert := flag.Bool("insert", false, "flag for inserting a value to the tree")
	flagFind := flag.Bool("find", false, "flag for find a value in the tree")
	flagTraverse := flag.Bool("traverse", false, "flag for traversing the tree")
	flagDeleteKey := flag.Bool("deletekey", false, "flag for deleting a key/value in the tree")
	flagDeleteTree := flag.Bool("deletetree", false, "flag for deleting a tree")

	flagToken := flag.String("token", "", "flag for token. necessary for all operations")
	flagID := flag.Int("id", 0, "flag for id. necessary for all operations")
	flagKey := flag.Int("key", 0, "key when inserting/deleting/finding values")
	flagValue := flag.String("value", "", "value when inserting a key")
	flagLeafSize := flag.Int("leafSize", 0, "leafSize")

	flag.Parse()

	remote.Start(*flagBind)

	var wg sync.WaitGroup

	props := actor.PropsFromProducer(func() actor.Actor {
		wg.Add(1)
		return &CLIActor{&wg}
	})
	context := actor.EmptyRootContext
	pid := context.Spawn(props)

	// message handling through flash here
	var msg interface{}
	switch {
	case *flagCreateTree:
		if *flagLeafSize < 2 {
			panic(errors.New("leaf size hast to be greater than 1"))
		}
		msg = &messages.CreateNewTreeForCLI{LeafSize: int32(*flagLeafSize)}
	case *flagInsert:
		msg = &messages.InsertCLI{
			Id:    int32(*flagID),
			Token: *flagToken,
			Key:   int32(*flagKey),
			Value: *flagValue,
		}
	case *flagFind:
		msg = &messages.SearchCLI{
			Id:    int32(*flagID),
			Token: *flagToken,
			Key:   int32(*flagKey),
		}
	case *flagTraverse:
		msg = &messages.TraverseCLI{
			Id:    int32(*flagID),
			Token: *flagToken,
		}
	case *flagDeleteKey:
		msg = &messages.DeleteCLI{
			Id:    int32(*flagID),
			Token: *flagToken,
			Key:   int32(*flagKey),
		}
	case *flagDeleteTree:
		msg = &messages.DeleteTreeCLI{
			Id:    int32(*flagID),
			Token: *flagToken,
		}
	default:
		msg = &messages.Trees{}
	}

	pidRemote, err := remote.SpawnNamed(*flagRemote, "remote", "hello", 5*time.Second)
	if err != nil {
		fmt.Printf("Could not create remote actor nodeservice!")
		panic(err)
	}

	remotePID := pidRemote.Pid

	// Now Sending message from CLI Remote Actor to NodeServiceActor
	context.RequestWithCustomSender(remotePID, msg, pid)
	wg.Wait()
}
