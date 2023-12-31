up_build:
	-docker compose up --build -d

up:
	-docker compose up -d

down:
	-docker compose down

install_hat:
	go get golang.org/x/sync/errgroup
	go get google.golang.org/grpc
	go get google.golang.org/grpc/codes
	go get google.golang.org/grpc/status
	go get google.golang.org/protobuf/reflect/protoreflect
	go get google.golang.org/protobuf/runtime/protoimpl

run_hat:
	make install_hat
	tail -f /dev/null

hat_exec:
	docker compose exec hat go run cmd/hat/main.go

illusionist_exec:
	docker compose exec illusionist go run cmd/illusionist/main.go

run_illusionist:
	go run cmd/illusionist/main.go

run_hat_exec:
	go run cmd/hat/main.go