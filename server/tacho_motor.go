package server

import (
	"fmt"

	"github.com/golang/glog"

	pb "github.com/ross-wu/ev3svc/proto"
	cpb "google.golang.org/genproto/googleapis/rpc/code"
)

// tachoMotor implements OutDevice interface.
type tachoMotor struct {
	dryRun      bool
	addr        string
	pos         int
	posSp       int
	maxSpeed    int
	speed       int
	stopAction  string
	stopActions []string
	cmds        []string
}

func newTachoMotor(port, driver string) (*tachoMotor, error) {
	return nil, fmt.Errorf("unimplemented")
}

func newMockTachoMotor(port, driver string) (*tachoMotor, error) {
	return &tachoMotor{
		dryRun:      true,
		addr:        fmt.Sprintf("/mock/tacho/motor/%s/%s", port, driver),
		pos:         0,
		posSp:       0,
		maxSpeed:    500,
		speed:       0,
		stopAction:  "coast",
		stopActions: []string{"coast", "brake", "hold"},
		cmds: []string{
			"run-forever",
			"run-to-abs-pos",
			"run-to-rel-pos",
			"run-timed",
			"run-direct",
			"stop",
			"reset",
		},
	}, nil
}

func (t *tachoMotor) Name() string {
	return "TachoMotor"
}

func (t *tachoMotor) GetAddress() *pb.Response_Result {
	return &pb.Response_Result{Value: &pb.Response_Result_Str{t.addr}}
}

func (t *tachoMotor) GetCommands() *pb.Response_Result {
	return &pb.Response_Result{List: t.cmds}
}

func (t *tachoMotor) GetPosition() *pb.Response_Result {
	return &pb.Response_Result{Value: &pb.Response_Result_Int{int32(t.pos)}}
}

func (t *tachoMotor) GetMaxSpeed() *pb.Response_Result {
	return &pb.Response_Result{Value: &pb.Response_Result_Int{int32(t.maxSpeed)}}
}

func (t *tachoMotor) GetSpeed() *pb.Response_Result {
	return &pb.Response_Result{Value: &pb.Response_Result_Int{int32(t.speed)}}
}

func (t *tachoMotor) GetStopActions() *pb.Response_Result {
	return &pb.Response_Result{List: t.stopActions}
}

func (t *tachoMotor) GetStopAction() *pb.Response_Result {
	return &pb.Response_Result{Value: &pb.Response_Result_Str{t.stopAction}}
}

func (t *tachoMotor) RunCommand(c string) cpb.Code {
	if t.dryRun {
		glog.Infof("command=%s", c)
	}
	return cpb.Code_OK
}

func (t *tachoMotor) SetPositionSp(sp int) cpb.Code {
	if t.dryRun {
		t.posSp = sp
	}
	return cpb.Code_OK
}

func (t *tachoMotor) SetSpeedSp(sp int) cpb.Code {
	if t.dryRun {
		t.speed = sp
	}
	return cpb.Code_OK
}

func (t *tachoMotor) SetStopAction(a string) cpb.Code {
	if t.dryRun {
		t.stopAction = a
	}
	return cpb.Code_OK
}

// func (t *tachoMotor) execImpl(c context.Context, op *pb.Request_Op) *pb.Response_Result {
// 	glog.Infof("Op: %+v", op)
// 	var ret *pb.Response_Result
// 	switch op.Attr {
// 	case pb.Request_ADDRESS:
// 		ret = t.getAddress()
// 	case pb.Request_POSITION:
// 		ret = t.getPosition()
// 	case pb.Request_POSITION_SP:
// 		ret = t.setPosition(int(op.GetInt()))
// 	case pb.Request_STOP_ACTIONS:
// 		ret = t.getStopActions()
// 	case pb.Request_STOP_ACTION:
// 		s, ok := op.GetValue().(*pb.Request_Op_Str)
// 		if ok {
// 			ret = t.setStopAction(s.Str)
// 		} else {
// 			ret = t.getStopAction()
// 		}
// 	case pb.Request_SPEED:
// 		ret = t.getSpeed()
// 	case pb.Request_SPEED_SP:
// 		ret = t.setSpeed(int(op.GetInt()))
// 	case pb.Request_COMMANDS:
// 		ret = t.getCommands()
// 	case pb.Request_COMMAND:
// 		ret = t.runCommand(op.GetStr())
// 	default:
// 		ret = &pb.Response_Result{
// 			Error: cpb.Code_UNIMPLEMENTED,
// 			Msg:   fmt.Sprintf("%v", op.Attr),
// 		}
// 	}
// 	return ret
// }

// func (t *tachoMotor) Exec(c context.Context, req *pb.Request) (*pb.Response, error) {
// 	resp := &pb.Response{}
// 	for i, op := range req.Ops {
// 		result := t.execImpl(c, op)
// 		if result.Error != cpb.Code_OK {
// 			glog.Errorf("i=%d execImpl(%s) error: %s %s",
// 				i, op.String(), result.Error.Enum(), result.Msg)
// 		}
// 		resp.Results = append(resp.Results, result)
// 	}
// 	return resp, nil
// }
