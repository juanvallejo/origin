package print

import (
	"io"

	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	kcmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

// VersionedPrintObject handles printing an object in the appropriate version by looking at 'output-version'
// on the command
func VersionedPrintObject(fn func(*cobra.Command, runtime.Object, io.Writer) error, c *cobra.Command, out io.Writer) func(runtime.Object) error {
	return func(obj runtime.Object) error {
		// TODO: fold into the core printer functionality (preferred output version)

		if items, err := meta.ExtractList(obj); err == nil {
			for i := range items {
				items[i] = kcmdutil.AsDefaultVersionedOrOriginal(items[i], nil)
			}
			if err := meta.SetList(obj, items); err != nil {
				return err
			}

		} else {
			obj = kcmdutil.AsDefaultVersionedOrOriginal(obj, nil)
		}
		return fn(c, obj, out)
	}
}
