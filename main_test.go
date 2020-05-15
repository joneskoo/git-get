package main

import "testing"

func Test_expand(t *testing.T) {
	type args struct {
		input         string
		defaultPrefix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"plain name without project or protocol",
			args{"hello", "git@github.com:"},
			"hello",
		},
		{
			"absolute HTTPS URI should be unmodified",
			args{"https://github.com/joneskoo/git-get", "git@github.com:"},
			"https://github.com/joneskoo/git-get",
		},
		{
			"absolute short ssh address should be unmodified",
			args{"git@github.com:joneskoo/git-get.git", "git@github.com:"},
			"git@github.com:joneskoo/git-get.git",
		},

		{
			"absolute ssh URI should be unmodified",
			args{"ssh://git@github.com:joneskoo/git-get.git", "git@github.com:"},
			"ssh://git@github.com:joneskoo/git-get.git",
		},
		{
			"relative github path should complete",
			args{"joneskoo/git-get", "git@github.com:"},
			"git@github.com:joneskoo/git-get.git",
		},
		{
			"relative https path should complete",
			args{"joneskoo/git-get", "https://example.com/"},
			"https://example.com/joneskoo/git-get.git",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expand(tt.args.input, tt.args.defaultPrefix); got != tt.want {
				t.Errorf("expand(%q, %q) = %q, want %q", tt.args.input, tt.args.defaultPrefix, got, tt.want)
			}
		})
	}
}

func Test_targetDir(t *testing.T) {
	type args struct {
		cloneURL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "non-URL should return error",
			args:    args{"foobar"},
			wantErr: true,
		},
		{
			name:    "plain hostname is not valid",
			args:    args{"https://github.com"},
			wantErr: true,
		},
		{
			name: "ssh short form address",
			args: args{"user@hostname:project/repo"},
			want: "hostname/project/repo",
		},
		{
			name: "ssh short form address with git suffix",
			args: args{"user@hostname:project/repo.git"},
			want: "hostname/project/repo",
		},
		{
			name: "ok https url",
			args: args{"https://github.com/joneskoo/git-get"},
			want: "github.com/joneskoo/git-get",
		},
		{
			name: "ok https url with git suffix",
			args: args{"https://github.com/joneskoo/git-get.git"},
			want: "github.com/joneskoo/git-get",
		},
		{
			name: "ok ssh url",
			args: args{"ssh://git@github.com/joneskoo/git-get"},
			want: "github.com/joneskoo/git-get",
		},
		{
			name: "ok ssh url with git suffix",
			args: args{"ssh://git@github.com/joneskoo/git-get.git"},
			want: "github.com/joneskoo/git-get",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := targetDir(tt.args.cloneURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("targetDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("targetDir(%q) = %v, want %v", tt.args.cloneURL, got, tt.want)
			}
		})
	}
}
