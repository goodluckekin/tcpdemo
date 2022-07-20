package bus

import (
	"demo/demo/core/player"
	"sync"
)

var Wm *worldManager

type worldManager struct {
	aoiManager *AoiManager
	players    map[int]*play.Player
	rwLock     sync.RWMutex
}

func init() {
	Wm = &worldManager{
		aoiManager: NewAoiManager(0, 500, 10, 0, 500, 10),
		players:    make(map[int]*play.Player),
	}
}

func (w *worldManager) AddPlayer(player *play.Player) {
	w.rwLock.Lock()
	defer w.rwLock.Unlock()
	w.aoiManager.AddPlayer2GridByPosition(player.Pid, player.X, player.Y)
	w.players[player.Pid] = player
}

func (w *worldManager) RemovePlayer(player *play.Player) {
	w.rwLock.Lock()
	defer w.rwLock.Unlock()
	w.aoiManager.RemovePlayer2GridByPosition(player.Pid, player.X, player.Y)
	delete(w.players, player.Pid)
}

func (w *worldManager) GetPlayer(pid int) *play.Player {
	w.rwLock.RLock()
	defer w.rwLock.RUnlock()
	return w.players[player.Pid]
}

func (w *worldManager) GetAllPlayers() []*play.Player {
	w.rwLock.RLock()
	defer w.rwLock.RUnlock()
	m := make([]*player.Player, len(w.players))
	for _, pl := range w.players {
		m = append(m, pl)
	}
	return m
}
