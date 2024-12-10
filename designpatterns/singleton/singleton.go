package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int // key: city name, value: city population
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

var (
	once     sync.Once
	instance *singletonDatabase
)

func readData(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := map[string]int{}

	for scanner.Scan() {
		city := scanner.Text()
		if !scanner.Scan() {
			return nil, fmt.Errorf("unexpected data format for city %s", city)
		}
		population, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("invalid population value for city %s: %v", city, err)
		}
		result[city] = population
	}

	return result, nil
}

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, err := readData("./capitals.txt")
		if err != nil {
			log.Fatalf("Failed to load data: %v", err)
		}
		instance = &singletonDatabase{capitals: caps}
	})
	return instance
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return d.dummyData[name]
}

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Println("Population of Seoul:", pop)

	cities := []string{"Seoul", "Mexico City"}
	tp := GetTotalPopulationEx(GetSingletonDatabase(), cities)
	ok := tp == (17500000 + 17400000)
	fmt.Println(ok)

	names := []string{"alpha", "gamma"}
	tp = GetTotalPopulationEx(&DummyDatabase{}, names)
	fmt.Println(tp == 4)
}
