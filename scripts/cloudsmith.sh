cd dist
for i in *.apk; do
    [ -f "$i" ] || break
    cloudsmith push alpine --republish 1-platform/apic/alpine/any-version $i
done

for i in *.deb; do
    [ -f "$i" ] || break
    cloudsmith push deb --republish 1-platform/apic/any-distro/any-version $i
done

for i in *.rpm; do
    [ -f "$i" ] || break
    cloudsmith push rpm --republish 1-platform/apic/any-distro/any-version $i
done
