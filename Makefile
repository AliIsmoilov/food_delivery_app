
build:  
	bazel build //...

tidy:
	go mod tidy

gazelle:
	bazel run //:gazelle

gazelle-synch:
	bazel run //:gazelle-update-repos

synch:
	go mod tidy
	bazel run //:gazelle
	bazel run //:gazelle-update-repos

protogen: #https://medium.com/goc0de/a-cute-bazel-proto-hack-for-golang-ides-2a4ef0415a7f
	bazel query 'kind("proto_link", //...)'  | xargs bazel run

list-proto-targets:
	bazel query 'kind("proto_link", //...)' 
	
test:
	bazel test //... --test_output=errors  
api:
	bazel run //src/api_gateway


proto-gen:
	bazel run //src/idl/auth_service:auth_service_go_proto_link
	bazel run //src/idl/customer_service:customer_go_proto_link
	bazel run //src/idl/restaurant_service:restaurant_go_proto_link