package quakeLogUsecase

import (
	"fmt"
	entity "quakelog/internal/entity"
	quakeLog "quakelog/internal/platform/http"
	"regexp"
	"sort"
	"strings"
)

type QuakeLogUsecase interface {
	GroupedByTypeOfDeath() (map[string]entity.QuakeLogKills, error)
	GroupByGame() (map[string]entity.QuakeLog, error)
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

func (ref quakeLogUsecase) GroupByGame() (map[string]entity.QuakeLog, error) {
	log, err := ref.quakeLog.Get()
	if err != nil {
		return map[string]entity.QuakeLog{}, err
	}

	var quakeLogsMatch []entity.QuakeLog

	games := strings.Split(log, "InitGame")
	for _, game := range games {
		players := ref.getPlayers(game)
		if len(players) == 0 {
			continue
		}

		kills := make(map[string]int, len(players))
		for _, player := range players {
			kills[player] = ref.getKillsByPlayer(game, player)
		}

		groupByGame := entity.QuakeLog{
			Players:    players,
			TotalKills: strings.Count(game, "killed"),
			Kills:      kills,
		}

		quakeLogsMatch = append(quakeLogsMatch, groupByGame)

	}

	return ref.orderQuakeLogByTotalKills(quakeLogsMatch), err
}

func (ref quakeLogUsecase) orderQuakeLogByTotalKills(quakeLogsMatch []entity.QuakeLog) map[string]entity.QuakeLog {
	sort.Slice(quakeLogsMatch, func(i, j int) bool {
		return quakeLogsMatch[i].TotalKills > quakeLogsMatch[j].TotalKills
	})

	quakeLogsMap := make(map[string]entity.QuakeLog)
	for key, quakeLog := range quakeLogsMatch {
		quakeLogsMap[ref.getNameGame(key)] = quakeLog

	}

	return quakeLogsMap
}

func (ref quakeLogUsecase) getNameGame(key int) string {
	if len(string(rune(key))) == 1 {
		return fmt.Sprintf("game-0%v", key+1)
	}

	return fmt.Sprintf("game-%v", key+1)
}

func (ref quakeLogUsecase) getKillsByPlayer(game string, player string) int {
	regexKill := regexp.MustCompile(fmt.Sprintf("Kill(.*)%s killed", player))
	playerKillNumbers := regexKill.FindAllString(game, 99)

	regexWasKilled := regexp.MustCompile(fmt.Sprintf("Kill(.*)<world> killed %s", player))
	playerWasKilledNumbers := regexWasKilled.FindAllString(game, 99)

	return len(playerKillNumbers) - len(playerWasKilledNumbers)
}

func (ref quakeLogUsecase) getPlayers(game string) []string {
	re := regexp.MustCompile(`n\\([^\\]+)`)
	quakeMatch := re.FindAllString(game, 99)

	if len(quakeMatch) == 0 {
		return nil
	}

	players := quakeMatch[1:]

	players = ref.removeDuplicates(players)
	for i := 0; i < len(players); i++ {
		player := players[i]
		players[i] = player[2:]
	}

	return players

}

func (ref quakeLogUsecase) removeDuplicates(arr []string) []string {
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
