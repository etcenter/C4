package id_test

import (
	// "bytes"

	"math/big"
	"strings"
	"testing"

	c4 "github.com/Avalanche-io/c4/id"
	"github.com/cheekybits/is"
)

func TestIDSliceSort(t *testing.T) {
	is := is.New(t)
	var b, s []byte
	for i := 0; i < 64; i++ {
		b = append(b, 0xFF)
		s = append(s, 0x00)
	}
	bigBig := big.NewInt(0)
	bigSmall := big.NewInt(0)
	bigBig = bigBig.SetBytes(b)
	bigSmall = bigSmall.SetBytes(s)
	bigID := c4.ID(*bigBig)
	smallID := c4.ID(*bigSmall)

	var idSlice c4.IDSlice

	idSlice.Push(&bigID)
	idSlice.Push(&smallID)
	is.Equal(idSlice[0].String(), `c467rpwLCuS5DGA8KGZXKsVQ7dnPb9goRLoKfgGbLfQg9WoLUgNY77E2jT11fem3coV9nAkguBACzrU1iyZM4B8roQ`)
	is.Equal(idSlice[1].String(), `c41111111111111111111111111111111111111111111111111111111111111111111111111111111111111111`)
	idSlice.Sort()
	is.Equal(idSlice[0].String(), `c41111111111111111111111111111111111111111111111111111111111111111111111111111111111111111`)
	is.Equal(idSlice[1].String(), `c467rpwLCuS5DGA8KGZXKsVQ7dnPb9goRLoKfgGbLfQg9WoLUgNY77E2jT11fem3coV9nAkguBACzrU1iyZM4B8roQ`)
}

func TestIDSliceString(t *testing.T) {
	is := is.New(t)

	var ids c4.IDSlice
	id1, err := c4.Identify(strings.NewReader("foo"))
	is.NoErr(err)
	id2, err := c4.Identify(strings.NewReader("bar"))
	is.NoErr(err)

	ids.Push(id1)
	ids.Push(id2)

	is.Equal(ids.String(), id1.String()+id2.String())
}

func TestIDSliceSearchIDs(t *testing.T) {
	is := is.New(t)

	var ids c4.IDSlice
	id1, err := c4.Identify(strings.NewReader("foo"))
	is.NoErr(err)
	id2, err := c4.Identify(strings.NewReader("bar"))
	is.NoErr(err)
	id3, err := c4.Identify(strings.NewReader("baz"))
	is.NoErr(err)

	ids.Push(id1)
	ids.Push(id2)
	ids.Push(id3)
	ids.Sort()

	is.True(id2.Less(id1))
	is.True(id3.Less(id2))

	is.Equal(c4.SearchIDs(ids, id1), 2)
	is.Equal(c4.SearchIDs(ids, id2), 1)
	is.Equal(c4.SearchIDs(ids, id3), 0)
}

func TestSliceIDFile(t *testing.T) {
	is := is.New(t)

	id, err := c4.Identify(errorReader(true))
	is.Err(err)
	is.Nil(id)
	is.Equal(err.Error(), "errorReader triggered error.")
}