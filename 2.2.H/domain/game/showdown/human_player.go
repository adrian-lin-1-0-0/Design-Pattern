package showdown

import "2.2.H/utils/pubsub"

type HumanPlayerCore struct {
	Subscriber pubsub.Subscriber
	PlayerCore
}
