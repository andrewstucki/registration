#Registration Server
##Mac

###Getting Started

As always the use of the [Homebrew](brew.sh) package manager is highly recommended. For the build itself you'll need a gcc compiler toolchain and automake (both of which you can get through [Xcode](developer.apple.com/xcode)) as well as a few other dependencies. OS X comes with most of what you'll need out of the box, but you'll need to additionally install [Mercurial](mercurial.selenic.com), [NodeJS](nodejs.org), and [Go](golang.org). With Homebrew you can simply:

	brew install mercurial
	brew install node
	brew install go

###Building

Once you have installed the above, change into the top-level directory of the project and run:

	make tools
	make deps
	make

The first two commands install the final required build tools and dependencies while the last actually builds the project. If you make any changes to the client or server source simply run `make` again. The output is found in `bin/register`.

##Windows

###Getting Started

Windows requires a little work to get it building. The following components are required before building

####A 64-bit [MinGW-w64](mingw-w64.sourceforge.net) toolchain

Currently the 32-bit version of mingw doesn't ship with the wlanapi library. This is the library used for automatic ad-hoc configuration on Windows. For development I use the mingw-builds toolchain that comes with gcc 4.7.3 found on [SourceForge](http://sourceforge.net/projects/mingwbuilds/files/host-windows/releases/4.7.3/64-bit/threads-posix/sjlj/x64-4.7.3-release-posix-sjlj-rev1.7z). I have had a few problems building with any of the 4.8.x branches by mingw-builds as go-sqlite3 fails to compile complaining about an undefined _localtime32 symbol. I imagine this has something to do with the differing MSVCRT symbol definitions between the mingw64 revisions used in the 4.8.x builds and the 4.7.x builds, but in any event, if you are having problems with the _localtime32 symbol and are using a different toolchain than me try using the patch for go-sqlite3 found in the `patches` directory. The main problem with the patch is that to redistribute the binary you'll have to figure out a way of packaging the msvcr10 redistributable.

####[Ruby](www.ruby-lang.org)

The client css is written in [Sass](sass-lang.com) and leverages [Compass](http://compass-style.org/) this is the only Ruby-dependant piece of the build, so chances are I will eventually convert all the files to use [Less](lesscss.org) in order to simplify the build process (the css is quite minimal anyway).

####[NodeJS](nodejs.org)

The client is built with the wonderful [Grunt](gruntjs.com) tool and utilizes tons of things like uglifyjs, bower, and coffeescript.

####64-bit [Go](golang.org)

Because this project utilizes a very small amount of C, cgo is required to build it. Go's build command has the C compiler hardcoded to the string "gcc" and can't cross-compile cgo code without some hackery. As a result, the go architecture needs to be 64-bit because the program can only compile with a 64-bit mingw.

####[Git](git-scm.com) and [Mercurial](mercurial.selenic.com)

The `go get` command fetches a bunch of dependencies out of both git and mercurial repositories, so both tools are needed.

####[MSYS](http://www.mingw.org/wiki/MSYS)

The whole MSYS package is not necessary, the only utility necessary for the build is msys-zip, a Windows port of the Linux `zip` command. It is used in embedding the server resources into the final binary.

###Building

Once you have all of the above dependencies installed and available on your path, build by running the following batch scripts in the top level directory of the project

	tools
	deps
	build

After initially installing the build tools and addiitonal dependencies with the above scripts, if you need to make any modifications to the server or client simply recompile with `build`. The output will be found in `bin/register.exe`.