package shccn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestNew(t *testing.T) {
	actual, err := New("../../test/test.sh")
	assert.NotNil(t, actual)
	assert.NoError(t, err)
}

func TestFileContents_TestGetLines(t *testing.T) {
	testContents, err := New("../../test/test.sh")
	assert.NoError(t, err)
	actual := testContents.GetLines()
	expected := 32
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestFileContents_TestGetBlanks(t *testing.T) {
	testContents, err := New("../../test/test.sh")
	assert.NoError(t, err)
	actual := testContents.GetBlankLines()
	expected := 4
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestFileContents_GetComments(t *testing.T) {
	testContents, err := New("../../test/test.sh")
	assert.NoError(t, err)
	actual := testContents.GetCommentLines()
	expected := 3
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestFileContents_GetCodes(t *testing.T) {
	testContents, err := New("../../test/test.sh")
	assert.NoError(t, err)
	actual := testContents.GetCodeLines()
	expected := 25
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestFileContents_GetFunctions(t *testing.T) {
	testContents, err := New("../../test/test.sh")
	assert.NoError(t, err)
	actual := testContents.GetFunctionLines()
	expected := 4
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestGetCodes(t *testing.T) {
	tc, err := New("../../test/test.sh")
	assert.NoError(t, err)
	actual := GetCodes(tc.Lines)
	expected := []string{
		`export aaa="aaa"`,
		`function test_method() {`,
		`    echo "test"  # comment test3(コメントではなく、code)`,
		`}`,
		`function test_method2() {`,
		`    echo "test"  # comment test3(コメントではなく、code)`,
		`}`,
		`test_method3() {`,
		`    if [[ "a" == "a" ]] || [[ "b" == "b" ]]; then`,
		`        echo "code with condition"`,
		`    fi`,
		`}`,
		`test_method4() {`,
		`    case $1 in`,
		`       "MON") echo "Monday.";;`,
		`       "TUE") echo "Tuesday.";;`,
		`       "WED") echo "Wednesday.";;`,
		`       "THU") echo "Thursday.";;`,
		`       "FRI") echo "Friday.";;`,
		`       "SAT") echo "Saturday.";;`,
		`       "SUN") echo "Sunday.";;`,
		`    esac`,
		`}`,
		`test_method3`,
	}
	assert.Equal(t, expected, actual)
}

func TestGetFunctionName(t *testing.T) {
	actual := GetFunctionName("function test_method() {")
	expected := "test_method"
	assert.Equal(t, expected, actual)
}

func TestGetFunctionName2(t *testing.T) {
	actual := GetFunctionName("hogehoge(){")
	expected := "hogehoge"
	assert.Equal(t, expected, actual)
}

func TestGetFunctions(t *testing.T) {
	tc, err := New("../../test/test.sh")
	assert.NoError(t, err)
	execCodes := GetCodes(tc.Lines)
	actual := GetFunctions(execCodes)

	expected := make(map[string][]string)
	expected["BARE_CODE"] = append(expected["BARE_CODE"], `export aaa="aaa"`)
	expected["BARE_CODE"] = append(expected["BARE_CODE"], `test_method3`)
	expected["test_method"] = append(expected["test_method"], `    echo "test"  # comment test3(コメントではなく、code)`)
	expected["test_method2"] = append(expected["test_method2"], `    echo "test"  # comment test3(コメントではなく、code)`)
	expected["test_method3"] = append(expected["test_method3"], `    if [[ "a" == "a" ]] || [[ "b" == "b" ]]; then`)
	expected["test_method3"] = append(expected["test_method3"], `        echo "code with condition"`)
	expected["test_method3"] = append(expected["test_method3"], `    fi`)
	expected["test_method4"] = append(expected["test_method4"], `    case $1 in`)
	expected["test_method4"] = append(expected["test_method4"], `       "MON") echo "Monday.";;`)
	expected["test_method4"] = append(expected["test_method4"], `       "TUE") echo "Tuesday.";;`)
	expected["test_method4"] = append(expected["test_method4"], `       "WED") echo "Wednesday.";;`)
	expected["test_method4"] = append(expected["test_method4"], `       "THU") echo "Thursday.";;`)
	expected["test_method4"] = append(expected["test_method4"], `       "FRI") echo "Friday.";;`)
	expected["test_method4"] = append(expected["test_method4"], `       "SAT") echo "Saturday.";;`)
	expected["test_method4"] = append(expected["test_method4"], `       "SUN") echo "Sunday.";;`)
	expected["test_method4"] = append(expected["test_method4"], `    esac`)
	assert.Equal(t, expected, actual)

}

func TestCalculateCCN(t *testing.T) {
	tc, err := New("../../test/test.sh")
	assert.NoError(t, err)
	execCodes := GetCodes(tc.Lines)
	functions := GetFunctions(execCodes)
	actual := CalculateCCN(functions["test_method3"])
	expected := 3
	assert.Equal(t, expected, actual)
}

func TestCalculateCCN2(t *testing.T) {
	tc, err := New("../../test/test.sh")
	assert.NoError(t, err)
	execCodes := GetCodes(tc.Lines)
	functions := GetFunctions(execCodes)
	actual := CalculateCCN(functions["test_method4"])
	expected := 8
	assert.Equal(t, expected, actual)
}

func TestIsBlankLine(t *testing.T) {
	actual := isBlankLine("    ")
	assert.True(t, actual)
}

func TestIsBlankLine2(t *testing.T) {
	actual := isBlankLine("  bbb")
	assert.False(t, actual)
}

func TestIsCommentLine(t *testing.T) {
	actual := isCommentLine("  #aaaaa")
	assert.True(t, actual)
}

func TestIsCommentLine2(t *testing.T) {
	actual := isCommentLine("  aaaaa # bbbb")
	assert.False(t, actual)
}

func TestIsCommentLine3(t *testing.T) {
	actual := isCommentLine("  '#aaaa'")
	assert.False(t, actual)
}

func TestIsFunctionLine(t *testing.T) {
	actual := isFunctionLine("function testA() {", true)
	assert.True(t, actual)
}

func TestIsFunctionLine2(t *testing.T) {
	actual := isFunctionLine("testA() {", true)
	assert.True(t, actual)
}

func TestIsFunctionLine3(t *testing.T) {
	actual := isFunctionLine("testA(){", true)
	assert.True(t, actual)
}

func TestIsFunctionLine4(t *testing.T) {
	actual := isFunctionLine("$(${aaaa})", true)
	assert.False(t, actual)
}

func TestIsFunctionLine5(t *testing.T) {
	actual := isFunctionLine("'testA(){'", true)
	assert.False(t, actual)
}

func TestRemoveQuote(t *testing.T) {
	actual := removeQuote(`'aaaaaaaaa'`)
	expected := ``
	assert.Equal(t, expected, actual)
}

func TestRemoveQuote2(t *testing.T) {
	actual := removeQuote(`'aaaaaaaaa'bbb'cccccccc'`)
	expected := `bbb`
	assert.Equal(t, expected, actual)
}

func TestRemoveQuote3(t *testing.T) {
	actual := removeQuote(`echo '${script_name} s|p'`)
	expected := `echo `
	assert.Equal(t, expected, actual)
}

func TestIntegration(t *testing.T) {
	tc, err := New("../../scripts/tshare.bash")
	assert.NoError(t, err)
	lines := tc.GetLines()
	codes := tc.GetCodeLines()
	comments := tc.GetCommentLines()
	blanks := tc.GetBlankLines()
	functions := tc.GetFunctionLines()

	execCodes := GetCodes(tc.Lines)
	functionCodes := GetFunctions(execCodes)

	sl := make([]string, len(functionCodes))
	i := 0
	for k := range functionCodes {
		sl[i] = k
		i++
	}
	sort.Strings(sl)

	assert.NoError(t, err)
	fmt.Print(BuildSummaryHeader())
	fmt.Print(BuildSummaryBody(tc.Name, lines, codes, comments, blanks, functions))
	fmt.Print(BuildFooter())

	fmt.Print(BuildFunctionHeader())
	for _, k := range sl {
		name := k
		code := len(functionCodes[k])
		ccn := CalculateCCN(functionCodes[k])
		fmt.Print(BuildFunctionBody(tc.Name, name, code, ccn))
	}
	fmt.Print(BuildFooter())
}
