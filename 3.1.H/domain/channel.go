package domain

import "fmt"

type Channel struct {
	ID          string
	Name        string
	subscribers map[string]*ChannelSubscriber
}

type ChannelOptions struct {
	Name string
}

func (channel *Channel) BeSubscribed(subscriber *ChannelSubscriber) {
	channel.subscribers[subscriber.ID] = subscriber
}

func (channel *Channel) BeUnSubscribed(subscriber *ChannelSubscriber) {
	delete(channel.subscribers, subscriber.ID)
}

func NewChannel(options *ChannelOptions) *Channel {
	return &Channel{
		ID:          GenerateUUID(),
		Name:        options.Name,
		subscribers: make(map[string]*ChannelSubscriber),
	}
}

func (channel *Channel) UploadVideo(video Video) {
	video.Uploader = channel
	fmt.Println("頻道", channel.Name, "上架了一則新影片", video.Title)
	channel.Notify(video)
}

func (channel *Channel) Notify(video Video) {
	for _, subscriber := range channel.subscribers {
		subscriber.Update(video)
	}
}
