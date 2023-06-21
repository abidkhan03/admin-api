package service

import (
	"context"
	p "github.com/IamFaizanKhalid/pointer"
	"github.com/spongeling/admin-api/internal/freeling"
	"github.com/spongeling/admin-api/internal/repo"
	"github.com/spongeling/admin-api/internal/request"
	"github.com/spongeling/admin-api/internal/service/model"
	"github.com/spongeling/admin-api/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_IsExampleValid(t *testing.T) {
	ctx := context.Background()
	trequire := require.New(t)
	tassert := assert.New(t)

	err := shared.LoadConfig("../../.env.test")
	trequire.NoError(err)

	db, err := repo.New(shared.GetDBConnectionString())
	trequire.NoError(err)

	fl, err := freeling.NewClient()
	trequire.NoError(err)

	svc := New(db, nil, fl)

	classId, err := svc.AddClass(ctx, request.Class{
		Name:        "Color",
		Description: "Color names",
		Words:       []string{"blanco", "negro", "rojo"},
	})
	trequire.NoError(err)

	defer func() {
		err := svc.DeleteClass(ctx, classId)
		trequire.NoError(err)
	}()

	testIt := func(pattern []model.Token, example string, expected bool) {
		ans := svc.IsExampleValid(example, pattern)
		tassert.Equal(expected, ans, "example -> %s\npattern -> %v", example, pattern)
		if ans != expected {
			tassert.FailNow("")
		}
	}

	examples := []string{
		"el libro blanco",
		"la casa blanca",
		"el coche negro",
		"el vestido rojo", // failing because FreeLing parse vestido (dress) as a verb (VMP00SM)
		// Also, rojo (red) is being treated as a noun (NCMS000)
		"el coche amarillo",
		"el vestido verde", // failing because verde (AQ0CS00) is common,
		// so also being returned for masculine (AQ0MS00), which is unable to match
		"el parque tranquilo y hermoso",
	}

	// pos only
	pattern := []model.Token{{Pos: p.String("DA0MS0")}, {Pos: p.String("NCMS000")}, {Pos: p.String("AQ0MS00")}}
	testIt(pattern, examples[0], true)
	testIt(pattern, examples[1], false) // incorrect because it is feminine
	testIt(pattern, examples[2], true)
	//testIt(pattern, examples[3], true)
	testIt(pattern, examples[4], true)
	//testIt(pattern, examples[5], true)
	testIt(pattern, examples[6], false) // different number of tokens

	// check irrespective of gender
	pattern = []model.Token{{Pos: p.String("DA00S0")}, {Pos: p.String("NC0S000")}, {Pos: p.String("AQ00S00")}}
	testIt(pattern, examples[0], true)
	testIt(pattern, examples[1], true)
	testIt(pattern, examples[2], true)
	//testIt(pattern, examples[3], true)
	testIt(pattern, examples[4], true)
	//testIt(pattern, examples[5], true)
	testIt(pattern, examples[6], false) // different number of tokens

	// with class ["blanco", "negro", "rojo"]
	pattern = []model.Token{{Pos: p.String("DA00S0")}, {Pos: p.String("NC0S000")}, {Class: p.String("Color")}}
	testIt(pattern, examples[0], true)
	testIt(pattern, examples[1], true)
	testIt(pattern, examples[2], true)
	//testIt(pattern, examples[3], true)
	testIt(pattern, examples[4], false) // `amarillo` not in the list
	testIt(pattern, examples[5], false) // `verde` not in the list
	testIt(pattern, examples[6], false) // different number of tokens

	// check with exact word
	// FIXME: Use {Word: "la"}
	pattern = []model.Token{{Word: p.String("el")}, {Pos: p.String("NC0S000")}, {Pos: p.String("AQ00S00")}}
	testIt(pattern, examples[0], true) // all these should be true
	testIt(pattern, examples[1], true) // because `la` and `el`
	testIt(pattern, examples[2], true) // both are different form
	//testIt(pattern, examples[3], true) // of the same word `el`
	testIt(pattern, examples[4], true)
	testIt(pattern, examples[5], true)
	testIt(pattern, examples[6], false) // different number of tokens

	// test with blank
	// FIXME: Use {Word: "la"}
	pattern = []model.Token{{Word: p.String("el")}, {}, {Pos: p.String("AQ00S00")}}
	testIt(pattern, examples[0], true)
	testIt(pattern, examples[1], true)
	testIt(pattern, examples[2], true)
	//testIt(pattern, examples[3], true)
	testIt(pattern, examples[4], true)
	testIt(pattern, examples[5], true)
	testIt(pattern, examples[6], true) // blank can be filled with any number of tokens

	// test with multiple blanks
	pattern = []model.Token{{}, {Pos: p.String("NC0S000")}, {}, {Pos: p.String("AQ00S00")}}
	testIt(pattern, examples[0], false) // Coordinating Conjunction
	testIt(pattern, examples[1], false) // not in the phrases
	testIt(pattern, examples[2], false)
	//testIt(pattern, examples[3], false)
	testIt(pattern, examples[4], false)
	testIt(pattern, examples[5], false)
	testIt(pattern, examples[6], true)

	// test with redundant blank
	// FIXME: Use {Word: "la"}
	pattern = []model.Token{{Word: p.String("el")}, {Pos: p.String("NC0S000")}, {}, {Pos: p.String("AQ00S00")}}
	testIt(pattern, examples[0], false)
	testIt(pattern, examples[1], false)
	testIt(pattern, examples[2], false)
	//testIt(pattern, examples[3], false)
	testIt(pattern, examples[4], false)
	testIt(pattern, examples[5], false)
	testIt(pattern, examples[6], true) // it has extra tokens to fill in the blank
}
