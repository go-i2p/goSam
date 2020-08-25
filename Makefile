
USER_GH=eyedeekay
VERSION=0.32.23
packagename=gosam

echo:
	@echo "type make version to do release $(VERSION)"

version:
	gothub release -s $(GITHUB_TOKEN) -u $(USER_GH) -r $(packagename) -t v$(VERSION) -d "version $(VERSION)"

del:
	gothub delete -s $(GITHUB_TOKEN) -u $(USER_GH) -r $(packagename) -t v$(VERSION)

tar:
	tar --exclude .git \
		--exclude .go \
		--exclude bin \
		--exclude examples \
		-cJvf ../$(packagename)_$(VERSION).orig.tar.xz .

link:
	rm -f ../goSam
	ln -sf . ../goSam
