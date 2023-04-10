run:
	@go run main.go

objectGameBuild:
	go build -o build/g5pt main.go
exeGameBuild:
	GOOS=windows go build -o build/g5pt.exe main.go
build: objectGameBuild exeGameBuild