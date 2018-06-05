FROM scratch

ADD ./view /view/
COPY main.out /

EXPOSE 9000
ENTRYPOINT ["/main.out"]