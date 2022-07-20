/**
 * @Author: ekin
 * @Description:
 * @File: core
 * @Version: 1.0.0
 * @Date: 2022/7/18 13:50
 */

package grid

import (
	"fmt"
	"sync"
)

type Grid struct {
	gid        int
	minX       int
	maxX       int
	minY       int
	maxY       int
	players    map[int]struct{}
	playerLock sync.Mutex
}

func NewGrid(gid int, minX int, maxX int, minY int, maxY int) *Grid {
	return &Grid{
		gid:     gid,
		minX:    minX,
		maxX:    maxX,
		minY:    minY,
		maxY:    maxY,
		players: make(map[int]struct{}),
	}
}

//增加玩家
func (g *Grid) AddPlayer(playerId int) {
	g.playerLock.Lock()
	defer g.playerLock.Unlock()
	g.players[playerId] = struct{}{}
}

//删除玩家
func (g *Grid) DelPlayer(playerId int) {
	g.playerLock.Lock()
	defer g.playerLock.Unlock()
	delete(g.players, playerId)
}

//获取所有玩家id
func (g *Grid) GetPlayerIds() []int {
	ids := make([]int, len(g.players))
	for id := range g.players {
		ids = append(ids, id)
	}
	return ids
}

func (g *Grid) String() string {
	fmt.Printf("core info gid:%d minX:%d maxX:%d minY:%d maxY:%d players:%+v\n", g.gid, g.minX, g.maxX, g.minY, g.maxY, g.players)
	return fmt.Sprintf("core info gid:%d minX:%d maxX:%d minY:%d maxY:%d players:%+v", g.gid, g.minX, g.maxX, g.minY, g.maxY, g.players)
}
