package main

import (
	"sort"
)

/*
*
设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近 10 条推文。

实现 Twitter 类：

Twitter() 初始化简易版推特对象
void postTweet(int userId, int tweetId) 根据给定的 tweetId 和 userId 创建一条新推文。每次调用此函数都会使用一个不同的 tweetId 。
List<Integer> getNewsFeed(int userId) 检索当前用户新闻推送中最近  10 条推文的 ID 。新闻推送中的每一项都必须是由用户关注的人或者是用户自己发布的推文。推文必须 按照时间顺序由最近到最远排序 。
void follow(int followerId, int followeeId) ID 为 followerId 的用户开始关注 ID 为 followeeId 的用户。
void unfollow(int followerId, int followeeId) ID 为 followerId 的用户不再关注 ID 为 followeeId 的用户。
*/
func main() {
}

type Twitter struct {
	users map[int]map[int]struct{} //用户关注列表
	posts map[int][]post           //用户发表推特列表
	t     int
}
type post struct {
	time int
	tid  int
}

func Constructor() Twitter {
	return Twitter{
		users: map[int]map[int]struct{}{},
		posts: map[int][]post{},
		t:     0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.posts[userId] = append(this.posts[userId], post{this.t, tweetId})
	this.t++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	p := this.posts[userId]                     //用户发表的推特列表
	for followUid := range this.users[userId] { //用户关注的用户列表
		p = append(p, this.posts[followUid]...) //关注的用户发表的推特列表，追加在自己发表的列表中
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].time > p[j].time
	})
	var r []int
	for i := 0; i < 10 && i < len(p); i++ {
		r = append(r, p[i].tid)
	}
	return r
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.users[followerId] == nil {
		this.users[followerId] = map[int]struct{}{}
	}
	this.users[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.users[followerId], followeeId)
}
