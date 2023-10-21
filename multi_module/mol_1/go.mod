module mol_1

replace (
	//source latest => target latest
	multi_module/mol_2 => ../mol_2
)

require (
	multi_module/mol_2 v0.0.0
)

go 1.20
