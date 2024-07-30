FROM golang:1.22.5-bullseye

# Set the Current Working Directory inside the container
WORKDIR /app

RUN apt update -y \
    && apt install -y make
#    && apt install -y supervisor

ENV NAME webdetect

#COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Add go mod and sum files
ADD . /app/

# Build the Go app
RUN go mod download
RUN make
RUN chmod +x /app/bin/webhook
EXPOSE 6969


# Command to run the executable
#CMD ["supervisord -c /etc/supervisor/supervisord.conf"]
ENTRYPOINT [ "/app/bin/webhook" ]