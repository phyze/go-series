module multi_module

replace (
	multi_module/mol_1 => ./mol_1
	multi_module/mol_2 => ./mol_2
)

require multi_module/mol_1 v0.0.0
require multi_module/mol_2 v0.0.0 // indirect

go 1.20
