package xcommon

import (
	"context"
	"time"
)

type RetryFunc func(ctx context.Context) error

//如果f方法出现错误，重试retryTimes次
func Retry(ctx context.Context, f RetryFunc, retryTimes uint8) error {
	for {
		if err := f(ctx); err != nil {
			if retryTimes == 0 {
				return err
			}

			retryTimes--
			time.Sleep(1 * time.Millisecond)
			continue
		}

		return nil
	}

	return nil
}
