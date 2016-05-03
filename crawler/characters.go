// Character List Crawler
package crawler

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/elct9620/granblue.api/granblue"
)

// Regexp Rules
var (
	nameRule = regexp.MustCompile(`(.*)\(.*\)`)
)

type CharacterParser struct {
	body *goquery.Selection
}

func (p *CharacterParser) SetBody(body *goquery.Selection) {
	p.body = body
}

func (p *CharacterParser) Parse() {
	characters := p.GetCharacters()

	for _, v := range characters {
		fmt.Printf("%v\n", v)
	}
}

// Get the character last update time
// TODO(elct9620): Convert string into time.Time for more control
func (p *CharacterParser) GetLastUpdateTime() string {
	title := p.body.Find("h2").Text()
	updateTime := regexp.MustCompile(`.*\((\d+\/\d+\/\d+)\)`)
	match := updateTime.FindStringSubmatch(title)
	return match[1]
}

// Get All character from wiki
func (p *CharacterParser) GetCharacters() []granblue.Character {
	characters := []granblue.Character{}

	p.body.Find("table>tbody>tr").Each(func(index int, node *goquery.Selection) {
		characters = append(characters, getCharacterFromNode(node))
	})

	return characters
}

func getCharacterFromNode(node *goquery.Selection) granblue.Character {
	character := granblue.Character{}

	node.Children().EachWithBreak(func(index int, node *goquery.Selection) bool {
		text := node.Text()
		switch index { // Assert which column must match which filed
		case 0: // Number
			character.Numbers, _ = strconv.Atoi(text)
		case 1: // Rank
			character.Rank = convertStringToRank(text)
		case 2: // Name
			character.Name = getCharacterName(text)
		case 3: // Element
			character.Element = convertStringToElement(text)
		case 4: // Job
			character.Job = convertStringToJob(text)
		case 5: // Race
			character.Race = convertStringToRace(text)
		case 6: // Gender
			character.Gender = convertStringToGender(text)
		case 7: // Preferred Weapon
			character.PreferredWeapon = convertStringToPreferredWeapon(text)
		case 8: // MinHP
			character.MinHP, _ = strconv.Atoi(text)
		case 9: // MinATK
			character.MinATK, _ = strconv.Atoi(text)
		case 10: // MaxHP
			character.MaxHP, _ = strconv.Atoi(text)
		case 11: // MaxATK
			character.MaxATK, _ = strconv.Atoi(text)
		case 12: // Not used exit
			return false
		}
		return true
	})

	return character
}

func convertStringToRank(rank string) granblue.Rank {
	switch rank {
	case "SSR":
		return granblue.SSR
	case "SR":
		return granblue.SR
	case "R":
		return granblue.R
	}
	return granblue.R // No found should be lower rank
}

func getCharacterName(name string) string {
	matched := nameRule.FindStringSubmatch(name)
	if len(matched) < 1 {
		return "N/A"
	}
	return matched[1]
}

func convertStringToElement(element string) granblue.Element {
	switch element {
	case "風":
		return granblue.Wind
	case "水":
		return granblue.Water
	case "土":
		return granblue.Earth
	case "火":
		return granblue.Fire
	case "光":
		return granblue.Light
	case "闇":
		return granblue.Dark
	}
	return granblue.Wind // No match must return "Wind" because GBF has most Wind element character
}

func convertStringToJob(job string) granblue.JobType {

	switch job {
	case "攻撃":
		return granblue.Attack
	case "防御":
		return granblue.Defense
	case "回復":
		return granblue.Heal
	case "特殊":
		return granblue.Special
	case "バランス":
		return granblue.Balance
	}

	return granblue.UnknownJob
}

func convertStringToRace(race string) granblue.Race {
	switch race {
	case "ヒューマン":
		return granblue.Human
	case "エルーン":
		return granblue.Erun
	case "ドラフ":
		return granblue.Doraf
	case "ハーヴィン":
		return granblue.Harbin
	}
	return granblue.UnknowRace
}

func convertStringToGender(gender string) granblue.Gender {
	switch gender {
	case "♂":
		return granblue.Male
	case "♀":
		return granblue.Female
	}

	return granblue.UnknownGender
}

func convertStringToPreferredWeapon(weapon string) []granblue.WeaponType {
	// TODO(elct9620): Multiple weapon assert
	weapons := []granblue.WeaponType{}
	switch weapon {
	case "剣":
		weapons = append(weapons, granblue.Sword)
	case "弓":
		weapons = append(weapons, granblue.Bow)
	case "斧":
		weapons = append(weapons, granblue.Axe)
	case "杖":
		weapons = append(weapons, granblue.Staff)
	case "格闘":
		weapons = append(weapons, granblue.Fist)
	case "楽器":
		weapons = append(weapons, granblue.Harp)
	case "槍":
		weapons = append(weapons, granblue.Spear)
	case "短剣":
		weapons = append(weapons, granblue.Dagger)
	case "銃":
		weapons = append(weapons, granblue.Gun)
	}

	if len(weapons) <= 0 {
		weapons = append(weapons, granblue.UnknowWeapon)
	}

	return weapons
}
