package providers_test

import (
	"testing"

	"github.com/sidan-lab/rum/interfaces"
	"github.com/sidan-lab/rum/providers"
)

func TestMaestroServiceImplementsIFetcher(t *testing.T) {
	var _ interfaces.IFetcher = (*providers.MaestroService)(nil)
}
func TestMaestroServiceImplementsISubmitter(t *testing.T) {
	var _ interfaces.ISubmitter = (*providers.MaestroService)(nil)
}
