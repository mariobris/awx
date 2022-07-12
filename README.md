# awx
Changes AWS_PROFILE environment variable in the current shell.

## Installation
You can clone repository and build binary with `go build` and copy it into directory in your $PATH. Add to your .zshrc, .bashrc, ... following section
```
ap() {
    eval $(awx)
}
```

## Usage
```
ap
```

## Configuration
- to configure interactive look, see https://github.com/peco/peco
- to change default collors create $HOME/.peco/config.json file
```
mkdir -p $HOME/.peco
echo "{
    "Prompt": "[awx]",
    "Style": {
        "Basic": ["on_black", "white"],
        "SavedSelection": ["bold", "on_black", "green"],
        "Selected": ["bold", "on_black", "yellow"],
        "Query": ["yellow", "bold"],
        "Matched": ["yellow", "on_black"]
    }
}" > $HOME/.peco/config.json
```
