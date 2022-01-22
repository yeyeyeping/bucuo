FROM alpine:latest
WORKDIR /app
COPY bucuo /app/bucuo
RUN chmod +x bucuo
COPY conf/* /app/conf/
ENTRYPOINT /app/bucuo
EXPOSE 3002

