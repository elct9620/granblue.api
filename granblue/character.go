// Character Model
package granblue

type Gender int

const (
	UnknownGender Gender = iota
	Male
	Female
)

type Character struct {
	Numbers int // From Granblue Fantasy Wiki
	Name    string
	// TODO(elct9620): Provide translatable string for Chinese name
	Rank Rank
	// TODO(elct9620): Add Skin with a set of image
	Element Element
	Race    Race
	Job     JobType
	Gender  Gender

	// Attribute
	MinHP  int
	MaxHP  int
	MinATK int
	MaxATK int

	PefferredWeapon WeaponType

	// TODO(elct9620): Add Ability and Support Ability
	// TODO(elct9620): Add Ougi(奧義) information
}
