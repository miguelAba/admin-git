@REM command to run the proto typescript generator
protoc --proto_path=./protos --ts_proto_out=./protos ./protos/notifications/*.proto --ts_proto_opt=esModuleInterop=false --ts_proto_opt=outputEncodeMethods=false --ts_proto_opt=outputJsonMethods=false --ts_proto_opt=outputPartialMethods=false --ts_proto_opt=outputClientImpl=false
protoc --proto_path=./protos --ts_proto_out=./protos ./protos/schedule/*.proto --ts_proto_opt=esModuleInterop=false --ts_proto_opt=outputEncodeMethods=false --ts_proto_opt=outputJsonMethods=false --ts_proto_opt=outputPartialMethods=false --ts_proto_opt=outputClientImpl=false


@REM command to run the proto ruby generator
protoc --proto_path=./protos --ruby_out=./protos ./protos/notifications/*.proto
protoc --proto_path=./protos --ruby_out=./protos ./protos/schedule/*.proto
