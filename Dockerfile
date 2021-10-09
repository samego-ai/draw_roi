FROM alpine:3.7
LABEL Maintainer="AlicFeng <a@samego.com>" \
      Description="draw_roi based on golang"

COPY bin/draw_roi /usr/local/sbin/draw_roi

COPY resources/font/* /usr/share/fonts/

RUN chmod a+x /usr/local/sbin/draw_roi

EXPOSE 1280

CMD ["draw_roi"]