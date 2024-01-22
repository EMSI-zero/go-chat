package boot

var bootHasFinished bool

func IsAfterBoot() bool {
	return bootHasFinished
}

func Boot() (err error) {
	if err = parseFlags(); err != nil {
		return err
	}

	if err == nil {
		bootHasFinished = true
	}
	return nil
}

func parseFlags() error{
	return nil
}