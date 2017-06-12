package main

import (
	"encoding/json"
	//"testing"
	"time"
	"com.heaven7.study/iso6801"
	"fmt"
)

func main() {
	//testISO8601Time()
	//testISO8601TimeInt()
	//testISO8601TimeWithoutSeconds()
}

func log4Iso6801(a ...interface{}){
	fmt.Println(a);
}
func logf4Iso6801(format string, a ...interface{}){
	fmt.Printf(format + "\n", a);
}

//go test ?
func testISO8601Time() {
	now := iso6801.NewISO6801Time(time.Now().UTC())

	data, err := json.Marshal(now)
	if err != nil {
		log4Iso6801(err);
	}

	_, err = time.Parse(`"`+ iso6801.FORMAT_ISO8601+`"`, string(data))
	if err != nil {
		log4Iso6801(err)
	}

	var now2 iso6801.ISO6801Time
	err = json.Unmarshal(data, &now2)
	if err != nil {
		log4Iso6801(err)
	}

	if now != now2 {
		logf4Iso6801("Time %s does not equal expected %s", now2, now)
	}

	if now.String() != now2.String() {
		logf4Iso6801("String format for %s does not equal expected %s", now2, now)
	}

	type TestTimeStruct struct {
		A int
		B *iso6801.ISO6801Time
	}
	var testValue TestTimeStruct
	err = json.Unmarshal([]byte("{\"A\": 1, \"B\":\"\"}"), &testValue)
	if err != nil {
		log4Iso6801(err)
	}
	logf4Iso6801("%v", testValue)
	if !testValue.B.IsDefault() {
		log4Iso6801("Invaid Unmarshal result for ISO6801Time from empty value")
	}
	logf4Iso6801("ISO6801Time String(): %s", now2.String())
}

func testISO8601TimeWithoutSeconds() {

	const dateStr = "\"2015-10-02T12:36Z\""

	var date iso6801.ISO6801Time

	err := json.Unmarshal([]byte(dateStr), &date)
	if err != nil {
		log4Iso6801(err)
	}

	const dateStr2 = "\"2015-10-02T12:36:00Z\""

	var date2 iso6801.ISO6801Time

	err = json.Unmarshal([]byte(dateStr2), &date2)
	if err != nil {
		log4Iso6801(err)
	}

	if date != date2 {
		log4Iso6801("The two dates shoudl be equal.")
	}

}

func testISO8601TimeInt() {

	const dateStr = "1405544146000"

	var date iso6801.ISO6801Time

	err := json.Unmarshal([]byte(dateStr), &date)
	if err != nil {
		log4Iso6801(err)
	}

	logf4Iso6801("date: %s", date)

}