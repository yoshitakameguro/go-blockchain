### Go Blockchain

### Start Up
```
docker-compose up
```

### Test Command
Run all test files 
```
docker-compose exec server go test -v ./...
```

Run specified directory test files
```
docker-compose exec server go test -v path/to/directory
ex) docker-compose exec server go test -v ./api/controllers/
```
