package cmd

import (
	"os/exec"
	"syscall"
)

func terminateProcess(pid int) error {
	// Signal the process group (-pid), not just the process, so that the process
	// and all its children are signaled. Else, child procs can keep running and
	// keep the stdout/stderr fd open and cause cmd.Wait to hang.	p := &os.Process{Pid: pid}
	for i := 0; i < 5; i++ {
		_ = p.Signal(os.Interrupt)
		_, err := os.FindProcess(pid)
		if err != nil {
			return nil
		}
		time.Sleep(300 * time.Millisecond)
	}
	return syscall.Kill(-pid, syscall.SIGTERM)
}

func setProcessGroupID(cmd *exec.Cmd) {
	// Set process group ID so the cmd and all its children become a new
	// process group. This allows Stop to SIGTERM the cmd's process group
	// without killing this process (i.e. this code here).
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}
