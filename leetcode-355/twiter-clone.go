package leetcode355

type Tweet struct {
	id        int
	timestamp int
	next      *Tweet
}

type Twitter struct {
	following map[int]map[int]bool
	tweets    map[int]*Tweet
	timestamp int
}

func Constructor() Twitter {
	return Twitter{
		tweets:    make(map[int]*Tweet, 0),
		following: make(map[int]map[int]bool, 0),
		timestamp: 0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	tweet := &Tweet{
		id:        tweetId,
		timestamp: this.timestamp,
	}
	this.timestamp++

	tweet.next = this.tweets[userId]
	this.tweets[userId] = tweet

}

func (this *Twitter) GetNewsFeed(userId int) []int {
	result := make([]int, 0, 10)
	heads := make([]*Tweet, 0, 501)

	if len(this.tweets) == 0 {
		return result
	}

	if tweet := this.tweets[userId]; tweet != nil {
		heads = append(heads, tweet)
	}

	if followings, ok := this.following[userId]; ok {
		for followeeId := range followings {
			if tweet := this.tweets[followeeId]; tweet != nil {
				heads = append(heads, tweet)
			}
		}
	}

	for len(result) < 10 && len(heads) > 0 {
		maxIdx := 0
		maxTs := -1

		for i := 0; i < len(heads); i++ {
			if heads[i].timestamp > maxTs {
				maxTs = heads[i].timestamp
				maxIdx = i
			}
		}
		if maxTs == -1 {
			break
		}
		result = append(result, heads[maxIdx].id)
		heads[maxIdx] = heads[maxIdx].next

		if heads[maxIdx] == nil {
			heads = append(heads[:maxIdx], heads[maxIdx+1:]...)
		}
	}

	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, exists := this.following[followerId]; !exists {
		this.following[followerId] = make(map[int]bool, 0)
	}

	this.following[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if following, ok := this.following[followerId]; ok {
		delete(following, followeeId)
	}
}
