package statistics

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func Test01(t *testing.T) {

	limit := 100000
	s := NewServer()
	s.Add("test01")
	m2 := make(map[int]string, limit)
	for j := 0; j < limit; j++ {
		m2[j] = strconv.Itoa(j) + "abccd"
	}

	time.Sleep(time.Second * 1)
	str, err := s.Print("test01")
	fmt.Println(str, err)

	s.Add("test02")
	time.Sleep(time.Second * 2)

	m := make(map[int]string)
	for j := 0; j < limit; j++ {
		m[j] = strconv.Itoa(j) + "abccd"
	}

	str2, err2 := s.Print("test02")
	fmt.Println(str2, err2)

	s.Add("test03")
	s.Print("test03")
	str3, err3 := s.Print("test03")
	fmt.Println(str3, err3)
}

func Test02(t *testing.T) {
	s := NewServer()
	limit := 100000
	var wg sync.WaitGroup
	for j := 0; j < limit; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			taskName := strconv.Itoa(j) + "_task"
			s.Add(taskName)
			if j == 1 {
				time.Sleep(time.Second * time.Duration(j))
			}
			str, err := s.Print(taskName)
			fmt.Println(str, err)
		}(j)
	}
	wg.Wait()

	str, err := s.All()
	fmt.Println("***all***", str, err)

	fmt.Println("*********************over*****************")
}

func Test03(t *testing.T) {
	s := NewServer()
	limit := 10
	var wg sync.WaitGroup
	for j := 0; j < limit; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			taskName := strconv.Itoa(j) + "_task"
			s.Add(taskName)
			if j == 1 {
				time.Sleep(time.Second * time.Duration(j))
			}
			str, err := s.Print(taskName)
			fmt.Println(str, err)
		}(j)
	}
	wg.Wait()

	//str, err := s.All()
	//fmt.Println("***all***", str, err)

	fmt.Println("*********************over*****************")
}
