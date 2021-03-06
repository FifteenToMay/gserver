package entity

import (
	"fmt"
	"testing"
)

// 常规测试1
func Test1(t *testing.T) {
	testPath(`
......................................................
..F...MM..........MMMMMMMMMMMMMM...................MMM
..............XXXXXXXXXXXXX..T........................
.................XXXXX................................
..........~~~~~..XXX..........XX......................
......................................................
`, t, 0)
}

// 常规测试2
func Test2(t *testing.T) {
	testPath(`
......................................................
..F...MM..........MMMMMMMMMMMMMM...................MMM
..............XXXXXXXXXXXXXXX.........................
.................XXXXX................................
..........~~~~~..XXX...T.......XX.....................
......................................................
`, t, 0)
}

// 常规测试3_
func Test3_(t *testing.T) {
	testPath(`
......................................................
..F...MM..........MMMMMMMMMMMMMM...................MMM
..............XXXXXXXXXXXXX...........................
.................XXXX..........XX.....................
..........~~~~~..XXX..........XX.T.XX.................
......................................................
`, t, 0)
}

// 常规错误3
func Test3(t *testing.T) {
	testPath(`
......................................................
..F...MM..........MMMMMMMMMMMMMM...................MMM
..............XXXXXXXXXXXXX...........................
.................XXXXX....XXXXXXXXXXXXXXXX............
..........~~~~~..XXX.......XXXX.T.XXXX................
...........................XXXXXXXXXXXX...............
`, t, 0)
}

// 选择了走水路而不是绕路
func Test4(t *testing.T) {
	testPath(`
......................................................
..F...MM..........MMMMMMMMMMMMMM...................MMM
~~~~~~~T.......XXXXXXXXXXXXXXXX.......................
.................XXXXX................................
..........~~~~~..XXX..........XX......................
......................................................
`, t, 0)
}

// 选择了绕路而不是走水路
func Test5(t *testing.T) {
	testPath(`
......................................................
..F...MM..........MMMMMMMMMMMMMM...................MMM
~~~~~~~~~T.......XXXXXXXXXXXXX........................
.................XXXXX................................
..........~~~~~..XXX..........XX......................
......................................................
`, t, 0)
}

func testPath(worldInput string, t *testing.T, expectedDist float64) {
	world := ParseWorld(worldInput)
	start := world.From()
	end := world.To()
	fmt.Println("起点(", start.X, ",", start.Y, ")===>终点(",
		end.X, ",", end.Y, ")")
	p, dist, found := Path(world.From(), world.To())
	if !found {
		//正常错误测试
		// t.Errorf("找不到路径嘛")
		t.Logf("找不到路径嘛")
	} else {
		t.Logf("渲染出路径嘛\n%s", world.RenderPath(p))
	}
	if found && dist != expectedDist {
		t.Fatalf("想消耗 %v 但实际上却消耗 %v", expectedDist, dist)
	}
}
