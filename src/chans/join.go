package chans

// JoinErrs waits for an error to recieve on either ch1 or ch2. If an error returns on either, returns it
func JoinErrs(ch1, ch2 <-chan error) error {
	select {
	case err := <-ch1:
		return err
	case err := <-ch2:
		return err
	}
}
