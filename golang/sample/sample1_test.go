package sample

import "testing"

// stTest:テスト用構造体
type stTest struct {
	x, y    int
	comment string
}

// testData:テストデータ
var testData = []stTest{
	stTest{-1, -1, "Err -x-"},
	stTest{0, -1, "Err 0x-"},
	stTest{-1, 0, "Err -x0"},
	stTest{1, 1, "Err 1"},
	stTest{2, 2, "Err 2"},
	stTest{10, 10, "Err 3"},
	stTest{100, 100, "Err 4"},
	stTest{1000, 1000, "Err 5"},
	stTest{10000, 10000, "Err 6"},
	stTest{100000, 100000, "Err 6"},
}

// TestS1:test for S1()
func TestS1(t *testing.T) {
	var x, y int
	var i interface{}

	for _, data := range testData {
		x = data.x
		y = data.y
		i = S1(x, y)
		_, err := i.(string)

		// ここでテスト結果を誤るとどうしようもない
		if !err {
			t.Errorf("is not True.%s", data.comment)
		}
	}
}

// TestS2:test for S2()
func TestS2(t *testing.T) {
	var x, y int
	var i interface{}

	for _, data := range testData {
		x = data.x
		y = data.y
		i = S2(x, y)
		i2, err := i.(int)

		// ここでテスト結果を誤るとどうしようもない
		if !err {
			t.Errorf("is not True1.%s", data.comment)
		}
		if i2 != 100 {
			t.Errorf("is not True2.%s", data.comment)
		}

	}
}

// TestS3:test for S3()
func TestS3(t *testing.T) {
	var x, y int
	var i interface{}

	for _, data := range testData {
		x = data.x
		y = data.y
		i = S3(x, y)
		_, err := i.(string)

		// ここでテスト結果を誤るとどうしようもない
		if !err {
			t.Errorf("is not True.%s", data.comment)
		}
	}
}
