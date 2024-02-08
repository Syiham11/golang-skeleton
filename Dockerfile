# builder image
FROM surnet/alpine-wkhtmltopdf:3.8-0.12.5-full as builder

# Image
FROM --platform=linux/amd64 golang:alpine 

# Install needed packages
ENV TZ=Asia/Jakarta
RUN apk add --update --no-cache git libc-dev tzdata
RUN apk update && apk add --no-cache git bash
RUN apk add --update make
RUN apk add mysql-client
RUN mkdir /greebel.core.be
RUN  echo "https://mirror.alwyzon.net/alpine/v3.8/main/" > /etc/apk/repositories \
     && echo "https://mirror.alwyzon.net/alpine/v3.8/community" >> /etc/apk/repositories \
     && apk update && apk add --no-cache \
      libstdc++ \
      libx11 \
      libxrender \
      libxext \
      libssl1.0 \
      ca-certificates \
      fontconfig \
      freetype \
      ttf-dejavu \
      ttf-droid \
      ttf-freefont \
      ttf-liberation \
      ttf-ubuntu-font-family \
    && apk add --no-cache --virtual .build-deps \
      msttcorefonts-installer \
    \
    # Install microsoft fonts
    && update-ms-fonts \
    && fc-cache -f \
    \
    # Clean up when done
    && rm -rf /var/cache/apk/* \
    && rm -rf /tmp/* \
    && apk del .build-deps

COPY --from=builder /bin/wkhtmltopdf /bin/wkhtmltopdf
COPY --from=builder /bin/wkhtmltoimage /bin/wkhtmltoimage

WORKDIR /greebel.core.be

COPY . /greebel.core.be

# COPY fonts/ /usr/share/fonts

 
RUN chmod 775 ./run.sh
RUN dos2unix ./run.sh

# download dependencies
RUN go mod tidy
# RUN go get github.com/swaggo/swag/cmd/swag

# generate swagger docs
# RUN swag init

ENTRYPOINT ./run.sh