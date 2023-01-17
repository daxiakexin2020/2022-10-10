package v1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"v1/random"
)

type Server struct {
	r      *random.Rand
	path   string
	smalls map[int][]string
}

func NewServer(path string) *Server {
	return &Server{
		r:      random.NewRand(),
		path:   path,
		smalls: make(map[int][]string),
	}
}

const (
	BIG_FILE         = "big_file.txt"
	SMALL_PREFIX     = "small_"
	BIG_LIMIT        = 5000000
	SMALL_FILE_LIMIT = 100
	TOP              = 100
)

func (s *Server) InitializeBigFile() {
	file, err := os.OpenFile(s.path+"/"+BIG_FILE, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var totalLine int32
	for j := 0; j < BIG_LIMIT; j++ {
		count := s.r.MakeCount(int32(j))
		var str strings.Builder
		var i int32
		for i = 1; i <= count; i++ {
			str.WriteString(strconv.FormatInt(int64(j), 10))
			str.WriteString("\n")
			totalLine++
		}
		_, err := file.WriteString(str.String())
		if err != nil {
			fmt.Println("	【file write error 】: ", err)
		}
	}
	fileInfo, _ := file.Stat()
	fmt.Printf("一共生成大文件=%d行，大小%dMib\n", totalLine, fileInfo.Size()/1024/1024)
}

func (s *Server) SpiltSmallFiles() {

	//先创建10个文件，0，1，2，3，4......
	var wg sync.WaitGroup
	for j := 0; j < SMALL_FILE_LIMIT; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			filePath := s.path + "/" + smallName(j)
			_, err := os.Create(filePath)
			if err != nil {
				panic(err)
			}
		}(j)
	}
	wg.Wait()

	for j := 0; j < SMALL_FILE_LIMIT; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			bfilePath := s.path + "/" + BIG_FILE
			bigFile, err := os.OpenFile(bfilePath, os.O_RDWR, 0666)
			if err != nil {
				panic(err)
			}
			defer bigFile.Close()

			sfilePath := s.path + "/" + smallName(j)
			smallFile, err := os.OpenFile(sfilePath, os.O_WRONLY, 0666)
			if err != nil {
				panic(err)
			}
			defer smallFile.Close()

			br := bufio.NewReader(bigFile) // 创建 Reader

			for {
				line, err := br.ReadString('\n')
				//去掉字符串首尾空白字符，返回字符串
				lineStr := strings.TrimSpace(line)

				if err != nil && err != io.EOF {
					panic(err)
				}
				if err == io.EOF {
					break
				}
				fdata, _ := strconv.ParseInt(lineStr, 10, 64)
				findex := int(fdata % SMALL_FILE_LIMIT)
				if findex == j {
					smallFile.WriteString(lineStr + "\n")
				}
			}
		}(j)
	}
	wg.Wait()
	fmt.Println("**********************切割结束**********************")
}

func (s *Server) EverySmallFileCount() {
	for j := 0; j < SMALL_FILE_LIMIT; j++ {
		filepath := s.path + "/" + smallName(j)
		file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		small := make(map[string]int32)
		sr := bufio.NewReader(file)
		for {
			line, err := sr.ReadString('\n')
			//去掉字符串首尾空白字符，返回字符串
			lineStr := strings.TrimSpace(line)

			if err != nil && err != io.EOF {
				panic(err)
			}
			if err == io.EOF {
				break
			}
			small[lineStr]++
		}

		//内部过滤，选择出来内部的Top，此处返回值数据量已经很低了，经过过滤的
		filteredSmall := s.smallFiltered(&small)

		//合并在大的集合中
		for count, elements := range filteredSmall {
			s.smalls[count] = append(s.smalls[count], elements...)
		}
	}
	s.ListSort()
}

func (s *Server) ListSort() {
	s.listSort()

	var counts []int
	for count, _ := range s.smalls {
		counts = append(counts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	for i, count := range counts {
		k := i
		elements := s.smalls[count]
		for _, element := range elements {
			fmt.Printf("%s一共出现%d次,排名第%d\n", element, count, k+1)
		}
	}
}

func (s *Server) listSort() {

	var keys []int
	for key, _ := range s.smalls {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	var innerCount int
	for _, count := range keys {
		//外层，已经超过Top
		if len(s.smalls) >= TOP {
			break
		}
		//内层，已经超过Top
		if innerCount >= TOP {
			break
		}
		elements, ok := s.smalls[count]
		if !ok {
			continue
		}

		destValueLen := TOP - innerCount
		currelementsLen := len(elements)
		if currelementsLen >= destValueLen {
			elements = elements[0:destValueLen]
		}
		s.smalls[count] = elements
		innerCount += len(elements)
	}
}

/*
*
@param  map[数量][]string{"元素"}
*/
func (s *Server) smallFiltered(data *map[string]int32) map[int][]string {

	dest := make(map[int32][]string, 0)

	//用来将map排序的keys
	var keys []int
	for e, count := range *data {
		elements, ok := dest[count]
		if len(elements) == TOP {
			continue
		}
		elements = append(elements, e)
		dest[count] = elements
		if !ok {
			keys = append(keys, int(count))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var innerCount int
	res := make(map[int][]string)
	for i, count := range keys {
		//外层，已经超过Top
		if i >= TOP {
			break
		}
		//内层，已经超过Top
		if innerCount >= TOP {
			break
		}
		destValueLen := TOP - innerCount
		elements := dest[int32(count)]
		currElementsLen := len(elements)
		if currElementsLen >= destValueLen {
			elements = elements[0:destValueLen]
		}
		res[count] = elements
		innerCount += len(elements)
	}
	return res
}

func smallName(i int) string {
	return SMALL_PREFIX + strconv.Itoa(i)
}
