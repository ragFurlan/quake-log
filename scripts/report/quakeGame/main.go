package main

import (
	"encoding/json"
	"fmt"
	"os"

	quakeLogController "quakelog/internal/controllers"
	quakeLog "quakelog/internal/platform/http"
	quakeLogUseCase "quakelog/internal/usecase/quakeLog"
)

var (
	controller quakeLogController.QuakeLog
)

func main() {
	url := os.Args[1]
	if url == "" {
		fmt.Println("we don't have url to continue!")
	}

	log := quakeLog.NewQuakeLog(url)
	usecase := quakeLogUseCase.NewQuakeLogUsecase(log)
	controller = quakeLogController.NewQuakeLog(usecase)

	groupedMatch, err := controller.GroupedMatch()
	if err != nil {
		fmt.Println("error on grouped match")
	}

	e, err := json.Marshal(groupedMatch)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))

}
