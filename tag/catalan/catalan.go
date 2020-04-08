package catalan

import (
	"github.com/kenkyu392/go-tm/tag"
)

/*
ca-AD,Catalan (Andorra)
ca-ES,Catalan (Spain)
ca-FR,Catalan (France)
ca-IT,Catalan (Italy)
*/

func init() {
	tag.Register("ca-AD", "Catalan (Andorra)")
	tag.Register("ca-ES", "Catalan (Spain)")
	tag.Register("ca-FR", "Catalan (France)")
	tag.Register("ca-IT", "Catalan (Italy)")
}
