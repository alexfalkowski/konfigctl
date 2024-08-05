include bin/build/make/client.mak
include bin/build/make/git.mak

# Generate the client.
generate:
	buf generate buf.build/afalkowski/konfig
