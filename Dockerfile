FROM alpine:3.7
LABEL Maintainer="AlicFeng <a@samego.com>" \
      Description="draw_roi based on golang"

ENV PATH "$PATH:/usr/local/draw_roi/bin"

WORKDIR /usr/local/draw_roi/

COPY bin /usr/local/draw_roi/bin
COPY resources/font/* /usr/local/draw_roi/fonts/

RUN chmod a+x /usr/local/draw_roi/bin/*

EXPOSE 1280

CMD ["draw_roi"]