package testutil

const (
	Alice = "cosmos1jqq546zgfhemhyztfsc7fpmcz44skzfxu0t7eq"
	Bob   = "cosmos1enyttm3m6pr9u458062vxm0xqnmrh9jxks7l69"
	Carol = "cosmos1uvh2k0gl3j3lc540l5yvgr5e502sdc0g0uedzv"
)

// alias checkersd=~/go/bin/checkersd
// export alice=$(checkersd keys show alice -a) export bob=$(checkersd keys show bob -a) export carol=$(checkersd keys show carol -a)

// To generate mocks
// ~/go/bin/mockgen -source=x/checkers/types/expected_keepers.go -package testutil -destination=x/checkers/testutil/expected_keepers_mocks.go
