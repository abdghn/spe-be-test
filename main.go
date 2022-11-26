package main

type SpeSkillTest interface {
	NarcissisticNumber(input int) bool
	NeedleInAHaystack(s []string, str string) int
	ParityOutlier(input int) interface{}
}

type speSkillTest struct {
}

func New() SpeSkillTest {
	return &speSkillTest{}
}

func (t speSkillTest) NarcissisticNumber(input int) bool {
	return isNarc(input)
}

func (t speSkillTest) NeedleInAHaystack(s []string, str string) int {
	return needleInAHayStack(s, str)
}

func (t speSkillTest) ParityOutlier(input int) interface{} {
	return parityOutlier(input)
}

func main() {

}
