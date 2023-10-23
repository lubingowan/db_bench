package main

// func BenchmarkRocksPut(b *testing.B) {
//	// 使用 gorocksdb 连接 RocksDB
//	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
//	bbto.SetBlockCache(gorocksdb.NewLRUCache(3 << 30))
//	opts := gorocksdb.NewDefaultOptions()
//	opts.SetBlockBasedTableFactory(bbto)
//	opts.SetCreateIfMissing(true)
//	// 设置输入目标数据库文件（可自行配置，./db 为当前测试文件的目录下的 db 文件夹）
//	db, _ := gorocksdb.OpenDb(opts, "./db")

//	// 创建输入输出流
//	// ro := gorocksdb.NewDefaultReadOptions()
//	wo := gorocksdb.NewDefaultWriteOptions()

//	for i := 0; i < b.N; i++ {
//		// 将键为 foo 值为 bar 的键值对写入文件中
//		_ = db.Put(wo, []byte("k"+fmt.Sprint(i)), []byte("v"+fmt.Sprint(i)))
//	}
//}

// func BenchmarkRocksGet(b *testing.B) {
//	// 使用 gorocksdb 连接 RocksDB
//	bbto := gorocksdb.NewDefaultBlockBasedTableOptions()
//	bbto.SetBlockCache(gorocksdb.NewLRUCache(3 << 30))
//	opts := gorocksdb.NewDefaultOptions()
//	opts.SetBlockBasedTableFactory(bbto)
//	opts.SetCreateIfMissing(true)
//	// 设置输入目标数据库文件（可自行配置，./db 为当前测试文件的目录下的 db 文件夹）
//	db, _ := gorocksdb.OpenDb(opts, "./db")

//	// 创建输入输出流
//	ro := gorocksdb.NewDefaultReadOptions()
//	// wo := gorocksdb.NewDefaultWriteOptions()

//	//for i := 0; i < b.N; i++ {
//	//	// 将键为 foo 值为 bar 的键值对写入文件中
//	//	_ = db.Put(wo, []byte("foo"), []byte("bar"))
//	//}

//	// 获取数据库中键为 foo 的值
//	for i := 0; i < b.N; i++ {
//		value, _ := db.Get(ro, []byte("k"+fmt.Sprint(i)))
//		defer value.Free()
//	}

//	// 打印数据库中键为 foo 的值
//	// fmt.Println("value: ", string(value.Data()[:]))
//	// 删除数据库中键为 foo 对应的键值对
//	//_ = db.Delete(wo, []byte("foo"))
//}
