/**
 * @Author: EDZ
 * @Description:
 * @File: manager_test.go
 * @Version: 1.0.0
 * @Date: 2022/7/19 14:48
 */

package grid

import (
	"fmt"
	"testing"
)

func TestNewManager(t *testing.T) {
	m := NewManager(0, 250, 5, 0, 250, 5)
	g := m.GetNearbyGridsByGid(4)
	fmt.Println("gid grids:", g)

	g1 := m.GetNearbyGridsByPos(80, 20)
	fmt.Println("pos grids:", g1)
}
