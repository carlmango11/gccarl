package parser

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeneratedParser(t *testing.T) {
	const gr = `main:
  main:statement*
statement:
  short:NUM
  sum:NUM PLUS NUM SEMI?
  group:LPAREN statement+ RPAREN
  pair:IDEN IDEN
  list:LSQUARE NUM* NUM RSQUARE
  ones:LESS_THAN NUM+ GREATER_THAN
`
	const tests = `package example
import ("testing"; "github.com/carlmango11/gccarl/gccarl/tokens")
func TestRuntime(t *testing.T) {
 if len(TokenDefinitions) == 0 || TokenDefinitions[0].Name != "IDEN" { t.Fatal("missing exported token definitions") }
 for _, text := range []string{"1 + 2;", "3 + 4"} {
  tree, err := Parse(tokens.NewReader(TokenDefinitions, text))
  if err != nil { t.Fatal(err) }
  if len(tree.Main.Statement) != 1 { t.Fatal(tree) }
  sum := tree.Main.Statement[0].Sum
  if sum == nil || string(sum.NUM0) != text[:1] || string(sum.NUM1) != text[4:5] { t.Fatalf("wrong sum: %#v", sum) }
 }
 tree, err := Parse(tokens.NewReader(TokenDefinitions, "(1 2)"))
 if err != nil || len(tree.Main.Statement[0].Group.Statement) != 2 { t.Fatalf("group: %#v %v", tree, err) }
 tree, err = Parse(tokens.NewReader(TokenDefinitions, "foo bar"))
 if err != nil || tree.Main.Statement[0].Pair.IDEN0 != "foo" || tree.Main.Statement[0].Pair.IDEN1 != "bar" { t.Fatalf("pair: %#v %v", tree, err) }
 tree, err = Parse(tokens.NewReader(TokenDefinitions, "[1 2 3]"))
 if err != nil || len(tree.Main.Statement[0].List.NUM0) != 2 || tree.Main.Statement[0].List.NUM1 != "3" {t.Fatalf("list: %#v %v",tree,err)}
 tree, err = Parse(tokens.NewReader(TokenDefinitions, "<1 2>"))
 if err != nil || len(tree.Main.Statement[0].Ones.NUM) != 2 {t.Fatalf("ones: %#v %v",tree,err)}
 tree, err = Parse(tokens.NewReader(TokenDefinitions, ""))
 if err != nil || len(tree.Main.Statement) != 0 { t.Fatalf("empty: %#v %v", tree, err) }
 for _, text := range []string{"1 +", "()", "1;", "@", "(1", "[]", "<>"} {
  if tree, err := Parse(tokens.NewReader(TokenDefinitions, text)); err == nil || tree != nil { t.Fatalf("accepted %q: %#v", text, tree) }
 }
}
func TestReaderPosition(t *testing.T) {
 reader := tokens.NewReader(TokenDefinitions, "skip 5 + 6")
 if _,err:=reader.Next();err!=nil {t.Fatal(err)}
 tree,err:=Parse(reader)
 if err!=nil || tree.Main.Statement[0].Sum.NUM1!="6" {t.Fatalf("reader position: %v %v",tree,err)}
 tree,err=Parse(reader)
 if err!=nil || len(tree.Main.Statement)!=0 {t.Fatalf("reader was not consumed: %v %v",tree,err)}
}
func TestIndependentCalls(t *testing.T) {
 for i:=0; i<8; i++ { t.Run("parse",func(t *testing.T){ t.Parallel(); tree,err:=Parse(tokens.NewReader(TokenDefinitions, "5 + 6")); if err!=nil || tree.Main.Statement[0].Sum.NUM1!="6" {t.Fatalf("%v %v",tree,err)} }) }
}
`
	runGeneratedTests(t, gr, tests)
}

func runGeneratedTests(t *testing.T, gr, tests string) {
	t.Helper()
	p, err := New(strings.NewReader(gr), false)
	require.NoError(t, err)
	out := t.TempDir()
	require.NoError(t, p.Generate(strings.NewReader(tokenDef), out, "example"))
	module, err := filepath.Abs("..")
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(filepath.Join(out, "go.mod"), []byte("module generatedtest\n\ngo 1.25\nrequire github.com/carlmango11/gccarl/gccarl v0.0.0\nreplace github.com/carlmango11/gccarl/gccarl => "+module+"\n"), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(out, "example", "parser_test.go"), []byte(tests), 0600))
	cmd := exec.Command("go", "test", "-race", "./...")
	cmd.Dir = out
	cmd.Env = append(os.Environ(), "GOWORK=off")
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "%s", output)
}

func TestRejectUnsupportedGrammars(t *testing.T) {
	for _, tc := range []struct{ name, grammar, want string }{
		{"left recursion", "main:\n  main:main NUM\n", "left recursion"},
		{"indirect left recursion", "main:\n  main:other\nother:\n  a:main\n", "left recursion"},
		{"nullable prefix recursion", "main:\n  main:item main NUM\nitem:\n  a:IDEN?\n", "left recursion"},
		{"nullable repetition", "main:\n  main:item*\nitem:\n  a:NUM?\n", "nullable"},
		{"missing main", "other:\n  a:NUM\n", "main rule"},
		{"unknown token", "main:\n  main:UNKNOWN\n", "undefined token"},
		{"colliding option", "main:\n  type:NUM\n", "option field"},
		{"colliding rules", "main:\n  main:main-type\nmain-type:\n  a:NUM\n", "duplicate generated identifier"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			p, err := New(strings.NewReader(tc.grammar), false)
			require.NoError(t, err)
			err = p.Generate(strings.NewReader(tokenDef), t.TempDir(), "example")
			require.ErrorContains(t, err, tc.want)
		})
	}
}

func TestGeneratedNullableAndAmbiguous(t *testing.T) {
	runGeneratedTests(t, `main:
  number:NUM
  duplicate:NUM
  empty:item
item:
  value:IDEN?
`, `package example
import ("testing"; "strings"; "github.com/carlmango11/gccarl/gccarl/tokens")
func TestNullableAndAmbiguous(t *testing.T) {
 tree,err:=Parse(tokens.NewReader(TokenDefinitions, ""))
 if err!=nil || tree.Empty==nil || tree.Empty.Item.Value.IDEN!="" {t.Fatalf("nullable: %#v %v",tree,err)}
 tree,err=Parse(tokens.NewReader(TokenDefinitions, "name"))
 if err!=nil || tree.Empty.Item.Value.IDEN!="name" {t.Fatalf("nonempty: %#v %v",tree,err)}
 if tree,err:=Parse(tokens.NewReader(TokenDefinitions, "1"));tree!=nil || err==nil || !strings.Contains(err.Error(),"ambiguous") {t.Fatalf("ambiguous: %#v %v",tree,err)}
}
`)
}

func TestGeneratedUnexpectedTokenLine(t *testing.T) {
	runGeneratedTests(t, "main:\n  main:NUM*\n", `package example
import ("strings"; "testing"; "github.com/carlmango11/gccarl/gccarl/tokens")
func TestSourceLine(t *testing.T) {
 for _,tc:=range []struct{text,line string}{
  {"  + 1\n", "  + 1"},
  {"1\n\t2 + 3  \n4", "\t2 + 3  "},
  {"1\n  2 +", "  2 +"},
  {"1\n  2 2\n  2 + 2", "  2 + 2"},
 } {
  _,err:=Parse(tokens.NewReader(TokenDefinitions, tc.text))
  if err==nil || !strings.Contains(err.Error(),"unexpected token") || !strings.HasSuffix(err.Error(),"\n"+tc.line) {t.Fatalf("expected source line %q in error: %v",tc.line,err)}
 }
}
`)
}

func TestGeneratedReaderInterface(t *testing.T) {
	runGeneratedTests(t, "main:\n  main:NUM\n", `package example
import ("errors"; "io"; "strings"; "testing"; "github.com/carlmango11/gccarl/gccarl/tokens")
type tokenReader struct { token *tokens.Token; err error }
func (r *tokenReader) Next() (*tokens.Token,error) {
 if r.token!=nil {token:=r.token;r.token=nil;return token,nil}
 if r.err!=nil {return nil,r.err}
 return nil,io.EOF
}
func TestReaderInterface(t *testing.T) {
 var reader tokens.Reader=&tokenReader{token:&tokens.Token{Name:"NUM",Val:"42"}}
 tree,err:=Parse(reader)
 if err!=nil || tree.Main.NUM!="42" {t.Fatalf("custom reader: %v %v",tree,err)}
 _,err=Parse(&tokenReader{token:&tokens.Token{Name:"PLUS",Val:"+"}})
 if err==nil || !strings.Contains(err.Error(),"unexpected token") {t.Fatalf("syntax error: %v",err)}
 failure:=errors.New("reader failed")
 _,err=Parse(&tokenReader{err:failure})
 if !errors.Is(err,failure) {t.Fatalf("reader error: %v",err)}
}
`)
}
