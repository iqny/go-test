package main

import (
	"fmt"
	proto "github.com/golang/protobuf/proto"
	ptypes "github.com/golang/protobuf/ptypes"
	"time"
)

func main() {
	msg := &WellKnownTypes{
		Now:  ptypes.TimestampNow(),
		Took: ptypes.DurationProto(10 * time.Minute),
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		panic(err)
	}
	err = proto.Unmarshal(data, msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(msg.Took)
}
