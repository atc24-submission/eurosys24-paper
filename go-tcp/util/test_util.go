package util

import pb "github.com/sosp23/replicated-store/go/multipaxos/network"

func MakeInstance(ballot int64, index int64) *pb.Instance {
	return &pb.Instance{Ballot: ballot, Command: &pb.Command{},
		Index: index, State: pb.Inprogress, ClientId: 0}
}

func MakeInstanceWithState(ballot int64, index int64,
	state pb.InstanceState) *pb.Instance {
	instance := MakeInstance(ballot, index)
	instance.State = state
	return instance
}

func MakeInstanceWithType(ballot int64, index int64,
	cmdType pb.CommandType) *pb.Instance {
	instance := MakeInstance(ballot, index)
	instance.Command.Type = cmdType
	return instance
}

func MakeInstanceWithAll(ballot int64, index int64, state pb.InstanceState,
	cmdType pb.CommandType) *pb.Instance {
	instance := MakeInstance(ballot, index)
	instance.State = state
	instance.Command.Type = cmdType
	return instance
}
