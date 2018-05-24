package geoip

import "time"

type Counter struct {
	data []int64
}

func (c *Counter) RecordHit() {
	c.data = append(c.data, time.Now().Unix())
}

func (c *Counter) Count() int {
	return len(c.data)
}

func (c *Counter) Cleanup() {
	currentTime := time.Now().Unix();
	for i := len(c.data)-1; i>=0; i-- {
		if currentTime - c.data[i] > 60 {
			c.data = c.data[i:]
			break
		}
	}
}


