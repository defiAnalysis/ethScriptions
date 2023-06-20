package utils

import (
	"crypto/rand"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/shopspring/decimal"
	"hash/fnv"
	"math"
	"math/big"
	time2 "time"
)

func GetRare(dna string) int64 {
	if dna == "0" {
		return 0
	}

	dnaDeci, _ := decimal.NewFromString(dna)
	baseDeci := decimal.NewFromFloat(math.Pow(10, 26))

	rareParser := dnaDeci.Div(baseDeci)
	rareParserInt := rareParser.BigInt().Int64()

	if rareParserInt < 5225 {
		return 1
	} else if rareParserInt < 7837 {
		return 2
	} else if rareParserInt < 8707 {
		return 3
	} else if rareParserInt < 9360 {
		return 4
	} else if rareParserInt < 9708 {
		return 5
	} else {
		return 6
	}
}

func GetLevel(_exp int64) int64 {
	if _exp < 200 {
		return 1
	} else if _exp < 700 {
		return 2
	} else if _exp < 2000 {
		return 3
	} else if _exp < 4000 {
		return 4
	} else if _exp < 8000 {
		return 5
	} else {
		return 6
	}
}

func GetWinRate(_tokenid int64, _monster int64) float64 {
	winRateRnd, err := Random(_tokenid, _monster, 1)
	if err != nil {
		beego.Error("====Random===", err)
		return 0
	}

	//if _monster == 0 {
	//	return winRateRnd + 70
	//} else if _monster == 1 {
	//	return winRateRnd + 60
	//} else if _monster == 2 {
	//	return winRateRnd + 50
	//} else if _monster == 3 {
	//	return winRateRnd + 10
	//}
	if _monster == 0 {
		return winRateRnd + 60
	} else if _monster == 1 {
		return winRateRnd + 50
	} else if _monster == 2 {
		return winRateRnd + 40
	} else if _monster == 3 {
		return winRateRnd + 10
	}

	return 0
}

func Random(_tokenid int64, _monster int64, len float64) (float64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(_tokenid))
	if err != nil {
		beego.Error("rand Int error", err)
		return 0, err
	}

	time := time2.Now().Unix()
	strFmt := fmt.Sprintf("%d%d%d%d", time, _tokenid, _monster, n)
	beego.Info("====strFmt====", strFmt)
	hx := Hash(strFmt)

	_data := math.Pow(10, len)
	beego.Info("====_data====", _data)
	//
	return math.Mod(float64(hx), _data), nil
}

func BattleReward(level, rate int64, winRate float64) float64 {
	x := 0.182

	beego.Info("level===", level)
	beego.Info("rate===", rate)
	beego.Info("winRate===", winRate)

	winRate += 10

	reward := x * float64(rate) * float64(level) * (100 - winRate)

	if level >= 1 && level <= 3 {
		reward = reward * 2
	}

	return reward
}

func Hash(input string) uint32 {
	//c := md5.New()
	//c.Write([]byte(input))
	//bytes := c.Sum(nil)
	//return bytes
	h := fnv.New32a()
	h.Write([]byte(input))
	return h.Sum32()
}

func GetEXP(MonsterLevel int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(10))
	if MonsterLevel == 0 {
		return 6 + n.Int64()%3
	} else if MonsterLevel == 1 {
		return 18 + n.Int64()%6
	} else if MonsterLevel == 2 {
		return 36 + n.Int64()%9
	} else if MonsterLevel == 3 {
		return 120 + n.Int64()%15
	}

	return 0
}
