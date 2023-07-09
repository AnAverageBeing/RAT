package antianalysis

import "os"

func Start() {
	done := make(chan struct{})
	defer close(done)

	checks := []func() bool{
		IsDebuggerPresent,
		IsNetworkAnalysisRunning,
		IsVm,
		SandBoxDetected,
	}

	for _, check := range checks {
		go func(chk func() bool) {
			if chk() {
				os.Exit(69)
			}
			done <- struct{}{}
		}(check)
	}

	for range checks {
		<-done
	}
}
