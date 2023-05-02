package domain

type Video struct {
	ID          string
	Title       string
	Description string
	Seconds     int
	Uploader    *Channel
}

type VideoOptions struct {
	Title       string
	Description string
	Seconds     int
}

func NewVideo(options *VideoOptions) *Video {
	return &Video{
		ID:          GenerateUUID(),
		Title:       options.Title,
		Description: options.Description,
		Seconds:     options.Seconds,
	}
}
