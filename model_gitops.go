package main

type ModelGitOps struct {
	github *GitHub
	git    *GitOperator
	plugin GitOpsPlugin
}

type GitOpsPrepareOutput struct {
	PullRequestID     string
	PullRequestNumber int
	Branch            string
	status            DeployStatus
}

func (self GitOpsPrepareOutput) Status() DeployStatus {
	return self.status
}

func (self GitOpsPrepareOutput) Message() string {
	return "Success to deploy"
}

func (self ModelGitOps) Commit(pullRequestID string) error {
	return self.github.MergePullRequest(pullRequestID)

}

func (self ModelGitOps) Deploy(pj DeployProject, phase string, option DeployOption) (do DeployOutput, err error) {
	o, err := self.plugin.Prepare(pj, phase, option.Branch, option.Assigner, option.Tag)
	if err != nil {
		return
	}
	if o.Status() == DeployStatusSuccess {
		err = self.Commit(o.PullRequestID)
		if err != nil {
			return
		}
	}
	return o, nil
}
