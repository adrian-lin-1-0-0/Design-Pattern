package main

import (
	"3.1.H/domain"
)

func main() {

	waterBallSoftwareAcademy := domain.NewChannel(&domain.ChannelOptions{Name: "水球軟體學院"})
	pewDiePie := domain.NewChannel(&domain.ChannelOptions{Name: "PewDiePie"})

	waterBall := domain.NewChannelSubscirber(&domain.ChannelSubscriberOptions{Name: "水球"}).
		SetUpdateEventHandler(
			domain.VideoUploadEventHandlerBuilder(
				func(video domain.Video) bool {
					return video.Seconds >= 3*60
				},
				func(subscriber *domain.ChannelSubscriber, video domain.Video) {
					subscriber.Likes(video)
				},
			),
		)

	fireBall := domain.NewChannelSubscirber(&domain.ChannelSubscriberOptions{Name: "火球"}).
		SetUpdateEventHandler(
			domain.VideoUploadEventHandlerBuilder(
				func(video domain.Video) bool {
					return video.Seconds <= 60
				},
				func(subscriber *domain.ChannelSubscriber, video domain.Video) {
					subscriber.UnSubscribe(video.Uploader)
				},
			),
		)

	waterBall.Subscribe(waterBallSoftwareAcademy)
	waterBall.Subscribe(pewDiePie)
	fireBall.Subscribe(waterBallSoftwareAcademy)
	fireBall.Subscribe(pewDiePie)

	waterBallSoftwareAcademy.UploadVideo(*domain.NewVideo(&domain.VideoOptions{
		Title:       "C1M1S2",
		Description: "這個世界正是物件導向的呢！",
		Seconds:     4 * 60,
	}))

	pewDiePie.UploadVideo(*domain.NewVideo(&domain.VideoOptions{
		Title:       "Hello guys",
		Description: "”Clickbait",
		Seconds:     30,
	}))

	waterBallSoftwareAcademy.UploadVideo(*domain.NewVideo(&domain.VideoOptions{
		Title:       "C1M1S3",
		Description: "物件 vs. 類別",
		Seconds:     60,
	}))

	pewDiePie.UploadVideo(*domain.NewVideo(&domain.VideoOptions{
		Title:       "Minecraft",
		Description: "Let's play Minecraft",
		Seconds:     30 * 60,
	}))

}
