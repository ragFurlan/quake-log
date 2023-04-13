# Project quake-log

## To run the report of quake games

- go to folder scripts/report/quakeGame
```
go run main.go https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log
```

## To run the report of kills by means

- go to folder scripts/report/killByMeans
```
go run main.go https://gist.githubusercontent.com/cloudwalk-tests/be1b636e58abff14088c8b5309f575d8/raw/df6ef4a9c0b326ce3760233ef24ae8bfa8e33940/qgames.log
```

## to create mock

### platform
```
~/go/bin/mockgen -source=internal/platform/http/quakeLog.go -destination=test/platform/quakeLog.go -package=quakelog
```

### usecase
```
~/go/bin/mockgen -source=internal/usecase/quakeLog/usecase.go -destination=test/usecase/usecase.go -package=quakelog
```

## To run test
go test ./...