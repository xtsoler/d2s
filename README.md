This is a fork of the fantastic d2s parser https://github.com/nokka/d2s which will try to handle Project Diablo 2 https://www.projectdiablo2.com/ .d2s save files.

How to jumpstart:

Since I am on windows I downloaded and installed go1.15.6.windows-amd64.msi in the default path.
I also downloaded and installed Visual Studio Code.

Now we are ready to fetch the repo.
From a command line or powershell

go get github.com/xtsoler/d2s

This should download the repo in the "go" folder in the user directory.
For windows it should be 
C:\Users\\{my_username}\go\src\github.com\xtsoler\d2s

This directory can be openned with git / gitExt etc.

Next we select "open folder" from Visual Studio Code or your favorite IDE.
We can make changes in the files, for example in d2s.go or item.go.
Then we can test our changes by running the following from the included in 
the repo corresponding directory with our CLI:

For windows batch:

set "GOARCH=amd64" && set "GOOS=windows" && go build d2s_pd2_run.go && d2s_pd2_run.exe ..\examples\pd2_thebow.d2s

This should give you the log output on the standard output of your cli.
The output json file should also be created in the same directory.

To build for web

batch:

set "GOARCH=wasm" && set "GOOS=js" && go build -o pd2_d2s_parser_v001.wasm pd2_d2s_web.go
