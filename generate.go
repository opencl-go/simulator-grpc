package simulator_grpc

//go:generate protoc --go_out=. --go_opt=paths=source_relative sim/platform.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative sim/devices.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative sim/devices.proto
//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative sim/platform.proto
