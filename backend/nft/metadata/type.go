package metadata

import (
	"errors"
	"orbit_nft/util"
	"strings"
)

const NftCharacter = "character"
const NftPet = "pet"

const ErrTokenTypeNotSupport = "token type does not support"

type baseToken struct {
	Name        string
	Description string
	Image       string
	Class       string
	Grade       string
}

type character struct {
	Base         baseToken
	GradeBonus   string
	AwakenLevel  string
	GenesisBonus string
	Damage       int
	Armor        int
	HP           int
	ATKSpeed     int
	CritChance   string
	Mana         string
	Aura         string
	Ultimate     string
	Abilities    string
}

func generateToken(tokenType, baseURIToken, imageCid string, record []string) (interface{}, error) {
	var imageLink string
	var err error
	var dame, armor, hp, atkSpeed int

	tokenType = strings.ToLower(tokenType)
	imageLink, err = util.ToLink(baseURIToken, imageCid)
	if err != nil {
		return nil, err
	}

	bToken := baseToken{
		Name:        record[0],
		Description: record[1],
		Image:       imageLink,
		Class:       record[3],
		Grade:       record[4],
	}
	// tính toán các chỉ số liên quan đến NFT Character
	if strings.EqualFold(tokenType, NftCharacter) {
		dame, err = util.RandomMinToMax(record[8], record[9])
		if err != nil {
			return nil, err
		}
		armor, err = util.RandomMinToMax(record[10], record[11])
		if err != nil {
			return nil, err
		}
		hp, err = util.RandomMinToMax(record[12], record[13])
		if err != nil {
			return nil, err
		}
		atkSpeed, err = util.RandomMinToMax(record[14], record[15])
		if err != nil {
			return nil, err
		}
	}

	switch tokenType {
	case NftCharacter:
		return character{
			Base:         bToken,
			GradeBonus:   record[5],
			AwakenLevel:  record[6],
			GenesisBonus: record[7],
			Damage:       dame,
			Armor:        armor,
			HP:           hp,
			ATKSpeed:     atkSpeed,
			CritChance:   record[16],
			Mana:         record[17],
			Aura:         record[18],
			Ultimate:     record[19],
			Abilities:    record[20],
		}, nil
	default:
		return nil, errors.New(ErrTokenTypeNotSupport)
	}
}
