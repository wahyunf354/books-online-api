 go test ../business/... -coverprofile=cover.out && go tool cover -html=cover.out

 go test ./business/... -coverprofile=cover.out && go tool cover -html=cover.out