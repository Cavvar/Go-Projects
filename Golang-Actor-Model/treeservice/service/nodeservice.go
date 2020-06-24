package service

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/ob-vss-ws19/blatt-3-angelosolo/messages"
	"github.com/ob-vss-ws19/blatt-3-angelosolo/tree"

	"log"
)

type NodeService struct {
	Trees  map[int32]TreeIdent
	NextID int32
}

type TreeIdent struct {
	token string
	pid   *actor.PID
	id    int32
}

func (treeService *NodeService) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.Trees:
		log.Printf("Currently available trees:\n")
		response := make([]*messages.TreesResponse, 0)
		for _, v := range treeService.Trees {
			response = append(response, &messages.TreesResponse{
				Token: v.token,
				Id:    v.id,
			})
		}
		context.Respond(&messages.SendBackTreeResponse{Trees: response})
	case *messages.CreateNewTreeForCLI:
		log.Printf("New tree-actor(root) will be created!")
		props := actor.PropsFromProducer(func() actor.Actor {
			return &tree.NodeActor{
				Parent:     nil,
				Left:       nil,
				Right:      nil,
				LeftMaxKey: 0,
				LeafSize:   msg.LeafSize,
				Values:     make(map[int32]string),
			}
		})
		pid := context.Spawn(props)
		newTree := TreeIdent{token: CreateToken(7), pid: pid, id: treeService.NextID}
		treeService.Trees[newTree.id] = newTree
		context.Respond(&messages.CreateNewTreeResponse{
			Id:    newTree.id,
			Token: newTree.token,
		})
		log.Printf("Tree with id: %v and token %v was successfully created!\n", newTree.id, newTree.token)
		treeService.NextID++
	case *messages.InsertCLI:
		log.Printf("Trying to insert key %v with value %v\n", msg.Key, msg.Value)
		treeIdent, ok := treeService.Trees[msg.Id]
		if CheckIDAndToken2(ok, msg.Id, msg.Token, treeIdent) {
			log.Printf("Adding key %v with value %v to the tree!\n", msg.Key, msg.Value)
			pid := treeIdent.pid
			// CLI sends the message to the actual NodeActor
			context.RequestWithCustomSender(pid, &messages.Add{Key: msg.Key, Value: msg.Value}, context.Sender())
			// Currently missing sending success message back
		} else {
			context.Respond(&messages.TreeTokenOrIDInvalid{Id: msg.Id, Token: msg.Token})
		}
	case *messages.DeleteCLI:
		log.Printf("Trying to delete key %v\n", msg.Key)
		treeIdent, ok := treeService.Trees[msg.Id]
		if CheckIDAndToken2(ok, msg.Id, msg.Token, treeIdent) {
			log.Printf("Deleting key %v!\n", msg.Key)
			pid := treeIdent.pid
			// CLI sends the message to the actual NodeActor
			context.RequestWithCustomSender(pid, &messages.DeleteKey{Key: msg.Key}, context.Sender())
			// Currently missing sending success message back
		} else {
			context.Respond(&messages.TreeTokenOrIDInvalid{Id: msg.Id, Token: msg.Token})
		}
	case *messages.SearchCLI:
		log.Printf("Trying to find a key %v\n", msg.Key)
		treeIdent, ok := treeService.Trees[msg.Id]
		if CheckIDAndToken2(ok, msg.Id, msg.Token, treeIdent) {
			log.Printf("Looking for key %v now!\n", msg.Key)
			pid := treeIdent.pid
			// CLI sends the message to the actual NodeActor
			context.RequestWithCustomSender(pid, &messages.Find{Key: msg.Key}, context.Sender())
			// Currently missing sending success message back
		} else {
			context.Respond(&messages.TreeTokenOrIDInvalid{Id: msg.Id, Token: msg.Token})
		}
	case *messages.TraverseCLI:
		log.Printf("Trying to traverse through a tree!\n")
		treeIdent, ok := treeService.Trees[msg.Id]
		if CheckIDAndToken2(ok, msg.Id, msg.Token, treeIdent) {
			log.Printf("Traversing through the tree with id %v now!\n", msg.Id)
			pid := treeIdent.pid
			// CLI sends the message to the actual NodeActor
			context.RequestWithCustomSender(pid, &messages.Traverse{}, context.Sender())
			// Currently missing sending success message back
		} else {
			context.Respond(&messages.TreeTokenOrIDInvalid{Id: msg.Id, Token: msg.Token})
		}
	case *messages.DeleteTreeCLI:
		log.Printf("Trying to delete a tree!\n")
		treeIdent, ok := treeService.Trees[msg.Id]
		if CheckIDAndToken2(ok, msg.Id, msg.Token, treeIdent) {
			log.Printf("Deleting the tree with id %v and token %v now!\n", msg.Id, msg.Token)
			pid := treeIdent.pid
			context.Stop(pid)
			delete(treeService.Trees, msg.Id)
			// NodeService directly responds to the remote actor
			context.Respond(&messages.SuccessfulTreeDelete{
				Id:    msg.Id,
				Token: msg.Token,
			})
			// Currently missing sending success message back
		} else {
			context.Respond(&messages.TreeTokenOrIDInvalid{Id: msg.Id, Token: msg.Token})
		}
	}
}
