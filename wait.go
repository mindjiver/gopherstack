package gopherstack

import (
	"fmt"
	"log"
	"time"
)

// waitForAsyncJob simply blocks until the the asynchronous job has
// executed or has timed out.
func (c CloudStackClient) WaitForAsyncJob(jobId string, timeout time.Duration) error {
	done := make(chan struct{})
	defer close(done)

	result := make(chan error, 1)
	go func() {
		attempts := 0
		for {
			attempts += 1

			log.Printf("Checking async job status... (attempt: %d)", attempts)
			status, err := c.QueryAsyncJobResult(jobId)
			if err != nil {
				result <- err
				return
			}

			// job is completed, we can exit
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

// waitForAsyncJob simply blocks until the virtual machine is in the
// specified state.
func (c CloudStackClient) WaitForVirtualMachineState(vmid string, wanted_state string, timeout time.Duration) error {
	done := make(chan struct{})
	defer close(done)

	result := make(chan error, 1)
	go func() {
		attempts := 0
		for {
			attempts += 1

			log.Printf("Checking virtual machine state... (attempt: %d)", attempts)
			_, current_state, err := c.VirtualMachineState(vmid)
			if err != nil {
				result <- err
				return
			}

			// check what the real state will be.
			log.Printf("current_state: %s", current_state)
			log.Printf("wanted_state:  %s", wanted_state)
			if current_state == wanted_state {
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

	log.Printf("Waiting for up to %d seconds for Virtual Machine state to converge", timeout)
	select {
	case err := <-result:
		return err
	case <-time.After(timeout):
		err := fmt.Errorf("Timeout while waiting to for Virtual Machine to converge")
		return err
	}
}
