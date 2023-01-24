package testutil

const (
	Alice = "cosmos146dsp3cf79y85mny8rnnnmpveg24kpqpntdces"
	Bob   = "cosmos1ae26g9qyldr5uun48asmulxs6zfpervktr9hrk"
	Carol = "cosmos1g8p0ax849hh2p9cgfmfd4xr93ytjzgzlppnfhd"
)

// alias checkersd=~/go/bin/checkersd
// export alice=$(checkersd keys show alice -a) export bob=$(checkersd keys show bob -a) export carol=$(checkersd keys show carol -a)

// To generate mocks
// ~/go/bin/mockgen -source=x/checkers/types/expected_keepers.go -package testutil -destination=x/checkers/testutil/expected_keepers_mocks.go
