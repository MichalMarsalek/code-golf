FROM alpine:3.20 AS builder

RUN apk add --no-cache build-base curl

RUN curl -L https://github.com/cisco/ChezScheme/releases/download/v10.0.0/csv10.0.0.tar.gz | tar xz --strip-components 1

RUN ./configure     \
  --disable-curses  \
  --disable-x11     \
 && make -j `nproc` \
 && make install

FROM codegolf/lang-base

COPY --from=0 /bin/cat /bin/sh         /bin/
COPY --from=0 /boot/pb/petite.boot     \
              /boot/pb/scheme.boot     /boot/pb/
COPY --from=0 /boot/ta6le/petite.boot  \
              /boot/ta6le/scheme.boot  /boot/ta6le/
COPY --from=0 /lib/ld-musl-x86_64.so.1 /lib/
COPY --from=0 /usr/bin/scheme          /usr/bin/

COPY /schemewrapper /usr/bin/

ENTRYPOINT ["schemewrapper"]

CMD ["--version"]
