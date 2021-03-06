syntax = "proto3"
package proto


service Greeter {
	rpc Hello(HelloRequest) returns (HelloResponse){}
}

message HelloRequest {
	string name = 1;
}

message HelloResponse {
	string greeting = 2
}