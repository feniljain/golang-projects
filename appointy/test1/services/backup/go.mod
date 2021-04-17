module go.saastack.io/project

go 1.16

replace go.uber.org/fx => github.com/appointy/fx v1.9.1-0.20190624110333-490d04d33ef6

replace go.uber.org/dig => github.com/paullen/dig v1.7.1-0.20190624104937-6e47ebbbdcf6

replace github.com/apache/thrift => github.com/apache/thrift v0.0.0-20190309152529-a9b748bb0e02

require (
	github.com/Shivam010/protoc-gen-validate v0.3.0
	github.com/elgris/sqrl v0.0.0-20190909141434-5a439265eeec
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	go.opencensus.io v0.23.0
	go.saastack.io/activity-log v0.0.0-20210301135449-26b28d7139d5
	go.saastack.io/chaku v0.0.0-20210215113324-b989abac4b46
	go.saastack.io/deployment/events v0.0.0-20210323121019-4485f107c320
	go.saastack.io/deployment/right v0.0.0-20210323121019-4485f107c320
	go.saastack.io/events v0.0.0-20200615072044-33e4fefd1624
	go.saastack.io/eventspush v0.0.0-20210114113847-5f737652968c
	go.saastack.io/idutil v0.0.0-20200122122527-bf060a213abc
	go.saastack.io/jaal v0.0.0-20210129082513-56f41dab727f
	go.saastack.io/modulerole v0.0.0-20210312052340-adf3fe388562
	go.saastack.io/pehredaar v0.0.0-20200810065542-af2466b1ff34
	go.saastack.io/protoc-gen-caw/convert v0.0.0-20210206112258-37a2c0aacda4
	go.saastack.io/protoc-gen-grpc-wrapper v0.0.0-20201124092054-65b8a2efbc5c
	go.saastack.io/protoc-gen-nakaab v0.0.0-20200124074048-9eedb7860ee1
	go.saastack.io/protos v0.0.0-20210317063919-79df979f70b7
	go.saastack.io/right v0.0.0-20201230125918-5d58ef42ac18
	go.saastack.io/userinfo v0.0.0-20200522114021-e48273ccb360
	go.uber.org/cadence v0.16.0
	go.uber.org/fx v1.12.0
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20210330142815-c8897c278d10
	google.golang.org/genproto v0.0.0-20210330181207-2295ebbda0c6
	google.golang.org/grpc v1.36.1
)
