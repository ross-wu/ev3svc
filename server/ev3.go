package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/glog"
	pb "github.com/ross-wu/ev3svc/proto"
	cpb "google.golang.org/genproto/googleapis/rpc/code"
)

var UnimplementedError = status.Error(codes.Unimplemented, "unimplemented")

type OutDevice interface {
	Name() string

	GetAddress() *pb.Response_Result
	GetCommands() *pb.Response_Result
	GetPosition() *pb.Response_Result
	GetMaxSpeed() *pb.Response_Result
	GetSpeed() *pb.Response_Result
	GetStopActions() *pb.Response_Result
	GetStopAction() *pb.Response_Result

	RunCommand(cmd string) cpb.Code
	SetPositionSp(sp int) cpb.Code
	SetSpeedSp(sp int) cpb.Code
	SetStopAction(a string) cpb.Code
}

type Ev3Server struct {
	dryRun  bool
	devices map[pb.Port]OutDevice
}

func NewMockEv3Server() (*Ev3Server, error) {
	return &Ev3Server{
		dryRun:  true,
		devices: map[pb.Port]OutDevice{},
	}, nil
}

func (e *Ev3Server) Init(c context.Context, req *pb.InitRequest) (*pb.InitResponse, error) {
	glog.Infof("Init(%+v)...", *req)

	var err error
	switch req.Device {
	case pb.InitRequest_TACHO_MOTOR:
		e.devices[req.Port], err = newMockTachoMotor("todo", req.Driver)
	default:
		return &pb.InitResponse{
			Error: cpb.Code_INVALID_ARGUMENT,
			Msg:   fmt.Sprintf("unknown device %v", req.Device),
		}, nil
	}
	if err != nil {
		glog.Errorf("create device on port %v error: %v", req.Port, err)
		return &pb.InitResponse{
			Error: cpb.Code_INTERNAL,
			Msg:   fmt.Sprintf("create device error: %v", err),
		}, nil
	}

	glog.Infof("Successfully created %s on port %s!", e.devices[req.Port], req.Port)
	return &pb.InitResponse{Error: cpb.Code_OK}, nil
}

func (e *Ev3Server) execOp(c context.Context, op *pb.Request_Op) *pb.Response_Result {
	d, ok := e.devices[op.Port]
	if !ok {
		glog.Errorf("port %v not found", op.Port)
		return &pb.Response_Result{
			Error: cpb.Code_FAILED_PRECONDITION,
			Msg:   fmt.Sprintf("port %v is not initialized, please call Init() first", op.Port),
		}
	}
	glog.Infof("Op: %+v", op)

	ret := &pb.Response_Result{}
	switch op.Attr {
	case pb.Request_ADDRESS:
		ret = d.GetAddress()
	case pb.Request_COMMANDS:
		ret = d.GetCommands()
	case pb.Request_POSITION:
		ret = d.GetPosition()
	case pb.Request_MAX_SPEED:
		ret = d.GetMaxSpeed()
	case pb.Request_SPEED:
		ret = d.GetSpeed()
	case pb.Request_STOP_ACTIONS:
		ret = d.GetStopActions()
	case pb.Request_STOP_ACTION:
		ret = d.GetStopAction()

	case pb.Request_COMMAND:
		ret.Error = d.RunCommand(op.GetStr())
	case pb.Request_POSITION_SP:
		ret.Error = d.SetPositionSp(int(op.GetInt()))
	case pb.Request_SPEED_SP:
		ret.Error = d.SetSpeedSp(int(op.GetInt()))
	case pb.Request_SET_STOP_ACTION:
		ret.Error = d.SetStopAction(op.GetStr())

	default:
		ret = &pb.Response_Result{
			Error: cpb.Code_UNIMPLEMENTED,
			Msg:   fmt.Sprintf("%v", op.Attr),
		}
	}

	return ret
}

func (e *Ev3Server) Exec(c context.Context, req *pb.Request) (*pb.Response, error) {
	resp := &pb.Response{}
	for i, op := range req.Ops {
		ret := e.execOp(c, op)
		if ret.Error != cpb.Code_OK {
			glog.Errorf("i=%d execOp(%v) error: %s %s", i, op, ret.Error.Enum(), ret.Msg)
		}
		resp.Results = append(resp.Results, ret)
	}
	return resp, nil
}
