package internal

type Arguments struct {
	*NoopCommand

	// INFO(celicoo): flags types extracted from:
	// https://github.com/golang/go/blob/master/src/cmd/go/internal/test/test.go
	Bench                string
	Benchmem             bool
	Benchtime            string
	Blockprofile         string
	Blockprofilerate     int
	Count                int
	Coverprofile         string
	Cpu                  string
	Cpuprofile           string
	Failfast             bool
	List                 string
	Memprofile           string
	Memprofilerate       int
	Mutexprofile         string
	Mutexprofilefraction int
	Outputdir            string
	Paniconexit0         bool
	Parallel             int
	R /* Run */          string
	Short                bool
	Shuffle              string
	Testlogfile          string
	Timeout              int
	Trace                string
	Verbose /* V */      bool
}

func (Arguments) Doc() string {
	return `
  -bench
  -benchmem
  -benchtime
  -blockprofile
  -blockprofilerate
  -count
  -coverprofile
  -cpu
  -cpuprofile
  -failfast
  -list
  -memprofile
  -memprofilerate
  -mutexprofile
  -mutexprofilefraction
  -outputdir
  -paniconexit0
  -parallel
  -r, run
  -short
  -shuffle
  -testlogfile
  -timeout
  -trace
  -v, verbose`
}
