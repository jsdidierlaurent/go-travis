// Copyright (c) 2015 Ableton AG, Berlin. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package travis

// MinimalCommit is a minimal representation of a GitHub commit as seen by Travis CI
//
// Travis CI API docs: https://developer.travis-ci.com/resource/commit#minimal-representation
type MinimalCommit struct {
	// Value uniquely identifying the commit
	Id uint `json:"id,omitempty"`
	// Checksum the commit has in git and is identified by
	Sha string `json:"sha,omitempty"`
	// Named reference the commit has in git.
	Ref string `json:"ref,omitempty"`
	// Commit mesage
	Message string `json:"message,omitempty"`
	// URL to the commit's diff on GitHub
	CompareUrl string `json:"compare_url,omitempty"`
	// Commit date from git
	CommittedAt string `json:"committed_at,omitempty"`
}

// StandardCommit is a standard representation of a GitHub commit as seen by Travis CI
//
// Travis CI API docs: https://developer.travis-ci.com/resource/commit#standard-representation
type StandardCommit struct {
	// Value uniquely identifying the commit
	Id uint `json:"id,omitempty"`
	// Checksum the commit has in git and is identified by
	Sha string `json:"sha,omitempty"`
	// Named reference the commit has in git.
	Ref string `json:"ref,omitempty"`
	// Commit mesage
	Message string `json:"message,omitempty"`
	// URL to the commit's diff on GitHub
	CompareUrl string `json:"compare_url,omitempty"`
	// Commit date from git
	CommittedAt string `json:"committed_at,omitempty"`
	// Committer of commit from git
	Committer Committer `json:"committer,omitempty"`
	// Author of commit from git
	Author Author `json:"author,omitempty"`
}

type Committer struct {
	// Name of the committer
	Name string `json:"name,omitempty"`
	// AvatarUrl of the committer
	AvatarUrl string `json:"avatar_url,omitempty"`
}

type Author struct {
	// Name of the author
	Name string `json:"name,omitempty"`
	// AvatarUrl of the author
	AvatarUrl string `json:"avatar_url,omitempty"`
}

