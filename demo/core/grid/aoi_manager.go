/**
 * @Author: 格子管理（一片区域）
 * @Description:
 * @File: manager
 * @Version: 1.0.0
 * @Date: 2022/7/18 14:23
 */

package grid

type AoiManager struct {
	minX  int           //最小的x
	maxX  int           //最大的x
	xNum  int           //x方向有多少格子
	minY  int           //最小边界的y
	maxY  int           //最大的y
	yNum  int           //y方向有多少格子
	grids map[int]*Grid //格子集合
}

func NewAoiManager(minX int, maxX int, xNum int, minY int, maxY int, yNum int) *AoiManager {
	m := &AoiManager{
		minX:  minX,
		maxX:  maxX,
		xNum:  xNum,
		minY:  minY,
		maxY:  maxY,
		yNum:  yNum,
		grids: make(map[int]*Grid, xNum*yNum),
	}

	//初始化格子
	for x := 0; x < xNum; x++ {
		for y := 0; y < yNum; y++ {
			gid := y*xNum + x
			m.grids[gid] = NewGrid(gid, x*(maxX/xNum), (x+1)*(maxX/xNum), y*(maxY/yNum), (y+1)*(maxY/yNum))
		}
	}
	return m
}

//根据gid获取九宫格集合
func (m *AoiManager) GetNearbyGridsByGid(gid int) []*Grid {
	grids := make([]*Grid, 0)

	//不存在这个格子
	if _, ok := m.grids[gid]; !ok {
		return nil
	}

	//判断左边界
	if g := m.GetLeftGrid(gid); g != nil {
		grids = append(grids, g)
	}

	//判断右边界
	if g := m.GetRightGrid(gid); g != nil {
		grids = append(grids, g)
	}

	//循环获取上方两个角
	for _, grid := range grids {

		//获取上方格子
		if g := m.GetTopGrid(grid.gid); g != nil {
			grids = append(grids, g)
		}

		//获取下方格子
		if g := m.GetBottomGrid(grid.gid); g != nil {
			grids = append(grids, g)
		}
	}

	//判断上边界
	if g := m.GetTopGrid(gid); g != nil {
		grids = append(grids, g)
	}

	//判断下边界
	if g := m.GetBottomGrid(gid); g != nil {
		grids = append(grids, g)
	}

	return grids
}

//根据gid获取上方格子gid
func (m *AoiManager) GetTopGrid(gid int) *Grid {
	idy := gid / m.yNum
	if idy > 0 {
		return m.grids[gid-m.xNum]
	}
	return nil
}

//根据gid获取下方格子gid
func (m *AoiManager) GetBottomGrid(gid int) *Grid {
	idy := gid / m.yNum
	if idy < m.yNum-1 {
		return m.grids[gid+m.xNum]
	}
	return nil
}

//根据gid获取左方格子gid
func (m *AoiManager) GetLeftGrid(gid int) *Grid {
	idx := gid % m.xNum
	if idx > 0 {
		return m.grids[gid-1]
	}
	return nil
}

//根据gid获取右方格子gid
func (m *AoiManager) GetRightGrid(gid int) *Grid {
	idx := gid % m.xNum
	if idx < m.xNum-1 {
		return m.grids[gid+1]
	}
	return nil
}

//获取x坐标每格宽度
func (m *AoiManager) GetGridWidth() int {
	return m.maxX / m.xNum
}

//获取y坐标每格高度
func (m *AoiManager) GetGridHeight() int {
	return m.maxY / m.yNum
}

//根据坐标获取gid
func (m *AoiManager) GetGidByPos(x, y int) int {
	idx := x / m.GetGridWidth()
	idy := y / m.GetGridHeight()
	gid := idy*m.xNum + idx
	return gid
}

//根据坐标获取九宫格
func (m *AoiManager) GetNearbyGridsByPos(x, y int) []*Grid {
	gid := m.GetGidByPos(x, y)
	return m.GetNearbyGridsByGid(gid)
}

//获取九宫格内的玩家id
func (m *AoiManager) GetNearbyPlayerIds(x, y int) []int {
	playIds := make([]int, 0)

	//获取九宫格
	grids := m.GetNearbyGridsByPos(x, y)
	if len(grids) > 0 {
		for _, g := range grids {
			playIds = append(playIds, g.GetPlayerIds()...)
		}
	}

	return playIds
}

//根据坐标新增玩家到格子里
func (m *AoiManager) AddPlayer2GridByPosition(pid int, x int, y int) {
	gid := m.GetGidByPos(x, y)
	grid := m.grids[gid]
	grid.AddPlayer(pid)
}

//根据坐标新增玩家到格子里
func (m *AoiManager) RemovePlayer2GridByPosition(pid int, x int, y int) {
	gid := m.GetGidByPos(x, y)
	grid := m.grids[gid]
	grid.DelPlayer(pid)
}
