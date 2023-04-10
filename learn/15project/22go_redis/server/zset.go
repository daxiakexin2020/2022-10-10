package server

import (
	"22go_redis/server/construct"
	"fmt"
	"github.com/gogo/protobuf/sortkeys"
)

func (gr *Gredis) Zadd(key string, data map[int64]interface{}) {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	vInterface, ok := gr.Data[key]
	var scores []int64
	flag := make(map[interface{}]int64)
	for k, v := range data {
		scores = append(scores, k)
		flag[v] = k
	}
	sortkeys.Int64s(scores)
	type fuck struct {
		val   interface{}
		score int64
		p     int //1 befor; 2 after
		index int
	}
	var fs []*fuck
	if !ok {
		val := construct.NewCZset(10)
		zset := val.GetVal().([]*construct.Czval)
		for _, score := range scores {
			czval := construct.NewCzval(data[score], score)
			zset = append(zset, czval)
		}
		val.SetVal(zset)
		gr.Data[key] = val
	} else {
		//append    update
		zset := vInterface.(*construct.CZset)
		czvals := zset.GetVal().([]*construct.Czval)
		copyCzvals := make([]*construct.Czval, len(czvals))
		copy(copyCzvals, czvals)
		for i := 0; i < len(czvals); {
			v := czvals[i]
			score := v.GetScore()
			val := v.GetVal()
			//update
			if newScore, ok := flag[val]; ok {
				if score == newScore {
					i++
					continue
				}
				f := &fuck{
					val:   val,
					score: newScore,
				}
				fs = append(fs, f)
				//update 找新位置插入，删除旧位置
				czvals = append(czvals[:i], czvals[i+1:]...)
				fmt.Println("append:", i, czvals)
			} else {
				//append 找新位置插入
				f := &fuck{
					val:   val,
					score: newScore,
				}
				fs = append(fs, f)
				i++
			}
		}
		//for _, v := range fs {
		//	fmt.Println("fs:", v.score, v.val)
		//}

		//for _, v := range czvals {
		//	fmt.Println("fs:", v.GetVal(), v.GetScore())
		//}
		for _, v := range fs {
			for k, dval := range czvals {
				if v.score < dval.GetScore() || k == len(czvals)-1 {
					fmt.Println("ccccC:", v.score, v.val, k, len(czvals)-1)
					if k == len(czvals)-1 {
						czvals = append(czvals[:k+1], construct.NewCzval(v.val, v.score))
					} else {
						czvals = append(czvals[:k], construct.NewCzval(v.val, v.score))
						czvals = append(czvals, czvals[k+1:]...)
					}
					for _, v := range czvals {
						fmt.Println("fs0000000000:", v.GetVal(), v.GetScore())
					}
					for _, v := range czvals {
						fmt.Println("fs11111:", v.GetVal(), v.GetScore())
					}
					continue
				}
			}
		}
		zset.SetVal(czvals)
		gr.Data[key] = zset
	}

}

func (gr *Gredis) Zrevrange(key string, start int64, end int64) ([]int64, []interface{}) {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	vInterface, ok := gr.Data[key]
	if !ok {
		return nil, nil
	}
	czvals, ok := vInterface.GetVal().([]*construct.Czval)
	if !ok {
		return nil, nil
	}
	var scores []int64
	var vals []interface{}
	for _, czv := range czvals {
		score := czv.GetScore()
		if score >= start && score <= end {
			scores = append(scores, score)
			vals = append(vals, czv.GetVal())
		}
	}
	return scores, vals
}
