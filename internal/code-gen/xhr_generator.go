package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log/slog"
	"slices"
	"unicode"

	"github.com/dave/jennifer/jen"
)

type ValueType struct {
	Value ValueTypes `json:"value"`
}

type ValueTypes struct {
	Values    []ValueType
	Value     *ValueType
	ValueName string
}

type ExtAttr struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Rhs  struct {
		Type  string     `json:"type"`
		Value ValueTypes `json:"value"`
	} `json:"rhs"`
}

type IdlType struct {
	Type     string    `json:"type"`
	ExtAttrs []ExtAttr `json:"extAttrs"`
	Generic  string    `json:"generic"`
	Nullable bool      `json:"nullable"`
	Union    bool      `json:"union"`
	IType    IdlTypes  `json:"idlType"`
}

type IdlTypes struct {
	Types    []IdlType
	IdlType  *IdlType
	TypeName string
}

func (i IdlTypes) String() string {
	if len(i.Types) > 0 {
		return fmt.Sprintf("%v", i.Types)
	}
	if i.IdlType != nil {
		return fmt.Sprintf("%v", *i.IdlType)
	}
	return i.TypeName
}

func (t *ValueTypes) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &t.Values)
	if err != nil {
		val := new(ValueType)
		err = json.Unmarshal(bytes, val)
		if err == nil {
			t.Value = val
		}
	}
	if err != nil {
		err = json.Unmarshal(bytes, &t.ValueName)
	}
	return err
}

func (t *IdlTypes) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &t.Types)
	if err != nil {
		typ := new(IdlType)
		err = json.Unmarshal(bytes, &typ)
		if err == nil {
			t.IdlType = typ
		}
	}
	if err != nil {
		err = json.Unmarshal(bytes, &t.TypeName)
	}
	return err

}

type Stuff struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	ExtAttrs []ExtAttr `json:"extAttrs"`
	IdlType  IdlTypes  `json:"idlType"`
}

type ArgumentType struct {
	Stuff
	Default  any  `json:"default"`
	Optional bool `json:"optional"`
	Variadic bool `json:"variadic"`
}

type IdlNameMember struct {
	Stuff
	Arguments []ArgumentType `json:"arguments"`
	Special   string         `json:"special"`
	Readonly  bool           `json:"readOnly"`
	Href      string         `json:"href"`
}

type IdlName struct {
	Type    string          `json:"type"`
	Name    string          `json:"name"`
	Members []IdlNameMember `json:"members"`
	Partial bool            `json:"partial"`
	Href    string          `json:"href"`
}

type IdlParsed struct {
	IdlNames map[string]IdlName
}

type ParsedIdlFile struct {
	IdlParsed `json:"idlParsed"`
}

//go:embed webref/curated/idlparsed/xhr.json
var xhrData []byte

func FindIdlType(idl IdlTypes) string {
	types := idl.Types
	if len(types) == 0 && idl.IdlType != nil {
		types = []IdlType{*idl.IdlType}
	}
	for _, t := range types {
		if t.Type == "return-type" {
			return t.IType.TypeName
		}
	}
	return ""

}

func createData() (ESConstructorData, error) {
	spec := ParsedIdlFile{}
	err := json.Unmarshal(xhrData, &spec)
	if err != nil {
		panic(err)
	}
	idlName := spec.IdlNames["XMLHttpRequest"]
	type tmp struct {
		Op ESOperation
		Ok bool
	}
	ops := []*tmp{}
	for _, member := range idlName.Members {
		if member.Type == "operation" {
			if -1 != slices.IndexFunc(
				ops,
				func(op *tmp) bool { return op.Op.Name == member.Name },
			) {
				slog.Warn("Function overloads", "Name", member.Name)
				continue
			}
			returnType := FindIdlType(member.IdlType)
			operation := &tmp{ESOperation{
				Name:       member.Name,
				ReturnType: returnType,
				Arguments:  []ESOperationArgument{},
			}, true}
			ops = append(ops, operation)
			for _, arg := range member.Arguments {
				esArg := ESOperationArgument{
					Name: arg.Name,
				}
				if len(arg.IdlType.Types) > 0 {
					slog.Warn(
						"Multiple argument types",
						"Operation",
						member.Name,
						"Argument",
						arg.Name,
					)
					operation.Ok = false
					break
				}
				if arg.IdlType.IdlType != nil {
					esArg.Type = arg.IdlType.IdlType.IType.TypeName
				}
				operation.Op.Arguments = append(operation.Op.Arguments, esArg)
			}
		}
	}
	// fmt.Println(operations)
	// os.Exit(1)

	operations := make([]ESOperation, 0, len(ops))
	for _, op := range ops {
		if op.Ok {
			operations = append(operations, op.Op)
		}
	}
	data := ESConstructorData{
		InnerTypeName:      "XmlHttpRequest",
		Operations:         operations,
		CreatesInnerType:   true,
		CustomConstruction: "scriptContext.Window().NewXmlHttpRequest()",
		IdlName:            idlName,
	}
	return data, nil
}

const br = "github.com/stroiman/go-dom/browser"
const sc = "github.com/stroiman/go-dom/scripting"
const v8 = "github.com/tommie/v8go"

func generateXhr(b *builder) error {

	file := jen.NewFilePath(sc)
	file.HeaderComment("This file is generated. Do not edit.")
	file.ImportName(br, "browser")
	file.ImportAlias(v8, "v8")
	data, err := createData()
	if err != nil {
		return err
	}
	writeFactory(file, data)
	return file.Render(b)
}

type ESOperationArgument struct {
	Name     string
	Type     string
	Optional bool
	Variadic bool
}

type ESOperation struct {
	Name       string
	ReturnType string
	Arguments  []ESOperationArgument
}

type ESConstructorData struct {
	CreatesInnerType   bool
	InnerTypeName      string
	CustomConstruction string
	Operations         []ESOperation
	IdlName
}

type Imports = [][][2]string

func WriteImports(b *builder, imports Imports) {
	b.Printf("import (\n")
	b.indent()
	defer b.unIndentF(")\n\n")
	for _, grp := range imports {
		for _, imp := range grp {
			alias := imp[0]
			pkg := imp[1]
			if alias == "" {
				b.Printf("\"%s\"\n", pkg)
			} else {
				b.Printf("%s \"%s\"\n", imp[0], imp[1])
			}
		}
	}
}

type st = *jen.Statement

type JSConstructor struct {
	argHost          st
	argInfo          st
	getThis          st
	varInstance      st
	varIso           st
	varScriptContext st
}

func CreateJSConstructor() JSConstructor {
	argInfo := jen.Id("info")
	argHost := jen.Id("host")
	getThis := argInfo.Clone().Dot("This").Call()
	varInstance := jen.Id("instance")
	varIso := jen.Id("iso")
	varScriptContext := jen.Id("scriptContext")
	return JSConstructor{
		argHost,
		argInfo,
		getThis,
		varInstance,
		varIso,
		varScriptContext,
	}
}

func (c JSConstructor) JSConstructorImpl(grp *jen.Group) {
	buildInstance := c.varScriptContext.Clone().
		Dot("Window").
		Call().
		Dot("NewXmlHttpRequest").
		Call()
	grp.Add(c.varScriptContext).
		Op(":=").
		Add(c.argHost).
		Dot("MustGetContext").
		Call(c.argInfo.Clone().Dot("Context").Call())
	grp.Add(c.varInstance).Op(":=").Add(buildInstance)
	grp.Return(c.varScriptContext.Clone().Dot("CacheNode").Call(
		c.getThis,
		c.varInstance,
	))
}

func (c JSConstructor) Run(f *jen.File, data ESConstructorData) {
	hostType := jen.Id("ScriptHost")
	hostPtr := jen.Add(jen.Op("*"), hostType)
	builder := jen.Id("builder")
	v8FunctionTemplatePtr := jen.Op("*").Qual(v8, "FunctionTemplate")
	v8Value := jen.Op("*").Qual(v8, "Value")
	errorT := jen.Id("error")
	g := Helper{&jen.Group{}}
	f.Func().
		Id(fmt.Sprintf("Create%sPrototype", data.InnerTypeName)).
		Params(c.argHost.Clone().Add(hostPtr)).Add(v8FunctionTemplatePtr).
		BlockFunc(func(grp *jen.Group) {
			grp.Add(c.varIso).Op(":=").Add(jen.Id("host")).Dot("iso").GoString()
			grp.Add(builder).
				Op(":=").
				Id("NewConstructorBuilder").Index(jen.Qual(br, data.InnerTypeName)).
				Call(c.argHost.Clone(), jen.Func().
					Params(c.argInfo.Clone().Add(g.v8FunctionCallbackInfoPtr())).
					Params(v8Value, errorT).
					BlockFunc(c.JSConstructorImpl))
			grp.Id("protoBuilder").Op(":=").Add(builder).Dot("NewPrototypeBuilder").Call()
			grp.Id("prototype").Op(":=").Id("protoBuilder").Dot("proto")
			grp.Line()
			for _, op := range data.Operations {
				c.CreateOperation(grp, op)
			}

			grp.Add(builder).Dot("SetDefaultInstanceLookup").Call()
			grp.Return(builder.Clone().Dot("constructor"))
		})
}

type JenGenerator interface {
	Generate() *jen.Statement
}

type GetArgStmt struct {
	Name    string
	ErrName string
	Getter  string
	Index   int
}

type IfStmt struct {
	Condition JenGenerator
	Block     JenGenerator
	Else      JenGenerator
}

func (s GetArgStmt) Generate() *jen.Statement {
	return AssignmentStmt{
		[]string{s.Name, s.ErrName},
		Stmt{jen.Id(s.Getter).Call(jen.Id("args"), jen.Lit(s.Index))},
	}.Generate()
}

type AssignmentStmt struct {
	VarNames   []string
	Expression JenGenerator
}

type StatementListStmt struct {
	Statements []JenGenerator
}

func (s AssignmentStmt) Generate() *jen.Statement {
	result := jen.Id(s.VarNames[0])
	for _, name := range s.VarNames[1:] {
		result.Op(",").Id(name)
	}
	result.Op(":=").Add(s.Expression.Generate())
	return result
}

func (s IfStmt) Generate() *jen.Statement {
	result := jen.If(s.Condition.Generate()).Block(s.Block.Generate())
	if s.Else != nil {
		result.Else().Block(s.Else.Generate())
	}
	return result
}

func GetSliceLength(gen JenGenerator) JenGenerator {
	return Stmt{jen.Len(gen.Generate())}
}

type GeneratorFunc func() *jen.Statement

func (g GeneratorFunc) Generate() *jen.Statement {
	return g.Generate()
}

func Statements(stmts ...JenGenerator) JenGenerator {
	return StatementListStmt{stmts}
}

func (s *StatementListStmt) Append(stmt JenGenerator) {
	s.Statements = append(s.Statements, stmt)
}

func (s StatementListStmt) Generate() *jen.Statement {
	result := []jen.Code{}
	g := jen.Group{}
	for i, s := range s.Statements {
		if i > 0 {
			result = append(result, jen.Line())
		}
		g.Add(s.Generate())
		jenStatement := s.Generate()
		if jenStatement != nil {
			result = append(result, jenStatement)
		}
	}
	// return g
	jenStmt := jen.Statement(result)
	return &jenStmt
}

type CallInstance struct {
	Name string
	Args []string
	Op   ESOperation
}

type GetGeneratorResult struct {
	Generator      JenGenerator
	RequireContext bool
}

func (c CallInstance) GetGenerator() GetGeneratorResult {
	var requireContext bool
	args := []jen.Code{}
	// jErr := jen.Id("err")
	// returnStmt := IfStmt{
	// 	Condition: Stmt{jErr.Clone().Op("!=").Nil()},
	// 	Block:     Stmt{jen.Return(jen.Nil(), jErr)},
	// }
	for _, a := range c.Args {
		args = append(args, jen.Id(a))
	}
	dest := jen.Id("err")
	if c.Op.ReturnType != "undefined" {
		dest = jen.Id("result").Op(",").Add(dest).Op(":=")
		// returnStmt.Else = Stmt{jen.Return(jen.Nil(), jen.Nil())}
	} else {
		dest = dest.Op("=")
	}
	list := StatementListStmt{}
	list.Append(Stmt{
		dest.Id("instance").Dot(camelCase(c.Name)).Call(args...),
	})
	if c.Op.ReturnType == "undefined" {
		list.Append(Stmt{jen.Return(jen.Nil(), jen.Id("err"))})
	} else {
		converter := "To" + c.Op.ReturnType
		requireContext = true
		list.Append(IfStmt{
			Condition: Stmt{jen.Id("err").Op("!=").Nil()},
			Block:     Stmt{jen.Return(jen.Nil(), jen.Id("err"))},
			Else:      Stmt{jen.Return(jen.Id(converter).Call(jen.Id("ctx"), jen.Id("result")))},
		})
	}
	// list.Append(returnStmt)
	return GetGeneratorResult{list, requireContext}
}

func (c JSConstructor) CreateOperation(grp *jen.Group, op ESOperation) {
	v8Value := jen.Op("*").Qual(v8, "Value")
	errorT := jen.Id("error")
	v8FunctionCallbackInfoPtr := jen.Op("*").Qual(v8, "FunctionCallbackInfo")
	f := jen.Func().
		Params(c.argInfo.Clone().Add(v8FunctionCallbackInfoPtr)).
		Params(v8Value, errorT).
		BlockFunc(func(grp *jen.Group) {
			grp.Add(jen.Id("instance")).Op(",").Id("err").
				Op(":=").
				Id("builder").
				Dot("GetInstance").
				Call(jen.Id("info"))
			WriteReturnOnError(grp)
			if len(op.Arguments) == 0 {
				generatorResult := CallInstance{
					Name: op.Name,
					Args: []string{},
					Op:   op,
				}.GetGenerator()
				if generatorResult.RequireContext {
					grp.Add(
						jen.Id("ctx").
							Op(":=").
							Id("host").
							Dot("MustGetGetContext").
							Call(jen.Id("host").Dot("Context").Call()),
					)
				}
				grp.Add(generatorResult.Generator.Generate())

				// grp.Add(
				// 	jen.Id("instance").Dot(camelCase(op.Name)).Call(),
				// )
				// WriteErrorHandler(grp, len(op.Arguments))
			} else {
				grp.Id("args").Op(":=").Id("info").Dot("Args").Call()
				grp.Add(
					AssignmentStmt{
						[]string{"argsLen"},
						GetSliceLength(Stmt{jen.Id("args")}),
					}.Generate(),
				)
				statements := &StatementListStmt{}
				outer := IfStmt{
					Block: statements,
				}
				inner := &outer

				argCount := len(op.Arguments)
				// opName := op.Name
				var args []string
				for i, arg := range op.Arguments {
					// targetBlock := statements
					// targetFunc := opName
					if arg.Optional {
						inner = new(IfStmt)
						statements.Append(inner)
						statements = &StatementListStmt{}
						inner.Block = statements
					}
					args = append(args, arg.Name)
					inner.Condition = (Stmt{jen.Id("argsLen").Op(">=").Lit(i + 1)})
					var errName string
					if argCount > 1 {
						errName = fmt.Sprintf("err%d", i)
					} else {
						errName = fmt.Sprintf("err")
					}
					stmt := GetArgStmt{
						Name:    arg.Name,
						ErrName: errName,
						Getter:  "GetArg" + arg.Type,
						Index:   i,
					}
					statements.Append(stmt)
					// statements.Append(stmt)
					// grp.Add(stmt.Generate())
				}
				statements.Append(genErrorHandler(len(op.Arguments)))
				generatorResult := CallInstance{
					Name: op.Name,
					Args: args,
					Op:   op,
				}.GetGenerator()
				statements.Append(generatorResult.Generator)
				if generatorResult.RequireContext {
					grp.Add(jen.Id("ctx").Op(":=").Id("host").Dot("MustGetContext").Call(jen.Id("info").Dot("Context").Call()))
				}
				grp.Add(outer.Generate())
				// WriteErrorHandler(grp, len(op.Arguments))
			}
			// WriteReturnOnError(grp)
			// grp.Id("instance").Dot(camelCase(opName)).Call(args...)
			// grp.Return(jen.Nil(), jen.Nil())
		})
	ft := jen.Qual(v8, "NewFunctionTemplateWithError").Call(c.varIso, f)
	grp.Id("prototype").Dot("Set").Call(jen.Lit(op.Name), ft)
	grp.Line()
}

func camelCase(s string) string {
	buffer := make([]rune, 0, len(s))
	buffer = append(buffer, unicode.ToUpper([]rune(s)[0]))
	buffer = append(buffer, []rune(s)[1:]...)
	return string(buffer)
}

type Stmt struct{ *jen.Statement }

func (s Stmt) Generate() *jen.Statement { return s.Statement }

func GenReturnOnError() JenGenerator {
	jErr := jen.Id("err")
	stmt := IfStmt{
		Condition: Stmt{jErr.Clone().Op("!=").Nil()},
		Block:     Stmt{jen.Return(jen.Nil(), jErr)},
	}
	return stmt
}

func WriteReturnOnError(grp *jen.Group) {
	grp.Add(GenReturnOnError().Generate())
}

func Noop() JenGenerator {
	return GeneratorFunc(func() *jen.Statement {
		empty := []jen.Code{}
		emptyStmt := jen.Statement(empty)
		return &emptyStmt
	})
}

func genErrorHandler(count int) JenGenerator {
	if count == 0 {
		return StatementListStmt{}
	}
	jErr := jen.Id("err")
	result := StatementListStmt{}
	if count > 1 {
		var args []jen.Code
		for i := 0; i < count; i++ {
			args = append(args, jen.Id(fmt.Sprintf("err%d", i)))
		}
		s := Stmt{jErr.Clone().Op(":=").Qual("errors", "Join").Call(args...)}
		result.Append(s)
	}
	result.Append(GenReturnOnError())
	return result

}

func WriteErrorHandler(grp *jen.Group, count int) {
	// if count == 0 {
	// 	return
	// }
	grp.Add(genErrorHandler(count).Generate())
}

func writeFactory(f *jen.File, data ESConstructorData) {
	CreateJSConstructor().Run(f, data)
}

type Helper struct{ *jen.Group }

func (h Helper) BuildInstance() *jen.Statement {
	return h.scriptContext().Dot("Window").Call().Dot("NewXmlHttpRequest").Call()
}

func (h Helper) v8FunctionCallbackInfoPtr() *jen.Statement {
	return h.Op("*").Qual(v8, "FunctionCallbackInfo")
}
func (h Helper) hostArg() *jen.Statement {
	return h.Id("info")
}
func (h Helper) infoArg() *jen.Statement {
	return h.Id("info")
}
func (h Helper) scriptContext() *jen.Statement {
	return h.Id("scriptContext")
}

// func WriteOperations(b *builder, data ESConstructorData) {
// 	for _, member := range data.IdlName.Members {
// 		if member.Type == "operation" {
// 			b.Printf("%s: %s\n", member.Name, member.Type)
//
// 			b.indent()
// 			for _, a := range member.Arguments {
// 				b.Printf("%s: %s - %v\n", a.Name, a.Type, a.IdlType)
// 			}
// 			b.unIndent()
// 		}
// 	}
// }
