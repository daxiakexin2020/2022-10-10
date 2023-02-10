package main

func main() {

}

/*
*
你需要设计一个包含验证码的验证系统。每一次验证中，用户会收到一个新的验证码，这个验证码在 currentTime 时刻之后 timeToLive 秒过期。如果验证码被更新了，那么它会在 currentTime （可能与之前的 currentTime 不同）时刻延长 timeToLive 秒。

请你实现 AuthenticationManager 类：

AuthenticationManager(int timeToLive) 构造 AuthenticationManager 并设置 timeToLive 参数。
generate(string tokenId, int currentTime) 给定 tokenId ，在当前时间 currentTime 生成一个新的验证码。
renew(string tokenId, int currentTime) 将给定 tokenId 且 未过期 的验证码在 currentTime 时刻更新。如果给定 tokenId 对应的验证码不存在或已过期，请你忽略该操作，不会有任何更新操作发生。
countUnexpiredTokens(int currentTime) 请返回在给定 currentTime 时刻，未过期 的验证码数目。
如果一个验证码在时刻 t 过期，且另一个操作恰好在时刻 t 发生（renew 或者 countUnexpiredTokens 操作），过期事件 优先于 其他操作
*/
type AuthenticationManager struct {
	mp  map[string]int
	ttl int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		mp:  make(map[string]int, 0),
		ttl: timeToLive,
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.mp[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if v, ok := this.mp[tokenId]; ok && v+this.ttl > currentTime {
		this.mp[tokenId] = currentTime
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	var count int
	for _, v := range this.mp {
		if v+this.ttl > currentTime {
			count++
		}
	}
	return count
}
