package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"gameserver/config"
	"gameserver/models"
	"gameserver/utils"
)

// 주사위 굴리기 핸들러
func RollDice(c echo.Context) error {
	// =============== 공통코드 =================================================================
	// 사용자 ID 가져오기 (AuthMiddleware에서 context에 저장됨)
	username, err := utils.GetUserIDFromToken(c)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "인증 실패"})
	}

	user, err := models.GetUser(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "사용자 정보 없음"})
	}

	coins := user.Coins
	if coins < 30 {
		return c.JSON(http.StatusForbidden, echo.Map{"error": "코인이 부족합니다"})
	}
	// ===========================================================================================

	// 주사위 굴리기 (1~6)
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1

	// 보상 계산
	reward := dice * 10
	newCoins := coins - 30 + reward

	key := "user:" + username
	// Redis에 업데이트
	err = config.Redis.HSet(config.Ctx, key, "coins", newCoins).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "코인 업데이트 실패"})
	}

	// 결과 응답
	return c.JSON(http.StatusOK, echo.Map{
		"message": "주사위 굴리기 성공",
		"dice":    dice,
		"reward":  reward,
		"coins":   newCoins,
	})
}

// 주사위 2개굴리기 핸들러
// 규칙
/*
- Small (2~6) : 2배
- Big(8~12) : 2배
- 11,22,33,44,55,66 : 6배
- 두수의합이짝수 : 2배
- 두수의합이홀수: 2배
*/
func DoubleRollDice(c echo.Context) error {
	// =============== 공통코드 =================================================================
	// 사용자 ID 가져오기 (AuthMiddleware에서 context에 저장됨)
	username, err := utils.GetUserIDFromToken(c)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "인증 실패"})
	}

	user, err := models.GetUser(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "사용자 정보 없음"})
	}

	coins := user.Coins
	// if coins < 30 {
	// 	return c.JSON(http.StatusForbidden, echo.Map{"error": "코인이 부족합니다"})
	// }
	// ===========================================================================================

	mapdice := map[string]int{}

	var req struct {
		Bets map[string]int `json:"bets"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	// 총 배팅금 계산
	totalBet := 0
	for _, amount := range req.Bets {
		if amount <= 0 {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid bet amount"})
		}
		totalBet += amount
	}

	if coins < totalBet {
		return c.JSON(http.StatusForbidden, echo.Map{"error": "코인이 부족합니다"})
	}

	user.Coins -= totalBet

	// 주사위 굴림
	rand.Seed(time.Now().UnixNano())
	d1 := rand.Intn(6) + 1
	d2 := rand.Intn(6) + 1
	sum := d1 + d2

	fmt.Println("d1:", d1, " d2:", d2)

	// 조건 평가
	reward := 0
	for condition, amount := range req.Bets {
		switch condition {
		case "double": // 같은 숫자 : 5배
			if d1 == d2 {
				reward += amount * 5
				mapdice[condition] = amount * 5
				fmt.Println("===double", reward)
			}
		case "dice11": // 11 : 32배
			if d1 == 1 && d2 == 1 {
				reward += amount * 32
				mapdice[condition] = amount * 32
				fmt.Println("===dice11", reward)
			}
		case "dice22": // 22 : 32배
			if d1 == 2 && d2 == 2 {
				reward += amount * 32
				mapdice[condition] = amount * 32
				fmt.Println("===dice22", reward)
			}
		case "dice33": // 33 : 32배
			if d1 == 3 && d2 == 3 {
				reward += amount * 32
				mapdice[condition] = amount * 32
				fmt.Println("===dice33", reward)
			}
		case "dice44": // 44 : 32배
			if d1 == 4 && d2 == 4 {
				reward += amount * 32
				mapdice[condition] = amount * 32
				fmt.Println("===dice44", reward)
			}
		case "dice55": // 55 : 32배
			if d1 == 5 && d2 == 5 {
				reward += amount * 32
				mapdice[condition] = amount * 32
				fmt.Println("===dice55", reward)
			}
		case "dice66": // 66 : 32배
			if d1 == 6 && d2 == 6 {
				reward += amount * 32
				mapdice[condition] = amount * 32
				fmt.Println("===dice66", reward)
			}
		case "small": // 합이 2~6 : 2배
			if sum >= 2 && sum <= 6 {
				reward += amount * 2
				mapdice[condition] = amount * 2
				fmt.Println("===small", reward)
			}
		case "big": // 합이 8~12 : 2배
			if sum >= 8 && sum <= 12 {
				reward += amount * 2
				mapdice[condition] = amount * 2
				fmt.Println("===big", reward)
			}
		case "even": // 합이 짝수 : 2배
			if sum%2 == 0 {
				reward += amount * 2
				mapdice[condition] = amount * 2
				fmt.Println("===even", reward)
			}
		case "odd": // 합이 홀수 ; 2배
			if sum%2 != 0 {
				reward += amount * 2
				mapdice[condition] = amount * 2
				fmt.Println("===odd", reward)
			}
		default:
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid bet condition: " + condition})
		}
	}

	// 보상 추가
	user.Coins += reward

	// 저장
	// newData, _ := json.Marshal(user)
	// rdb.Set(ctx, "user:"+userID, newData, 0)
	key := "user:" + username
	// Redis에 업데이트
	err = config.Redis.HSet(config.Ctx, key, "coins", user.Coins).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "코인 업데이트 실패"})
	}

	jsonbytes, err := json.Marshal(mapdice)
	if err != nil {
		fmt.Println("JSON 변환 실패:", err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"dice_1":        d1,
		"dice_2":        d2,
		"sum":           sum,
		"reward":        reward,
		"current_coins": user.Coins,
		"bets":          string(jsonbytes),
		//"bets":          req.Bets,
	})

}

func DailyReward(c echo.Context) error {
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "인증 실패"})
	}

	user, err := models.GetUser(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "사용자 정보 없음"})
	}

	// 오늘 날짜 문자열
	today := time.Now().Format("2006-01-02")

	if user.LastLogin == today {
		// 이미 오늘 보상 받음
		return c.JSON(http.StatusOK, echo.Map{
			"message": "오늘 이미 보상을 받았습니다",
			"coins":   user.Coins,
		})
	}

	// 보상 지급
	const reward = 100
	user.Coins += reward
	user.LastLogin = today

	err = models.SaveUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "보상 저장 실패"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "일일 보상을 받았습니다",
		"reward":  reward,
		"coins":   user.Coins,
	})
}

// 보상
// 100코인배팅 => 10,40.90,160,250,360 (확률 : 50,30,10,6,3,1)
// 한번 던지는 비용은 50
func Roulette(c echo.Context) error {
	// =============== 공통코드 =================================================================
	// 사용자 ID 가져오기 (AuthMiddleware에서 context에 저장됨)
	username, err := utils.GetUserIDFromToken(c)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "인증 실패"})
	}

	user, err := models.GetUser(username)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "사용자 정보 없음"})
	}

	bet := 50 // 룰렛 배팅 비용
	if user.Coins < bet {
		return c.JSON(http.StatusForbidden, echo.Map{"error": "코인이 부족합니다"})
	}
	// ===========================================================================================

	rewards := []int{10, 40, 90, 160, 250, 360}
	weights := []int{50, 30, 10, 6, 3, 1} // 총합 = 100 → 문제 없고, 500이어도 OK
	reward := _roulette(rewards, weights) // 룰렛 돌리기
	fmt.Println("룰렛 보상:", reward)

	user.Coins += reward - bet // 배팅 비용 차감 후 보상 추가
	key := "user:" + username
	// Redis에 업데이트
	err = config.Redis.HSet(config.Ctx, key, "coins", user.Coins).Err()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "코인 업데이트 실패"})
	}
	// 결과 응답
	if reward > 0 {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "룰렛 돌리기 성공",
			"reward":  reward,
			"coins":   user.Coins,
		})
	}

	return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Roulette failed"})
}

// 룰렛 돌리기 함수
// rewards: 보상 목록, weights: 각 보상의 가중치
// 외부에서 확률표를 제공받아야 함
func _roulette(rewards []int, weights []int) int {
	if len(rewards) != len(weights) || len(rewards) == 0 {
		return 0 // 잘못된 입력 처리
	}

	// 전체 가중치 합 계산
	totalWeight := 0
	for _, w := range weights {
		totalWeight += w
	}

	// 랜덤 시드 설정
	rand.Seed(time.Now().UnixNano())

	// 1~100 중 무작위 정수 생성
	r := rand.Intn(totalWeight) + 1

	// 누적 확률 기반 선택
	sum := 0
	for i, w := range weights {
		sum += w
		if r <= sum {
			fmt.Println("룰렛 보상 선택:", rewards[i], "확률:", w)
			return rewards[i]
		}
	}
	return 0 // 이론적으로 도달하지 않음
}
