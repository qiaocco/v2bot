package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"time"
)

type sendMessageReqBody struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Node  struct {
		Title string `json:"title"`
	} `json:"node"`
}
type Posts []Post

func (posts Posts) IDList() []int {
	var list []int
	for _, post := range posts {
		list = append(list, post.ID)
	}
	return list
}

var ids []int
var posts = Posts{}

func getList() ([]int, error) {
	url := "https://www.v2ex.com/api/topics/latest.json"
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("get failed, err:%v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ReadAll failed, err:%v\n", err)
		return nil, err
	}
	//fmt.Println(string(body))
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Printf("ReadAll failed, err:%v\n", err)
		return nil, err
	}

	return posts.IDList(), nil
}

// 获得a有b没有的数字
func difference(a, b []int) (diff []int) {
	m := make(map[int]bool)
	for _, item := range b {
		m[item] = true
	}
	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	diff = append(diff, 783130)
	log.Printf("diff=%v\n", diff)
	return
}

func init() {
	ids, _ = getList()
}

func push(id int) {
	for _, post := range posts {
		if post.ID == id {
			node := post.Node.Title
			title := post.Title
			link := post.Url
			msg := fmt.Sprintf("#%s %s %s", node, title, link)
			log.Printf("msg=%v\n", msg)
			url := fmt.Sprintf("https://msg.qiaocco.com?msg=%v", url2.QueryEscape(msg))
			_, err := http.Get(url)
			if err != nil {
				log.Printf("http get failed: err:%v\n", err)
				return
			}

		}
	}
}

func main() {
	for {
		fetchIds, err := getList()
		if err != nil {
			time.Sleep(60 * time.Second)
			continue
		}
		log.Printf("fetchIds=%v\nids=%v\n", fetchIds, ids)
		newIds := difference(fetchIds, ids)
		for _, id := range newIds {
			log.Println(id)
			go push(id)
		}
		ids = fetchIds
		time.Sleep(30 * time.Second)
	}

}
