package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func main() {

	product := "Kayak"
	fmt.Println("Product:", product)

	f16_4()
	f16_ex()
	f16_5()
	f16_6()
	f16_7()
	f16_8()
	f16_9()
	f16_10()
	f16_11()
	f16_12()
	f16_13()
	f16_14()
	f16_15()
	f16_16()
	f16_17()
	f16_18()
	f16_19()
	f16_20()
	f16_21()
	f16_22()
	f16_23()
	f16_24()
	f16_25()
	f16_26()
	f16_27()
	f16_28()
	f16_29()
	f16_30()
	f16_31()
	f16_32()
	f16_33()
}

func f16_4() {
	product := "Kayak"
	// yak문자열을 포함하고 있으므로 true
	fmt.Println("Contains:", strings.Contains(product, "yak"))
	// a,b,c 중에 한개라도 포함하고 있기때문에 true
	fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
	// 문자열 product가 특정 rune을 포함하고 있으면 true
	fmt.Println("ContainsRune:", strings.ContainsRune(product, 'K'))
	// 대소문자를 구분하지 않고 비교하는데, 서로 문자열이 동일하면 true를 반환한다.
	fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
	// Product 문자열이 "Ka"로 시작하는 경우라면 true를 반환한다.
	fmt.Println("HasPrefix:", strings.HasPrefix(product, "Ka"))
	// Product 문자열이 "yak"로 끝나는 경우라면 true를 반환한다.
	fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))
}

func f16_ex() {
	price := "€100"
	// strings 패키지의 모든함수는 bytes에서 동작하는 대응함수가 있다.
	fmt.Println("Strings Prefix:", strings.HasPrefix(price, "€"))
	fmt.Println("Bytes Prefix:", bytes.HasPrefix([]byte(price), []byte{226, 130}))
}

func f16_5() {
	fmt.Println("========== f16_5 ==================")
	description := "A boat for sailing"
	fmt.Println("Original:", description)
	// 각 단어의 첫번째 문자가 대문자이고, 나머지가 소문자
	fmt.Println("Title:", strings.Title(description))
}

func f16_6() {
	fmt.Println("========== f16_6 ==================")

	specialChar := "\u01c9"
	fmt.Println("Original:", specialChar, []byte(specialChar))
	upperChar := strings.ToUpper(specialChar)
	fmt.Println("Upper:", upperChar, []byte(upperChar))
	titleChar := strings.ToTitle(specialChar)
	fmt.Println("Title:", titleChar, []byte(titleChar))
}

func f16_7() {
	fmt.Println("========== f16_7 ==================")
	product := "Kayak"
	for _, char := range product {
		fmt.Println(string(char), "Upper case:", unicode.IsUpper(char))
	}
}

func f16_8() {
	fmt.Println("========== f16_8 ==================")

	description := "A boat for one person"
	// 해당 문자열이 몇개인지 카운트한다.
	fmt.Println("Count:", strings.Count(description, "o"))
	// 해당 문자열이 처음으로 나타나는 인덱스를 반환한다.(없으면 -1)
	fmt.Println("Index:", strings.Index(description, "o"))
	// 해당 문자열이 마지막으로 나타나는 인덱스를 반환한다.(없으면 -1)
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	// 해당 문자열이 모든 문자에 대해서 처음으로 나타나는 인덱스를 반환한다.(없으면 -1)
	// abcd중 b가처음으로 나타나서 2가 나온다.
	fmt.Println("IndexAny:", strings.IndexAny(description, "abcd"))
	// 해당 문자열이 마지막으로 나타나는 인덱스를 반환한다.(없으면 -1)
	fmt.Println("LastIndex:", strings.LastIndex(description, "o"))
	// 해당 문자열이 모든 문자에 대해서 마지막으로 나타나는 인덱스를 반환한다.(없으면 -1)
	fmt.Println("LastIndexAny:", strings.LastIndexAny(description, "abcd"))
}

func f16_9() {
	fmt.Println("======== list_16_9 ========")
	description := "A boat for one person"
	// 익명함수(람다)를 정의해서 isLetterB 라는 변수에 담는다.
	// r이 B,b라면 true이다.
	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}
	// IndexFunc() : 사용자정의 함수를 이용해서 지정한 함수가 true를 반환하는 문자열의 인덱스를 출력한다.
	fmt.Println("IndexFunc:", strings.IndexFunc(description, isLetterB))
}

func f16_10() {
	fmt.Println("======== list_16_10 ========")
	description := "A boat for one person"
	// 공백을 기준으로 문자열을 분리하여 문자열 슬라이스를 반환한다.
	splits := strings.Split(description, " ")
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}
	// 공백을 포함하여 문자열을 분리하여 문자열 슬라이스를 반환한다.
	splitsAfter := strings.SplitAfter(description, " ")
	for _, x := range splitsAfter {
		fmt.Println("SplitAfter >>" + x + "<<")
	}
}

func f16_11() {
	fmt.Println("======== list_16_11 ========")
	description := "A boat for one person"
	// 결과를 3개까지만 분할한다.
	splits := strings.SplitN(description, " ", 3)
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}
	// splitsAfter := strings.SplitAfter(description, " ")
	// for _, x := range splitsAfter {
	//     fmt.Println("SplitAfter >>" + x + "<<")
	// }
}

func f16_12() {
	fmt.Println("======== list_16_12 ========")
	description := "This  is  double  spaced"
	// 공백이 2개이상일경우 제대로 문자열을 분할하지 않는 문제가 있다.
	splits := strings.SplitN(description, " ", 3)
	for _, x := range splits {
		fmt.Println("Split >>" + x + "<<")
	}
}

func f16_13() {
	fmt.Println("======== list_16_13 ========")
	description := "This  is  double  spaced"
	// Split 은 구분자를 정확히 그대로 사용하기때문에 공백이 2개 이상이면 문제가 된다.
	// Fields는 연속된 공백을 하나의 구분자로 취급한다. (중복된 공백을 무시한다.)
	splits := strings.Fields(description)
	for _, x := range splits {
		fmt.Println("Field >>" + x + "<<")
	}
}

func f16_14() {
	fmt.Println("======== list_16_14 ========")
	description := "This  is  double  spaced"
	splitter := func(r rune) bool {
		return r == ' '
	}
	// 사용자 정의 함수를 사용하여, 구분자에 대한 처리를 직접 수행할수있다.
	splits := strings.FieldsFunc(description, splitter)
	for _, x := range splits {
		fmt.Println("Field >>" + x + "<<")
	}
}

func f16_15() {
	fmt.Println("======== list_16_15 ========")
	username := " Alice"
	// 문자열의 앞,뒤 공백을 모두 제거한다.
	trimmed := strings.TrimSpace(username)
	fmt.Println("Trimmed:", ">>"+trimmed+"<<")
}

func f16_16() {
	fmt.Println("======== list_16_16 ========")
	description := "A boat for one person"
	// s 문자열의 앞뒤에서, cutset에 포함된 문자를 모두 자른다.
	trimmed := strings.Trim(description, "Asno ")
	fmt.Println("Trimmed:", trimmed)
}

func f16_17() {
	fmt.Println("======== list_16_17 ========")
	description := "A boat for one person"
	// 문자열에서 앞부분이 "A boat "과 일치하면 해당 문자열을 제거한다.
	prefixTrimmed := strings.TrimPrefix(description, "A boat ")
	wrongPrefix := strings.TrimPrefix(description, "A hat ")
	fmt.Println("Trimmed:", prefixTrimmed)
	fmt.Println("Not trimmed:", wrongPrefix)
}

func f16_18() {
	fmt.Println("======== list_16_18 ========")
	description := "A boat for one person"
	trimmer := func(r rune) bool {
		return r == 'A' || r == 'n'
	}
	// 사용자 정의 함수를 사용하여 해당 조건이 만족하면 문자열의 앞뒤에서 문자를 제거한다.
	trimmed := strings.TrimFunc(description, trimmer)
	fmt.Println("Trimmed:", trimmed)
}

func f16_19() {
	fmt.Println("======== list_16_19 ========")
	text := "It was a boat. A small boat."
	// "boat" 를 "canoe" 으로 변경한다. 최대 1회까지 가능
	replace := strings.Replace(text, "boat", "canoe", 1)
	// 문자열에서 "boat" 를  "truck"으로 모두 교체한다.
	replaceAll := strings.ReplaceAll(text, "boat", "truck")
	fmt.Println("Replace:", replace)
	fmt.Println("Replace All:", replaceAll)
}

func f16_20() {
	fmt.Println("======== list_16_20 ========")
	text := "It was a boat. A small boat."
	// b가 오면 c를 리턴하는 함수
	mapper := func(r rune) rune {
		if r == 'b' {
			return 'c'
		}
		return r
	}
	// Map()함수를 사용하여 b를 c로 치환한다.
	// 책에서는 매핑함수라고 하는데 사용자 정의함수와 같은뜻아닌가??
	mapped := strings.Map(mapper, text)
	fmt.Println("Mapped:", mapped)
}

func f16_21() {
	fmt.Println("======== list_16_21 ========")
	text := "It was a boat. A small boat."
	// Replacer를 만든다.
	// 인자는 pair로 되어야한다.(쌍으로 있어야한다.)
	replacer := strings.NewReplacer("boat", "kayak", "small", "huge")
	// Replacer를 사용하여 문자열 text를 치환한다. (boar => kayak, small => huge)
	replaced := replacer.Replace(text)
	fmt.Println("Replaced:", replaced)
}

func f16_22() {
	fmt.Println("======== list_16_22 ========")
	text := "It was a boat. A small boat."
	// text를 공백을 기준으로 분리하여 슬라이스([]string)으로 변환한다.
	elements := strings.Fields(text)
	// 문자열 슬라이스를 sep구분자로 모두 연결하여 이어 붙인다.
	joined := strings.Join(elements, "--")
	fmt.Println("Joined:", joined)
}

func f16_23() {
	fmt.Println("======== list_16_23 ========")
	text := "It was a boat. A small boat."
	// strings.Builder 를 사용하여 문자열을 합친다.
	// 문자열을 반복적으로 합칠때 효율적이다.
	var builder strings.Builder
	// strings.Fields를 사용하여 문자열을 분리하여 슬라이스로 만든다.
	for _, sub := range strings.Fields(text) {
		if sub == "small" {
			builder.WriteString("very ")
		}
		builder.WriteString(sub)
		builder.WriteRune(' ')
	}
	fmt.Println("String:", builder.String())
}

func f16_24() {
	fmt.Println("======== list_16_24 ========")
	description := "A boat for one person"
	// [A-z] → 대문자 A~Z와 소문자 a~z 범위 내의 문자 중 하나
	// [A~z까지의 문자들중하나]oat => 이 문자열이 있으면 true
	match, err := regexp.MatchString("[A-z]oat", description)
	if err == nil {
		fmt.Println("Match:", match)
	} else {
		fmt.Println("Error:", err)
	}
}

func f16_25() {
	fmt.Println("======== list_16_25 ========")
	// 정규식 패턴을 컴파일해서 *Regexp 타입의 객체를 반환한다.
	pattern, compileErr := regexp.Compile("[A-z]oat")
	description := "A boat for one person"
	question := "Is that a goat?"
	preference := "I like oats"
	if compileErr == nil {
		// 위에서 구한 정규식객체를 사용하여 3개의 문자열을 각각 패턴 일치여부를 검사한다.
		fmt.Println("Description:", pattern.MatchString(description))
		fmt.Println("Question:", pattern.MatchString(question))
		fmt.Println("Preference:", pattern.MatchString(preference))
	} else {
		fmt.Println("Error:", compileErr)
	}
}

// 인엑스범위로 문자열을 잘라내는 함수
func getSubstring_26(s string, indices []int) string {
	return string(s[indices[0]:indices[1]])
}

func f16_26() {
	fmt.Println("======== list_16_26 ========")
	// "K" + 소문자 알파벳 4개 (예: Kayak)
	// | => or 연산자
	// "[A-z]oat" → 알파벳 한 글자 + "oat" (boat, goat 등)
	pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description := "Kayak. A boat for one person."
	// 첫 번째 매치된 문자열의 [시작, 끝] 인덱스 반환
	// Kayak 문자열이 조건에 만족해서 [0,5] 반환
	firstIndex := pattern.FindStringIndex(description)
	// 문자열 전체에서 모든 매치 위치를 찾음 ,  (-1 : 갯수제한없음)
	// [[0,5],[8,12]]
	allIndices := pattern.FindAllStringIndex(description, -1)
	// [0,5]를 기준으로 문자열을 자른다.
	fmt.Println("First index", firstIndex[0], "-", firstIndex[1],
		"=", getSubstring_26(description, firstIndex))
	// allIndices 안에 있는 값들을 순회하며 해당 문자열을 잘라서 출력한다.
	for i, idx := range allIndices {
		fmt.Println("Index", i, "=", idx[0], "-",
			idx[1], "=", getSubstring_26(description, idx))
	}
}

func f16_27() {
	fmt.Println("======== list_16_27 ========")
	pattern := regexp.MustCompile("K[a-z]{4}|[A-z]oat")
	description := "Kayak. A boat for one person."
	// 해당 조건에 맞는 첫번째 문자열을 반환한다. string
	firstMatch := pattern.FindString(description)
	// 해당 정규식에 해당하는 모든 문자열을 반환한다. []string
	allMatches := pattern.FindAllString(description, -1)
	fmt.Println("First match:", firstMatch)
	for i, m := range allMatches {
		fmt.Println("Match", i, "=", m)
	}
}

func f16_28() {
	fmt.Println("======== list_16_28 ========")
	// " " => 공백문자를 기준으로 분리
	// "boat" → 단어 "boat"를 기준으로 분리
	// "one" → 단어 "one"을 기준으로 분리
	pattern := regexp.MustCompile(" |boat|one")
	description := "Kayak. A boat for one person."
	// 해당 정규식의 조건에  만족하는 문자열을 분리하여 []string 을 리턴한다.
	split := pattern.Split(description, -1)
	for _, s := range split {
		if s != "" {
			fmt.Println("Substring:", s)
		}
	}
}

func f16_29() {
	fmt.Println("======== list_16_29 ========")
	// "A " => A와공백
	// "[A-z]*" => A~z로 시작하고, 뒤에는 0개이상의 문자 (여기서는 boat가 매치됨)
	// " for " => 해당 문자열과 그대로 매치
	// [A-z]*" => 위와 같음
	// " person" => 해당 문자열과 그대로 매치
	pattern := regexp.MustCompile("A [A-z]* for [A-z]* person")
	// description := "Kayak. A boat for one person. aaa bbb ccc"
	description := "Kayak. A boat for one person."
	// 정규식과 일치하는 첫번째 문자열을 반환
	str := pattern.FindString(description)
	fmt.Println("Match:", str)
}

func f16_30() {
	fmt.Println("======== list_16_30 ========")
	// 하위 표현식 사용
	// 이전의 소스와 비교하여, 정규식에서 ([A-z]*) 항목이 수정되었다.
	// 하위표현식은 ()로 표현한다.
	pattern := regexp.MustCompile("A ([A-z]*) for ([A-z]*) person")
	description := "Kayak. A boat for one person."
	// A boat for one person. 해당 구조를 찾고, 그 안에서 "boat" 와 "one" 을 따로 추출한다.
	subs := pattern.FindStringSubmatch(description)
	for _, s := range subs {
		fmt.Println("Match:", s)
	}
}

func f16_31() {
	fmt.Println("======== list_16_31 ========")
	// "A "
	// (?P<type>[A-z]*) → 이름이 type 인 캡처 그룹
	// " for "
	// (?P<capacity>[A-z]*) → 이름이 capacity 인 캡처 그룹
	// " person"
	pattern := regexp.MustCompile(
		"A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description := "Kayak. A boat for one person."
	// subs에는 "A boat for one person", "boat", "one" 이 저장된다.
	subs := pattern.FindStringSubmatch(description)
	// pattern.SubexpIndex(name) => 해당 그룹 이름에 해당 하는 인덱스를 반환한다.
	for _, name := range []string{"type", "capacity"} {
		fmt.Println(name, "=", subs[pattern.SubexpIndex(name)])
	}
}

func f16_32() {
	fmt.Println("======== list_16_32 ========")
	pattern := regexp.MustCompile(
		"A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description := "Kayak. A boat for one person."
	// ${type} 와 ${capacity} 는 정규식에서 캡처한 그룹 이름을 참조
	template := "(type: ${type}, capacity: ${capacity})"
	// description 문자열에서 pattern에 맞는 부분을 찾아 template 형태로 변환한다.
	// "A boat for one person."" 문자열이 해당 패턴에 일치한다.
	// "(type: boat, capacity: one)."" 으로 변경된다.
	replaced := pattern.ReplaceAllString(description, template)
	fmt.Println(replaced)
}

func f16_33() {
	fmt.Println("======== list_16_33 ========")
	pattern := regexp.MustCompile(
		"A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")
	description := "Kayak. A boat for one person."
	// ReplaceAllStringFunc 함수는 정규식에 매칭된 부분을 찾아서, 그 부분을 사용자 정의 함수로 바꾼다.
	// 인자로 들어온 s는 "A boat for one person."
	// 해당 값을 "This is the replacement content" 으로 치환한다.
	replaced := pattern.ReplaceAllStringFunc(description, func(s string) string {
		return "This is the replacement content"
	})
	fmt.Println(replaced)
}
