FROM golang:1.15.10-alpine3.12

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ JST-9
ENV MECAB_VERSION 0.996
ENV IPADIC_VERSION 2.7.0-20070801
ENV mecab_url https://drive.google.com/uc?export=download&id=0B4y35FiV1wh7cENtOXlicTFaRUE
ENV ipadic_url https://drive.google.com/uc?export=download&id=0B4y35FiV1wh7MWVlSDBCSXZMTXM

RUN apk update && \
    apk upgrade && \
    apk add --no-cache build-base curl git g++ bash file openssl openssh-client sudo && \
    # Install MeCab
    curl -SL -o mecab-${MECAB_VERSION}.tar.gz ${mecab_url} && \
    tar zxf mecab-${MECAB_VERSION}.tar.gz && \
    cd mecab-${MECAB_VERSION} && \
    ./configure --enable-utf8-only --with-charset=utf8 && \
    make && \
    make install && \
    cd && \
    # Install IPA dic
    curl -SL -o mecab-ipadic-${IPADIC_VERSION}.tar.gz ${ipadic_url} && \
    tar zxf mecab-ipadic-${IPADIC_VERSION}.tar.gz && \
    cd mecab-ipadic-${IPADIC_VERSION} && \
    ./configure --with-charset=utf8 && \
    # compile
    make && \
    # install
    make install && \
    cd && \
    # Install Neologd
    git clone --depth 1 https://github.com/neologd/mecab-ipadic-neologd.git /mecab-ipadic-neologd && \
    /mecab-ipadic-neologd/bin/install-mecab-ipadic-neologd -n -a -y -p /usr/local/lib/mecab/dic/mecab-ipadic-neologd && \
    echo "dicdir = /usr/local/lib/mecab/dic/mecab-ipadic-neologd" > /usr/local/etc/mecabrc && \
    # Clean up
    rm -rf \
      mecab-${MECAB_VERSION}* \
      mecab-${IPADIC_VERSION}* \
      /mecab-ipadic-neologd

COPY startup.sh /startup.sh
RUN chmod 744 /startup.sh
CMD ["/startup.sh"]