package gonan

import "fmt"

type Truth struct {
	culprit Character
	cause   string
}

type TruthOption func(*Truth) error

func NewTruth(culprit Character, cause string) (*Truth, error) {
	t := &Truth{
		culprit: culprit,
		cause:   cause,
	}

	return t, nil
}

type Case struct {
	victime      Character
	causeOfDeath string
	location     string
	suspects     map[Character]bool
}

type CaseOption func(*Case) error

func WithVictim(v Character) CaseOption {
	return func(c *Case) error {
		c.victime = v
		return nil
	}
}

func WithCauseOfDeath(s string) CaseOption {
	return func(c *Case) error {
		c.causeOfDeath = s
		return nil
	}
}

func WithLocation(s string) CaseOption {
	return func(c *Case) error {
		c.location = s
		return nil
	}
}

func AddSuspect(s Character, isCulprit bool) CaseOption {
	return func(c *Case) error {
		c.suspects[s] = isCulprit
		return nil
	}
}

type Detective interface {
	solve() bool
}

func (c *Case) Close(detective Character, t *Truth) error {
	d, ok := detective.(Detective)
	if !ok {
		return fmt.Errorf("%s is not detective", detective.Name())
	}

	if !d.solve() {
		return fmt.Errorf("%s is don't solve", detective.Name())
	}

	Say(
		detective,
		"やっと分かったんですよ...",
		fmt.Sprintf("%sで%sさんを殺害した犯人がね...", c.location, c.victime.DisplayName()),
		fmt.Sprintf("犯人は%sさん、あなただ!", t.culprit.DisplayName()),
		fmt.Sprintf("犯行の動機はおそらく%sでしょう", t.cause),
	)
	return nil
}

func NewCase(optons ...CaseOption) (*Case, error) {
	c := new(Case)
	c.suspects = make(map[Character]bool)
	for _, opt := range optons {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}
