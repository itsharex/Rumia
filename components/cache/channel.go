package cache

import (
	"Rumia/models"
	"sync"
)

var PushChannel = make(chan *models.Data)

var RetryChannel = make(chan *models.Data)

var FeedbackResult sync.Map
