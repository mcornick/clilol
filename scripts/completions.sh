set -e
rm -rf completions
mkdir completions
for sh in bash zsh fish powershell; do
  go run . completion "$sh" >"completions/clilol.$sh"
done
