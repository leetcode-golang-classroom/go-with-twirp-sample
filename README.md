# go-with-twirp-sample

This repository is for implementation twirp with golang 

## define the proto

```proto
syntax = "proto3";

package twirp.example.haberdasher;
option go_package = "github.com/leetcode-golang-classroom/go-with-twirp-sample/rpc/haberdasher";

// Haberdasher service makes hats for clients.
service Haberdasher {
  // MakeHat produces a hat of mysterious, randomly-selected color!
  rpc MakeHat(Size) returns (Hat);
}

// Size of a Hat, in inches.
message Size {
  int32 inches = 1; // must be > 0
}

// A Hat is a piece of headwear made by a Haberdasher.
message Hat {
  int32 inches = 1;
  string color = 2; // anything but "invisible"
  string name = 3; // i.e. "bowler"
}
```

## install twirp plugin

```shell
go install github.com/twitchtv/twirp/protoc-gen-twirp@latest
```

## test with curl

```shell
curl -X POST http://localhost:8080/twirp/twirp.example.haberdasher.Haberdasher/MakeHat -H 'Content-Type: application/json' -d '{"inches": 15}'
```

## twirp vs gRPC

twirp support HTTP/1.1 while gRPC only allow HTTP/2.0 from client side