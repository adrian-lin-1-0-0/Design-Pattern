package domain

import (
	"fmt"
	"time"
)

type ChannelSubscriber struct {
	ID                 string
	Name               string
	updateEventHandler func(video Video)
}

type ChannelSubscriberOptions struct {
	Name string
}

func (subscriber *ChannelSubscriber) Likes(video Video) {
	like := Like{
		UserID:    subscriber.ID,
		Timestamp: time.Now(),
	}

	likesMap[video.ID] = append(likesMap[video.ID], like)
	fmt.Println(subscriber.Name, "對影片", video.Title, "按讚")
}

func (subscriber *ChannelSubscriber) Subscribe(channel *Channel) {
	channel.BeSubscribed(subscriber)
	fmt.Println(subscriber.Name, "訂閱了", channel.Name)
}

func (subscriber *ChannelSubscriber) UnSubscribe(channel *Channel) {
	channel.BeUnSubscribed(subscriber)
	fmt.Println(subscriber.Name, "解除訂閱了", channel.Name)
}

func (subscriber *ChannelSubscriber) SetUpdateEventHandler(handlerBuilder func(*ChannelSubscriber) func(Video)) *ChannelSubscriber {
	subscriber.updateEventHandler = handlerBuilder(subscriber)
	return subscriber
}

func (subscriber *ChannelSubscriber) Update(video Video) {
	subscriber.updateEventHandler(video)
}

func VideoUploadEventHandlerBuilder(condition func(Video) bool, action func(*ChannelSubscriber, Video)) func(*ChannelSubscriber) func(Video) {
	return func(subscriber *ChannelSubscriber) func(Video) {
		return func(video Video) {
			if condition(video) {
				action(subscriber, video)
			}
		}
	}
}

func NewChannelSubscirber(options *ChannelSubscriberOptions) *ChannelSubscriber {
	return &ChannelSubscriber{
		ID:                 GenerateUUID(),
		Name:               options.Name,
		updateEventHandler: func(video Video) {},
	}
}

func VideosLongerThan(seconds int) func(Video) bool {
	return func(video Video) bool {
		return video.Seconds >= seconds
	}
}

func VideosShorterThan(seconds int) func(Video) bool {
	return func(video Video) bool {
		return video.Seconds < seconds
	}
}

func UnSubscribeChannel(subscriber *ChannelSubscriber, video Video) {
	subscriber.UnSubscribe(video.Uploader)
}

func LikeVideo(subscriber *ChannelSubscriber, video Video) {
	subscriber.Likes(video)
}
