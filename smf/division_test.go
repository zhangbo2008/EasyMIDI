package smf

///////////////////////////////
//写下如何调用

//下载别人的代码,然后把他clone到goroot/src里面才行,才能调用!
//因为别人的import 都是github.com/作者这么写的.所以也在自己的上面src目录里面放这个目录才行.
//dsfdsfsdf
import "testing"

func TestCheckSMTPE(t *testing.T) {

	//arrange
	testData := []int8{NOSMTPE, SMTPE24, SMTPE25, SMTPE29, SMTPE30}

	for _, smpte := range testData {

		//act
		result := CheckSMTPE(smpte)

		//assert
		if !result {
			t.Errorf("Test case %b wait true but was false", smpte)
		}
	}
}

func TestCheckWrongSMTPE(t *testing.T) {

	//arrange
	smpte := NOSMTPE + 1

	//act
	result := CheckSMTPE(NOSMTPE + 1)

	//assert
	if result {
		t.Errorf("Test case %b wait false but was true", smpte)
	}
}

func TestNewDivision(t *testing.T) {

	//arrange
	smpte := NOSMTPE
	ticks := TicksMaxValue

	//act
	result, err := NewDivision(ticks, smpte)

	//assert
	if err != nil {
		t.Errorf("%s", err)
	}

	if result.GetSMTPE() != smpte {
		t.Errorf("Wrong SMPTE")
	}

	if result.GetTicks() != ticks {
		t.Errorf("Wrong ticks")
	}
}

func TestNewDivisionSMPTE(t *testing.T) {

	//arrange
	smpte := SMTPE30
	ticks := SMTPETicksMaxValue

	//act
	result, err := NewDivision(ticks, smpte)

	//assert
	if err != nil {
		t.Errorf("%s", err)
	}

	if result.GetSMTPE() != smpte {
		t.Errorf("Wrong SMPTE")
	}

	if result.GetTicks() != ticks {
		t.Errorf("Wrong ticks")
	}
}

func TestNewDivisionTickOutRange(t *testing.T) {

	//arrange
	smpte := NOSMTPE
	ticks := TicksMaxValue + 1

	//act
	_, err := NewDivision(ticks, smpte)

	//assert
	if err == nil {
		t.Errorf("Wait for error but was nil")
	}
}

func TestNewDivisionTickOutRangeSMPTE(t *testing.T) {

	//arrange
	smpte := SMTPE25
	ticks := TicksMaxValue

	//act
	_, err := NewDivision(ticks, smpte)

	//assert
	if err == nil {
		t.Errorf("Wait for error but was nil")
	}
}

func TestNewDivisionWrongSMPTE(t *testing.T) {

	//arrange
	smpte := NOSMTPE + 1

	//act
	_, err := NewDivision(0, smpte)

	//assert
	if err == nil {
		t.Errorf("Wait for error but was nil")
	}
}

func TestIsSMPTE(t *testing.T) {

	//arrange
	testData := []int8{NOSMTPE, SMTPE24, SMTPE25, SMTPE29, SMTPE30}

	for _, smpte := range testData {

		//act
		division, _ := NewDivision(0, smpte)
		result := division.IsSMTPE()

		//assert
		if result != (smpte != NOSMTPE) {
			t.Errorf("Test case %x wait true but was false", smpte)
		}
	}
}
