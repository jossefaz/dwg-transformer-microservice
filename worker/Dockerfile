FROM osgeo/gdal:alpine-normal-latest as base

FROM base as builder

ENV PYTHONUNBUFFERED=1



RUN echo "**** install Python ****" && \
    apk add --no-cache python3 && \
    if [ ! -e /usr/bin/python ]; then ln -sf python3 /usr/bin/python ; fi && \
    \
    apk --no-cache add --virtual .builddeps gcc gfortran musl-dev && \
    echo "**** install pip ****" && \
    python3 -m ensurepip && \
    rm -r /usr/lib/python*/ensurepip && \
    pip3 install --no-cache --upgrade pip setuptools wheel && \

    if [ ! -e /usr/bin/pip ]; then ln -s pip3 /usr/bin/pip ; fi && \
    apk del .builddeps     && rm -rf /root/.cache


FROM base

WORKDIR /app
RUN mkdir shared && \
    chmod 777 shared
COPY requirements.txt /requirements.txt
RUN     pip3 install -r /requirements.txt

COPY main.py ./

CMD ["python", "./main.py"]

