package tree

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/ob-vss-ws19/blatt-3-angelosolo/messages"
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

func TestCreateEmptyTree(t *testing.T) {
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{
			Left:       nil,
			Right:      nil,
			LeftMaxKey: 0,
			LeafSize:   1,
			Values:     make(map[int32]string),
		}
	})
	assert.NotPanics(t, func() { context.Spawn(props) }, "The Code did panic!")
}

func TestAdd(t *testing.T) {
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{
			Left:       nil,
			Right:      nil,
			LeftMaxKey: 0,
			LeafSize:   2,
			Values:     make(map[int32]string),
		}
	})
	root := context.Spawn(props)
	context.Send(root, &messages.Add{
		Key:   7,
		Value: "a",
	})
	context.Send(root, &messages.Add{
		Key:   4,
		Value: "b",
	})
	context.Send(root, &messages.Add{
		Key:   1,
		Value: "c",
	})
	context.Send(root, &messages.Add{
		Key:   5,
		Value: "d",
	})
	context.Send(root, &messages.Find{Key: 5})
	context.Send(root, &messages.Find{Key: 4})
}

func TestAddTwo(t *testing.T) {
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{
			Left:       nil,
			Right:      nil,
			LeftMaxKey: 0,
			LeafSize:   2,
			Values:     make(map[int32]string),
		}
	})
	root := context.Spawn(props)
	context.Send(root, &messages.Add{
		Key:   7,
		Value: "a",
	})
	context.Send(root, &messages.Add{
		Key:   4,
		Value: "b",
	})
	context.Send(root, &messages.Add{
		Key:   5,
		Value: "d",
	})
	context.Send(root, &messages.Add{
		Key:   1,
		Value: "c",
	})
	context.Send(root, &messages.Find{Key: 4})
	context.Send(root, &messages.Find{Key: 5})
	context.Send(root, &messages.Find{Key: 2})
}

func TestFind(t *testing.T) {
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{
			Left:       nil,
			Right:      nil,
			LeftMaxKey: 0,
			LeafSize:   2,
			Values:     make(map[int32]string),
		}
	})
	root := context.Spawn(props)
	context.Send(root, &messages.Add{
		Key:   7,
		Value: "a",
	})
	context.Send(root, &messages.Add{
		Key:   4,
		Value: "b",
	})
	context.Send(root, &messages.Find{Key: 4})
}

func TestFindError(t *testing.T) {
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{
			Left:       nil,
			Right:      nil,
			LeftMaxKey: 0,
			LeafSize:   1,
			Values:     make(map[int32]string),
		}
	})
	root := context.Spawn(props)
	context.Send(root, &messages.Add{Key: 3, Value: "This is a test!"})
	future := context.RequestFuture(root, &messages.Find{Key: 2}, 1*time.Second)

	result, err := future.Result()
	if err == nil {
		response, ok := result.(*messages.ErrorFindingValue)
		if !ok {
			t.Errorf("Did not get the expected error message! %v\n", response)
		}
	} else {
		t.Errorf("Could not get future: %v\n", err)
	}
}

func TestTraverse(t *testing.T) {
	want := map[int32]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight"}

	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{
			Left:       nil,
			Right:      nil,
			LeftMaxKey: 0,
			LeafSize:   2,
			Values:     make(map[int32]string),
		}
	})
	root := context.Spawn(props)
	context.Send(root, &messages.Add{Key: 1, Value: "one"})
	context.Send(root, &messages.Add{Key: 2, Value: "two"})
	context.Send(root, &messages.Add{Key: 3, Value: "three"})
	context.Send(root, &messages.Add{Key: 4, Value: "four"})
	context.Send(root, &messages.Add{Key: 5, Value: "five"})
	context.Send(root, &messages.Add{Key: 6, Value: "six"})
	context.Send(root, &messages.Add{Key: 7, Value: "seven"})
	context.Send(root, &messages.Add{Key: 8, Value: "eight"})

	future := context.RequestFuture(root, &messages.Traverse{}, 1*time.Second)
	result, err := future.Result()

	if err == nil {
		response, ok := result.(*messages.TraverseResponse)
		if ok {
			for i := range response.KvPair {
				if response.KvPair[i].Value != want[response.KvPair[i].Key] {
					t.Errorf("Got: %v - Wanted: %v\n", response.KvPair[i].Value, want[int32(i)])
				}
			}
		} else {
			t.Errorf("Expected other message type than %v\n", response.GoString())
		}
	} else {
		t.Errorf("Could not get future: %v\n", err)
	}
}

func TestDeleteFromRoot(t *testing.T) {
	want := map[int32]string{}
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{
			Left:       nil,
			Right:      nil,
			LeftMaxKey: 0,
			LeafSize:   3,
			Values:     make(map[int32]string),
		}
	})
	root := context.Spawn(props)
	context.Send(root, &messages.Add{Key: 2, Value: "b"})
	context.Send(root, &messages.DeleteKey{Key: 2})

	future := context.RequestFuture(root, &messages.Traverse{}, 1*time.Second)
	result, err := future.Result()
	if err == nil {
		response, ok := result.(*messages.TraverseResponse)
		if ok {
			for i := range response.KvPair {
				if response.KvPair[i].Value != want[response.KvPair[i].Key] {
					t.Errorf("Got: %v - Wanted: %v\n", response.KvPair[i].Value, want[int32(i)])
				}
			}
		} else {
			t.Errorf("Expected other message type than %v\n", response.GoString())
		}
	} else {
		t.Errorf("Could not get future: %v\n", err)
	}
}

func TestDeleteFromEmptyRoot(t *testing.T) {
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{
			Left:       nil,
			Right:      nil,
			LeftMaxKey: 0,
			LeafSize:   1,
			Values:     make(map[int32]string),
		}
	})
	root := context.Spawn(props)
	future := context.RequestFuture(root, &messages.DeleteKey{Key: 2}, 1*time.Second)
	result, err := future.Result()
	if err == nil {
		response, ok := result.(*messages.CouldNotFindKey)
		if !ok {
			t.Errorf("Did not get the expected error message! %v\n", response)
		}
	} else {
		t.Errorf("Could not get future: %v\n", err)
	}
}

func TestDeleteFromRightSubTreeEmptyValues(t *testing.T) {
	wantOne := map[int32]string{1: "a"}
	context := actor.EmptyRootContext
	props := actor.PropsFromProducer(func() actor.Actor {
		return &NodeActor{
			Left:       nil,
			Right:      nil,
			LeftMaxKey: 0,
			LeafSize:   1,
			Values:     make(map[int32]string),
		}
	})
	root := context.Spawn(props)
	context.Send(root, &messages.Add{Key: 1, Value: "a"})
	context.Send(root, &messages.Add{Key: 3, Value: "c"})
	context.Send(root, &messages.DeleteKey{Key: 3})
	future := context.RequestFuture(root, &messages.Traverse{}, 1*time.Second)
	result, err := future.Result()
	if err == nil {
		response, ok := result.(*messages.TraverseResponse)
		if ok {
			for i := range response.KvPair {
				if response.KvPair[i].Value != wantOne[response.KvPair[i].Key] {
					t.Errorf("Got: %v - Wanted: %v\n", response.KvPair[i].Value, wantOne[int32(i)])
				}
			}
		} else {
			t.Errorf("Expected other message type than %v\n", response.GoString())
		}
	} else {
		t.Errorf("Could not get future: %v\n", err)
	}
}
