package gopherstack

import (
	"fmt"
	"log"
	"time"
)

// waitForAsyncJob simply blocks until the the asynchronous job has
// executed or has timed out.
func WaitForAsyncJob(jobId string, client *CloudStackClient, timeout time.Duration) error {
	done := make(chan struct{})
	defer close(done)

	result := make(chan error, 1)
	go func() {
		attempts := 0
		for {
			attempts += 1

			log.Printf("Checking async job status... (attempt: %d)", attempts)
			status, err := client.QueryAsyncJobResult(jobId)
			if err != nil {
				result <- err
				return
			}

			// check what the real state will be.
			if status == 1 {
				result <- nil
				return
			}

			// Wait 3 seconds in between
			time.Sleep(3 * time.Second)

			// Verify we shouldn't exit
			select {
			case <-done:
				// We finished, so just exit the goroutine
				return
			default:
				// Keep going
			}
		}
	}()

	log.Printf("Waiting for up to %d seconds for async job %s", timeout, jobId)
	select {
	case err := <-result:
		return err
	case <-time.After(timeout):
		err := fmt.Errorf("Timeout while waiting to for async job to finish")
		return err
	}
}
