package validator

type Command struct {
	WriteCommand  string
	VerifyCommand string
}

func initCommands() []Command {
	return []Command{
		// 基本case
		{
			WriteCommand:  "DEL keyToDelete",
			VerifyCommand: "GET keyToDelete",
		},
		{
			WriteCommand:  "EXPIRE keyToExpire 10",
			VerifyCommand: "TTL keyToExpire",
		},
		{
			WriteCommand:  "EXPIREAT keyToExpireAt 1724813932", // Example timestamp
			VerifyCommand: "TTL keyToExpireAt",
		},
		{
			WriteCommand:  "PERSIST keyToPersist",
			VerifyCommand: "TTL keyToPersist",
		},
		{
			WriteCommand:  "PEXPIRE keyToPexpire 10000", // TTL in milliseconds
			VerifyCommand: "PTTL keyToPexpire",
		},
		{
			WriteCommand:  "PEXPIREAT keyToPexpireAt 1724813932000", // Example timestamp in milliseconds
			VerifyCommand: "PTTL keyToPexpireAt",
		},
		{
			WriteCommand:  "UNLINK keyToUnlink",
			VerifyCommand: "GET keyToUnlink",
		},

		// String Commands
		{
			WriteCommand:  "DECR key11",
			VerifyCommand: "GET key11",
		},
		{
			WriteCommand:  "DECRBY key12 5",
			VerifyCommand: "GET key12",
		},
		{
			WriteCommand:  "INCR key13",
			VerifyCommand: "GET key13",
		},
		{
			WriteCommand:  "INCRBY key14 10",
			VerifyCommand: "GET key14",
		},
		{
			WriteCommand:  "INCRBYFLOAT key15 1.5",
			VerifyCommand: "GET key15",
		},
		{
			WriteCommand:  "MSET key16 value16 key17 value17",
			VerifyCommand: "MGET key16 key17",
		},
		{
			WriteCommand:  "SET key18 value18",
			VerifyCommand: "GET key18",
		},
		{
			WriteCommand:  "SETEX key19 60 value19", // TTL of 60 seconds
			VerifyCommand: "GET key19",
		},
		{
			WriteCommand:  "SETNX key20 value20",
			VerifyCommand: "GET key20",
		},

		// Set Commands
		{
			WriteCommand:  "SADD setKey1 member1",
			VerifyCommand: "SMEMBERS setKey1",
		},
		{
			WriteCommand:  "SPOP setKey2",
			VerifyCommand: "SMEMBERS setKey2",
		},
		{
			WriteCommand:  "SREM setKey3 member3",
			VerifyCommand: "SMEMBERS setKey3",
		},

		// Hash Commands
		{
			WriteCommand:  "HDEL hashKey1 field1",
			VerifyCommand: "HGET hashKey1 field1",
		},
		{
			WriteCommand:  "HINCRBY hashKey2 field2 10",
			VerifyCommand: "HGET hashKey2 field2",
		},
		{
			WriteCommand:  "HINCRBYFLOAT hashKey3 field3 1.5",
			VerifyCommand: "HGET hashKey3 field3",
		},
		{
			WriteCommand:  "HMSET hashKey4 field4 value4 field5 value5",
			VerifyCommand: "HMGET hashKey4 field4 field5",
		},
		{
			WriteCommand:  "HSET hashKey5 field6 value6",
			VerifyCommand: "HGET hashKey5 field6",
		},
		{
			WriteCommand:  "HSETNX hashKey6 field7 value7",
			VerifyCommand: "HGET hashKey6 field7",
		},

		// Sorted Set Commands
		{
			WriteCommand:  "ZADD sortedSetKey1 1 member1",
			VerifyCommand: "ZRANGE sortedSetKey1 0 -1 WITHSCORES",
		},
		{
			WriteCommand:  "ZPOPMAX sortedSetKey2",
			VerifyCommand: "ZRANGE sortedSetKey2 0 -1 WITHSCORES",
		},
		{
			WriteCommand:  "ZPOPMIN sortedSetKey3",
			VerifyCommand: "ZRANGE sortedSetKey3 0 -1 WITHSCORES",
		},
		{
			WriteCommand:  "ZREM sortedSetKey4 member4",
			VerifyCommand: "ZRANGE sortedSetKey4 0 -1 WITHSCORES",
		},
		{
			WriteCommand:  "ZREMRANGEBYRANK sortedSetKey5 0 1",
			VerifyCommand: "ZRANGE sortedSetKey5 0 -1 WITHSCORES",
		},
		{
			WriteCommand:  "ZREMRANGEBYSCORE sortedSetKey6 1 2",
			VerifyCommand: "ZRANGE sortedSetKey6 0 -1 WITHSCORES",
		},

		// 更新 case
		// String Commands
		{
			WriteCommand:  "SET key18 newValue181; SET key18 newValue182",
			VerifyCommand: "GET key18",
		},
		{
			WriteCommand:  "INCR key13; INCR key13",
			VerifyCommand: "GET key13",
		},
		{
			WriteCommand:  "INCRBY key14 20; INCRBY key14 30",
			VerifyCommand: "GET key14",
		},
		{
			WriteCommand:  "INCRBYFLOAT key15 2.5; INCRBYFLOAT key15 3.5",
			VerifyCommand: "GET key15",
		},
		{
			WriteCommand:  "MSET key16 newValue161 key17 newValue171; MSET key16 newValue162 key17 newValue172",
			VerifyCommand: "MGET key16 key17",
		},
		{
			WriteCommand:  "SETEX key19 120 newValue191; SETEX key19 180 newValue192", // 更新值并设置新的TTL
			VerifyCommand: "GET key19",
		},

		// Hash Commands
		{
			WriteCommand:  "HSET hashKey5 field6 newValue61; HSET hashKey5 field6 newValue62",
			VerifyCommand: "HGET hashKey5 field6",
		},
		{
			WriteCommand:  "HMSET hashKey4 field4 newValue41 field5 newValue51; HMSET hashKey4 field4 newValue42 field5 newValue52",
			VerifyCommand: "HMGET hashKey4 field4 field5",
		},
		{
			WriteCommand:  "HINCRBY hashKey2 field2 20; HINCRBY hashKey2 field2 30",
			VerifyCommand: "HGET hashKey2 field2",
		},
		{
			WriteCommand:  "HINCRBYFLOAT hashKey3 field3 2.5; HINCRBYFLOAT hashKey3 field3 3.5",
			VerifyCommand: "HGET hashKey3 field3",
		},

		// Set Commands
		{
			WriteCommand:  "SADD setKey1 newMember11; SADD setKey1 newMember12",
			VerifyCommand: "SMEMBERS setKey1",
		},
		{
			WriteCommand:  "SREM setKey3 member31; SREM setKey3 member32",
			VerifyCommand: "SMEMBERS setKey3",
		},

		// Sorted Set Commands
		{
			WriteCommand:  "ZADD sortedSetKey1 2 newMember11; ZADD sortedSetKey1 3 newMember12",
			VerifyCommand: "ZRANGE sortedSetKey1 0 -1 WITHSCORES",
		},
		{
			WriteCommand:  "ZREM sortedSetKey4 member41; ZREM sortedSetKey4 member42",
			VerifyCommand: "ZRANGE sortedSetKey4 0 -1 WITHSCORES",
		},
		{
			WriteCommand:  "ZREMRANGEBYRANK sortedSetKey5 0 2; ZREMRANGEBYRANK sortedSetKey5 3 4",
			VerifyCommand: "ZRANGE sortedSetKey5 0 -1 WITHSCORES",
		},

		// ttl case
		// String Commands with TTL
		{
			WriteCommand:  "SETEX key21 100 value211; SETEX key21 200 value212", // 使用 SETEX 设置不同的TTL
			VerifyCommand: "TTL key21",
		},
		{
			WriteCommand:  "PSETEX key22 100000 value221; PSETEX key22 200000 value222", // 使用 PSETEX 设置不同的TTL（毫秒）
			VerifyCommand: "PTTL key22",
		},
		{
			WriteCommand:  "SET key23 value231 EX 100; SET key23 value232 EX 200", // 使用 SET 命令并设置EX（秒）TTL
			VerifyCommand: "TTL key23",
		},
		{
			WriteCommand:  "SET key24 value241 PX 100000; SET key24 value242 PX 200000", // 使用 SET 命令并设置PX（毫秒）TTL
			VerifyCommand: "PTTL key24",
		},
		{
			WriteCommand:  "SET key25 value251; EXPIRE key25 100; SET key25 value252; EXPIRE key25 200", // 设置TTL并更新值
			VerifyCommand: "TTL key25",
		},
		{
			WriteCommand:  "SET key26 value261; PEXPIRE key26 100000; SET key26 value262; PEXPIRE key26 200000", // 设置毫秒TTL并更新值
			VerifyCommand: "PTTL key26",
		},

		// Hash Commands with TTL
		{
			WriteCommand:  "HSET hashKey7 field71 value71; EXPIRE hashKey7 100; HSET hashKey7 field71 value72; EXPIRE hashKey7 200", // 为Hash键设置TTL并更新字段
			VerifyCommand: "TTL hashKey7",
		},
		{
			WriteCommand:  "HMSET hashKey8 field81 value81 field82 value82; EXPIRE hashKey8 100; HMSET hashKey8 field81 value83 field82 value84; EXPIRE hashKey8 200", // 为多个字段设置TTL并更新
			VerifyCommand: "TTL hashKey8",
		},

		// Set Commands with TTL
		{
			WriteCommand:  "SADD setKey4 member41; EXPIRE setKey4 100; SADD setKey4 member42; EXPIRE setKey4 200", // 为集合设置TTL并添加新成员
			VerifyCommand: "TTL setKey4",
		},

		// Sorted Set Commands with TTL
		{
			WriteCommand:  "ZADD sortedSetKey7 1 member71; EXPIRE sortedSetKey7 100; ZADD sortedSetKey7 2 member72; EXPIRE sortedSetKey7 200", // 为排序集合设置TTL并添加新成员
			VerifyCommand: "TTL sortedSetKey7",
		},
	}
}
