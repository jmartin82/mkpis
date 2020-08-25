package vcs

import "time"

type Client interface {
	GetMergedPRList(owner string, repo string, from time.Time, to time.Time, base string) ([]PR, error)
}

type PR struct {
	Number         int
	CreatedAt      time.Time
	MergedAt       time.Time
	Base           string
	ChangedFiles   int
	ChangedLines   int
	ReviewComments int
	Commits        int
	Head           string
	FirstCommitAt  time.Time
	LastCommitAt   time.Time
	FirstCommentAt time.Time
	LastCommentAt  time.Time
}

func (pr *PR) PRLeadTime() time.Duration {
	return pr.MergedAt.Sub(pr.CreatedAt)
}

func (pr *PR) TimeToMerge() time.Duration {
	return pr.MergedAt.Sub(pr.FirstCommitAt)
}

func (pr *PR) TimeToReview() time.Duration {
	return pr.LastCommentAt.Sub(pr.FirstCommentAt)
}
func (pr *PR) TimeToFirstReview() time.Duration {
	if pr.FirstCommentAt.IsZero() {
		return 0
	}
	return pr.FirstCommentAt.Sub(pr.CreatedAt)
}
func (pr *PR) LastReviewToMerge() time.Duration {
	if pr.LastCommentAt.IsZero() {
		return 0
	}
	return pr.MergedAt.Sub(pr.LastCommentAt)
}
