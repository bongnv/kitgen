package addsvc

//go:generate protoc -I./pb -I$P/googleapis --go_out=plugins=grpc:./pb ./pb/addsvc.proto
//go:generate protoc -I./pb -I$P/googleapis --grpc-gateway_out=logtostderr=true:./pb ./pb/addsvc.proto
//go:generate kitgen -d=$TESTDATA/pb -i=AddServer -o=$TESTDATA/server/endpoints.go -t=../../templates/tmpl/endpoints.tmpl
