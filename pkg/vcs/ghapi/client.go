package ghapi

import (
	"context"
	"log"
	"time"

	"github.com/jmartin82/mkpis/pkg/vcs"

	"github.com/google/go-github/v32/github" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"golang.org/x/oauth2"
)

type Client struct {
	c   *github.Client
	ctx context.Context
}

func (cli *Client) connect(accessToken string) {

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(cli.ctx, ts)

	cli.c = github.NewClient(tc)
}

func NewClient(accessToken string) *Client {
	cli := &Client{}
	cli.ctx = context.Background()
	cli.connect(accessToken)
	return cli
}

func (cli *Client) getFirstAndLastCommitTime(owner string, repo string, prNum int) (first time.Time, last time.Time) {
	log.Printf("Getting first and last commit from %d", prNum)
	var commits []*github.RepositoryCommit
	commits, resp, err := cli.c.PullRequests.ListCommits(cli.ctx, owner, repo, prNum, &github.ListOptions{PerPage: 50})
	if err != nil {
		log.Printf("Error getting first commit date: %s\n", err)
		return time.Time{}, time.Time{}
	}
	fcommit := commits[0]
	first = fcommit.GetCommit().Committer.GetDate()
	if resp.NextPage != 0 {
		lastCommits, _, err := cli.c.PullRequests.ListCommits(cli.ctx, owner, repo, prNum, &github.ListOptions{PerPage: 1, Page: resp.LastPage})
		if err != nil {
			log.Printf("Error getting last commit date: %s\n", err)
			return time.Time{}, time.Time{}
		}
		commits = lastCommits
	}

	lcommit := commits[len(commits)-1]

	last = lcommit.GetCommit().Committer.GetDate()
	return
}

func (cli *Client) getFirstAndLastReviewCommentTime(owner string, repo string, prNum int) (first time.Time, last time.Time) {
	log.Printf("Getting first and last comment from %d", prNum)
	var comments []*github.PullRequestReview
	comments, resp, err := cli.c.PullRequests.ListReviews(cli.ctx, owner, repo, prNum, &github.ListOptions{PerPage: 50})
	if err != nil {
		log.Printf("Error getting first review date: %s\n", err)
		return time.Time{}, time.Time{}
	}

	if len(comments) == 0 {
		return time.Time{}, time.Time{}
	}

	fComment := comments[0]
	first = fComment.GetSubmittedAt()

	if resp.NextPage != 0 {
		lastComments, _, err := cli.c.PullRequests.ListReviews(cli.ctx, owner, repo, prNum, &github.ListOptions{PerPage: 1, Page: resp.LastPage})
		if err != nil {
			log.Printf("Error getting last review date: %s\n", err)
			return time.Time{}, time.Time{}
		}
		comments = lastComments
	}

	lComment := comments[len(comments)-1]

	last = lComment.GetSubmittedAt()
	return
}

func (cli *Client) GetMergedPRList(owner string, repo string, from time.Time, to time.Time, base string) ([]vcs.PR, error) {
	var pRNums []int
	var pRList []vcs.PR
	opt := &github.PullRequestListOptions{State: "closed", Base: base /*"devel"*/}
	opt.PerPage = 100
	log.Printf("Fetching Closed PR List from: %s to: %s", from.Format("2006-01-02"), to.Format("2006-01-02"))
pagination:
	for {

		prs, resp, err := cli.c.PullRequests.List(cli.ctx, owner, repo, opt)
		if err != nil {
			return nil, err
		}

		for _, pr := range prs {
			if pr.ClosedAt.Before(from) {
				log.Printf("Discarded PR: %d out of the date range", pr.GetNumber())
				//shortcut to avoid more call I know I don't need the rest
				break pagination
			}

			if pr.GetMergedAt().IsZero() || pr.GetClosedAt().After(to) {
				continue
			}

			pRNums = append(pRNums, pr.GetNumber())
		}
		if resp.NextPage == 0 {
			break
		}

		opt.Page = resp.NextPage
	}
	for _, prNum := range pRNums {
		pr, _, err := cli.c.PullRequests.Get(cli.ctx, owner, repo, prNum)
		if err != nil {
			return nil, err
		}

		log.Printf("Fetching info from PR %d", pr.GetNumber())
		fc, lc := cli.getFirstAndLastCommitTime(owner, repo, pr.GetNumber())
		fr, lr := cli.getFirstAndLastReviewCommentTime(owner, repo, pr.GetNumber())
		pRList = append(pRList, vcs.PR{
			Number:         pr.GetNumber(),
			CreatedAt:      pr.GetCreatedAt(),
			MergedAt:       pr.GetMergedAt(),
			ChangedFiles:   pr.GetChangedFiles(),
			ChangedLines:   pr.GetDeletions() + pr.GetAdditions(),
			ReviewComments: pr.GetReviewComments(),
			Base:           pr.GetBase().GetRef(),
			Head:           pr.GetHead().GetSHA(),
			Commits:        pr.GetCommits(),
			FirstCommitAt:  fc,
			LastCommitAt:   lc,
			FirstCommentAt: fr,
			LastCommentAt:  lr,
		})
	}
	return pRList, nil
}
