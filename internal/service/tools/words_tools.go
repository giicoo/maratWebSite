package tools_service

import (
	"math/rand"
	"time"

	"github.com/giicoo/maratWebSite/models"
)

func RandomWords(w []*models.WordDB, i_root int) []*models.WordDB {
	words := make([]*models.WordDB, len(w))
	copy(words, w)

	words = words[:]
	words[i_root], words[len(words)-1] = words[len(words)-1], words[i_root]

	words = words[:len(words)-1]
	rand.Seed(time.Now().Unix())
	i1 := rand.Intn(len(words))
	w1 := words[i1]

	words[i1], words[len(words)-1] = words[len(words)-1], words[i1]

	words = words[:len(words)-1]
	i2 := rand.Intn(len(words))
	w2 := words[i2]

	words[i2], words[len(words)-1] = words[len(words)-1], words[i2]

	words = words[:len(words)-1]
	i3 := rand.Intn(len(words))
	w3 := words[i3]

	words[i3], words[len(words)-1] = words[len(words)-1], words[i3]

	return []*models.WordDB{w1, w2, w3}
}

func Insert(a []*models.WordDB, index int, value *models.WordDB) []*models.WordDB {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}
