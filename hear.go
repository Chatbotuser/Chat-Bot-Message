package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type hearStruct struct {
	regex string
	text  string
	echo  bool // phân biệt người gửi là user hay fanpage
}

type ListQuestions struct {
	Id       int
	Question string
	Replay   string
	Sendid   string
}

type CallSendApiSimSimResponse struct {
	Messages string
}

// func crawlNews
func crawlNews() {

}

func hear(word string, echo bool, userid string, repid string) string {
	if echo {

	} else {
		url := "https://fb.siblog.net/api/simsimi.php?key=sibendz&text=" + word

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("Accept", "application/json")
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("Postman-Token", "ddd7e433-300f-43e1-8410-e2f3e3aabb58")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		var data struct {
			Messages []struct {
				Text string `json:"text"`
			} `json:"messages"`
		}

		err := json.Unmarshal([]byte(body), &data)
		fmt.Println(err, data.Messages[0].Text)
		SendMessage(userid, data.Messages[0].Text)

		// End SIMSIM
	}
	// db, errDB := gorm.Open("mysql", "root:@/chatbot?charset=utf8&parseTime=True&loc=Local")
	// if errDB != nil {
	// 	log.Fatal("Cannot connect DB", errDB)
	// 	panic(errDB)
	// }
	// defer db.Close()
	// var s ListQuestions
	// if echo {
	// 	// Page trả lời.
	// 	errQR := db.Where("sendid = ? and replay = ''", repid).First(&s)
	// 	if errQR != nil {
	// 		db.Model(&s).Where("id = ?", s.Id).Update("replay", word)
	// 		SendMessage(userid, s.Replay)
	// 	}

	// 	return "Page reply"
	// } else {
	// 	// User gửi

	// 	errQR := db.Where("question LIKE ?", "%"+word+"%").Find(&s)

	// 	if errQR.RecordNotFound() {
	// 		// Không có dữ liệu tiến hành insert question
	// 		s = ListQuestions{Question: word, Sendid: userid}
	// 		db.Create(&s)

	// 		// ddg(userid)
	// 		// SendMessage(userid, "Not found DB")
	// 		// errQR = db.Where("MATCH(question) AGAINST ('?' IN NATURAL LANGUAGE MODE)", word).Find(&s)
	// 		// fmt.Println(errQR)
	// 		// if errQR.RecordNotFound() {
	// 		// 	fmt.Println("not found")
	// 		// 	ddg(userid)
	// 		// 	SendMessage(userid, "Not found DB")
	// 		// 	return "PostBack"
	// 		// } else {
	// 		// 	SendMessage(userid, s.Replay)
	// 		// 	return "Reply"
	// 		// }
	// 	} else {
	// 		SendMessage(userid, s.Replay)
	// 		return "Reply"
	// 	}
	// }

	return ""
}

/**
userid Id user gửi tin
repid Id user nhận nhận
*/
func (h *hearStruct) listen(userid string, repid string) {
	if h.text == "" && h.regex == "" {
		panic("Oops! Nothing to listen for")
	}

	if h.regex != "" {
		hear(h.regex, h.echo, userid, repid)
		fmt.Printf("REGEX PASSED:::%s\n\n", h.regex)
	} else if h.text != "" {
		hear(h.text, h.echo, userid, repid)
		fmt.Printf("TEXT PASSED:::%s\n\n", h.text)
	}
}
