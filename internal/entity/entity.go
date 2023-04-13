package entity

type MeansOfDeath string

const (
	ModeUnknown       MeansOfDeath = "MOD_UNKNOWN"
	ModeShotGum       MeansOfDeath = "MOD_SHOTGUN"
	ModeGauntlet      MeansOfDeath = "MOD_GAUNTLET"
	ModeMachineGun    MeansOfDeath = "MOD_MACHINEGUN"
	ModeGrenade       MeansOfDeath = "MOD_GRENADE"
	ModeGrenadeSplash MeansOfDeath = "MOD_GRENADE_SPLASH"
	ModeRocket        MeansOfDeath = "MOD_ROCKET"
	ModeRocketSplash  MeansOfDeath = "MOD_ROCKET_SPLASH"
	ModePlasma        MeansOfDeath = "MOD_PLASMA"
	ModePlasmaSplash  MeansOfDeath = "MOD_PLASMA_SPLASH"
	ModeRailGun       MeansOfDeath = "MOD_RAILGUN"
	ModeLightning     MeansOfDeath = "MOD_LIGHTNING"
	ModeBfg           MeansOfDeath = "MOD_BFG"
	ModeBfgSplash     MeansOfDeath = "MOD_BFG_SPLASH"
	ModeWater         MeansOfDeath = "MOD_WATER"
	ModeLava          MeansOfDeath = "MOD_LAVA"
	ModeCrush         MeansOfDeath = "MOD_CRUSH"
	ModeTelefrag      MeansOfDeath = "MOD_TELEFRAG"
	ModeFalling       MeansOfDeath = "MOD_FALLING"
	ModeSuicide       MeansOfDeath = "MOD_SUICIDE"
	ModeTargetLaser   MeansOfDeath = "MOD_TARGET_LASER"
	ModeTargetHurt    MeansOfDeath = "MOD_TRIGGER_HURT"
)

var AllMeansOFDeath = map[MeansOfDeath]string{
	ModeUnknown:       "MOD_UNKNOWN",
	ModeShotGum:       "MOD_SHOTGUN",
	ModeGauntlet:      "MOD_GAUNTLET",
	ModeMachineGun:    "MOD_MACHINEGUN",
	ModeGrenade:       "MOD_GRENADE",
	ModeGrenadeSplash: "MOD_GRENADE_SPLASH",
	ModeRocket:        "MOD_ROCKET",
	ModeRocketSplash:  "MOD_ROCKET_SPLASH",
	ModePlasma:        "MOD_PLASMA",
	ModePlasmaSplash:  "MOD_PLASMA_SPLASH",
	ModeRailGun:       "MOD_RAILGUN",
	ModeLightning:     "MOD_LIGHTNING",
	ModeBfg:           "MOD_BFG",
	ModeBfgSplash:     "MOD_BFG_SPLASH",
	ModeWater:         "MOD_WATER",
	ModeLava:          "MOD_LAVA",
	ModeCrush:         "MOD_CRUSH",
	ModeTelefrag:      "MOD_TELEFRAG",
	ModeFalling:       "MOD_FALLING",
	ModeSuicide:       "MOD_SUICIDE",
	ModeTargetLaser:   "MOD_TARGET_LASER",
	ModeTargetHurt:    "MOD_TRIGGER_HURT",
}

type QuakeLog struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

type QuakeLogKills struct {
	KillsByMeans map[string]int `json:"kills_by_means "`
}
