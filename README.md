# message-in-proto

## commands to generate a profile 

`protoc --go_out=. ./proto/basic/*.proto` 

if you want folder statucture like wihout module name use this command

`protoc --go_opt=module=my-protobuf --go_out=. ./proto/basic/*.proto`