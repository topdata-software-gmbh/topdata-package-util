package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-util/commands/cache_commands"
	"github.com/topdata-software-gmbh/topdata-package-util/commands/localgit_commands"
	"github.com/topdata-software-gmbh/topdata-package-util/commands/pkg_commands"
	"github.com/topdata-software-gmbh/topdata-package-util/config"
	"github.com/topdata-software-gmbh/topdata-package-util/globals"
	"os"
)

var appRootCommand = &cobra.Command{
	Use:   "main",
	Short: "The entrypoint",
}

func Execute() {
	if err := appRootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var PackagePortfolioFile string

func init() {
	appRootCommand.PersistentFlags().StringVar(&PackagePortfolioFile, "portfolio-file", "/topdata/topdata-package-portfolio/portfolio.yaml", "config file") // FIXME: make the default path configurable (~/.config/topdata-package-util.config ?)
	pkg_commands.Register(appRootCommand)
	cache_commands.Register(appRootCommand)
	localgit_commands.Register(appRootCommand)

	globals.PkgConfigList = config.LoadPackagePortfolioFile(PackagePortfolioFile)
}
