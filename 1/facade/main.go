package main

type IDEfacade struct {
	textEditor *TextEditor
	compiler *Compiler
	clr *CLR
}

func newIDEFacade() *IDEfacade {
	return &IDEfacade{
		textEditor: &TextEditor{},
		compiler: &Compiler{},
		clr: &CLR{},
	}
}

func (idef *IDEfacade) Start() {
	idef.textEditor.writeCode()
	idef.textEditor.saveCode()
	idef.compiler.Compile()
	idef.clr.Execute()
}

func (idef *IDEfacade) Stop() {
	idef.clr.Finish()
}

func main() {
	code := newIDEFacade()
	code.Start()
	code.Stop()
}
