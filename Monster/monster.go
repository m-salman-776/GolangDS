package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	//stack.New()
	totalMonster := 60
	monsterChan := make(chan *Monster, totalMonster)
	wg := &sync.WaitGroup{}
	for i := 1; i <= totalMonster; i++ {
		wg.Add(1)
		go fetchMonsterById(i, monsterChan, wg)
	}
	wg.Wait()
	close(monsterChan)
	for monster := range monsterChan {
		monster.print()
		fmt.Println()
	}
}

type Monster struct {
	Name    string    `json:"name"`
	Ailment []Ailment `json:"ailments"`
}

func (m *Monster) print() {
	name := m.Name
	x := ""
	fmt.Print("Name : ", name)
	if len(m.Ailment) > 0 {
		x += " => Ailments  are : "
		for _, ailment := range m.Ailment {
			if len(ailment.Name) == 0 {
				continue
			}
			x = fmt.Sprintf("%s ,%s", x, ailment.Name)
		}
	}

	fmt.Print(x)
}

type Ailment struct {
	Name string `json:"name"`
}

func fetchMonsterById(id int, monsterChan chan *Monster, wg *sync.WaitGroup) {
	defer wg.Done()
	baseUrl := "https://mhw-db.com/monsters"
	url := fmt.Sprintf("%s/%d", baseUrl, id)
	//fmt.Println(url)
	response, err := http.Get(url)
	//fmt.Println(response)
	// todo response code check pending
	//if response.Status != http.StatusOK {
	//
	//}
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var monster *Monster
	if err := json.Unmarshal(responseData, &monster); err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(monster.Name)
	monsterChan <- monster
	//return monster
}
