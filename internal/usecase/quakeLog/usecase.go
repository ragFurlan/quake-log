package quakeLogUsecase

import (
	"fmt"
	"regexp"
	"strings"

	entity "quakelog/internal/entity"
	quakeLog "quakelog/internal/platform/http"
)

type QuakeLogUsecase interface {
	GroupedByTypeOfDeath() (map[string]entity.QuakeLogKills, error)
	GroupedMatch() (map[string]entity.QuakeLog, error)
}

type quakeLogUsecase struct {
	quakeLog quakeLog.QuakeLog
}

func NewQuakeLogUsecase(quakeLog quakeLog.QuakeLog) *quakeLogUsecase {
	return &quakeLogUsecase{quakeLog: quakeLog}
}

func (ref quakeLogUsecase) GroupedByTypeOfDeath() (map[string]entity.QuakeLogKills, error) {
	log, err := ref.quakeLog.Get()
	if err != nil {
		return map[string]entity.QuakeLogKills{}, err
	}

	var quakeLogsKills = make(map[string]entity.QuakeLogKills)
	games := strings.Split(log, "InitGame")

	for key, game := range games {
		quakeLogKill := entity.QuakeLogKills{
			KillsByMeans: getKillsByMeans(game),
		}
		nameOfGame := fmt.Sprintf("game-%v", key)
		quakeLogsKills[nameOfGame] = quakeLogKill
	}

	return quakeLogsKills, nil
}

func getKillsByMeans(game string) map[string]int {
	killsByMeans := make(map[string]int)
	for _, meansOfDeath := range entity.AllMeansOFDeath {
		regex := regexp.MustCompile(fmt.Sprintf("%s", meansOfDeath))
		meansOfDeathNumber := regex.FindAllString(game, 99)
		if len(meansOfDeathNumber) == 0 {
			continue
		}

		killsByMeans[meansOfDeath] = len(meansOfDeathNumber)
	}

	return killsByMeans
}

func (ref quakeLogUsecase) GroupedMatch() (map[string]entity.QuakeLog, error) {
	log, err := ref.quakeLog.Get()
	if err != nil {
		return map[string]entity.QuakeLog{}, err
	}

	quakeLogsMatch := make(map[string]entity.QuakeLog)

	games := strings.Split(log, "InitGame")
	for key, game := range games {
		players := getPlayers(game)
		if len(players) == 0 {
			continue
		}

		kills := make(map[string]int, len(players))
		for _, player := range players {
			kills[player] = getKillsByPlayer(game, player)
		}

		groupedMatch := entity.QuakeLog{
			Players:    players,
			TotalKills: strings.Count(game, "killed"),
			Kills:      kills,
		}

		nameOfGame := fmt.Sprintf("game-%v", key)
		quakeLogsMatch[nameOfGame] = groupedMatch
	}

	return quakeLogsMatch, err
}

func getKillsByPlayer(game string, player string) int {
	regexKill := regexp.MustCompile(fmt.Sprintf("Kill(.*)%s killed", player))
	playerKillNumbers := regexKill.FindAllString(game, 99)

	regexWasKilled := regexp.MustCompile(fmt.Sprintf("Kill(.*)<world> killed %s", player))
	playerWasKilledNumbers := regexWasKilled.FindAllString(game, 99)

	return len(playerKillNumbers) - len(playerWasKilledNumbers)
}

func getPlayers(game string) []string {
	re := regexp.MustCompile(`n\\([^\\]+)`)
	quakeMatch := re.FindAllString(game, 99)

	if len(quakeMatch) == 0 {
		return nil
	}

	players := quakeMatch[1:]

	players = removeDuplicates(players)
	for i := 0; i < len(players); i++ {
		player := players[i]
		players[i] = player[2:]
	}

	return players

}

func removeDuplicates(arr []string) []string {
	unique := make(map[string]bool)
	for _, val := range arr {
		unique[val] = true
	}
	result := make([]string, 0, len(unique))
	for key := range unique {
		result = append(result, key)
	}
	return result
}
