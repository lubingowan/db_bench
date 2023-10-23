goos: darwin
goarch: amd64
pkg: suite/golang/db/bench
cpu: VirtualApple @ 2.50GHz
BenchmarkBoltPut-8                                   142           8504517 ns/op
BenchmarkBoltPut2-8                                  141           8069721 ns/op
BenchmarkBoltView-8                                  148           8098907 ns/op
BenchmarkEtcdPut-8                                   244           4730437 ns/op
BenchmarkEtcdGet-8                                 14010             80722 ns/op
BenchmarkMysqlInsert-8                              4825            285321 ns/op
BenchmarkMysqlInsertMultiColumn-8                   3880            288398 ns/op
BenchmarkMysqlSelect-8                                74          15289404 ns/op
BenchmarkMysqlSelectByPrimary-8                    12916             91909 ns/op
BenchmarkMysqlSelectByIndex-8                         75          15300187 ns/op
BenchmarkPostgreSQLInsert-8                         7807            158941 ns/op
BenchmarkPostgreSQLInsertMultiColumn-8              6338            180769 ns/op
BenchmarkPostgreSQLSelect-8                          502           2502049 ns/op
BenchmarkPostgreSQLSelectByPrimary-8               14233             83200 ns/op
BenchmarkRedisPut-8                                39427             28952 ns/op
BenchmarkRedisPut2-8                               40465             29644 ns/op
BenchmarkRedisHset-8                               39747             29822 ns/op
BenchmarkZKSet-8                                     290           4584684 ns/op
BenchmarkZKGet-8                                   15018             68699 ns/op