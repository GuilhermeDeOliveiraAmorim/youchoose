package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListFavorite_Equals(t *testing.T) {
	listFavorite1 := NewListFavorite("chooser1", "list1")
	listFavorite2 := NewListFavorite("chooser1", "list1")

	equals := listFavorite1.Equals(listFavorite2)
	assert.True(t, equals, "ListFavorites iguais devem retornar true")

	listFavorite3 := NewListFavorite("chooser1", "list1")
	listFavorite4 := NewListFavorite("chooser2", "list1")

	equals = listFavorite3.Equals(listFavorite4)
	assert.False(t, equals, "ListFavorites com ChooserID diferente devem retornar false")

	listFavorite5 := NewListFavorite("chooser1", "list1")
	listFavorite6 := NewListFavorite("chooser1", "list2")

	equals = listFavorite5.Equals(listFavorite6)
	assert.False(t, equals, "ListFavorites com ListID diferente devem retornar false")

	listFavorite7 := NewListFavorite("chooser1", "list1")
	listFavorite8 := NewListFavorite("chooser2", "list2")

	equals = listFavorite7.Equals(listFavorite8)
	assert.False(t, equals, "ListFavorites com ChooserID e ListID diferentes devem retornar false")
}
