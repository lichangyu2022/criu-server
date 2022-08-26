FROM criu-base

WORKDIR /app

ENV PATH="${pwd}:${PATH}"

ADD criu* /app/

EXPOSE 8080

EXPOSE 3000

ENTRYPOINT ./criu-server
