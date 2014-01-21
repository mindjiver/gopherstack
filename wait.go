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
			response, err := c.QueryAsyncJobResult(jobId)
			if err != nil {
				result <- err
				return
			}

			// job is completed, we can exit
			status := response.Queryasyncjobresultresponse.Jobstatus
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
func (c CloudStackClient) WaitForVirtualMachineState(vmid string, wantedState string, timeout time.Duration) error {
	done := make(chan struct{})
	defer close(done)

	result := make(chan error, 1)
	go func() {
		attempts := 0
		for {
			attempts += 1

			log.Printf("Checking virtual machine state... (attempt: %d)", attempts)
			response, err := c.ListVirtualMachines(vmid)
			if err != nil {
				result <- err
				return
			}

			count := response.Listvirtualmachinesresponse.Count
			if count != 1 {
				result <- err
				return
			}

			currentState := response.Listvirtualmachinesresponse.Virtualmachine[0].State
			// check what the real state will be.
			log.Printf("current state: %s", currentState)
			log.Printf("wanted state:  %s", wantedState)
			if currentState == wantedState {
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
