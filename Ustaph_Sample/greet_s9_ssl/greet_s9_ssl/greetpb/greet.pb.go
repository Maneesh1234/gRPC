
syntax = "proto3";

package greet_s8_bid;

//THIS IS FOR CODE GENERATION
option go_package="greet_s8_bid/greetpb";





message Greeting {
    string first_name = 1;
    string last_name = 2;
}

message GreetRequest {
    Greeting greeting = 1;
}

message GreetResponse {
    string result = 1;
}

message GreetManyTimesRequest {
    Greeting greeting = 1;
}

message GreetManytimesResponse {
    string result = 1;
}

message LongGreetRequest {
    Greeting greeting = 1;
}

message LongGreetResponse {
    string result = 1;
}
message GreetEveryoneRequest {
    Greeting greeting = 1;
}

message GreetEveryoneResponse {
    string result = 1;
}

service GreetService{
    // Unary
    rpc Greet(GreetRequest) returns (GreetResponse) {};

     // Server Streaming
     rpc GreetManyTimes(GreetManyTimesRequest) returns (stream GreetManytimesResponse) {};

      // Client Streaming
    rpc LongGreet(stream LongGreetRequest) returns (LongGreetResponse) {};

     // BiDi Streaming
     rpc GreetEveryone(stream GreetEveryoneRequest) returns (stream GreetEveryoneResponse) {};

}
