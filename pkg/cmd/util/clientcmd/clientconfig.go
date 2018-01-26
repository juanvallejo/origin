package clientcmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/openshift/origin/pkg/cmd/cli/config"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	clientcmdapi "k8s.io/kubernetes/pkg/client/unversioned/clientcmd/api"
)

func DefaultClientConfig(flags *pflag.FlagSet) clientcmd.ClientConfig {
	loadingRules := config.NewOpenShiftClientConfigLoadingRules()
	flags.StringVar(&loadingRules.ExplicitPath, config.OpenShiftConfigFlagName, "", "Path to the config file to use for CLI requests.")
	cobra.MarkFlagFilename(flags, config.OpenShiftConfigFlagName)

	// set our explicit defaults
	defaultOverrides := &clientcmd.ConfigOverrides{ClusterDefaults: clientcmdapi.Cluster{Server: os.Getenv("KUBERNETES_MASTER")}}
	loadingRules.DefaultClientConfig = clientcmd.NewDefaultClientConfig(clientcmdapi.Config{}, defaultOverrides)

	fmt.Printf("%v: clientcmd/clientconfig.go: Overriding default flags.\n", time.Now().Format("15:14:13.123456"))
	overrides := &clientcmd.ConfigOverrides{ClusterDefaults: defaultOverrides.ClusterDefaults}
	overrideFlags := clientcmd.RecommendedConfigOverrideFlags("")
	overrideFlags.ContextOverrideFlags.Namespace.ShortName = "n"
	overrideFlags.AuthOverrideFlags.Username.LongName = ""
	overrideFlags.AuthOverrideFlags.Password.LongName = ""
	clientcmd.BindOverrideFlags(overrides, flags, overrideFlags)
	cobra.MarkFlagFilename(flags, overrideFlags.AuthOverrideFlags.ClientCertificate.LongName)
	cobra.MarkFlagFilename(flags, overrideFlags.AuthOverrideFlags.ClientKey.LongName)
	cobra.MarkFlagFilename(flags, overrideFlags.ClusterOverrideFlags.CertificateAuthority.LongName)

	fmt.Printf("%v: clientcmd/clientconfig.go: Instatiating NonInteractiveDeferredLoadingClientConfig with loadingRules: %v.\n", time.Now().Format("15:14:13.123456"), loadingRules)
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, overrides)

	return clientConfig
}
