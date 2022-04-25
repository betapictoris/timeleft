# Time Left [![Go](https://github.com/BetaPictoris/timeleft/actions/workflows/go.yml/badge.svg)](https://github.com/BetaPictoris/timeleft/actions/workflows/go.yml)
Display the time left of the day as a progress bar. 

![timeleft-colour](https://user-images.githubusercontent.com/65696362/155674372-203151f6-bf40-42ac-87a4-b5cb6b512a2d.png)
<!--![timeleft-nocolour](https://user-images.githubusercontent.com/65696362/155673053-091749e2-a455-4ad7-8019-a89969dcb6e4.png)-->

## Installation
### From release
```bash
curl -LO https://github.com/BetaPictoris/timeleft/releases/latest/download/timeleft    # Download the latest binary.
sudo install -Dt /usr/local/bin -m 755 timeleft                                        # Install Time Left to "/usr/local/bin" with the mode "755"
```

### Using [Goblin](https://github.com/barelyhuman/goblin)
```bash
curl -sf https://goblin.reaper.im/github.com/BetaPictoris/timeleft/src | OUT=timeleft sh
```

### Build from source 
```bash
git clone git@github.com:BetaPictoris/timeleft.git      # Clone the repository
cd timeleft                                             # Change into the repository's directory
bash build.sh                                           # Build Time Left
sudo install -Dt /usr/local/bin -m 755 timeleft         # Install Time Left to "/usr/local/bin" with the mode "755"
```

### User install
If you don't have access to `sudo` on your system you can install to your user's `~/.local/bin` directory with this command: 
```bash
install -Dt ~/.local/bin -m 755 timeleft
```
