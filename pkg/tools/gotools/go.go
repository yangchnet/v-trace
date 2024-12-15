package gotools

import "gitee.com/qciip-icp/v-trace/pkg/logger"

// Go start a goroutine with recover.
func Go(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Errorf("recover from error: %+v", r)
			}
		}()
		fn()
	}()
}
