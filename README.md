# Autobrew

Automate homebrew formula publishing for your CLI tools, regardless of programming language.

✅ Discover facts about your project such as name, version, description, homepage, and more.  
✅ Find a latest binary release for OSX and run a SHA256 over the content.  
✅ Compose and render an appropriate Homebrew formula file.  
✅ Push a commit with the new (or updated) formula.  




## Quick Start

You should have:

* A Github project that publishes binary releases (e.g. [rawsort](https://github.com/jondot/rawsort))
* A Github repo that serves as your Homebrew tap (e.g. [jondot/homebrew-tap](https://github.com/jondot/homebrew-tap))
* A Github token (set scope to public only, for your open-source projects, which is a good practice)

And now to publish a new Homebrew formula:

```
$ export AUTOBREW_GITHUB_TOKEN=xxxx
$ autobrew --user jondot --project=rawsort --tap=homebrew-tap
```

For more options see:

```
$ autobrew --help
usage: autobrew --user=USER --project=PROJECT --tap=TAP --github-token=GITHUB-TOKEN [<flags>]

Flags:
      --help               Show context-sensitive help (also try --help-long and --help-man).
  -u, --user=USER          Github user
  -p, --project=PROJECT    Github project
  -b, --tap=TAP            Homebrew tap
      --tap-user=TAP-USER  Homebrew tap user (default to user)
  -t, --github-token=GITHUB-TOKEN
                           Github token
      --version            Show application version.
```

### Thanks:

To all [Contributors](https://github.com/jondot/autobrew/graphs/contributors) - you make this happen, thanks!

# Copyright

Copyright (c) 2018 [@jondot](http://twitter.com/jondot). See [LICENSE](LICENSE.txt) for further details.