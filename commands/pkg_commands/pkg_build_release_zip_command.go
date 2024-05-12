package pkg_commands

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/topdata-software-gmbh/topdata-package-service/app_constants"
	"github.com/topdata-software-gmbh/topdata-package-service/factory"
	"github.com/topdata-software-gmbh/topdata-package-service/git_cli_wrapper"
	"github.com/topdata-software-gmbh/topdata-package-service/globals"
	"github.com/topdata-software-gmbh/topdata-package-service/pkg_zipper"
	"github.com/topdata-software-gmbh/topdata-package-service/util"
	"path/filepath"
	"time"
)

var buildReleaseZipCommand = &cobra.Command{
	Use:   "build-release-zip [packageName] [releaseBranchName]",
	Short: "Builds a release zip for uploading to the shopware6 plugin store",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		packageName := args[0]
		releaseBranchName := args[1]

		// pathPackagePortfolioFile, _ := cmd.Flags().GetString("portfolio-file")
		// pkgConfigList := config.LoadPackagePortfolioFile(pathPackagePortfolioFile)
		pkgConfig := globals.PkgConfigList.FindOneByNameOrFail(packageName)

		gitBranchInfo := factory.NewGitBranchInfo(*pkgConfig, releaseBranchName)

		// -- switch to the release branch
		git_cli_wrapper.SwitchBranch(*pkgConfig, releaseBranchName)

		//  -- update local git repository
		// TODO... git_cli_wrapper.UpdateRepo(*pkgConfig)

		// -- write release_info.txt

		pathReleaseInfoFile := filepath.Join(pkgConfig.GetLocalGitRepoDir(), app_constants.ReleaseInfoFileName)
		color.Blue("Writing release info to " + pathReleaseInfoFile + " ...")
		now := time.Now().Format(time.RFC3339)
		//		releaseInfo := strings.TrimSpace(util.RenderString(`
		//Commit ID: {commitId}
		//Branch:    {branch}
		//Created:   {created}
		//`, map[string]string{
		//			"commitId": gitBranchInfo.CommitId,
		//			"branch":   gitBranchInfo.Name,
		//			"created":  now,
		//		}))
		releaseInfo := util.MapToTable(map[string]string{
			"Version":   gitBranchInfo.PackageVersion,
			"Created":   now,
			"Branch":    gitBranchInfo.Name,
			"Commit ID": gitBranchInfo.CommitId,
		})
		util.WriteToFile(pathReleaseInfoFile, releaseInfo)

		// -- copy files from pkgConfig.GetLocalGitRepoDir() to temporary folder with foldername same as shopware6 store technical name
		tmpReleaseDir := "/tmp/releases-tmp/" + pkgConfig.Shopware6StoreTechnicalName
		util.RunCommand("mkdir", "-p", tmpReleaseDir)
		util.RsyncDirectory(pkgConfig.GetLocalGitRepoDir(), tmpReleaseDir, []string{".git"})

		//  -- create a zip file
		pathDestZipFile := app_constants.PathReleaseZipsDir + "/" + pkgConfig.Shopware6StoreTechnicalName + "-" + gitBranchInfo.PackageVersion + ".zip"
		color.Blue("Creating zip file " + pathDestZipFile + "...")
		pkg_zipper.CreateZipArchive(tmpReleaseDir, pathDestZipFile)
		// -- TODO: upload the zip file to the shopware6 plugin store

	},
}

func init() {
	pkgRootCommand.AddCommand(buildReleaseZipCommand)
}
