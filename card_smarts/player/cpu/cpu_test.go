package cpu

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CPUSuite struct {
	suite.Suite
	comp CPU
}

func (suite *CPUSuite) SetupTest() {
	suite.comp.Hand = []int{1, 2, 2, 3, 3, 4, 4, 5}
}

func (suite *CPUSuite) TestPlayCard() {
	played := []int{}
	for i := 0; i < 8; i++ {
		played = append(played, suite.comp.PlayCard())
	}
	sort.Ints(played)
	suite.Equal([]int{1, 2, 2, 3, 3, 4, 4, 5}, played)
	suite.Equal(0, len(suite.comp.Hand))
}

func TestCPUSuite(t *testing.T) {
	suite.Run(t, new(CPUSuite))
}



