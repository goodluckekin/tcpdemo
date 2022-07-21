package bus

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"sync"
	"zinx/demo/core/grid"
	"zinx/demo/core/player"
)

var Wm *worldManager

type worldManager struct {
	aoiManager *grid.AoiManager
	players    map[int]*player.Player
	rwLock     sync.RWMutex
}

func init() {
	Wm = &worldManager{
		aoiManager: grid.NewAoiManager(0, 500, 10, 0, 500, 10),
		players:    make(map[int]*player.Player),
	}
}

func (w *worldManager) AddPlayer(p *player.Player) {
	w.rwLock.Lock()
	defer w.rwLock.Unlock()
	w.aoiManager.AddPlayer2GridByPosition(int(p.Pid), int(p.X), int(p.Y))
	w.players[int(p.Pid)] = p
}

func (w *worldManager) RemovePlayer(p *player.Player) {
	w.rwLock.Lock()
	defer w.rwLock.Unlock()
	w.aoiManager.RemovePlayer2GridByPosition(int(p.Pid), int(p.X), int(p.Y))
	delete(w.players, int(p.Pid))
}

func (w *worldManager) GetPlayer(pid int) *player.Player {
	w.rwLock.RLock()
	defer w.rwLock.RUnlock()
	return w.players[pid]
}

func (w *worldManager) GetAllPlayers() []*player.Player {
	w.rwLock.RLock()
	defer w.rwLock.RUnlock()
	m := make([]*player.Player, len(w.players))
	for _, pl := range w.players {
		m = append(m, pl)
	}
	return m
}

//广播消息
func (w *worldManager) Brocast(msgId uint32, message proto.Message) {
	for _, p := range w.players {
		if err := p.SendMsg(msgId, message); err != nil {
			fmt.Printf("【brocast】 msgId:%d err:%v", msgId, err)
		}
	}
}
