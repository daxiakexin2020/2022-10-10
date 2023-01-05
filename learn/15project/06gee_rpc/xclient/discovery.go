package xclient

import (
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

type SelectMode int

const (
	RandomSelect     SelectMode = iota //随机
	RoundRobinSelect                   //轮询
)

type Discovery interface {
	Refresh() error                      //从注册中心更新服务列表
	Update(servers []string) error       //手动更新服务列表
	Get(mode SelectMode) (string, error) //根据负载均衡策略，选择一个服务实例
	GetAll() ([]string, error)           //返回所有服务实例
}

type MultiServersDiscovery struct {
	r       *rand.Rand   //r 是一个产生随机数的实例，初始化时使用时间戳设定随机数种子，避免每次产生相同的随机数序列。
	mu      sync.RWMutex //读写锁
	servers []string     //服务实例集合
	index   int          //index 记录 Round Robin 算法已经轮询到的位置，为了避免每次从 0 开始，初始化时随机设定一个值。
}

func NewMultiServerDiscovery(servers []string) *MultiServersDiscovery {
	d := &MultiServersDiscovery{
		servers: servers,
		r:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	d.index = d.r.Intn(math.MaxInt32 - 1) //index 记录 Round Robin 算法已经轮询到的位置，为了避免每次从 0 开始，初始化时随机设定一个值。
	return d
}

var _ Discovery = (*MultiServersDiscovery)(nil)

func (d *MultiServersDiscovery) Refresh() error {
	return nil
}

func (d *MultiServersDiscovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers
	return nil
}

func (d *MultiServersDiscovery) Get(mode SelectMode) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	n := len(d.servers)
	if n == 0 {
		return "", errors.New("rpc discovery: no available servers")
	}
	switch mode {
	case RandomSelect:
		return d.servers[d.r.Intn(n)], nil //d.r.Intn(n) 返回一个0-范围的随机数，在slice中随机拿一个
	case RoundRobinSelect:
		/**
		data := []int{1, 2, 3, 4, 5, 6}
		n := len(data)
		index := 0
		for i := 0; i < 10; i++ {
			key := index % n
			res := data[key]
			fmt.Printf("i=%d,index=%d,key=%d,res=%d\n", i, index, key, res)
			index = (index + 1) % n
		}
		i=0,index=0,key=0,res=1
		i=1,index=1,key=1,res=2
		i=2,index=2,key=2,res=3
		i=3,index=3,key=3,res=4
		i=4,index=4,key=4,res=5
		i=5,index=5,key=5,res=6

		todo  按照顺序轮询，一圈之后，又重归为0开始...

		i=6,index=0,key=0,res=1
		i=7,index=1,key=1,res=2
		i=8,index=2,key=2,res=3
		i=9,index=3,key=3,res=4
		*/
		s := d.servers[d.index%n]
		d.index = (d.index + 1) % n //后移一位，轮询下一位
		return s, nil
	default:
		return "", errors.New("rpc discovery: not supported select mode")
	}
}

func (d *MultiServersDiscovery) GetAll() ([]string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	servers := make([]string, len(d.servers), len(d.servers))
	copy(servers, d.servers)
	return servers, nil
}
