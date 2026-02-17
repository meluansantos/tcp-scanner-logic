package scanner

type ScannerResult struct {
	IP  string
	Port uint16
	Open bool
	Banner string
	Err error
}
