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
				domain.VideosLongerThan(3*domain.Minutes),
				domain.LikeVideo,
			),
		)

	fireBall := domain.NewChannelSubscirber(&domain.ChannelSubscriberOptions{Name: "火球"}).
		SetUpdateEventHandler(
			domain.VideoUploadEventHandlerBuilder(
				domain.VideosShorterThan(1*domain.Minutes),
				domain.UnSubscribeChannel,
			),
		)

	waterBall.Subscribe(waterBallSoftwareAcademy)
	waterBall.Subscribe(pewDiePie)
	fireBall.Subscribe(waterBallSoftwareAcademy)
	fireBall.Subscribe(pewDiePie)

	waterBallSoftwareAcademy.UploadVideo(*domain.NewVideo(&domain.VideoOptions{
		Title:       "C1M1S2",
		Description: "這個世界正是物件導向的呢！",
		Length:      4 * domain.Minutes,
	}))

	pewDiePie.UploadVideo(*domain.NewVideo(&domain.VideoOptions{
		Title:       "Hello guys",
		Description: "”Clickbait",
		Length:      30 * domain.Seconds,
	}))

	waterBallSoftwareAcademy.UploadVideo(*domain.NewVideo(&domain.VideoOptions{
		Title:       "C1M1S3",
		Description: "物件 vs. 類別",
		Length:      1 * domain.Minutes,
	}))

	pewDiePie.UploadVideo(*domain.NewVideo(&domain.VideoOptions{
		Title:       "Minecraft",
		Description: "Let's play Minecraft",
		Length:      30 * domain.Minutes,
	}))

}
