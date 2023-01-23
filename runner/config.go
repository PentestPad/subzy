package runner

type Settings struct {
	Targets     string
	HTTPS       bool
	Concurrency int
	Target      string
	Timeout     int
	VerifySSL   bool
	Emoji       bool
	HideFails   bool
}
