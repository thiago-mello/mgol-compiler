package e

type EndOfFileReachedError string

func (e EndOfFileReachedError) Error() string {
	return string(e)
}
