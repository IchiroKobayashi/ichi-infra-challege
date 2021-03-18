FROM golang:1.15.10-alpine3.12

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ JST-9

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git g++ build-base bash curl xz openssl openssh-client sudo openblas-dev

RUN /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

RUN brew install mecab && brew install mecab-ipadic
RUN git clone --depth 1 git@github.com:neologd/mecab-ipadic-neologd.git /mecab-ipadic-neologd && \
    /mecab-ipadic-neologd/bin/install-mecab-ipadic-neologd -n -a -y -p /usr/local/lib/mecab/dic/mecab-ipadic-neologd

RUN echo "dicdir = /usr/local/lib/mecab/dic/mecab-ipadic-neologd" > /usr/local/etc/mecabrc

RUN rm -rf /mecab-ipadic-neologd

COPY startup.sh /startup.sh
RUN chmod 744 /startup.sh
CMD ["/startup.sh"]