#!/bin/sh

# remove any old stuff lying around
rm -rf .cache site

# build the docs
mkdocs build

# replace google fonts CDN with bunny CDN
# (attempting to not violate GDPR, and also to hell with Google)
for x in $(find site -type f | xargs grep fonts.googleapis.com | cut -d: -f1); do
  sed -E 's/fonts.(googleapis|gstatic).com/fonts.bunny.net/g' $x > $x.tmp
  mv $x.tmp $x
  chmod 644 $x
done

# send the docs off to the server
rsync -a --progress --delete site/ hel1.mcornick.dev:/srv/www/clilol

# clean up after ourselves
rm -rf .cache site
