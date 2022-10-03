package controllers

import (
	"github.com/google/go-github/v47/github"
	"github.com/haiau/actions-runner-controller/api/v1alpha1"
)

func (autoscaler *HorizontalRunnerAutoscalerGitHubWebhook) MatchPushEvent(event *github.PushEvent) func(scaleUpTrigger v1alpha1.ScaleUpTrigger) bool {
	return func(scaleUpTrigger v1alpha1.ScaleUpTrigger) bool {
		g := scaleUpTrigger.GitHubEvent

		if g == nil {
			return false
		}

		push := g.Push

		return push != nil
	}
}
