package Day13

import (
	"AOD2024/Utils"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
)

const day = 13
const title = "--- Day 13: Claw Contraption ---"

func extractUnit(re *regexp.Regexp, line string) (int, int) {
	matches := re.FindAllStringSubmatch(line, -1)
	x, _ := strconv.Atoi(matches[0][1])
	y, _ := strconv.Atoi(matches[1][1])
	return x, y
}

func totalToken(buttonAX, buttonAY, buttonBX, buttonBY, X, Y int) int {
	a := 1
	for buttonAX*a <= X {
		fmt.Println(buttonAX * a)
		if (X-buttonAX*a)%buttonBX == 0 {
			leftB := (X - buttonAX*a) / buttonBX
			if buttonAY*a+buttonBY*leftB == Y {
				return a*3 + leftB
			}
		}
		a++
	}
	return 0
}

func ExtendedGCD(a, b int64) (int64, int64, int64) {
	if b == 0 {
		return a, 1, 0
	}
	gcd, x1, y1 := ExtendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return gcd, x, y
}

func SolveDiophantine(a, b, c int64) (int64, int64, bool) {
	gcd, x, y := ExtendedGCD(a, b)
	if c%gcd != 0 {
		return 0, 0, false
	}
	scale := c / gcd
	return x * scale, y * scale, true
}

func totalTokenPart2(buttonAX, buttonAY, buttonBX, buttonBY, x, y int) int64 {
	bigButtonAX := big.NewInt(int64(buttonAX))
	bigButtonAY := big.NewInt(int64(buttonAY))
	bigButtonBX := big.NewInt(int64(buttonBX))
	bigButtonBY := big.NewInt(int64(buttonBY))
	bigX := big.NewInt(int64(x))
	bigY := big.NewInt(int64(y))

	// Compute X and Y

	X := big.NewInt(int64(Utils.Pow10(13)))
	X.Add(X, bigX)

	Y := big.NewInt(int64(Utils.Pow10(13)))
	Y.Add(Y, bigY)

	CoeffB := new(big.Int).Sub(
		new(big.Int).Mul(bigButtonBX, bigButtonAY),
		new(big.Int).Mul(bigButtonBY, bigButtonAX),
	)

	// Calculate ConstC = X*buttonAY - Y*buttonAX
	ConstC := new(big.Int).Sub(
		new(big.Int).Mul(X, bigButtonAY),
		new(big.Int).Mul(Y, bigButtonAX),
	)

	if new(big.Int).Mod(ConstC, CoeffB).Sign() != 0 {
		return 0
	}
	b := new(big.Int).Div(ConstC, CoeffB)

	ConstA := new(big.Int).Sub(X, new(big.Int).Mul(bigButtonBX, b))

	if new(big.Int).Mod(ConstA, bigButtonAX).Sign() != 0 {
		return 0
	}
	a := new(big.Int).Div(ConstA, bigButtonAX)

	result := new(big.Int).Add(new(big.Int).Mul(big.NewInt(3), a), b)
	return result.Int64()
}

func Part1(submit bool) {
	lines, err := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	if submit {
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/%v.txt", day))
		if err != nil {
			fmt.Println("Input not already fetch today")
			Utils.ReadInput(day)
		}
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/%v.txt", day))
	}
	ans := 0
	re := regexp.MustCompile(`[XY][+=](?P<i>\d+)`)
	for i := 0; i < len(lines); i = i + 4 {
		buttonAX, buttonAY := extractUnit(re, lines[i])
		buttonBX, buttonBY := extractUnit(re, lines[i+1])
		X, Y := extractUnit(re, lines[i+2])
		ans += totalToken(buttonAX, buttonAY, buttonBX, buttonBY, X, Y)
	}
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem .", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 1, ans)
	}
}

func Part2(submit bool) {
	lines, err := Utils.ReadFileLines(fmt.Sprintf("Day%v/sample.txt", day))
	if submit {
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/%v.txt", day))
		if err != nil {
			fmt.Println("Input not already fetch today")
			Utils.ReadInput(day)
		}
		lines, err = Utils.ReadFileLines(fmt.Sprintf("Input/%v.txt", day))
	}
	ans := int64(0)
	re := regexp.MustCompile(`[XY][+=](?P<i>\d+)`)
	for i := 0; i < len(lines); i = i + 4 {
		buttonAX, buttonAY := extractUnit(re, lines[i])
		buttonBX, buttonBY := extractUnit(re, lines[i+1])
		x, y := extractUnit(re, lines[i+2])
		ans += totalTokenPart2(buttonAX, buttonAY, buttonBX, buttonBY, x, y)
	}
	fmt.Println(Utils.Blue + fmt.Sprintf("%s", title) + Utils.Reset)
	fmt.Println(Utils.Yellow + fmt.Sprintf("Answers this part: %v. Let's submit this problem.", ans) + Utils.Reset)
	if submit {
		Utils.Submit(day, 2, int(ans))
	}
}
