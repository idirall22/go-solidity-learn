solc:
	solc --abi --bin --overwrite contract/todo.sol -o build
abigen:
	abigen --bin=build/Todo.bin --abi=build/Todo.abi --overwrite --pkg=todo --out=gen/todo.go

deploy:
	go run deploy/main.go