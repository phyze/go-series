package mol_1

import "multi_module/mol_2"

func GetValueModule2() string {
	return mol_2.NewModule("this is module 2").Field
}
