package xkafka

import "context"

func RunWithContext(ctx context.Context, f func() error) error {
	errChan := make(chan error)
	go func() {
		errChan <- f()
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}
