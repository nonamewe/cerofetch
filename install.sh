set -e
REPO="https://github.com/nonamewe/cerofetch.git"
DIR="/tmp/cerofetch"
echo "Installing cerofetch..."

if [ "$EUID" -ne 0 ]; then
    echo "Please run as root"
    exit 1
fi

command -v git >/dev/null || {
    echo "Git is required"
    exit 1
}
rm -rf "$DIR"
git clone "$REPO" "$DIR"

mkdir -p /usr/share/cerofetch
cp -r "$DIR/logo" /usr/share/cerofetch/

cd "$DIR"
go build -o cerofetch .
install -m 755 cerofetch /usr/bin/cerofetch
rm -rf "$DIR"

echo "cerofetch installed successfully!"
