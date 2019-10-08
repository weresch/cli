package v7

import (
	"strconv"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/v7/shared"
	"code.cloudfoundry.org/cli/util/ui"
	"code.cloudfoundry.org/clock"
)

//go:generate counterfeiter . BuildpacksActor

type BuildpacksActor interface {
	GetBuildpacks(labelSelector string) ([]v7action.Buildpack, v7action.Warnings, error)
}

type BuildpacksCommand struct {
	usage           interface{} `usage:"CF_NAME buildpacks [--labels SELECTOR]\n\nEXAMPLES:\n   CF_NAME buildpacks\n   CF_NAME buildpacks --labels 'environment in (production,staging),tier in (backend)'\n   CF_NAME buildpacks --labels 'env=dev,!chargeback-code,tier in (backend,worker)'"`
	relatedCommands interface{} `related_commands:"create-buildpack, delete-buildpack, rename-buildpack, update-buildpack"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       BuildpacksActor
	Labels      string `long:"labels" description:"Selector to filter buildpacks by labels"`
}

func (cmd *BuildpacksCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	sharedActor := sharedaction.NewActor(config)
	cmd.SharedActor = sharedActor

	ccClient, uaaClient, err := shared.GetNewClientsAndConnectToCF(config, ui, "")
	if err != nil {
		return err
	}
	cmd.Actor = v7action.NewActor(ccClient, config, sharedActor, uaaClient, clock.NewClock())

	return nil
}

func (cmd BuildpacksCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(false, false)
	if err != nil {
		return err
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

	cmd.UI.DisplayTextWithFlavor("Getting buildpacks as {{.Username}}...", map[string]interface{}{
		"Username": user.Name,
	})
	cmd.UI.DisplayNewline()

	buildpacks, warnings, err := cmd.Actor.GetBuildpacks(cmd.Labels)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}

	if len(buildpacks) == 0 {
		cmd.UI.DisplayTextWithFlavor("No buildpacks found")
	} else {
		cmd.displayTable(buildpacks)
	}
	return nil
}

func (cmd BuildpacksCommand) displayTable(buildpacks []v7action.Buildpack) {
	if len(buildpacks) > 0 {
		var keyValueTable = [][]string{
			{"position", "name", "stack", "enabled", "locked", "filename"},
		}
		for _, buildpack := range buildpacks {
			keyValueTable = append(keyValueTable, []string{
				strconv.Itoa(buildpack.Position.Value),
				buildpack.Name,
				buildpack.Stack,
				strconv.FormatBool(buildpack.Enabled.Value),
				strconv.FormatBool(buildpack.Locked.Value),
				buildpack.Filename,
			})
		}

		cmd.UI.DisplayTableWithHeader("", keyValueTable, ui.DefaultTableSpacePadding)
	}
}