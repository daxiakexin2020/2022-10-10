package items

import (
	"22go_redis/server/construct"
	"errors"
	"sort"
)

func (gr *Gredis) Zadd(key string, data map[interface{}]int64) {
	gr.Lock()
	defer gr.Unlock()
	gr.zadd(key, data)
}

func (gr *Gredis) Zrevrange(key string, start int64, end int64) ([]interface{}, []int64) {
	gr.Lock()
	defer gr.Unlock()
	vInterface, ok := gr.isData(key)
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
	return vals, scores
}

func (gr *Gredis) Zcard(key string) int {
	gr.RLock()
	defer gr.RUnlock()
	vInterface, ok := gr.isData(key)
	if !ok {
		return 0
	}
	return len(vInterface.GetVal().([]*construct.Czval))
}

func (gr *Gredis) Zcount(key string, min, max int64) int {
	gr.RLock()
	defer gr.RUnlock()
	vInterface, ok := gr.isData(key)
	if !ok {
		return 0
	}
	if vInterface.Type() != construct.ZSET {
		return 0
	}
	var count int
	czvals := vInterface.GetVal().([]*construct.Czval)
	for _, v := range czvals {
		if v.GetScore() >= min && v.GetScore() <= max {
			count++
		}
	}
	return count
}

func (gr *Gredis) Zincrby(key string, val interface{}, score int64) int64 {
	gr.RLock()
	defer gr.RUnlock()
	vInterface, ok := gr.isData(key)
	if !ok {
		m := make(map[interface{}]int64)
		m[val] = score
		gr.zadd(key, m)
		return score
	} else {
		czvals := vInterface.GetVal().([]*construct.Czval)
		for _, v := range czvals {
			if v.GetVal() == val {
				score += v.GetScore()
				break
			}
		}
		m := make(map[interface{}]int64)
		m[val] = score
		gr.zadd(key, m)
		return score
	}
}

func (gr *Gredis) Zrank(key string, m interface{}) (int64, error) {
	data, b := gr.isData(key)
	if !b {
		return 0, errors.New("key is not set")
	}
	czvals := data.GetVal().([]*construct.Czval)
	for i, v := range czvals {
		if v.GetVal() == m {
			return int64(i) + 1, nil
		}
	}
	return 0, errors.New("key is not set")
}

func (gr *Gredis) Zrem(key string, m ...interface{}) int {
	gr.Lock()
	defer gr.Unlock()
	vInterface, b := gr.isData(key)
	if !b {
		return 0
	}
	flag := make(map[interface{}]struct{})
	for _, v := range m {
		flag[v] = struct{}{}
	}
	var count int
	czvals := vInterface.GetVal().([]*construct.Czval)
	for i := 0; i < len(czvals); i++ {
		if len(flag) == 0 {
			break
		}
		v := czvals[i].GetVal()
		if _, ok := flag[v]; ok {
			delete(flag, v)
			czvals = append(czvals[:i], czvals[i+1:]...)
			count++
			i--
		}
	}
	vInterface.SetVal(czvals)
	return count
}

func (gr *Gredis) Zscore(key string, m interface{}) (int64, error) {
	data, b := gr.isData(key)
	if !b {
		return 0, errors.New("this key is not set")
	}
	czvals := data.GetVal().([]*construct.Czval)
	for _, v := range czvals {
		if v.GetVal() == m {
			return v.GetScore(), nil
		}
	}
	return 0, errors.New("this member is not set")
}

func (gr *Gredis) zadd(key string, data map[interface{}]int64) {
	vInterface, ok := gr.isData(key)
	var scores []int64
	var fs []*construct.Czval
	for _, score := range data {
		scores = append(scores, score)
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[j] <= scores[i]
	})

	if !ok {
		val := construct.NewCZset(10)
		zset := val.GetVal().([]*construct.Czval)
		for _, score := range scores {
			for dv, dscore := range data {
				if dscore == score {
					czval := construct.NewCzval(dv, score)
					zset = append(zset, czval)
				}
			}
		}
		val.SetVal(zset)
		gr.CGOSet(key, val)
	} else {
		zset := vInterface.(*construct.CZset)
		czvals := zset.GetVal().([]*construct.Czval)
		for i := 0; i < len(czvals); i++ {
			val := czvals[i].GetVal()
			score := czvals[i].GetScore()
			if newScore, ok := data[val]; ok {
				delete(data, val)
				if score == newScore {
					continue
				}
				//update
				fs = append(fs, construct.NewCzval(val, newScore))
				czvals = append(czvals[:i], czvals[i+1:]...)
				i--
			}
		}
		//add
		for val, score := range data {
			fs = append(fs, construct.NewCzval(val, score))
		}
		for _, v := range fs {
			if len(czvals) == 0 {
				czvals = append(czvals, v)
				continue
			}
			for k, dval := range czvals {
				if v.GetScore() > dval.GetScore() {
					left := make([]*construct.Czval, len(czvals[k:]))
					copy(left, czvals[k:])
					czvals = append(czvals[:k], v)
					czvals = append(czvals, left...)
					break
				} else if k == len(czvals)-1 {
					czvals = append(czvals[:k+1], v)
					break
				}
			}
		}
		zset.SetVal(czvals)
		gr.CGOSet(key, zset)
	}
}
