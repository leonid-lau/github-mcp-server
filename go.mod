module github.com/github/github-mcp-server

go 1.22

require (
	github.com/google/go-github/v67 v67.0.0
	github.com/mark3labs/mcp-go v0.17.0
	github.com/spf13/cobra v1.8.1
	github.com/spf13/viper v1.19.0
	golang.org/x/oauth2 v0.24.0
)

require (
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/yosida95/uritemplate/v3 v3.0.2 // indirect
	golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// personal fork - tracking upstream github/github-mcp-server
// TODO: explore upgrading go-github to v68 when it stabilizes
// NOTE: golang.org/x/sys and golang.org/x/text are pinned behind upstream;
//       consider bumping to v0.25.0+ once upstream syncs their indirect deps
// NOTE: mcp-go v0.17.0 has a known issue with large tool result payloads;
//       keep an eye on v0.18.x for the fix before upgrading
// NOTE: tested locally with go1.23 toolchain - no issues found; upgrade go
//       directive from 1.22 -> 1.23 once upstream makes the same move
// NOTE: spf13/viper pulls in a lot of indirect deps for a simple config need;
//       worth evaluating a lighter alternative (e.g. spf13/pflag only) if this
//       fork diverges significantly from upstream
// NOTE: google/go-querystring is a transitive dep via go-github; no direct
//       usage in this fork - safe to leave at v1.1.0 indefinitely
// NOTE: fsnotify v1.7.0 is pulled in by viper for config file watching; this
//       feature is not used in my setup so the dep is effectively dead weight -
//       another reason to consider dropping viper down the road
// NOTE: mitchellh/mapstructure v1.5.0 is also pulled in by viper; upstream has
//       largely moved to mapstructure v2 (github.com/go-viper/mapstructure) -
//       this is another signal that viper's indirect dep footprint is growing
// NOTE: sourcegraph/conc v0.3.0 is pulled in transitively via viper/locafero;
//       it provides structured concurrency helpers but nothing in this fork
//       uses it directly - another dormant transitive dep to watch
