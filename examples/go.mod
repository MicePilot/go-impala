module examples

go 1.18

require github.com/MicePilot/go-impala v0.0.0

require github.com/MicePilot/go-impala/hive v0.0.0

require github.com/MicePilot/go-impala/sasl v0.0.0

require github.com/MicePilot/go-impala/services/cli_service v0.0.0

require github.com/apache/thrift v0.12.0 // indirect

replace github.com/MicePilot/go-impala => ../

replace github.com/MicePilot/go-impala/hive => ../hive

replace github.com/MicePilot/go-impala/sasl => ../sasl

replace github.com/MicePilot/go-impala/services/cli_service => ../services/cli_service
