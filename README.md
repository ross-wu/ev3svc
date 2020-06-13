# ev3svc
Ev3 service runs on LEGO EV3 brick (6009996) and allows clients to send commands to it from your computer.


## Build

### Build proto

```
$ cd ./proto/
$ protoc -I. -I$HOME/github/googleapis --go_out=plugins=grpc:../../../../ service.proto
```

### Build and run `ev3_server`

```
$ go build ev3_server.go

$ ./ev3_server --logtostderr
```

**Test RPC**

```
# init device:
$ grpc_cli call localhost:10000 ev3proto.Ev3.Init '
  device: TACHO_MOTOR port: OUT_A driver: "lego-ev3-l-motor"'

# send operations:
$ grpc_cli call localhost:10000 ev3proto.Ev3.Exec '
  ops { port:OUT_A attr:COMMANDS }
  ops { port:OUT_A attr:MAX_SPEED }
  ops { port:OUT_A attr:POSITION_SP int:100 }
  ops { port:OUT_A attr:SPEED_SP int:200 }
  ops { port:OUT_A attr:COMMAND str:"run-to-abs-pos" }
'
```
