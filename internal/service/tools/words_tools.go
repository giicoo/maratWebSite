package tools

import (
	"math/rand"
	"time"

	"github.com/giicoo/maratWebSite/models"
)

func RandomWords(w []*models.Word, i_root int) []*models.Word {
	// copy the slice so that it does not change
	words := make([]*models.Word, len(w))
	copy(words, w)

	// do swap root_item with last element and discard it
	words = words[:]
	words[i_root], words[len(words)-1] = words[len(words)-1], words[i_root]
	words = words[:len(words)-1]

	// get 1 words and discard it
	rand.Seed(time.Now().UnixNano())
	i1 := rand.Intn(len(words))
	w1 := words[i1]

	words[i1], words[len(words)-1] = words[len(words)-1], words[i1]

	// get 2 words and discard it
	words = words[:len(words)-1]
	i2 := rand.Intn(len(words))
	w2 := words[i2]

	words[i2], words[len(words)-1] = words[len(words)-1], words[i2]

	// get 3 words and discard it
	words = words[:len(words)-1]
	i3 := rand.Intn(len(words))
	w3 := words[i3]

	words[i3], words[len(words)-1] = words[len(words)-1], words[i3]

	return []*models.Word{w1, w2, w3}
}

func InsertByIndex(a []*models.Word, index int, value *models.Word) []*models.Word {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}
