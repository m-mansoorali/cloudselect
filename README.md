# CloudSelect

An interactive CLI utility written in Go to select an OpenStack cloud from your `clouds.yaml` and automatically export it to your shell environment.

---

## 🚀 Features

- Parses `~/.config/openstack/clouds.yaml`
- Interactive cloud selection using `promptui`
- Outputs `export OS_CLOUD=<selected>` to stdout
- Updates `.zshrc` with the selected cloud
- Works on macOS and Linux

---

## 📦 Installation

```bash
git clone https://github.com/<your-username>/cloudselect.git
cd cloudselect
go build -o cloudselect

---
## 🧪 Usage
```bash
./cloudselect 

Then run:
```bash
eval $(./cloudselect 2>/dev/null | grep "^export OS_CLOUD")
Or open a new terminal after sourcing .zshrc
