package quakeLogUsecase

import (
	"errors"
	quakeLogMock "quakelog/test/platform"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var anyError = errors.New("Error")

func TestGroupByGame_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	log := quakeLogMock.NewMockQuakeLog(controller)
	service := NewQuakeLogUsecase(log)

	game := `
	0:00 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0
	0:25 ClientConnect: 2
	0:25 ClientUserinfoChanged: 2 n\Dono da Bola\t\0\model\sarge/krusade\hmodel\sarge/krusade\g_redteam\\g_blueteam\\c1\5\c2\5\hc\95\w\0\l\0\tt\0\tl\0
	0:27 ClientUserinfoChanged: 2 n\Mocinha\t\0\model\sarge\hmodel\sarge\g_redteam\\g_blueteam\\c1\4\c2\5\hc\95\w\0\l\0\tt\0\tl\0  
	1:41 Kill: 1022 2 19: <world> killed Dono da Bola by MOD_FALLING 
	1:47 ShutdownGame:
	`

	log.EXPECT().Get().Return(game, nil)
	_, err := service.GroupByGame()

	assert.NoError(t, err)

}

func TestGroupByGame_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	log := quakeLogMock.NewMockQuakeLog(controller)
	service := NewQuakeLogUsecase(log)

	game := `
	`

	log.EXPECT().Get().Return(game, anyError)
	_, err := service.GroupByGame()

	assert.Error(t, err)

}

func TestGroupedByTypeOfDeath_Success(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	log := quakeLogMock.NewMockQuakeLog(controller)
	service := NewQuakeLogUsecase(log)

	game := `
	0:00 InitGame: \sv_floodProtect\1\sv_maxPing\0\sv_minPing\0\sv_maxRate\10000\sv_minRate\0\sv_hostname\Code Miner Server\g_gametype\0\sv_privateClients\2\sv_maxclients\16\sv_allowDownload\0\dmflags\0\fraglimit\20\timelimit\15\g_maxGameClients\0\capturelimit\8\version\ioq3 1.36 linux-x86_64 Apr 12 2009\protocol\68\mapname\q3dm17\gamename\baseq3\g_needpass\0
	0:25 ClientConnect: 2
	0:25 ClientUserinfoChanged: 2 n\Dono da Bola\t\0\model\sarge/krusade\hmodel\sarge/krusade\g_redteam\\g_blueteam\\c1\5\c2\5\hc\95\w\0\l\0\tt\0\tl\0
	0:27 ClientUserinfoChanged: 2 n\Mocinha\t\0\model\sarge\hmodel\sarge\g_redteam\\g_blueteam\\c1\4\c2\5\hc\95\w\0\l\0\tt\0\tl\0  
	1:41 Kill: 1022 2 19: <world> killed Dono da Bola by MOD_FALLING 
	1:47 ShutdownGame:
	`

	log.EXPECT().Get().Return(game, nil)
	_, err := service.GroupedByTypeOfDeath()

	assert.NoError(t, err)

}

func TestGroupedByTypeOfDeath_Error(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	log := quakeLogMock.NewMockQuakeLog(controller)
	service := NewQuakeLogUsecase(log)

	game := `
	`

	log.EXPECT().Get().Return(game, anyError)
	_, err := service.GroupedByTypeOfDeath()

	assert.Error(t, err)

}
