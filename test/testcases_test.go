package test

import (
	math "math"
	"math/rand"
	"time"
	"unsafe"

	"github.com/billyplus/fastproto"
	pb "github.com/billyplus/fastproto/test/pb"
)

var (
	allTests        = getTestCases()
	fullProtoTest   = fullProtoMsg()
	arrStringTest   = &Strings{Val: []string{"this is test string", "\n", "", "\r\f", "我是大赢家！@#\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E;"}}
	arrBytesTest    = &Bytess{Val: [][]byte{[]byte("this is test []byte"), nil, []byte("\n"), []byte(""), []byte("\r\f"), []byte("我是大赢家！@#\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E;")}}
	arrInt32Test    = &Int32S{Val: append(randArr(20, rand.Int31), []int32{0, 0, 0, math.MaxInt32, math.MinInt32}...)}
	arrSint64Test   = &Sint64S{Val: append(randArr(20, rand.Int63), []int64{0, 0, 0, math.MaxInt64, math.MinInt64}...)}
	arrSfixed32Test = &Sfixed32S{Val: append(randArr(20, rand.Int31), []int32{0, 0, 0, math.MaxInt32, math.MinInt32}...)}
	arrSfixed64Test = &Sfixed64S{Val: append(randArr(20, rand.Int63), []int64{0, 0, 0, math.MaxInt64, math.MinInt64}...)}
)

func getTestCases() []fastproto.Message {
	testcase := []fastproto.Message{

		&Float{Val: 3036602},
		&Float{Val: 3036.32899602},
		&Float{Val: -3036.32899602},
		&Float{Val: 0.342438888},
		&Float{Val: -0.234342438888},
		&Float{Val: -467775643},
		&Float{Val: 0},
		&Float{Val: math.MaxFloat32},
		&Float{Val: -math.MaxFloat32},

		&Double{Val: 8933602},
		&Double{Val: 5736.32823602},
		&Double{Val: -3036.32899602},
		&Double{Val: 0.342438888},
		&Double{Val: -0.234342438888},
		&Double{Val: -467775643},
		&Double{Val: 0},
		&Double{Val: math.MaxFloat64},
		&Double{Val: -math.MaxFloat64},

		&Int32{Val: 3036602},
		&Int32{Val: -467775643},
		&Int32{Val: 0},
		&Int32{Val: math.MaxInt32},
		&Int32{Val: math.MinInt32},

		&Int64{Val: 3033202},
		&Int64{Val: -4534373},
		&Int64{Val: 0},
		&Int64{Val: math.MaxInt32},
		&Int64{Val: math.MinInt32},

		&Sint32{Val: 34755602},
		&Sint32{Val: -434633643},
		&Sint32{Val: 0},
		&Sint32{Val: math.MaxInt32},
		&Sint32{Val: math.MinInt32},

		&Sint64{Val: 760644202},
		&Sint64{Val: -4532644473},
		&Sint64{Val: 0},
		&Sint64{Val: math.MaxInt32},
		&Sint64{Val: math.MinInt32},

		&Uint32{Val: 34755602},
		&Uint32{Val: 0},
		&Uint32{Val: math.MaxUint32},

		&Uint64{Val: 57855602},
		&Uint64{Val: 0},
		&Uint64{Val: math.MaxUint64},

		&Fixed32{Val: 34755602},
		&Fixed32{Val: 0},
		&Fixed32{Val: math.MaxUint32},

		&Fixed64{Val: 57855602},
		&Fixed64{Val: 0},
		&Fixed64{Val: math.MaxUint64},

		&Sfixed32{Val: 34755602},
		&Sfixed32{Val: -434633643},
		&Sfixed32{Val: 0},
		&Sfixed32{Val: math.MaxInt32},
		&Sfixed32{Val: math.MinInt32},

		&Sfixed64{Val: 760644202},
		&Sfixed64{Val: -4532644473},
		&Sfixed64{Val: 0},
		&Sfixed64{Val: math.MaxInt32},
		&Sfixed64{Val: math.MinInt32},

		&Bool{Val: true},
		&Bool{Val: false},

		&String{Val: "string"},
		&String{Val: ""},
		&String{Val: "\n"},
		&String{Val: "\r\f"},
		&String{Val: "\f"},
		&String{Val: "\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E"},
		&String{Val: "我是大赢家！@#\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E;ereriufk;lJIJHIj32323^%^^$^&%*^&"},

		&Bytes{Val: []byte("Bytes")},
		&Bytes{Val: nil}, // []byte("") - empty bytes will mashaled as nil, which will cause tests failed to compare nil and []byte("")
		&Bytes{Val: []byte("\n")},
		&Bytes{Val: []byte("\r\f")},
		&Bytes{Val: []byte("\f")},
		&Bytes{Val: []byte("我是大赢家！@#\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E;ereriufk;lJIJHIj32323^%^^$^&%*^&")},

		&OtherMessage{
			Eid:    math.MaxInt64 - 1,
			OpenId: "\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E",
			Name:   "this.is.test.string",
			Job:    83804044,
			Sex:    833333,
		},

		&TestProtoMsg{
			Val: &OtherMessage{
				Eid:    math.MaxInt64 - 1,
				OpenId: "\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E",
				Name:   "this.is.test.string",
				Job:    83804044,
				Sex:    833333,
			},
			Val2: math.MaxInt32},

		&Floats{Val: []float32{3.0038344, 4.00000001, -3434.22283, 0, 0, 0, math.MaxFloat32, -math.MaxFloat32}},
		&Floats{Val: []float32{-3.0038344, 4.00000001, -3434.22283, 0, 0, 0, math.MaxFloat32, -math.MaxFloat32}},
		&Floats{Val: []float32{0, 0, 0, 0, 0, 0, 0, 0}},
		&Floats{Val: randArr(20, rand.Float32)},

		&Doubles{Val: []float64{33.00382222344, 4.00000000001, -343344.22283, 0, 0, 0, math.MaxFloat64, -math.MaxFloat64}},
		&Doubles{Val: []float64{-33.00382222344, 4.00000000001, -343344.22283, 0, 0, 0, math.MaxFloat64, -math.MaxFloat64}},
		&Doubles{Val: []float64{0, 0, 0, 0, 0, 0, 0, 0}},
		&Doubles{Val: randArr(20, rand.Float64)},

		&Int32S{Val: []int32{23434, -389894, 0, 0, 0, math.MaxInt32, math.MinInt32}},
		&Int32S{Val: []int32{-389894, 0, 236666, 0, math.MaxInt32, math.MinInt32, 0}},
		&Int32S{Val: []int32{0, 0, 0, 0, 0, 0, 0, 0}},
		&Int32S{Val: randArr(20, rand.Int31)},

		&Int64S{Val: []int64{23434, -389894, 0, 0, 0, math.MaxInt64, math.MinInt64}},
		&Int64S{Val: []int64{-389894, 0, 236666, 0, math.MaxInt64, math.MinInt64, 0}},
		&Int64S{Val: []int64{0, 0, 0, 0, 0, 0, 0, 0}},
		&Int64S{Val: randArr(20, rand.Int63)},

		&Uint32S{Val: []uint32{23434, 3844339894, 0, 0, 0, math.MaxInt32, 432222333}},
		&Uint32S{Val: []uint32{3833339894, 0, 236666, 0, math.MaxInt32, 432222333, 0}},
		&Uint32S{Val: []uint32{0, 0, 0, 0, 0, 0, 0, 0}},
		&Uint32S{Val: randArr(20, rand.Uint32)},

		&Uint64S{Val: []uint64{23434, 385549894, 0, 0, 0, math.MaxInt64, 1123322222223}},
		&Uint64S{Val: []uint64{3898333494, 0, 236666, 0, math.MaxInt64, 323444444444, 0}},
		&Uint64S{Val: []uint64{0, 0, 0, 0, 0, 0, 0, 0}},
		&Uint64S{Val: randArr(20, rand.Uint64)},

		&Sint32S{Val: []int32{23434, -389894, 0, 0, 0, math.MaxInt32, math.MinInt32}},
		&Sint32S{Val: []int32{-389894, 0, 236666, 0, math.MaxInt32, math.MinInt32, 0}},
		&Sint32S{Val: []int32{0, 0, 0, 0, 0, 0, 0, 0}},
		&Sint32S{Val: randArr(20, rand.Int31)},

		&Sint64S{Val: []int64{23434, -389894, 0, 3333, 0, 0, math.MaxInt64, math.MinInt64}},
		&Sint64S{Val: []int64{-389894, 0, 236666, 0, math.MaxInt64, math.MinInt64, 0}},
		&Sint64S{Val: []int64{0, 0, 0, 0, 0, 0, 0, 0}},
		&Sint64S{Val: randArr(20, rand.Int63)},

		&Fixed32S{Val: []uint32{23434, 38933894, 0, 0, 0, math.MaxInt32}},
		&Fixed32S{Val: []uint32{3333893894, 0, 236666, 0, math.MaxInt32, 0}},
		&Fixed32S{Val: []uint32{0, 0, 0, 0, 0, 0, 0, 0}},
		&Fixed32S{Val: randArr(20, rand.Uint32)},

		&Fixed64S{Val: []uint64{23434, 644389894, 0, 0, 0, math.MaxInt64}},
		&Fixed64S{Val: []uint64{894389894, 0, 236666, 0, math.MaxInt64, 0}},
		&Fixed64S{Val: []uint64{0, 0, 0, 0, 0, 0, 0, 0}},
		&Fixed64S{Val: randArr(20, rand.Uint64)},

		&Sfixed32S{Val: []int32{23434, -389894, 0, 0, 0, math.MaxInt32, math.MinInt32}},
		&Sfixed32S{Val: []int32{-389894, 0, 236666, 0, math.MaxInt32, math.MinInt32, 0}},
		&Sfixed32S{Val: []int32{0, 0, 0, 0, 0, 0, 0, 0}},
		&Sfixed32S{Val: randArr(20, rand.Int31)},

		&Sfixed64S{Val: []int64{23434, -389894, 0, 0, 0, math.MaxInt64, math.MinInt64}},
		&Sfixed64S{Val: []int64{-389894, 0, 236666, 0, math.MaxInt64, math.MinInt64, 0}},
		&Sfixed64S{Val: []int64{0, 0, 0, 0, 0, 0, 0, 0}},
		&Sfixed64S{Val: randArr(20, rand.Int63)},

		&Bools{Val: []bool{true, false, true, false, false}},
		&Bools{Val: []bool{true, true, true, true, true}},
		&Bools{Val: []bool{false, false, false, false, false}},
		&Bools{Val: nil}, // []bool{} would become nil after marshaled
		&Bools{Val: []bool{true, false, false, true, true, false}},

		// proto.Marshal will check if string is valid utf-8. So we need a valid utf-8 string
		&Strings{Val: []string{"this is test string", "\n", "", "\r\f", "我是大赢家！@#\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E;"}},
		&Strings{Val: []string{"", "", "", "", ""}},
		&Strings{Val: nil}, // []string{} would become nil

		&Bytess{Val: [][]byte{[]byte("this is test []byte"), []byte(""), []byte("\n"), []byte(""), []byte("\r\f"), []byte("我是大赢家！@#\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E;")}},
		&Bytess{Val: [][]byte{[]byte(""), []byte(""), []byte(""), []byte(""), []byte("")}}, // nil in []byte slice would be unmarshaled to []byte("")
		// &Bytess{Val: [][]byte{nil, nil, nil, nil, nil, nil}},
		&Bytess{Val: nil},

		&TestEnums{Val: []TestEnum{TestEnum_Test_Cmd, TestEnum_Test_None, TestEnum_Test_Push, TestEnum_Test_Cmd, TestEnum_Test_Max, TestEnum_Test_Push}},
	}

	// test message with slice of custom messages
	msgs := make([]*OtherMessage, 8)
	for i := range msgs {
		v := OtherMessage{
			Eid:    rand.Int63n(math.MaxInt64),
			OpenId: "\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E",
			Name:   "this.is.test.string",
			Job:    83804044,
			Sex:    rand.Int31n(math.MaxInt32),
		}
		msgs[i] = &v
	}
	testcase = append(testcase, &TestProtoMsgs{Val: msgs})

	m1 := &Mapint32Fixed64{
		Val: randMap(20, rand.Int31, rand.Uint64),
	}
	m1.Val[0] = rand.Uint64()
	m1.Val[math.MaxInt32] = rand.Uint64()
	m1.Val[math.MinInt32] = math.MaxUint64
	testcase = append(testcase, m1)

	m2 := &MapSint32Int64{
		Val: randMap(20, rand.Int31, rand.Int63),
	}
	m2.Val[0] = rand.Int63()
	m2.Val[math.MinInt32] = math.MinInt64
	m2.Val[math.MaxInt32] = math.MaxInt64
	testcase = append(testcase, m2)

	m3 := &MapSint64Int32{
		Val: randMap(20, rand.Int63, rand.Int31),
	}
	m3.Val[0] = rand.Int31()
	m3.Val[math.MinInt64] = math.MinInt32
	m3.Val[math.MaxInt64] = math.MaxInt32
	testcase = append(testcase, m3)

	m4 := &Mapint64Fixed32{
		Val: randMap(20, rand.Int63, rand.Uint32),
	}
	m4.Val[0] = rand.Uint32()
	m4.Val[math.MinInt64] = math.MaxUint32
	m4.Val[math.MaxInt64] = 0
	testcase = append(testcase, m4)

	testcase = append(testcase, &Mapuint32Sint64{Val: randMap(40, rand.Uint32, rand.Int63)})
	testcase = append(testcase, &Mapuint64Sint32{Val: randMap(30, rand.Uint64, rand.Int31)})
	testcase = append(testcase, &MapFixed32Double{Val: randMap(30, rand.Uint32, rand.Float64)})
	testcase = append(testcase, &MapFixed64Float{Val: randMap(30, rand.Uint64, rand.Float32)})
	testcase = append(testcase, &Mapsfixed32Uint64{Val: randMap(30, rand.Int31, rand.Uint64)})
	testcase = append(testcase, &Mapsfixed64Uint32{Val: randMap(30, rand.Int63, rand.Uint32)})
	testcase = append(testcase, &Mapstringsfixed32{Val: randMap(30, randStringN(16), rand.Int31)})
	testcase = append(testcase, &Mapstringsfixed64{Val: randMap(30, randStringN(16), rand.Int63)})

	// testcase = append(testcase, fullProtoMsg())

	testcase = append(testcase, &FullProto{MapInt32Bool: randMap(20, rand.Int31, func() bool { return rand.Int31()%2 == 1 })})
	testcase = append(testcase, &FullProto{MapInt32Int32: randMap(20, rand.Int31, rand.Int31)})
	testcase = append(testcase, &FullProto{MapStringEnum: randMap(20, randStringN(20), func() TestEnum { return TestEnum(rand.Int31() % int32(TestEnum_Test_Max+1)) })})
	testcase = append(testcase, &FullProto{MapStringActor: randMap(20, randStringN(20), func() *OtherMessage {
		return &OtherMessage{
			Eid:    rand.Int63(),
			OpenId: "\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E",
			Name:   randStringN(20)(),
			Job:    rand.Int31(),
			Sex:    rand.Int31(),
		}
	})})
	testcase = append(testcase, &FullProto{ArrActor: randArr(20, func() *OtherMessage {
		return &OtherMessage{
			Eid:    rand.Int63(),
			OpenId: "\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E",
			Name:   randStringN(20)(),
			Job:    rand.Int31(),
			Sex:    rand.Int31(),
		}
	})})
	testcase = append(testcase, &FullProto{Outer: &pb.OuterMsg{
		Eid:    rand.Int63(),
		OpenId: randStringN(20)(),
		Name:   randStringN(20)(),
		Job:    rand.Int31(),
		Sex:    rand.Int31(),
	}})

	return testcase
}

func fullProtoMsg() *FullProto {
	m := &FullProto{}
	m.VInt32 = rand.Int31()
	m.VBool = true
	m.VBytes = []byte("我是大赢家！@#\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E;ereriufk;lJIJHIj32323^%^^$^&%*^&" + randStringN(12)())
	m.VInt64 = rand.Int63()
	m.VUint32 = rand.Uint32()
	m.VUint64 = rand.Uint64()
	m.VString = "我是大赢家！@#\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E;ereriufk;lJIJHIj32323^%^^$^&%*^&" + randStringN(20)()

	m.SInt32 = rand.Int31()
	m.SInt64 = rand.Int63()
	m.Sfixed32 = rand.Int31()
	m.Sfixed64 = rand.Int63()
	m.Fixed32 = rand.Uint32()
	m.Fixed64 = rand.Uint64()
	m.ArrInt32 = randArr(20, rand.Int31)
	m.ArrInt64 = randArr(20, rand.Int63)
	m.ArrUint32 = randArr(20, rand.Uint32)
	m.ArrUint64 = randArr(20, rand.Uint64)
	m.ArrBool = randArr(20, func() bool { return rand.Int31()%2 == 1 })
	m.ArrString = randArr(20, randStringN(32))
	m.ArrBytes = randArr(20, func() []byte { return []byte(randStringN(32)()) })
	m.ArrActor = randArr(20, func() *OtherMessage {
		return &OtherMessage{
			Eid:    rand.Int63(),
			OpenId: "\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E",
			Name:   randStringN(20)(),
			Job:    rand.Int31(),
			Sex:    rand.Int31(),
		}
	})
	m.MapInt32Bool = randMap(20, rand.Int31, func() bool { return rand.Int31()%2 == 1 })
	m.MapInt32Int32 = randMap(20, rand.Int31, rand.Int31)
	m.MapInt32String = randMap(20, rand.Int31, randStringN(24))
	m.MapSfixed32Fixed64 = randMap(20, rand.Int31, rand.Uint64)
	m.MapStringSint64 = randMap(20, randStringN(20), rand.Int63)
	m.MapStringSfixed64 = randMap(20, randStringN(20), rand.Int63)
	m.MapStringActor = randMap(20, randStringN(20), func() *OtherMessage {
		return &OtherMessage{
			Eid:    rand.Int63(),
			OpenId: "\u0020\xe2\x8c\x98\b\u2318⌘\u65E5\u8A9E",
			Name:   randStringN(20)(),
			Job:    rand.Int31(),
			Sex:    rand.Int31(),
		}
	})

	return m
}

func randMap[K int32 | int64 | uint32 | uint64 | string, V any](n int, randKey func() K, randV func() V) map[K]V {
	arr := make(map[K]V, n)
	for i := 0; i < n; i++ {
		arr[randKey()] = randV()
	}
	return arr
}

func randArr[T any](n int, randV func() T) []T {
	arr := make([]T, n)
	for i := 0; i < n; i++ {
		arr[i] = randV()
	}
	return arr
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// https://stackoverflow.com/a/31832326/10737552
func randStringN(n int) func() string {
	return func() string {
		b := make([]byte, n)
		var src = rand.NewSource(time.Now().UnixNano())
		// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
		for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
			if remain == 0 {
				cache, remain = src.Int63(), letterIdxMax
			}
			if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
				b[i] = letterBytes[idx]
				i--
			}
			cache >>= letterIdxBits
			remain--
		}

		return *(*string)(unsafe.Pointer(&b))
	}
}
