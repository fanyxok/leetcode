package main

type void struct{}

var VOID void

var timer = 0

func getTime() int {
	timer++
	return timer
}

type Set[T comparable] map[T]void

type Tweet struct {
	tweetId int
	userId  int
	time    int
}

type Twitter struct {
	FollowedSet map[int]Set[int]
	FollowerSet map[int]Set[int]
	Tweets      map[int]Tweet
	PostList    map[int][]int
	NewsQueue   map[int][]int
}

func Constructor() Twitter {
	return Twitter{
		FollowedSet: map[int]Set[int]{},
		FollowerSet: map[int]Set[int]{},
		NewsQueue:   map[int][]int{},
	}
}

func (ct *Twitter) PostTweet(userId int, tweetId int) {
	if ct.FollowerSet[userId] == nil {
		ct.FollowerSet[userId] = make(Set[int])
	}
	ct.appendToQueue(userId, tweetId)
	ct.Tweets[tweetId] = Tweet{
		tweetId: tweetId,
		userId:  userId,
		time:    getTime(),
	}
	followers := ct.FollowerSet[userId]
	for i := range followers {
		ct.appendToQueue(i, tweetId)
	}
}
func (ct *Twitter) appendToQueue(userId int, tweetId int) {
	if ct.NewsQueue[userId] == nil {
		ct.NewsQueue[userId] = []int{}
	}
	ct.NewsQueue[userId] = append(ct.NewsQueue[userId], tweetId)
}
func (ct *Twitter) GetNewsFeed(userId int) []int {
	if ct.NewsQueue[userId] == nil {
		ct.NewsQueue[userId] = []int{}
	}
	que := ct.NewsQueue[userId]
	var news []int
	if len(que) <= 10 {
		news = que
	} else {
		news = que[len(que)-11 : len(que)-1]
	}
	return news
}

func (ct *Twitter) Follow(followerId int, followeeId int) {
	if ct.FollowedSet[followerId] == nil {
		ct.FollowedSet[followerId] = make(Set[int])
	}
	ct.FollowedSet[followerId][followeeId] = VOID
	ct.insertTweetsToQueue(followerId, followeeId)
}

func (ct *Twitter) insertTweetsToQueue(followerId, followeeId int) {
	if ct.PostList[followeeId] == nil {
		ct.PostList[followeeId] = []int{}
	}
	posts := ct.PostList[followeeId]
	queue := ct.NewsQueue[followerId]

}

func (ct *Twitter) Unfollow(followerId int, followeeId int) {
	if ct.FollowedSet[followerId] == nil {
		ct.FollowedSet[followerId] = make(Set[int])
	}
	delete(ct.FollowedSet[followerId], followeeId)
	deleteTweetsFromQueue(followerId, followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
