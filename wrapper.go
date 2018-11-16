package koromo

import (
	eb "github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// LoadImage wraps ebitenutil.NewImageFromFile
func LoadImage(path string) (*eb.Image, error) {
	i, _, err := ebitenutil.NewImageFromFile(path, eb.FilterDefault)
	return i, err
}
